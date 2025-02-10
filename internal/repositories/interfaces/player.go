package interfaces

import "github.com/pabloantipan/hobe-locations-api/internal/models"

type PlayerRepository interface {
	Create(player models.Player) (models.Player, error)
	GetByID(id string) (models.Player, error)
	GetAll() ([]models.Player, error)
	Update(player models.Player) (models.Player, error)
	Delete(id string) error
}
