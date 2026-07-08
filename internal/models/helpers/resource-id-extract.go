package helpers

import (
	"strconv"
	"strings"
)

func GetResourceId(p_url string) int {
	url := strings.TrimSuffix(p_url, "/")
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return -1
	}
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(idStr)
		return -128
	}
	return id
}
