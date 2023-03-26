package controllers

import (
	"net/http"

	"github.com/Twiddit/Twiddit_auth_ms/initializers"
	"github.com/Twiddit/Twiddit_auth_ms/models"
	"github.com/Twiddit/Twiddit_auth_ms/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(ctx *gin.Context) {
	// Get the requests data with the new password
	Previous, _ := ctx.Get("Previous")
	NewPass, _ := ctx.Get("NewPass")

	// Get the data from the user that is making the request
	userID, _ := ctx.Get("user")

	// Find the user
	var user models.User
	initializers.DB.First(&user, userID)

	// Check if the password sent as the original matches the one on the record
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Previous.(string)))

	// If it doesnt we return an error to the user
	if err != nil {
		utils.APIResponse(ctx, "Invalid credentials", http.StatusOK, http.MethodPost, nil)
		return
	}

	// If it does we now hash the new password
	hash, err := bcrypt.GenerateFromPassword([]byte(NewPass.(string)), 10)

	if err != nil {
		utils.APIResponse(ctx, "Failed to hash password", http.StatusOK, http.MethodPost, nil)
	}

	// And we save the hash into the database
	user.Password = string(hash)
	initializers.DB.Save(&user)

	// Produce a response
	utils.APIResponse(ctx, "Successful Password Change", http.StatusOK, http.MethodPost, nil)
}
