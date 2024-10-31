package handler

import (
	"eco-smart-api/controller"

	"github.com/gorilla/mux"
)

type CollectionPointHandler struct {
	collectionPointController *controller.CollectionPointController
}

func NewCollectionPointHandler(collectionPointController *controller.CollectionPointController) *CollectionPointHandler {
	return &CollectionPointHandler{collectionPointController}
}

func (cph *CollectionPointHandler) RegisterRoutes(mux *mux.Router) {
	mux.HandleFunc("/collection-points", cph.collectionPointController.GetCollectionPointsHandler).Methods("GET")
	mux.HandleFunc("/collection-points/{id}", cph.collectionPointController.GetCollectionPointHandler).Methods("GET")
}
