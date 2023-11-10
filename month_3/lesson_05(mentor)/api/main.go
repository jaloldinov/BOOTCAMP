package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/pkg/helper"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	// r.Use(helper.AuthMiddleWare)

	r.POST("/branch", helper.AuthMiddleWare, h.CreateBranch)
	r.GET("/branch/:id", h.GetBranch)
	r.GET("/branch", h.GetAllBranch)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/tariff", h.CreateStaffTarif)
	r.GET("/tariff/:id", h.GetStaffTarif)
	r.GET("/tariff", h.GetAllStaffTarif)
	r.PUT("/tariff/:id", h.UpdateStaffTarif)
	r.DELETE("/tariff/:id", h.DeleteStaffTarif)

	// middlewares, vor validating phone number, password and  username
	r.POST("/staff", h.CreateStaff)
	r.POST("/login", h.GetByUsername)
	r.POST("/staff/change-password/:id", h.UpdateStaffPassword)

	r.GET("/staff/:id", h.GetStaff)
	r.GET("/staff", h.GetAllStaff)
	r.PUT("/staff/:id", h.UpdateStaff)
	r.DELETE("/staff/:id", h.DeleteStaff)

	r.POST("/sale", h.CreateSale)
	r.GET("/sale/:id", h.GetSale)
	r.GET("/sale", h.GetAllSale)
	r.PUT("/sale/:id", h.UpdateSale)
	r.DELETE("/sale/:id", h.DeleteSale)

	r.POST("/transaction", h.CreateTransaction)
	r.GET("/transaction/:id", h.GetTransaction)
	r.GET("/transaction", h.GetAllTransaction)
	r.PUT("/transaction/:id", h.UpdateTransaction)
	r.DELETE("/transaction/:id", h.DeleteTransaction)

	r.GET("/get_top_staff", h.GetTopStaff)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
