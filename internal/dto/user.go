package dto

type UserRegistration struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
