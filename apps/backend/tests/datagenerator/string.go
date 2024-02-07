package datagenerator

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func randomLowercaseChar() byte {
	return byte(rand.Intn(26) + 97) // lowercase starts from 65
}

func RandomEmail(usernameLength, domainLength int) string {
	rand.Seed(time.Now().UnixNano())

	// Generate random username
	username := make([]byte, usernameLength)
	for i := range username {
		username[i] = randomLowercaseChar()
	}

	// Generate random domain
	domain := make([]byte, domainLength)
	for i := range domain {
		domain[i] = randomLowercaseChar()
	}

	extensions := []string{"com", "net", "org", "edu", "hu"}
	extension := extensions[rand.Intn(len(extensions))]

	email := fmt.Sprintf("%s@%s.%s", string(username), string(domain), extension)
	return email
}
