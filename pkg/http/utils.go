package http

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bacalhau-project/lilypad/pkg/web3"
	"github.com/rs/zerolog/log"
)

// write some string constants for x-lilypad headers
// this is the address of the user
const X_LILYPAD_USER_HEADER = "X-Lilypad-User"

// this is the signature of the message
const X_LILYPAD_SIGNATURE_HEADER = "X-Lilypad-Signature"

// the context name we keep the address
const CONTEXT_ADDRESS = "address"

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

func decodeUserAddress(userPayload string, signature string) (string, error) {
	address, err := web3.GetAddressFromSignedMessage([]byte(userPayload), []byte(signature))
	if err != nil {
		return "", err
	}
	return address.String(), nil
}

// returns userPayload and signature as strings ready to be written into request headers
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
	return string(userBytes), string(userSignature), nil
}

func AddHeaders(
	req *http.Request,
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
	var authUser AuthUser
	err := json.Unmarshal([]byte(userHeader), &authUser)
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

type httpWrapper[T any] func(res http.ResponseWriter, req *http.Request) (T, error)

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
func Wrapper[T any](handler httpWrapper[T]) func(res http.ResponseWriter, req *http.Request) {
	ret := func(res http.ResponseWriter, req *http.Request) {
		data, err := handler(res, req)
		if err != nil {
			log.Ctx(req.Context()).Error().Msgf("error for route: %s", err.Error())
			httpError, ok := err.(HTTPError)
			if ok {
				http.Error(res, httpError.Error(), httpError.StatusCode)
			} else {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			}
			return
		} else {
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

func Get[ResultType any](
	options ClientOptions,
	path string,
) (ResultType, error) {
	var result ResultType
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", options.URL, path), nil)
	if err != nil {
		return result, err
	}

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

func Post[RequestType any, ResultType any](
	options ClientOptions,
	path string,
	data RequestType,
) (ResultType, error) {
	var result ResultType
	client := &http.Client{}
	privateKey, err := web3.ParsePrivateKey(options.PrivateKey)
	if err != nil {
		return result, err
	}
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return result, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", options.URL, path), bytes.NewBuffer(dataBytes))
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
