package timelines

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `GET /v1/timelines/public`
func (h *handler) Index(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
	max_id, _ := strconv.Atoi(r.FormValue("max_id"))
    since_id, _ := strconv.Atoi(r.FormValue("since_id"))
    limit, _ := strconv.Atoi(r.FormValue("limit"))

    if statuses, err := h.app.Dao.Timeline().AllStatuses(ctx, max_id, since_id, limit); err != nil {
        httperror.InternalServerError(w, err)
        return
    } else if statuses == nil {
        httperror.Error(w, http.StatusNotFound)
        return
    } else {
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(statuses); err != nil {
            httperror.InternalServerError(w, err)
            return
        }
    }
}