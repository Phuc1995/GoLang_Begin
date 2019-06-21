package error

import "errors"

type ValidationError error

var(
	ErrNoUsername       = ValidationError(errors.New("You must supply a username"))
	ErrNoEmail          = ValidationError(errors.New("You must supply an email"))
	ErrNoPassword       = ValidationError(errors.New("You must supply a password"))
	ErrPasswordTooShort = ValidationError(errors.New("Your password is too short"))
	ErrUsernameExists   = ValidationError(errors.New("That username is taken"))
	ErrEmailExists      = ValidationError(errors.New("That email address has an account"))
	ErrCredentialsIncorrect = ValidationError(errors.New("We couldn't find a user with the supplied and pass word combination"))
	)

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}