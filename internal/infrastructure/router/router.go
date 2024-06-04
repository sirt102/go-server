package router

import (
	"go-server/internal/api/handler"
	"go-server/internal/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Interview Be Earning API
//	@version		1.0
//	@description	A employee attandance management service API in Go using Gin framework..
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.email	trisnm102@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func Initialize(h handler.AppHandler) {
	router := gin.Default()
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	router.Use(configSwagger).GET("/swagger/*any", swaggerHandler)

	adminGroup := router.Group("/admin")
	{
		adminGroup.POST("/employees", h.EmployeemManagementHandler.CreateNewEmployee)
	}

	appVersion1Group := router.Group("/v1")
	{
		employeePublicGroup := appVersion1Group.Group("employees/:employee_id")
		{
			actionGroup := employeePublicGroup.Group("/actions")
			{
				actionGroup.POST("", h.UserDoActionHandler.UserCreateAction)
			}

			attendanceGroup := employeePublicGroup.Group("/attendances")
			{
				attendanceGroup.GET("", h.UserGetInfoHandler.UserGetAttendanceList)
			}
		}
	}

	router.Run(":8080")
}

func configSwagger(c *gin.Context) {
	docs.SwaggerInfo.Host = c.Request.Host
}
