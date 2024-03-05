package dto

type UserRegistration struct {
	Email    string `validate:"required,email"`
	Login    string `validate:"required"`
	Password string `validate:"required"`
}
