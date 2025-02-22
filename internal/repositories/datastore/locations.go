package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
	"github.com/pabloantipan/hobe-locations-api/internal/models"
)

type DatastoreLocationRepo struct {
	ctx    context.Context
	kind   string
	client *datastore.Client
}

func NewDatastoreLocationRepository(ctx context.Context, client *datastore.Client) LocationRepository {
	kind := "Location"
	return &DatastoreLocationRepo{
		ctx:    ctx,
		client: client,
		kind:   kind,
	}
}

type LocationRepository interface {
	Add(location models.Location) (models.Location, error)
	GetByID(id string) (models.Location, error)
	GetAll() ([]models.Location, error)
	GetThemByEmail(email string) (*[]models.Location, error)
	Update(location models.Location) (models.Location, error)
	Delete(id string) error
}

func (r *DatastoreLocationRepo) Add(location models.Location) (models.Location, error) {

	if location.ID == "" {
		location.ID = uuid.New().String()
	}

	key := datastore.NameKey(r.kind, location.ID, nil)

	newKey, err := r.client.Put(r.ctx, key, &location)
	if err != nil {
		return location, err
	}
	location.ID = newKey.Name
	return location, nil
}

func (r *DatastoreLocationRepo) GetByID(id string) (models.Location, error) {
	key := datastore.NameKey(r.kind, id, nil)
	location := &models.Location{}

	if err := r.client.Get(r.ctx, key, location); err != nil {
		return models.Location{}, err
	}

	location.ID = id
	return *location, nil
}

func (r *DatastoreLocationRepo) GetAll() ([]models.Location, error) {

	var locations []models.Location
	query := datastore.NewQuery(r.kind)

	keys, err := r.client.GetAll(r.ctx, query, &locations)
	if err != nil {
		return nil, err
	}

	for i, key := range keys {
		locations[i].ID = key.Name
	}

	return locations, nil
}

func (r *DatastoreLocationRepo) GetThemByEmail(email string) (*[]models.Location, error) {
	query := datastore.NewQuery(r.kind).FilterField("UserEmail", "=", email)
	var locations []models.Location
	_, err := r.client.GetAll(r.ctx, query, &locations)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch locations: %v", err)
	}
	if locations == nil {
		locations = []models.Location{}
	}
	return &locations, nil
}

func (r *DatastoreLocationRepo) Update(location models.Location) (models.Location, error) {

	key := datastore.NameKey(r.kind, location.ID, nil)
	_, err := r.client.Put(r.ctx, key, &location)
	return location, err
}

func (r *DatastoreLocationRepo) Delete(id string) error {
	key := datastore.NameKey(r.kind, id, nil)
	return r.client.Delete(r.ctx, key)
}
