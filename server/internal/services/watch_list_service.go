package services

import (
	logg "github.com/SwanHtetAungPhyo/stockAggregation/internal/log"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/models"
	"github.com/SwanHtetAungPhyo/stockAggregation/internal/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"strconv"
)

var (
	logger = logg.GetLogger()
)

type (
	WatchListService interface {
	}

	WatchListServiceImpl struct {
		WatchListRepo *repo.WatchRepo
	}
)

func NewWatchListService(watchListRepo *repo.WatchRepo) *WatchListServiceImpl {
	return &WatchListServiceImpl{
		WatchListRepo: watchListRepo,
	}
}

func (w *WatchListServiceImpl) AddWatchList(ctx *fiber.Ctx) error {
	var watchList []models.StockWatchList
	userId, _ := strconv.Atoi(ctx.Params("id"))
	if err := ctx.BodyParser(&watchList); err != nil {
		return sendError(ctx, fiber.StatusBadRequest, err.Error())
	}
	for _, watch := range watchList {
		if err := watch.Validate(); err != nil {
			log.Infof("Current WatchList is not valid %v", watch.Stock)
			log.Error(err)
			return sendError(ctx, fiber.StatusBadRequest, err.Error())
		}
	}

	if len(watchList) == 0 {
		return sendError(ctx, fiber.StatusBadRequest, "WatchList is empty")
	}

	for i := range watchList {
		watchList[i].UserID = uint(userId)
	}
	err := w.WatchListRepo.AddWatch(&watchList, userId)
	if err != nil {
		return sendError(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully added to watchlist",
		"list":    watchList,
	})
}
