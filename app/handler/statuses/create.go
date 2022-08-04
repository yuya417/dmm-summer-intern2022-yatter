package statuses

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Status string
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	status := new(object.Status)
	status.Content = &req.Status

	account := auth.AccountOf(r)
	status.Account = account

	ctx := r.Context()
    if created_status, err := h.app.Dao.Status().CreateStatus(ctx, account, status); err != nil {
        httperror.InternalServerError(w, err)
        return
    } else if created_status == nil {
        httperror.Error(w, http.StatusUnauthorized)
        return
    } else {
        w.Header().Set("Content-Type", "application/json")

        if err := json.NewEncoder(w).Encode(created_status); err != nil {
            httperror.InternalServerError(w, err)
            return
        }
    }
}