package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 时间格式校验 yyyy-mm-dd HH:ii:ss yyyy-MM-dd 都可以
func ValidateDateTimeFormat(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	// 匹配 yyyy-mm-dd
	matched, err := regexp.MatchString(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-3][0-9])$`, val)
	if err != nil {
		return false
	}
	if matched {
		return matched
	}
	// 匹配 yyyy-mm-dd HH:ii:ss
	matched, err = regexp.MatchString(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-3][0-9])\s+[0-2][0-9]:[0-5][0-9]:[0-5][0-9]$`, val)
	if err != nil {
		return false
	}
	return matched
}

// 时间格式校验格式必须为 yyyy-MM-dd
func ValidateDateFormat(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	matched, err := regexp.MatchString(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-3][0-9])$`, val)
	if err != nil {
		return false
	}
	return matched
}

// 时间格式校验必须为 HH:ii:ss
func ValidateTimeFormat(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	matched, err := regexp.MatchString(`^[0-2][0-9]:[0-5][0-9]:[0-5][0-9]$`, val)
	if err != nil {
		return false
	}
	return matched
}