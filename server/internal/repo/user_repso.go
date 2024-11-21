package repo

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/db"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/log"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var logging  = *log.GetLogger()
type UserRepo struct {

}

func (u *UserRepo) Create(user *models.User)  error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logging.Error(err.Error())
		return err
	}
	user.Password = string(hashedPassword);
	err = db.DB.Create(&user).Error;
	if err != nil {
		logging.Error(err.Error())
		return err
	}
	return nil
}

func (u *UserRepo) Login(user *models.User) error  {
	loginUser := &models.User{}
	err := db.DB.Where("email = ?", user.Email).First(&loginUser).Error
	if err != nil {
		logging.Error(err.Error())
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(loginUser.Password), []byte(user.Password))
	if err != nil {
		logging.Error(err.Error())
		return err
	}
	return nil
}