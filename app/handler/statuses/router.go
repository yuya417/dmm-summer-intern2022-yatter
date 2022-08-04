package statuses

import (
	"net/http"

	"yatter-backend-go/app/app"

	"github.com/go-chi/chi"
	"yatter-backend-go/app/handler/auth"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/statuses/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()
	h := &handler{app: app}
	
	//r.Get("/{username}", h.Show)
	r.With(auth.Middleware(h.app)).Post("/", h.Create)

	return r
}