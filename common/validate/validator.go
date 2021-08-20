package validate

import (
	"errors"
	. "github.com/go-playground/validator/v10"
)

var validate *Validate

func init() {
	validate = New()
	// 注册自定义的验证
	_ = RegisterValidation("datetime-format", ValidateDateTimeFormat)
	_ = RegisterValidation("date-format", ValidateDateFormat)
	_ = RegisterValidation("time-format", ValidateTimeFormat)
}

// 注册自定义验证方法
func RegisterValidation(tag string, fn Func,callValidationEvenIfNull ...bool) error {
	return validate.RegisterValidation(tag, fn,callValidationEvenIfNull...)
}

// 结构体验证
func Struct(s interface{}) error {
	err := validate.Struct(s)
	// 如果没有验证错误直接返回
	if err == nil {
		return err
	}

	for _, err := range err.(ValidationErrors) {
		return errors.New(Uncapitalize(err.Field()) + " " + err.ActualTag())
		//fmt.Println(err.Namespace())
		//fmt.Println(err.Field())
		//fmt.Println(err.StructNamespace())
		//fmt.Println(err.StructField())
		//fmt.Println(err.Tag())
		//fmt.Println(err.ActualTag())
		//fmt.Println(err.Kind())
		//fmt.Println(err.Type())
		//fmt.Println(err.Value())
		//fmt.Println(err.Param())
		//fmt.Println()
	}
	return err
}

// 将字符首字母小写
func Uncapitalize(str string) string {
	var str1 string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 90 {
				vv[i] += 32 // string的码表相差32位
				str1 += string(vv[i])
			} else {
				//fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			str1 += string(vv[i])
		}
	}
	return str1
}