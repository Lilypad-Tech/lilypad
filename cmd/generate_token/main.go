package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type DurationFlag time.Duration

func (d *DurationFlag) String() string {
	return time.Duration(*d).String()
}

func (d *DurationFlag) Set(value string) error {
	if strings.HasSuffix(value, "y") {
		years, err := time.ParseDuration(strings.TrimSuffix(value, "y") + "h")
		if err != nil {
			return err
		}
		*d = DurationFlag(years * 24 * 365) // approximate year as 365 days
		return nil
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		return err
	}
	*d = DurationFlag(duration)
	return nil
}

type JWTKey struct {
	ID     string
	Secret []byte
}

func readJWTKey(keyID string) (*JWTKey, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	// Convert key ID to environment variable format (e.g., "local" -> "JWT_KEY_LOCAL_SECRET")
	envKeyPrefix := fmt.Sprintf("JWT_KEY_%s", strings.ToUpper(keyID))

	// Get the actual key ID from env
	actualKeyID := os.Getenv(envKeyPrefix + "_ID")
	if actualKeyID == "" {
		return nil, fmt.Errorf("%s_ID not found in .env file", envKeyPrefix)
	}

	// Get the secret
	secret := os.Getenv(envKeyPrefix + "_SECRET")
	if secret == "" {
		return nil, fmt.Errorf("%s_SECRET not found in .env file", envKeyPrefix)
	}

	return &JWTKey{
		ID:     actualKeyID,
		Secret: []byte(secret),
	}, nil
}

// GenerateToken creates a new JWT token for the given subject with a specified expiration duration, secret and key ID
func GenerateToken(subject string, expiry time.Duration, secret []byte, keyID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   subject,
		"iss":   "kafka-auth",
		"aud":   "kafka-broker",
		"scope": "kafka-cluster",
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(expiry).Unix(),
		"jti":   uuid.New().String(),
	})

	// Add the key ID to the token header
	token.Header["kid"] = keyID

	// Sign the token with the secret
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func main() {
	subject := flag.String("subject", "", "Subject for the JWT token")
	keyID := flag.String("kid", "", "Key ID for the JWT token (e.g., local, test)")
	var duration DurationFlag = DurationFlag(24 * time.Hour)
	flag.Var(&duration, "duration", "Token expiry duration (e.g., 1y, 24h, 1h30m, 90m)")
	flag.Parse()

	if *subject == "" {
		fmt.Println("Error: subject is required")
		flag.Usage()
		os.Exit(1)
	}

	if *keyID == "" {
		fmt.Println("Error: kid (Key ID) is required")
		flag.Usage()
		os.Exit(1)
	}

	jwtKey, err := readJWTKey(*keyID)
	if err != nil {
		fmt.Printf("Error reading JWT key: %v\n", err)
		os.Exit(1)
	}

	token, err := GenerateToken(*subject, time.Duration(duration), jwtKey.Secret, jwtKey.ID)
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated JWT Token:\n%s\n", token)
}
