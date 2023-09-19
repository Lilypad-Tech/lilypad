package http

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	stdlog "log"
	"net/http"
	"net/url"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/rs/zerolog/log"
)

// write some string constants for x-lilypad headers
// this is the address of the user
const X_LILYPAD_USER_HEADER = "X-Lilypad-User"

// this is the signature of the message
const X_LILYPAD_SIGNATURE_HEADER = "X-Lilypad-Signature"

// the context name we keep the address
const CONTEXT_ADDRESS = "address"

// the sub path any API's are served over
const API_SUB_PATH = "/api/v1"

// the sub path the websocket server is mounted on
const WEBSOCKET_SUB_PATH = "/ws"

type HTTPError struct {
	Message    string
	StatusCode int
}

type AuthUser struct {
	Address string `json:"address"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func getWsURL(url string) string {
	// replace http(s) with ws(s)
	// e.g. return strings.Replace(s, old, new, n)
	url = strings.Replace(url, "https://", "wss://", 1)
	url = strings.Replace(url, "http://", "ws://", 1)
	return url
}

// we assume both userPayload and signature are base64 encoded
func decodeUserAddress(userPayload string, signature string) (string, error) {
	decodedUserPayload, err := base64.StdEncoding.DecodeString(userPayload)
	if err != nil {
		return "", err
	}
	decodedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return "", err
	}
	address, err := web3.GetAddressFromSignedMessage(decodedUserPayload, decodedSignature)
	if err != nil {
		return "", err
	}
	return address.String(), nil
}

// returns userPayload and signature as strings ready to be written into request headers
// we encode these both as base64 so they can be included in http headers
func encodeUserAddress(privateKey *ecdsa.PrivateKey, address string) (string, string, error) {
	user := AuthUser{
		Address: address,
	}
	userBytes, err := json.Marshal(user)
	if err != nil {
		return "", "", err
	}
	userSignature, err := web3.SignMessage(privateKey, userBytes)
	if err != nil {
		return "", "", err
	}

	return base64.StdEncoding.EncodeToString(userBytes), base64.StdEncoding.EncodeToString(userSignature), nil
}

func AddHeaders(
	req *retryablehttp.Request,
	privateKey *ecdsa.PrivateKey,
	address string,
) error {
	userPayload, userSignature, err := encodeUserAddress(privateKey, address)
	if err != nil {
		return err
	}
	req.Header.Add(X_LILYPAD_USER_HEADER, userPayload)
	req.Header.Add(X_LILYPAD_SIGNATURE_HEADER, userSignature)
	return nil
}

// this will use the client headers to ensure that a message was signed
// by the holder of a private key for a specific address
// there is a "X-Lilypad-User" header that will contain the address
// there is a "X-Lilypad-Signature" header that will contain the signature
// we use the signature to verify that the message was signed by the private key
func GetAddressFromHeaders(req *http.Request) (string, error) {
	userHeader := req.Header.Get(X_LILYPAD_USER_HEADER)
	if userHeader == "" {
		return "", HTTPError{
			Message:    "missing user header",
			StatusCode: http.StatusUnauthorized,
		}
	}
	userSignature := req.Header.Get(X_LILYPAD_SIGNATURE_HEADER)
	if userSignature == "" {
		return "", HTTPError{
			Message:    "missing signature header",
			StatusCode: http.StatusUnauthorized,
		}
	}

	// decode the userHeader into a AuthUser struct
	// let's remember this is in base64 format
	decodedUserHeader, err := base64.StdEncoding.DecodeString(userHeader)
	if err != nil {
		return "", HTTPError{
			Message:    fmt.Sprintf("invalid user header %s", err.Error()),
			StatusCode: http.StatusUnauthorized,
		}
	}

	var authUser AuthUser
	err = json.Unmarshal(decodedUserHeader, &authUser)
	if err != nil {
		return "", HTTPError{
			Message:    fmt.Sprintf("invalid user header %s", err.Error()),
			StatusCode: http.StatusUnauthorized,
		}
	}

	signatureAddress, err := decodeUserAddress(userHeader, userSignature)
	if err != nil {
		return "", HTTPError{
			Message:    fmt.Sprintf("invalid user header or signature %s", err.Error()),
			StatusCode: http.StatusUnauthorized,
		}
	}

	if signatureAddress != authUser.Address {
		return "", HTTPError{
			Message:    "invalid signature",
			StatusCode: http.StatusUnauthorized,
		}
	}

	return signatureAddress, nil
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(res, req)
	})
}

func URL(options ClientOptions, path string) string {
	return fmt.Sprintf("%s%s%s", options.URL, API_SUB_PATH, path)
}

func WebsocketURL(options ClientOptions, path string) string {
	return getWsURL(URL(options, path))
}

type httpGetWrapper[ResultType any] func(res http.ResponseWriter, req *http.Request) (ResultType, error)
type httpPostWrapper[RequestType any, ResultType any] func(data RequestType, res http.ResponseWriter, req *http.Request) (ResultType, error)

func ReadBody[T any](req *http.Request) (T, error) {
	var data T
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// wrap a http handler with some error handling
// so if it returns an error we handle it
func GetHandler[T any](handler httpGetWrapper[T]) func(res http.ResponseWriter, req *http.Request) {
	ret := func(res http.ResponseWriter, req *http.Request) {
		data, err := handler(res, req)
		if err != nil {
			log.Error().
				Str("method GET", req.URL.String()).
				Err(err).
				Msgf("")
			httpError, ok := err.(HTTPError)
			if ok {
				http.Error(res, httpError.Error(), httpError.StatusCode)
			} else {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			}
			return
		} else {
			// get is trace because it does not mutate
			log.Trace().
				Str("method GET", req.URL.String()).
				Str("res", fmt.Sprintf("%+v", data)).
				Msgf("")
			err = json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Ctx(req.Context()).Error().Msgf("error for json encoding: %s", err.Error())
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	return ret
}

func PostHandler[RequestType any, ResultType any](handler httpPostWrapper[RequestType, ResultType]) func(res http.ResponseWriter, req *http.Request) {
	ret := func(res http.ResponseWriter, req *http.Request) {
		requestBody, err := ReadBody[RequestType](req)
		if err != nil {
			http.Error(res, fmt.Sprintf("Error parsing request body"), http.StatusBadRequest)
			return
		}
		data, err := handler(requestBody, res, req)
		if err != nil {
			log.Error().
				Str("method POST", req.URL.String()).
				Err(err).
				Msgf("")
			httpError, ok := err.(HTTPError)
			if ok {
				http.Error(res, httpError.Error(), httpError.StatusCode)
			} else {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			}
			return
		} else {
			// post is debug because it does mutate
			log.Debug().
				Str("method POST", req.URL.String()).
				Str("req", fmt.Sprintf("%+v", requestBody)).
				Str("res", fmt.Sprintf("%+v", data)).
				Msgf("")
			err = json.NewEncoder(res).Encode(data)
			if err != nil {
				log.Ctx(req.Context()).Error().Msgf("error for json encoding: %s", err.Error())
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	return ret
}

func GetRequest[ResultType any](
	options ClientOptions,
	path string,
	queryParams map[string]string,
) (ResultType, error) {
	var result ResultType
	buf, err := GetRequestBuffer(
		options,
		path,
		queryParams,
	)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(buf.Bytes(), &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetRequestBuffer(
	options ClientOptions,
	path string,
	queryParams map[string]string,
) (*bytes.Buffer, error) {
	client := newRetryClient()

	parsedURL, err := url.Parse(URL(options, path))
	if err != nil {
		return nil, err
	}

	urlValues := url.Values{}
	for key, value := range queryParams {
		urlValues.Add(key, value)
	}
	parsedURL.RawQuery = urlValues.Encode()

	req, err := retryablehttp.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func PostRequest[RequestType any, ResultType any](
	options ClientOptions,
	path string,
	data RequestType,
) (ResultType, error) {
	var result ResultType
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return result, fmt.Errorf("THIS IS A JOSN ERROR: %s", err.Error())
	}
	return PostRequestBuffer[ResultType](
		options,
		path,
		bytes.NewBuffer(dataBytes),
	)
}

func PostRequestBuffer[ResultType any](
	options ClientOptions,
	path string,
	data *bytes.Buffer,
) (ResultType, error) {
	var result ResultType
	client := newRetryClient()
	privateKey, err := web3.ParsePrivateKey(options.PrivateKey)
	if err != nil {
		return result, err
	}
	req, err := retryablehttp.NewRequest("POST", URL(options, path), data)
	if err != nil {
		return result, err
	}
	AddHeaders(req, privateKey, web3.GetAddress(privateKey).String())
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// parse body as json into result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func newRetryClient() *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	retryClient.Logger = stdlog.New(io.Discard, "", stdlog.LstdFlags)
	retryClient.RequestLogHook = func(_ retryablehttp.Logger, req *http.Request, attempt int) {
		switch {
		case req.Method == "POST":
			log.Debug().
				Str(req.Method, req.URL.String()).
				Int("attempt", attempt).
				Msgf("")
		default:
			// GET, PUT, DELETE, etc.
			log.Trace().
				Str(req.Method, req.URL.String()).
				Int("attempt", attempt).
				Msgf("")
		}
	}
	return retryClient
}
