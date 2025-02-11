package datastore

import (
	"context"

	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/repositories/interfaces"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

type DatastorePlayerRepo struct {
	kind   string
	client *datastore.Client
}

func NewDatastorePlayerRepository(client *datastore.Client) interfaces.PlayerRepository {
	kind := "Player"
	return &DatastorePlayerRepo{
		client: client,
		kind:   kind,
	}
}

func (p *DatastorePlayerRepo) Create(player models.Player) (models.Player, error) {
	ctx := context.Background()

	if player.ID == "" {
		player.ID = uuid.New().String()
	}

	// Create new key
	key := datastore.NameKey(p.kind, player.ID, nil)

	// Save entity
	newKey, err := p.client.Put(ctx, key, &player)
	if err != nil {
		return player, err
	}

	// Update player ID with the generated key
	player.ID = newKey.Name
	return player, nil
}

func (p *DatastorePlayerRepo) GetByID(id string) (models.Player, error) {
	ctx := context.Background()

	key := datastore.NameKey(p.kind, id, nil)
	player := &models.Player{}

	if err := p.client.Get(ctx, key, player); err != nil {
		return models.Player{}, err
	}

	player.ID = id
	return *player, nil
}

func (p *DatastorePlayerRepo) GetAll() ([]models.Player, error) {
	ctx := context.Background()

	var players []models.Player
	q := datastore.NewQuery(p.kind)

	_, err := p.client.GetAll(ctx, q, &players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func (p *DatastorePlayerRepo) Update(player models.Player) (models.Player, error) {
	ctx := context.Background()

	key := datastore.NameKey(p.kind, player.ID, nil)
	_, err := p.client.Put(ctx, key, &player)
	return player, err
}

func (p *DatastorePlayerRepo) Delete(id string) error {
	ctx := context.Background()

	key := datastore.NameKey(p.kind, id, nil)
	return p.client.Delete(ctx, key)
}

// func (p *DatastorePlayerRepo) List() ([]*models.Player, error) {
// 	ctx := context.Background()

// 	var players []*models.Player
// 	q := datastore.NewQuery(p.kind)

// 	keys, err := p.client.GetAll(ctx, q, &players)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Set IDs from keys
// 	for i, key := range keys {
// 		players[i].ID = key.Name
// 	}

// 	return players, nil
// }

// Query example with filters
// func (p *DatastorePlayerRepo) FindByPosition(position string) ([]*models.Player, error) {
// 	ctx := context.Background()

// 	var players []*models.Player
// 	q := datastore.NewQuery(p.kind).
// 		Filter("position =", position)

// 	keys, err := p.client.GetAll(ctx, q, &players)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for i, key := range keys {
// 		players[i].ID = key.Name
// 	}

// 	return players, nil
// }
