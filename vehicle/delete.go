package vehicle

import (
	"net/http"
	"strconv"

	"github.com/MatRouillard/vehicle-server/storage"
	"go.uber.org/zap"
)

type DeleteHandler struct {
	store  storage.Store
	logger *zap.Logger
}

func NewDeleteHandler(store storage.Store, logger *zap.Logger) *DeleteHandler {
	return &DeleteHandler{
		store:  store,
		logger: logger.With(zap.String("handler", "delete_vehicles")),
	}
}

func (d *DeleteHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		http.Error(rw, "Failed to parse ID", http.StatusInternalServerError)
		return
	}

	result, err := d.store.Vehicle().Delete(r.Context(), id)
	if err != nil {
		http.Error(rw, "Failed to delete vehicle", http.StatusInternalServerError)
		return
	}

	if result {
		rw.WriteHeader(http.StatusNoContent)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}

}
