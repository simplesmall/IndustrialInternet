package v1

import (
	middlewareJWT "IndustrialInternet/common/middleware/jwt"
	"IndustrialInternet/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 登录测试接口
// @Description  登录测试接口描述信息
// @Success 200 {string} string    "ok"
// @Router /api/v1/login [get]
// @Tags 登录
// BaseURL /
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		//res:=response.ResponseStructure{
		//	Code: 0,
		//	Msg:  "Hello login!!",
		//	Data: nil,
		//}
		//c.JSON(http.StatusOK,gin.H{
		//	"data":res,
		//})
		var loginInput middlewareJWT.LoginInput
		err := c.ShouldBindJSON(&loginInput)
		if err != nil {
			//Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		}

		//参数校验
		//TODO: 参数检验

		token, err := middlewareJWT.Login(loginInput.Username, loginInput.Password)
		if err != nil {
			if err.Error() == "record not found" {
				//Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("该用户不存在"))
				return
			} else if err.Error() == "invalid password" {
				//Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("密码校验不通过"))
				return
			} else {
				//Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("登录错误"))
				return
			}
		}
		//Middlewares.ResponseSuccess(c, token, "登录成功")
		fmt.Println(token)
		return
	}
}

// @Summary notfound测试接口
// @Description  notfound测试接口描述信息
// @Success 200 {string} string    "ok"
// @Router /api/v1/notfound [get]
// @Tags NotFound
// BaseURL /
func NotFoundPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"data":response.ResponseStructure.NotFound(response.ResponseStructure{}),
		})
	}
}
