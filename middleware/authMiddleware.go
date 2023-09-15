package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequiredAuth(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Unauthorized",
		})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": err.Error(),
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("jwtClaims", claims)
	}

	c.Next()

}

func RequiredAuthLoan(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Unauthorized",
		})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": err.Error(),
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.Set("jwtClaims", claims)
	}

	// Mengekstrak JWT claims dari konteks
	jwtClaims, _ := c.Get("jwtClaims")
	claims, ok := jwtClaims.(jwt.MapClaims)
	fmt.Println(jwtClaims)
	// Memeriksa apakah pengguna memiliki peran "admin"
	rolesID, _ := claims["roles"].(float64)
	fmt.Println(claims)
	fmt.Println(rolesID)
	if !ok || rolesID != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errors": "Access denied",
		})
		c.Abort()
		return
	}

	c.Next()

}
