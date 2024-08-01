package servers

import (
	"app-library/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

func router(hndlr *handlers.ConnectionHandler) *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	r.HandleFunc("/api/members", hndlr.GetAllMembers).Methods(http.MethodGet)
	r.HandleFunc("/api/books", hndlr.GetAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/api/borrow", hndlr.BorrowBook).Methods(http.MethodPost)
	r.HandleFunc("/api/return", hndlr.ReturnBook).Methods(http.MethodPost)
	return r
}
