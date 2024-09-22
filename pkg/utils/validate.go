package utils

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"regexp"
)

// ValidatePassword 验证密码，长度至少8位，并且包含字母和数字
func ValidatePassword(password string) error {
	// 检查密码长度
	if !govalidator.StringLength(password, "8", "100") {
		return fmt.Errorf("密码长度至少为8位")
	}

	// 使用正则表达式检查是否包含字母和数字
	hasLetter := regexp.MustCompile(`[a-zA-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	if !hasLetter || !hasDigit {
		return fmt.Errorf("密码必须包含字母和数字")
	}

	return nil
}
