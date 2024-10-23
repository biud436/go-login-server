// validators/custom_validators.go
package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegisterCustomValidators(v *validator.Validate) {
    // 사용자 이름에 특수 문자가 포함되지 않도록 검증
    v.RegisterValidation("no_special_chars", func(fl validator.FieldLevel) bool {
        username := fl.Field().String()
        matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, username)
        return matched
    })
}
