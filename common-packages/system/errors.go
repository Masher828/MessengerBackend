package system

import "errors"

var (
	InvalidNameErr           = errors.New("please enter a valid name of minimum length of 2")
	InvalidEmailErr          = errors.New("please enter a valid email")
	InvalidPasswordFormatErr = errors.New("please enter a password of length (8,20)")
	InvalidContactNumberErr  = errors.New("please enter a valid contact number")
	InvalidCredentialsErr    = errors.New("invalid email id & password")
)
