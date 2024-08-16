package usecases

import (
	repositories "task-manager/Repositories"
	"task-manager/domain"
	"task-manager/infrastructure"
)

type UserUseCase interface {
	RegisterUser(username, password string) (*domain.User, error)
	LoginUser(username, password string) (string, error)
	PromoteUser(userID string) error
}

type userUseCase struct {
	userRepo repositories.UserRepository
	passwordSvc   infrastructure.PasswordService
	jwtSvc        infrastructure.JWTService
}

func NewUserUseCase (userRepo repositories.UserRepository, passwordSvc infrastructure.PasswordService, jwtSvc infrastructure.JWTService) UserUseCase {
	return &userUseCase {
		userRepo: userRepo,
		passwordSvc: passwordSvc,
		jwtSvc: jwtSvc,
	}
}


func (uc *userUseCase) RegisterUser(username, password string) (*domain.User, error) {
	hashedPassword, err := uc.passwordSvc.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username: username,
		Password: hashedPassword,
		Role:     "user",
	}

	// Check if first user and promote to admin
	if err := uc.userRepo.RegisterUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *userUseCase) LoginUser(username, password string) (string, error) {
	user, err := uc.userRepo.LoginUser(username)
	if err != nil {
		return "", err
	}

	if err := uc.passwordSvc.CheckPasswordHash(password, user.Password); err != nil {
		return "", err
	}

	token, err := uc.jwtSvc.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *userUseCase) PromoteUser(userID string) error {
	return uc.userRepo.PromoteUser(userID)
}


