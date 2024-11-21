package repo

import (
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/db"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/models"
)

type WatchRepo struct {
}

func NewWatchRepo() *WatchRepo {
	return &WatchRepo{}
}

func (w *WatchRepo) GetWatchList() (*models.StockWatchList, error) {
	return nil, nil
}

func (w *WatchRepo) AddWatch(watchList *[]models.StockWatchList, id int) error {
	useRepo := new(UserRepo)
	ok, err := useRepo.FindById(id)
	if err != nil {
		return err
	}

	if !ok {
		logging.Error("User not found")
		return err
	} else {
		if err := db.DB.Create(&watchList).Error; err != nil {
			logging.Error("Error creating watch list")
			return err
		}
	}
	logging.Info("Watch list created")
	return nil
}

func (w *WatchRepo) RemoveWatch() (bool, error) {
	return false, nil
}
