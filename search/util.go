package search

import (
	"fmt"
)

func stringifyID(id int) string {
	return fmt.Sprintf("%d", id)
}

func toStringSlice(arr []interface{}) []string {
	result := make([]string, len(arr))
	for _, item := range arr {
		result = append(result, item.(string))
	}
	return result
}
