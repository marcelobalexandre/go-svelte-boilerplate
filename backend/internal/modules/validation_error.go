package modules

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Message string              `json:"message"`
	Details map[string][]string `json:"details"`
}

func (e ValidationError) Error() string {
	var builder strings.Builder
	for field, messages := range e.Details {
		builder.WriteString(fmt.Sprintf("%s: %v\n", field, messages))
	}

	return fmt.Sprintf("Validation error: %s", builder.String())
}
