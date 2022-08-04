package statuses

import (
    "encoding/json"
    "net/http"

    "yatter-backend-go/app/handler/httperror"

    "github.com/go-chi/chi"
)

// Handle request for GET /v1/statuses/{id}
func (h handler) Show(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    ctx := r.Context()
    if status, err := h.app.Dao.Status().FindByStatusID(ctx, id); err != nil {
        httperror.InternalServerError(w, err)
        return
    } else if status == nil {
        httperror.Error(w, http.StatusNotFound)
        return
    } else {
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(status); err != nil {
            httperror.InternalServerError(w, err)
            return
        }
    }
}