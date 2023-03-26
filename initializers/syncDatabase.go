package initializers

import "github.com/Twiddit/Twiddit_auth_ms/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
