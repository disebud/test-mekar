package repositories

import "github.com/disebud/api-users-test/main/master/models"

// UserRepository kumpulan fungsi
type UserRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetJob() ([]*models.Job, error)
	GetEducation() ([]*models.Education, error)
	CreateUser(user models.User) error
	GetUserById(userId string) (*models.User, error)
	DeleteUserById(userId string) (*models.User, error)
	UpdateUser(userId string, User models.User) error
}
