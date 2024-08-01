package servers

import (
	"app-library/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func router(hndlr *handlers.ConnectionHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/members", hndlr.GetAllMembers).Methods(http.MethodGet)
	r.HandleFunc("/api/books", hndlr.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/api/borrow", hndlr.BorrowBook).Methods(http.MethodPost)
	r.HandleFunc("/api/return", hndlr.ReturnBook).Methods(http.MethodPost)
	return r
}
