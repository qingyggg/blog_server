package utils

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// ConvertStringToOrdered  将字符串转换为指定类型的值
func ConvertStringToOrdered[T constraints.Ordered](s string, v *T) error {
	var temp T
	switch any(temp).(type) {
	case int:
		_, err := fmt.Sscan(s, &temp)
		*v = temp
		return err
	case float64:
		_, err := fmt.Sscan(s, &temp)
		*v = temp
		return err
	case string:
		return nil
	default:
		return fmt.Errorf("unsupported type")
	}
}
