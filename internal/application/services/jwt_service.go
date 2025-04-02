package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // Replace with a secure secret key

func GenerateToken(userID uint, role string) (string, string, error) {
	// Define access token claims
	accessTokenClaims := jwt.MapClaims{
		"user":    userID,
		"role":    role,
		"refresh": false,
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // Access token expires in 15 Minutes
	}

	// Create access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	// Define refresh token claims
	refreshTokenClaims := jwt.MapClaims{
		"user":    userID,
		"refresh": true,
		"exp":     time.Now().Add(time.Hour * 24 * 3).Unix(), // Refresh token expires in 3 days
	}

	// Create refresh token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func RefreshToken(refreshTokenString string) (string, string, error) {
	// Validate the refresh token
	token, err := ValidateToken(refreshTokenString)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", errors.New("invalid refresh token claims")
	}

	// Ensure the token is a refresh token
	isRefresh, ok := claims["refresh"].(bool)
	if !ok || !isRefresh {
		return "", "", errors.New("provided token is not a refresh token")
	}

	// Check if the token has expired
	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		return "", "", errors.New("refresh token has expired")
	}

	// Extract user ID from claims
	userIDFloat, ok := claims["user"].(float64)
	if !ok {
		return "", "", errors.New("invalid user ID in refresh token")
	}
	userID := uint(userIDFloat)

	// Generate new tokens
	return GenerateToken(userID, claims["role"].(string))
}
