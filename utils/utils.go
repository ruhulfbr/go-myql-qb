package utils

import "log"

var AllowedOperators = map[string]bool{
	"=":  true,
	"!=": true,
	"<":  true,
	"<=": true,
	">":  true,
	">=": true,
}

// IsValidOperator Exported function
func IsValidOperator(operator string) bool {

	if !AllowedOperators[operator] {
		log.Fatalf("Invalid operator: %s", operator)
	}

	return true
}
