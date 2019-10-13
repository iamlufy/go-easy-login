package facade

import (
	"oneday-infrastructure/internal/pkg/token"
)

func GenerateToken(uniqueCode string, effectiveSeconds int) string {
	return token.Generate(uniqueCode, effectiveSeconds)
}
