// validators/gin_validator.go
package validators

import (
	"github.com/go-playground/validator/v10" // 수정됨
)

// GinValidator는 Gin의 binding.Validator 인터페이스를 구현합니다.
type GinValidator struct {
    validator *validator.Validate
}

// ValidateStruct는 구조체를 검증합니다.
func (v *GinValidator) ValidateStruct(obj interface{}) error {
    return v.validator.Struct(obj)
}

// Engine는 GinValidator를 반환합니다.
func (v *GinValidator) Engine() interface{} {
    return v.validator
}

// NewGinValidator는 새로운 GinValidator 인스턴스를 생성합니다.
func NewGinValidator() *GinValidator {
    v := validator.New()
    RegisterCustomValidators(v)
    return &GinValidator{validator: v}
}
