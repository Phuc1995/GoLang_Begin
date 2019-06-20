package user

import (
	error2 "error"
	"fmt"
	"generateId"
	"github.com/giantswarm/go.crypto/bcrypt"
)

type User struct {
	ID             string
	Email          string
	HashedPassword string
	Username       string
}

const (
	hashCost       = 10
	passwordLength = 6
	userIDLength   = 16
)

func NewUser(username, email, password string) (User, error) {
	user := User{
		Email: email,
		Username: username,
	}
	if username == ""{
		return user, error2.ErrNoUsername
	}
	if email == "" {
		return user, error2.ErrNoEmail
	}

	if password == "" {
		return user, error2.ErrNoPassword
	}

	if len(password) < passwordLength {
		return user, error2.ErrPasswordTooShort
	}

	//Check if the username exists
	existingUser, err := GlobalUserStore.FindByUsername(username)
	fmt.Print("existingUser: ",existingUser)
	if err != nil{
		return user, err
	}
	if existingUser != nil{
		return user, error2.ErrEmailExists
	}

	// Check if the email exists
	existingUser, err = GlobalUserStore.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, error2.ErrEmailExists
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	user.HashedPassword = string(hashedPassword)
	user.ID = generateId.GenerateID("usr", userIDLength)
	return user, err
}