package internal

import (
	"strings"
)

func ValidateAPIKey(header, apiKey string) bool {

	parts := strings.SplitN(header, " ", 2)

	if len(parts) != 2 {
		return false
	}
	return parts[1] == apiKey

}
