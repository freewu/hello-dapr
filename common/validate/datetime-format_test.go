package validate

import (
	"fmt"
	"regexp"
	"testing"
)

func TestValidateDateFormat(t *testing.T) {
	matched, err := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, "2021-06-01")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)

	matched, err = regexp.MatchString(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-3][0-9])$`, "2021-06-01")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)
}


func TestValidateDateTimeFormat(t *testing.T) {
	matched, err := regexp.MatchString(`^\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[1-3][0-9])\s+[0-2][0-9]:[0-5][0-9]:[0-5][0-9]$`, "2021-06-01 01:02:59")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)
}

func TestValidateTimeFormat(t *testing.T) {
	matched, err := regexp.MatchString(`^[0-2][0-9]:[0-5][0-9]:[0-5][0-9]$`, "01:02:59")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matched)
}