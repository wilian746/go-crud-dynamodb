package health

import (
	"errors"
	"github.com/wilian746/go-crud-dynamodb/internal/handlers"
	"github.com/wilian746/go-crud-dynamodb/internal/repository/adapter"
	HttpStatus "github.com/wilian746/go-crud-dynamodb/utils/http"
	"net/http"
)

type Handler struct {
	handlers.Interface
	Repository adapter.Interface
}

func NewHandler(repository adapter.Interface) handlers.Interface {
	return &Handler{
		Repository: repository,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	if !h.Repository.Health() {
		HttpStatus.StatusInternalServerError(w, r, errors.New("Relational database not alive"))
		return
	}

	HttpStatus.StatusOK(w, r, "Service OK")
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusMethodNotAllowed(w, r)
}

func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	HttpStatus.StatusNoContent(w, r)
}
