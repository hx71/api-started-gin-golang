package usecase

import (
	"log"

	"github.com/hx71/api-started-gin-golang/app/auth"
	"github.com/hx71/api-started-gin-golang/app/dto"
	"github.com/hx71/api-started-gin-golang/app/user"
	"github.com/hx71/api-started-gin-golang/models"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase struct {
	repo user.Repository
}

func NewAuthUsecase(repo user.Repository) auth.Usecase {
	return &authUsecase{
		repo: repo,
	}
}

func (r *authUsecase) VerifyCredential(email string, password string) interface{} {
	res := r.repo.VerifyCredential(email, password)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (r *authUsecase) CreateUser(user dto.RegisterValidation) error {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	return r.repo.Create(userToCreate)
}

func (r *authUsecase) FindByEmail(email string) bool {
	return r.repo.FindByEmail(email)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		return false
	}
	return true
}
