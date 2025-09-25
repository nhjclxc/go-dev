package handler

import (
	"fmt"
	"gin_casbin/config"
	"gin_casbin/middleware"
	"gin_casbin/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var json struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := config.DB.Preload("Roles").Where("username = ?", json.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	roles := make([]string, 0, len(user.Roles))
	for _, role := range user.Roles {
		roles = append(roles, role.Name)
	}

	m := map[string]any{
		"username": user.Username,
		"roles":    roles,
	}
	token, _ := middleware.GenerateToken(m)
	c.JSON(200, gin.H{"token": token})
}

func encryption(password string) (string, error) {
	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("加密失败！！", err)
		return "", err
	}

	//fmt.Println("加密成功：", string(bcryptedPassword))
	return string(bcryptedPassword), nil
}
