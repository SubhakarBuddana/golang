package validations

import (
	"errors"

	"github.com/SubhakarBuddana/golang/types"
)

func SignUpValidation(person types.Person) error {
	if person.Username == "" {
		return errors.New("Invalid Username")
	}
	if person.Firstname == "" {
		return errors.New("Invalid Firstname")
	}
	if person.Lastname == "" {
		return errors.New("Invalid Lastname")
	}
	if person.Email == "" {
		return errors.New("Invalid Email")
	}
	if person.Password == "" {
		return errors.New("Password must not be empty")
	}
	if person.Password != person.Confirmpassword {
		return errors.New("Password does not match")
	}
	return nil

}

func ChangePasswordRequestValidation(request types.ChangePasswordRequest) error {
	if request.Username == "" {
		return errors.New("Username missing")
	}
	if request.Oldpassword == "" {
		return errors.New("Password missing")
	}
	if request.Newpassword == "" {
		return errors.New("New Password missing")
	}
	return nil

}

func LoginRequestValidation(loginuser types.Login) error {
	if loginuser.Username == "" {
		return errors.New(("Username missing"))
	}
	if loginuser.Password == "" {
		return errors.New("Password missing")
	}
	return nil
}
