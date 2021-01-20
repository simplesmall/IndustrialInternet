package v1

import (
	middlewareJWT "IndustrialInternet/common/MiddleWare/jwt"
	"IndustrialInternet/model/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 登录测试接口
// @Description  登录测试接口描述信息
// @Accept application/json
// @Produce application/json
// @Success 200 {object} Response.ResponseStructure   {"code":200,"msg":"","data":null}
// @Router /api/v1/auth/login [get]
// @Param token header string false "用户令牌"
// @Param username path integer true "角色ID"
// @Tags 登录
// BaseURL /
func Login(c *gin.Context) {
		var loginInput middlewareJWT.LoginInput
		err := c.ShouldBindJSON(&loginInput)
		if err != nil {
			Response.ResponseStructure{}.FailResWithMsg("数据获取失败","")
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
		c.JSON(200, Response.ResponseStructure{}.OKResultWithMsg("登录成功",token))
}


// @Summary 普通GET测试
// @Description  普通GET测试接口描述信息
// @Success 200 {object} Response.ResponseStructure {"code":200,"data":null,"msg":""}
// @Router /api/login [get]
// @Tags Test
// BaseURL /
func LoginTest(c *gin.Context) {
	c.JSON(http.StatusOK, &Response.ResponseStructure{
		230,
		"Hello login!!",
		nil,
	})
}

// @Summary notfound测试接口
// @Description  notfound测试接口描述信息
// @Success 200 {string} string    "ok"
// @Router /api/v1/notfound [get]
// @Tags NotFound
// BaseURL /
func NotFoundPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, Response.ResponseStructure.NotFound(Response.ResponseStructure{}))
	}
}
// @Summary NormalTest测试接口
// @Description  NormalTest测试接口描述信息
// @Success 200 {object} Response.ResponseStructure {"code":200,"data":null,"msg":""}
// @Router /api/normal [get]
// @Tags Test
// BaseURL /
func NormalTest(c *gin.Context) {
	data:="hello"
	//structure := response.ResponseStructure{}
	c.JSON(http.StatusOK,  Response.ResponseStructure{}.OKResult(data))
}
// @Summary NormalOKTest测试接口
// @Description  NormalOKTest测试接口描述信息
// @Success 200 {object} Response.ResponseStructure {"code":200,"data":null,"msg":""}
// @Router /api/msgOK [get]
// @Tags Test
// BaseURL /
func NormalOKTest(c *gin.Context) {
	data:="JIUSHI OKKKK"
	msg:="简单的消息"
	//structure := response.ResponseStructure{}
	c.JSON(http.StatusOK, Response.ResponseStructure{}.OKResultWithMsg(msg,data))
}
