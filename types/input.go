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

type PostInput struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type CommentInput struct {
	Content string `json:"content" validate:"required"`
	PostID  string `json:"post_id" validate:"required"`
}
