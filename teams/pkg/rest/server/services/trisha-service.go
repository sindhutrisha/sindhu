package services

import (
	"github.com/sindhutrisha/sindhu/teams/pkg/rest/server/daos"
	"github.com/sindhutrisha/sindhu/teams/pkg/rest/server/models"
)

type TrishaService struct {
	trishaDao *daos.TrishaDao
}

func NewTrishaService() (*TrishaService, error) {
	trishaDao, err := daos.NewTrishaDao()
	if err != nil {
		return nil, err
	}
	return &TrishaService{
		trishaDao: trishaDao,
	}, nil
}

func (trishaService *TrishaService) CreateTrisha(trisha *models.Trisha) (*models.Trisha, error) {
	return trishaService.trishaDao.CreateTrisha(trisha)
}

func (trishaService *TrishaService) UpdateTrisha(id int64, trisha *models.Trisha) (*models.Trisha, error) {
	return trishaService.trishaDao.UpdateTrisha(id, trisha)
}

func (trishaService *TrishaService) DeleteTrisha(id int64) error {
	return trishaService.trishaDao.DeleteTrisha(id)
}

func (trishaService *TrishaService) ListTrishas() ([]*models.Trisha, error) {
	return trishaService.trishaDao.ListTrishas()
}

func (trishaService *TrishaService) GetTrisha(id int64) (*models.Trisha, error) {
	return trishaService.trishaDao.GetTrisha(id)
}
