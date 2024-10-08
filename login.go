package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var tokens []string

var tokenLength = 4
var bearerLength = 6

func generateToken(w http.ResponseWriter, r *http.Request) {
	token, _ := randomHex(tokenLength)
	tokens = append(tokens, token)
	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(token)
	if encodeErr != nil {
		fmt.Printf("error (generateToken): %s", encodeErr.Error())
		http.Error(w, "Error!", http.StatusSeeOther)
	}
}

func validateToken(r *http.Request) bool {
	bearerToken := r.Header.Get("Authorization")
	fmt.Println(bearerToken)
	if len(bearerToken) < ((tokenLength * 2) + bearerLength + 1) { //+1 is for the space
		return false
	}
	reqToken := strings.Split(bearerToken, " ")[1]
	for _, token := range tokens {
		if token == reqToken {
			return true
		}
	}
	return false
}

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
