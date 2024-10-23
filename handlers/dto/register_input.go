// handlers/dto/register_input.go
package dto

type RegisterInput struct {
    Username string `json:"username" binding:"required,min=3,max=16,no_special_chars" validate:"required,min=3,max=50,no_special_chars"`
    Password string `json:"password" binding:"required,min=8" validate:"required,min=8"`
}
