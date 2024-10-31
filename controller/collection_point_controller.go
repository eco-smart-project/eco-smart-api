package controller

import (
	"eco-smart-api/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CollectionPointController struct {
	collectionPointRepo *repository.CollectionPointRepository
}

func NewCollectionPointController(collectionPointRepo *repository.CollectionPointRepository) *CollectionPointController {
	return &CollectionPointController{collectionPointRepo}
}

func (cpc *CollectionPointController) GetCollectionPointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	//cast id to int64
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid collection point ID", http.StatusBadRequest)
		return
	}

	point, err := cpc.collectionPointRepo.GetCollectionPointByID(id)

	if err != nil {
		if err == repository.ErrCollectionPointNotFound {
			http.Error(w, "CollectionPoint not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to get collection point", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(point)
}

func (cpc *CollectionPointController) GetCollectionPointsHandler(w http.ResponseWriter, r *http.Request) {
	points, err := cpc.collectionPointRepo.GetCollectionPoints()

	if len(points) == 0 {
		http.Error(w, "No collection points found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Failed to get collection points", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(points)
}
