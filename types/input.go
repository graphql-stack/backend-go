package types

// RegisterInput is form binding struct for register
type RegisterInput struct {
	Name     string `json:"name" validate:"required,min=6,max=16"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=16"`
	Avatar   string `json:"avatar" validate:"required,url"`
}

// LoginInput is form binding struct for login
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=16"`
}
