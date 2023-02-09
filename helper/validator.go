package helper

import "errors"

func ValidateRegisterRequest(nama string, email string, password string) error {
	if nama == "" {
		return errors.New("nama tidak boleh kosomg")
	}
	if email == "" {
		return errors.New("email tidak boleh kosomg")
	}
	if password == "" {
		return errors.New("password tidak boleh kosomg")
	}
	// if !IsValidEmail(email) {
	// 	return errors.New("Invalid email address")
	// }
	// if !IsStrongPassword(password) {
	// 	return errors.New("Password is not strong enough")
	// }
	return nil
}
