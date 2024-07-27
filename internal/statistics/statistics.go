package statistics

import (
	"fmt"

	"github.com/SurkovIlya/message-handler-app/internal/model"
)

type Storage interface {
	GetStat() (model.Statistic, error)
}

type StatisticHandler struct {
	Storage Storage
}

func New(query Storage) *StatisticHandler {
	return &StatisticHandler{
		Storage: query,
	}
}

func (sh *StatisticHandler) GetStat() (model.Statistic, error) {
	stat, err := sh.Storage.GetStat()
	if err != nil {
		return model.Statistic{}, fmt.Errorf("error GetStatistics query: %s", err)
	}

	return stat, nil
}
