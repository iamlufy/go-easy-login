package authenticate_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"oneday-infrastructure/api/authenticate"
	"oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/tools"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	router := gin.Default()
	authenticate.InitAuthenticateApi(router)

	w := httptest.NewRecorder()
	cmd := domain.LoginCmd{
		Username:         "zzf",
		EffectiveSeconds: 60 * 60,
		PassCode:         "zzf",
		LoginMode:        "PASSWORD",
		EncryptWay:       "MD5",
		UniqueCode:       "code",
	}

	req, _ := http.NewRequest(
		"POST",
		"login",
		strings.NewReader(tools.JsonString(cmd)))

	req.Header.Set("Content-Type", "Application/json")
	req.Header.Set("tenantCode", "tttt")

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.True(t, true, (tools.JsonStringToMap(w.Body.String()))["token"].(string) != "")
}

func TestResetPassword(t *testing.T) {
	router := gin.Default()
	authenticate.InitAuthenticateApi(router)

	w := httptest.NewRecorder()
	cmd := domain.ResetPasswordCmd{
		Username:    "zzf",
		EncryptWay:  "MD5",
		NewPassword: "zzf",
		OldPassword: "zzf",
	}

	req, _ := http.NewRequest(
		"POST",
		"resetPassword",
		strings.NewReader(tools.JsonString(cmd)))

	req.Header.Set("Content-Type", "Application/json")
	req.Header.Set("tenantCode", "tttt")

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, domain.ResetPasswordSuccess, (tools.JsonStringToMap(w.Body.String()))["result"].(string))

}
