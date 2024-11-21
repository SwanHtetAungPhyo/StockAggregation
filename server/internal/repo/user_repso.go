package repo

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/db"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/log"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/models"
)

var logging  = *log.GetLogger()
type UserRepo struct {

}

func (u *UserRepo) Create(user *models.User)  error {
	err := db.DB.Create(&user).Error;
	if err != nil {
		logging.Error(err.Error())
		return err
	}
	return nil
}