package facade

import (
	"oneday-infrastructure/internal/pkg/token"
)

func IsUserExisting(tenantCode, username string) bool {
	return true
}

func GenerateToken(uniqueCode string, effectiveSeconds int) string {
	return token.Generate(uniqueCode, effectiveSeconds)
}
