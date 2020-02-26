package usecase

import (
	"CRUD-Project-user_management/user"
	"CRUD-Project-user_management/user/common"
	"CRUD-Project-user_management/user/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UserUsecaseImpl use for get a repo connection
type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

// CreateUsecase use for get connection from repository
func CreateUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

// CheckMail use for business logic when a new account want to regsiter he mail
// this function will check the mail is already register or nah
func (call *UserUsecaseImpl) CheckMail(user *models.User) bool {
	return call.userRepo.CheckMail(user)
}

// Login use for business logic when use trying to loggin in
func (call *UserUsecaseImpl) Login(user *models.User) (*models.User, error) {
	model, err := call.userRepo.Login(user)
	if err != nil {
		return nil, errors.New("Email not registered")
	}

	err = common.VerifyPassword(model.Password, user.Password)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New("Error invalid password")
	}

	return model, nil
}

// Create use for business logic when a new account is create
func (call *UserUsecaseImpl) Create(user *models.User) (*models.User, error) {
	return call.userRepo.Create(user)
}

// Update use for business logic when update a account
func (call *UserUsecaseImpl) Update(id int, user *models.User) (*models.User, error) {
	return call.userRepo.Update(id, user)
}

// FindAll use for business logic when you want to print all user account to the List
func (call *UserUsecaseImpl) FindAll() ([]*models.UserWrapper, error) {
	return call.userRepo.FindAll()
}

// FindByID use for business logic when you want to find account by id
func (call *UserUsecaseImpl) FindByID(id int) (*models.UserWrapper, error) {
	return call.userRepo.FindByID(id)
}

func (call *UserUsecaseImpl) Delete(id int) error {
	return call.userRepo.Delete(id)
}