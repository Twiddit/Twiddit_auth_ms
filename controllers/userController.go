package controllers

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
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	// Get the data from the req body
	var body struct {
		Email        string
		Password     string
		Birthday     string
		Phone        string
		ProfilePhoto string
		Description  string
		Username     string
	}

	if ctx.Bind(&body) != nil {
		utils.APIResponse(ctx, "Failed to read requests body", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		utils.APIResponse(ctx, "Failed to hash password", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	// Birthday parsing
	date, err := time.Parse("02-01-2006", body.Birthday)

	// Check for errors
	if err != nil {
		fmt.Println(err)
		return
	}
	// Create the user
	user := models.User{Email: body.Email,
		Password:     string(hash),
		Birthday:     date,
		Phone:        body.Phone,
		ProfilePhoto: body.ProfilePhoto,
		Description:  body.Description,
		Username:     body.Username,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		utils.APIResponse(ctx, "Failed to create user", http.StatusBadRequest, http.MethodPost, nil)
	}

	// Respond

	utils.APIResponse(ctx, "User created successfully", http.StatusOK, http.MethodPost, nil)
}

func Login(ctx *gin.Context) {
	// Get the data from the req body
	var body struct {
		Email    string
		Password string
	}

	if ctx.Bind(&body) != nil {
		utils.APIResponse(ctx, "Failed to read requests body", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		utils.APIResponse(ctx, "User is not registered. Please Sign Up first", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		utils.APIResponse(ctx, "Invalid user or password", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	// Sign it
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		utils.APIResponse(ctx, "Failed to create token", http.StatusBadRequest, http.MethodPost, nil)
		return
	}

	// Send it back
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	utils.APIResponse(ctx, "Successful Login", http.StatusOK, http.MethodPost, map[string]string{"accessToken": tokenString})
}

func Validate(ctx *gin.Context) {
	// Validate the token using the middleware
	// validAuth, _ := ctx.Get("auth")
	utils.APIResponse(ctx, "Successful Validation", http.StatusOK, http.MethodPost, map[string]string{"accessToken": "True"})
}
