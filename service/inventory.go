package service

import (
	"github.com/i-jonathan/pharmacy-api/repository"
)

type inventoryService struct {
	repo repository.PharmacyRepository
}

func NewInventoryService(r repository.PharmacyRepository) *inventoryService {
	return &inventoryService{r}
}
