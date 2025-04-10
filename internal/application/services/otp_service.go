package services

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// Generate OTP and store it in Redis
func GenerateOTP(email string) (string, string, error) {
	var otpInt uint32
	if err := binary.Read(rand.Reader, binary.LittleEndian, &otpInt); err != nil {
		return "", "", fmt.Errorf("failed to generate OTP: %v", err)
	}
	otp := fmt.Sprintf("%06d", otpInt%1000000)

	// Generate a random reference string for the OTP
	refBytes := make([]byte, 4)
	if _, err := rand.Read(refBytes); err != nil {
		return "", "", fmt.Errorf("failed to generate reference code: %v", err)
	}
	refCode := fmt.Sprintf("%X", refBytes)

	// Store OTP and reference code in Redis with a 5-minute expiration
	otpWithRef := fmt.Sprintf("%s|%s", otp, refCode)
	err := redisClient.Set(ctx, email, otpWithRef, 5*time.Minute).Err()
	if err != nil {
		return "", "", fmt.Errorf("failed to store OTP and reference code: %v", err)
	}

	return otp, refCode, nil
}

// Validate OTP from Redis
func ValidateOTP(email, otp string) (bool, error) {
	// Retrieve OTP from Redis
	storedOTP, err := redisClient.Get(ctx, email).Result()
	if err == redis.Nil {
		return false, fmt.Errorf("OTP not found or expired")
	} else if err != nil {
		return false, fmt.Errorf("failed to retrieve OTP: %v", err)
	}

	// Check if the provided OTP matches
	if storedOTP != otp {
		return false, fmt.Errorf("provided OTP does not match")
	}

	// Delete OTP from Redis after successful validation
	redisClient.Del(ctx, email)

	return true, nil
}
