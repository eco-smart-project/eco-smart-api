package model

type CollectionPointCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CollectionPointCategory struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CollectionPoint struct {
	ID          uint64                     `json:"id"`
	Position    CollectionPointCoordinates `json:"position"`
	Title       string                     `json:"title"`
	Description string                     `json:"description"`
	Icon        string                     `json:"icon"`
	Category    CollectionPointCategory    `json:"category"`
	CreatedAt   string                     `json:"created_at"`
	UpdatedAt   string                     `json:"updated_at"`
	DeletedAt   string                     `json:"deleted_at"`
}
