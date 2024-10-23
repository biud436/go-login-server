// validators/custom_validators.go
package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10" // 수정됨
)

// NoSpecialChars는 문자열에 특수 문자가 포함되지 않았는지 검증합니다.
func NoSpecialChars(fl validator.FieldLevel) bool {
    str := fl.Field().String()
    matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, str)
    return matched
}

// RegisterCustomValidators는 커스텀 검증기를 등록합니다.
func RegisterCustomValidators(v *validator.Validate) {
    v.RegisterValidation("no_special_chars", NoSpecialChars)
}
