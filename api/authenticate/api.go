package authenticate

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oneday-infrastructure/internal/pkg/authenticate/domain"
	"oneday-infrastructure/internal/pkg/tenant/base"
	"oneday-infrastructure/pkg/authenticate"
	"oneday-infrastructure/tools"
)

func InitAuthenticateApi(r *gin.Engine) {

	r.POST("/login", Login)
	r.POST("/resetPassword", ResetPassword)
}

func checkTenantCode(tenantCode string) bool {
	_, exist := base.NewTenantRepo(tools.OpenDB).
		FindOne(tools.NewMap("tenant_code", tenantCode))
	return exist
}

func Login(c *gin.Context) {
	var cmd domain.LoginCmd
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tenantCode := c.GetHeader("tenantCode")
	if !checkTenantCode(tenantCode) {
		c.JSON(http.StatusOK, gin.H{
			"code":    "1000",
			"message": "Invalid tenant code",
		})
	}
	token, err := authenticate.Login(&cmd, tenantCode)
	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": err,
	})
}

func ResetPassword(c *gin.Context) {
	var cmd domain.ResetPasswordCmd
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tenantCode := c.GetHeader("tenantCode")
	if !checkTenantCode(tenantCode) {
		c.JSON(http.StatusOK, gin.H{
			"code":    "1000",
			"message": "Invalid tenant code",
		})
	}
	result := authenticate.ReSetPassword(&cmd, tenantCode)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}
