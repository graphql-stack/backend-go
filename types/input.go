package types

// RegisterInput is form binding struct for register
type RegisterInput struct {
	Name     string `json:"name" binding:"omitempty,min=6,max=16"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

// LoginInput is form binding struct for login
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
