package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
	"github.com/pabloantipan/hobe-locations-api/internal/repositories/interfaces"
)

type DatastoreLocationRepo struct {
	kind   string
	client *datastore.Client
}

func NewDatastoreLocationRepository(client *datastore.Client) interfaces.LocationRepository {
	kind := "Location"
	return &DatastoreLocationRepo{
		client: client,
		kind:   kind,
	}
}

func (r *DatastoreLocationRepo) Add(location models.Location) (models.Location, error) {
	ctx := context.Background()

	if location.ID == "" {
		location.ID = uuid.New().String()
	}
	key := datastore.NameKey(r.kind, location.ID, nil)

	newKey, err := r.client.Put(ctx, key, &location)
	if err != nil {
		return location, err
	}
	location.ID = newKey.Name
	return location, nil
}

func (r *DatastoreLocationRepo) GetByID(id string) (models.Location, error) {
	ctx := context.Background()

	key := datastore.NameKey(r.kind, id, nil)
	location := &models.Location{}

	if err := r.client.Get(ctx, key, location); err != nil {
		return models.Location{}, err
	}

	location.ID = id
	return *location, nil
}

func (r *DatastoreLocationRepo) GetAll() ([]models.Location, error) {
	ctx := context.Background()

	var locations []models.Location
	query := datastore.NewQuery(r.kind)

	keys, err := r.client.GetAll(ctx, query, &locations)
	if err != nil {
		return nil, err
	}

	for i, key := range keys {
		locations[i].ID = key.Name
	}

	return locations, nil
}

func (r *DatastoreLocationRepo) Update(location models.Location) (models.Location, error) {
	ctx := context.Background()

	key := datastore.NameKey(r.kind, location.ID, nil)
	_, err := r.client.Put(ctx, key, &location)
	return location, err
}

func (r *DatastoreLocationRepo) Delete(id string) error {
	ctx := context.Background()

	key := datastore.NameKey(r.kind, id, nil)
	return r.client.Delete(ctx, key)
}
