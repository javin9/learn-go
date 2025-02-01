package util

import "fmt"

func GenerateTokenKey(user_id int) string {
	return fmt.Sprintf("token_%d", user_id)
}

func GenerateTokenValueKey(token string) string {
	return fmt.Sprintf("auth:%s", token)
}
