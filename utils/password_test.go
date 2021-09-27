package utils

import "testing"

func TestPassword(t *testing.T) {
	p := "laslk3432k"

	// Generate the hashed password
	hashedPassword, err := PasswordHash(p)
	if err != nil {
		t.Errorf("There was an error hashing the password: %v", err)
	}
	if hashedPassword == "" {
		t.Errorf("The hashed password can not be empty")
	}
	t.Logf("Password hashed succesfully! :) \n HashedPassword: %v", hashedPassword)

	// Check the password
	err = PasswordCheck(p, hashedPassword)
	if err != nil {
		t.Errorf("The password check failed: %v", err)
	}
	t.Log("Password checked succesfully! :)")

}
