package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"unicode"
)

func main() {

	randomPassword, err := passwordGenerator(8)
	if err != nil {
		panic(err)
	}

	fmt.Println("Password Generated: ", string(randomPassword))
}

const DefaultPasswordLength = 8
const SpecialChars = "!@#$%^&*()-_=+,.?/:;{}[]~"
const Numbers = "0123456789"
const Letters = "abcdfghijklmnopqrstuvwxyzABCDFGHIJKLMNOPQRSTUVWXYZ"

func passwordGenerator(passwordLength int) (string, error) {
	if passwordLength == 0 {
		passwordLength = DefaultPasswordLength
	}

	validChars := Letters + Numbers + SpecialChars

	password := make([]byte, passwordLength)

	//Fill the password with valid chars from Letters, Number & SpecialChars
	for i := 0; i < passwordLength; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(validChars))))
		if err != nil {
			return "", err
		}
		password[i] = validChars[int(index.Int64())]
	}

	if len(password) != passwordLength {
		return "", errors.New("password length is wrong")
	}

	if !containsUpperCase(password) || !containsLowerCase(password) || !containsNumbers(password) || !containsSpecialChar(password) {
		return passwordGenerator(passwordLength)
	}

	return string(password), nil
}

func containsUpperCase(s []byte) bool {
	for _, b := range s {
		if unicode.IsUpper(rune(b)) {
			fmt.Println("containsUpperCase: ", string(b))
			return true
		}
	}
	return false
}

func containsLowerCase(s []byte) bool {
	for _, b := range s {
		if unicode.IsLower(rune(b)) {
			fmt.Println("containsLowerCase: ", string(b))
			return true
		}
	}
	return false
}

func containsNumbers(s []byte) bool {
	for _, b := range s {
		if unicode.IsNumber(rune(b)) {
			fmt.Println("containsNumbers: ", string(b))
			return true
		}
	}
	return false
}

func containsSpecialChar(s []byte) bool {
	for _, b := range s {
		if unicode.IsPunct(rune(b)) {
			fmt.Println("containsSpecialChar: ", string(b))
			return true
		}
	}
	return false
}
