package util

import (
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	usernameLength = 10
	passwordLength = 12
)

var letters = "abcdefghijklmnopqrstuvwxyz"
var digits = "0123456789"
var symbols = "!@#$%^&*"

// RandomUUID generates a random UUID string
func RandomUUID() string {
	return uuid.New().String()
}

// RandomUsername generates a random username consisting of lowercase letters
func RandomUsername() string {
	return randomString(usernameLength, letters)
}

// RandomEmail generates a random email address
func RandomEmail() string {
	username := RandomUsername()
	return fmt.Sprintf("%s@example.com", username)
}

// RandomPassword generates a random password with letters, digits, and symbols
func RandomPassword() string {
	charset := letters + strings.ToUpper(letters) + digits + symbols
	return randomString(passwordLength, charset)
}

// randomString generates a random string of given length from the charset
func randomString(length int, charset string) string {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))
	for i := range result {
		randomIndex, _ := rand.Int(rand.Reader, charsetLen)
		result[i] = charset[randomIndex.Int64()]
	}
	return string(result)
}

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPasswordHash compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns nil on success, or an error if not matching.
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// RandomTitle generates a random question title
func RandomTitle() string {
	titles := []string{
		"Two Sum",
		"Reverse String",
		"Palindrome Check",
		"Fibonacci Sequence",
		"Binary Search",
		"Sorting Algorithm",
		"Graph Traversal",
		"Dynamic Programming",
		"Array Manipulation",
		"String Processing",
	}
	return titles[randInt(len(titles))]
}

// RandomDescription generates a random question description
func RandomDescription() string {
	descriptions := []string{
		"Given an array of integers, return indices of the two numbers such that they add up to a specific target.",
		"Write a function that reverses a string. The input string is given as an array of characters.",
		"Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.",
		"Compute the nth Fibonacci number using dynamic programming.",
		"Implement binary search to find the target value in a sorted array.",
		"Sort an array of integers using quicksort algorithm.",
		"Traverse a graph using breadth-first search (BFS).",
		"Solve the knapsack problem using dynamic programming.",
		"Manipulate arrays to perform operations like insertion and deletion.",
		"Process strings to find substrings or patterns.",
	}
	return descriptions[randInt(len(descriptions))]
}

// RandomDifficulty generates a random difficulty level
func RandomDifficulty() string {
	difficulties := []string{"easy", "medium", "hard"}
	return difficulties[randInt(len(difficulties))]
}

// RandomTags generates random tags as a comma-separated string or null
func RandomTags() sql.NullString {
	if randInt(2) == 0 { // 50% chance to be null
		return sql.NullString{Valid: false}
	}
	tags := []string{
		"array", "string", "dynamic-programming", "math", "graph", "tree", "sorting", "searching",
	}
	numTags := randInt(3) + 1 // 1 to 3 tags
	selectedTags := make([]string, numTags)
	for i := 0; i < numTags; i++ {
		selectedTags[i] = tags[randInt(len(tags))]
	}
	return sql.NullString{String: strings.Join(selectedTags, ","), Valid: true}
}

// randInt generates a random integer between 0 and max-1
func randInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}
