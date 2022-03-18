package repository

import (
	"log"
	"srp-golang/app/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Index() []models.User
	Create(model models.User) models.User
	Show(id uint64) models.User
	Update(model models.User) models.User
	Delete(user models.User) models.User
	// Pagination(*models.Pagination) (RepositoryResult, int)

	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Index() []models.User {
	var user []models.User
	db.connection.Find(&user)
	return user
}

func (db *userConnection) Create(model models.User) models.User {
	model.Password = hashAndSalt([]byte(model.Password))
	db.connection.Save(&model)
	return model
}

func (db *userConnection) Show(id uint64) models.User {
	var user models.User
	db.connection.Find(&user, id)
	return user
}

func (db *userConnection) Update(model models.User) models.User {
	db.connection.Updates(&model)
	db.connection.Find(&model)
	return model
}

func (db *userConnection) Delete(user models.User) models.User {
	db.connection.Delete(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) models.User {
	var user models.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
