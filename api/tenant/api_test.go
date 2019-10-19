package tenant_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"oneday-infrastructure/api/tenant"
	"oneday-infrastructure/internal/pkg/tenant/domain"
	"oneday-infrastructure/tools"
	"strconv"
	"strings"
	"testing"
	"time"
)

var genUniqueCode = func() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(1000000))
}

func TestAddTenant(t *testing.T) {
	router := gin.Default()
	tenant.InitTenantApi(router)

	w := httptest.NewRecorder()

	cmd := domain.AddTenantCmd{TenantName: "1`tttt", TenantCode: "tttt"}
	req, _ := http.NewRequest(
		"POST",
		"tenant",
		strings.NewReader(tools.JsonString(cmd)))
	req.Header.Set("Content-Type", "Application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestAddTenantUser(t *testing.T) {
	router := gin.Default()
	tenant.InitTenantApi(router)

	w := httptest.NewRecorder()

	cmd := domain.AddUserCmd{
		TenantCode: "tttt",
		Username:   "zzf",
		Mobile:     "12345678901",
		Password:   "zzf",
		EncryptWay: "MD5",
	}

	req, _ := http.NewRequest(
		"POST",
		"tenant/user",
		strings.NewReader(tools.JsonString(cmd)))
	req.Header.Set("Content-Type", "Application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
