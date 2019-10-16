package domain

import (
	"oneday-infrastructure/internal/pkg/token"
)

//TODO move out
func GenerateToken(uniqueCode string, effectiveSeconds int) string {
	return token.Generate(uniqueCode, effectiveSeconds)
}
