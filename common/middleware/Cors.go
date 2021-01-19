package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var AuthToken string

// 跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-encoding", "UTF-8")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		token := c.Request.Header.Get("token")
		AuthToken = token
		c.Next()
	}
}

/* 文档注解参考
@Summary 接口概要说明
@Description 接口详细描述信息
@Tags 用户信息   //swagger API分类标签, 同一个tag为一组
@accept json  //浏览器可处理数据类型，浏览器默认发 Accept: *
@Produce  json  //设置返回数据的类型和编码
@Param id path int true "ID"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
@Param name query string false "name"
@Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
@Failure 400 {object} Res {"code":200,"data":null,"msg":""}
@Router /test/{id} [get]    //路由信息，一定要写上
*/
