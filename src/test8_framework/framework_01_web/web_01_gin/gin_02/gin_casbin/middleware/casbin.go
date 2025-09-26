package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func CasbinMiddleware2(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _ := c.Get("username")
		obj := c.Request.URL.Path
		act := c.Request.Method

		ok, err := e.Enforce(username, obj, act)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesTemp, _ := c.Get("roles")
		roles, _ := rolesTemp.([]any)
		obj := c.Request.URL.Path
		act := c.Request.Method

		for _, role := range roles {
			if role == "" {
				continue
			}
			// 用内存中的策略判断是否允许
			// 检查某个人角色role是否有资源obj的act权限
			ok, err := e.Enforce(role, obj, act)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
			if ok {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
		c.Abort()
		return
	}
}
