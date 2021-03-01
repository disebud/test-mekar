package usecases

import (
	"github.com/disebud/api-users-test/main/master/models"
	"github.com/disebud/api-users-test/main/master/repositories"
)

type UserUsecaseImpl struct {
	UserRepo repositories.UserRepository
}

func InitUserUsecase(UserRepo repositories.UserRepository) UserUseCase {
	return &UserUsecaseImpl{UserRepo}
}

func (s UserUsecaseImpl) GetAllUsers() ([]*models.User, error) {
	User, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (s UserUsecaseImpl) GetJob() ([]*models.Job, error) {
	Job, err := s.UserRepo.GetJob()
	if err != nil {
		return nil, err
	}
	return Job, nil
}

func (s UserUsecaseImpl) GetEducation() ([]*models.Education, error) {
	Education, err := s.UserRepo.GetEducation()
	if err != nil {
		return nil, err
	}
	return Education, nil
}

func (s UserUsecaseImpl) GetUserById(userId string) (*models.User, error) {
	User, err := s.UserRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (s UserUsecaseImpl) DeleteUserById(userId string) (*models.User, error) {
	User, err := s.UserRepo.DeleteUserById(userId)
	if err != nil {
		return nil, err
	}
	return User, nil
}

func (s UserUsecaseImpl) CreateUser(User models.User) error {
	// err := utils.ValidateInputLenCharacter(3, 6, User.IdUser)
	// if err != nil {
	// 	return err
	// }
	// err1 := utils.ValidateInputNotNil(User.NameUser, User.Location, User.Price)
	// if err1 != nil {
	// 	return err
	// }
	err := s.UserRepo.CreateUser(User)
	if err != nil {
		return err
	}

	return nil
}

func (s UserUsecaseImpl) UpdateUser(userId string, User models.User) error {
	err := s.UserRepo.UpdateUser(userId, User)
	if err != nil {
		return err
	}
	return nil
}

// func (s UserUsecaseImpl) GetCountUserType(StatusKey string) ([]*models.UserType, error) {
// 	User, err := s.UserRepo.GetCountUserType(StatusKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return User, nil
// }
