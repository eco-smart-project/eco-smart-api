package repository

import (
	"database/sql"
	"eco-smart-api/model"
	"errors"
)

type CollectionPointRepository struct {
	db *sql.DB
}

var ErrCollectionPointNotFound = errors.New("collection point not found")

func NewCollectionPointRepository(db *sql.DB) *CollectionPointRepository {
	return &CollectionPointRepository{db}
}

func (ur *CollectionPointRepository) CreateCollectionPoint(point *model.CollectionPoint) error {
	_, err := ur.db.Exec(`
		insert into ecosmart.collection_points(id)
						values (nextval('ecosmart.seq_collection_points'))
	`)

	return err
}

func (ur *CollectionPointRepository) GetCollectionPointByID(id uint64) (*model.CollectionPoint, error) {
	row := ur.db.QueryRow(`
		select cp.id
		      ,cp.position_latitude
			  ,cp.position_longitude
			  ,cp.title
			  ,cp.description
			  ,coalesce(cp.icon, cpc.icon) as icon
			  ,cpc.id as category_id
			  ,cpc.name as category_name
			  ,cpc.description as category_description
			  ,cp.created_at
			  ,cp.updated_at
		 from ecosmart.collection_points cp
		     ,ecosmart.collection_categories cpc
		where cp.category_id = cpc.id
		  and cp.id = $1::bigint
		  and cp.status = 'A'
	`, id)

	point := &model.CollectionPoint{}

	err := row.Scan(&point.ID, &point.Position.Latitude, &point.Position.Longitude, &point.Title, &point.Description, &point.Icon, &point.Category.ID, &point.Category.Name, &point.Category.Description, &point.CreatedAt, &point.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, ErrCollectionPointNotFound
	}

	return point, err
}

func (ur *CollectionPointRepository) GetCollectionPoints() ([]model.CollectionPoint, error) {
	rows, err := ur.db.Query(`
		select cp.id
			  ,cp.position_latitude
			  ,cp.position_longitude
			  ,cp.title
			  ,cp.description
			  ,coalesce(cp.icon, cpc.icon) as icon
			  ,cpc.id as category_id
			  ,cpc.name as category_name
			  ,cpc.description as category_description
			  ,cp.created_at
			  ,cp.updated_at
		  from ecosmart.collection_points cp
		  	  ,ecosmart.collection_categories cpc
		 where cp.category_id = cpc.id
		   and cp.status = 'A'
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	points := []model.CollectionPoint{}

	for rows.Next() {
		point := model.CollectionPoint{}
		err := rows.Scan(&point.ID, &point.Position.Latitude, &point.Position.Longitude, &point.Title, &point.Description, &point.Icon, &point.Category.ID, &point.Category.Name, &point.Category.Description, &point.CreatedAt, &point.UpdatedAt)
		if err != nil {
			return nil, err
		}
		points = append(points, point)
	}

	return points, nil
}
