package repository

import (
	"log"

	models "github.com/adewumi0550/pro_work/m/v2/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	LoginCrendtial(user models.User) models.User
	UpdateUser(user models.User) models.User
	// VerifyCredential(email string, password string) interface{}
	// IsDuplicateEmail(email string) (tx *gorm.DB)
	// FindByEmail(email string) models.User
	// ProfileUser(userID string) models.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) LoginCrendtial(user models.User) models.User {
	user.Password = hashAndSakt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user models.User) models.User {
	user.Password = hashAndSakt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func hashAndSakt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)

	if err != nil {
		log.Panicln(err)
		panic("Failed to hased a paaword")
	}

	return string(hash)
}
