package strUtils

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type New struct {
	V string
}

func (str New) CamelToSnake() string {
	// Replace all capital letters with an underscore followed by the lowercase letter
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(str.V, "${1}_${2}")
	// Convert all characters to lowercase
	snake = strings.ToLower(snake)
	return snake
}

func CamelToSnake(str string) string {
	// Replace all capital letters with an underscore followed by the lowercase letter
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(str, "${1}_${2}")
	// Convert all characters to lowercase
	snake = strings.ToLower(snake)
	return snake
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func StringRandom(length int) string {
	return stringWithCharset(length, charset)
}

var letters = []rune("0123456789")

func RandomNumberString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[seededRand.Intn(len(letters))]
	}
	return string(b)
}
