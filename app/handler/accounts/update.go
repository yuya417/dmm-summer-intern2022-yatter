package accounts

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts/update_credentials`
type UpdateRequest struct {
	DisplayName string
	Note string
	Avatar string
	Header string
}

type Queries struct {
    DisplayName string
	Note string
	Avatar string
	Header string
}

// Handle request for `POST /v1/accounts/update_credentials`
func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	var req UpdateRequest
	queries := []string{"display_name", "note", "avatar", "header"}

	account := auth.AccountOf(r)
	account.DisplayName = &req.DisplayName
	account.Note = &req.Note
	account.Avatar = &req.Avatar
	account.Header = &req.Header

	for _, v := range queries {
        switch v {
        case "display_name":
			value := r.FormValue(v)
            account.DisplayName = &value
        case "note":
			value := r.FormValue(v)
            account.Note = &value
        case "avatar":
			file, file_header, err := r.FormFile(v)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()
			file_dir := "assets/account/avatar/" + file_header.Filename
			dst, err := os.Create(file_dir)
			if err != nil {
				log.Printf("err %v", err)
			}
			defer dst.Close()
			if _, err = io.Copy(dst, file); err != nil {
				log.Printf("err %v", err)
			}
            account.Avatar = &file_header.Filename
		case "header":
			file, file_header, err := r.FormFile(v)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()
			file_dir := "assets/account/header/" + file_header.Filename
			dst, err := os.Create(file_dir)
			if err != nil {
				log.Printf("err %v", err)
			}
			defer dst.Close()
			if _, err = io.Copy(dst, file); err != nil {
				log.Printf("err %v", err)
			}
            account.Header = &file_header.Filename
        default:
            httperror.Error(w, http.StatusBadRequest)
            return
        }
    }

	ctx := r.Context()
    if created_user, err := h.app.Dao.Account().UpdateUser(ctx, account); err != nil {
        httperror.InternalServerError(w, err)
        return
    } else if created_user == nil {
        httperror.Error(w, http.StatusUnauthorized)
        return
    } else {
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(created_user); err != nil {
            httperror.InternalServerError(w, err)
            return
        }
    }
}