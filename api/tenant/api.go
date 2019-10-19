package tenant

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oneday-infrastructure/internal/pkg/tenant/domain"
	"oneday-infrastructure/pkg/tenant"
)

func InitTenantApi(r *gin.Engine) {

	r.POST("/tenant", AddTenant)
	r.POST("/tenant/user", AddTenantUser)

}

func AddTenant(c *gin.Context) {
	var cmd domain.AddTenantCmd
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tenantCO, result := tenant.Add(&cmd)
	c.JSON(http.StatusOK, gin.H{
		"tenant": tenantCO,
		"result": result,
	})
}

func AddTenantUser(c *gin.Context) {
	var cmd domain.AddUserCmd
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tenant.AddTenantUser(&cmd)
	c.JSON(http.StatusOK, gin.H{})

}
