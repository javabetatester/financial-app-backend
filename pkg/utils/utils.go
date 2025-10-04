package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GenerateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return emailRegex.MatchString(strings.ToLower(email))
}

func FormatCurrency(amount float64) string {
	return "R$ " + formatFloat(amount, 2)
}

func formatFloat(num float64, precision int) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", num), "0"), ".")
}

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func StringPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func BoolPtr(b bool) *bool {
	return &b
}