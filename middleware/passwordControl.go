package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Twiddit/Twiddit_auth_ms/initializers"
	"github.com/Twiddit/Twiddit_auth_ms/models"
	"github.com/Twiddit/Twiddit_auth_ms/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func PasswordControl(ctx *gin.Context) {
	// Get the cookie off the request
	var body struct {
		Previous      string
		NewPass       string
		Authorization string
	}

	if ctx.ShouldBindJSON(&body) != nil {
		utils.APIResponse(ctx, "Failed to read requests body", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	tokenString := body.Authorization

	// Decode it and validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusForbidden)
		}

		// Find the user
		var user models.User
		initializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusForbidden)
		}
		// Attach it to the request
		ctx.Set("auth", "True")
		ctx.Set("Previous", body.Previous)
		ctx.Set("NewPass", body.NewPass)
		ctx.Set("user", user.ID)
		// Continue
		ctx.Next()
	} else {
		fmt.Println(err)
	}
}
