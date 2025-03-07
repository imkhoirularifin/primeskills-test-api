package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a plain text password and returns a hashed version of it using bcrypt.
// It returns the hashed password as a string and an error if the hashing process fails.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// ComparePassword takes a plain text password and a hashed password,
// and compares them using bcrypt. It returns an error if the passwords
// do not match or if there is an issue with the comparison process.
func ComparePassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return bcrypt.ErrMismatchedHashAndPassword
	}

	return nil
}
