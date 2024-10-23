// handlers/dto/login_input.go
package dto

type LoginInput struct {
    Username string `json:"username" binding:"required,min=3,max=16" validate:"required,min=3,max=16"`
    Password string `json:"password" binding:"required"`
}
