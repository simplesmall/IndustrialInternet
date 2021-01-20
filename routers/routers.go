package routers

import (
	"IndustrialInternet/common/MiddleWare"
	MiddleJWT "IndustrialInternet/common/MiddleWare/jwt"
	_ "IndustrialInternet/common/MiddleWare/jwt"
	"IndustrialInternet/config"
	apiv1 "IndustrialInternet/controller/api/v1"
	_ "IndustrialInternet/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer() {
	server := gin.Default()
	// 初始化数据库连接
	config.InitConnect()

	// 配置swagger
	server.Use(MiddleWare.Cors())
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 路由分组
	api := server.Group("api")
	{
		// 响应体返回测试
		api.GET("/login", apiv1.LoginTest)
		api.GET("/normal", apiv1.NormalTest)
		api.GET("/msgOK", apiv1.NormalOKTest)
		v1 := api.Group("v1")
		{
			// 测试登录登出
			auth:=v1.Group("auth")
			{
				auth.GET("/login", apiv1.Login)
				auth.GET("/logout", apiv1.Login)
				auth.GET("/notfound", apiv1.NotFoundPage())
			}
		}

		// JWT模块测试
		v2 := api.Group("v2")
		{
			v2.GET("/login",MiddleJWT.JWTAuth(), apiv1.NotFoundPage())
			v2.POST("/auth")
		}

		// orm练手测试
	}
	server.Run(":8090")
}
