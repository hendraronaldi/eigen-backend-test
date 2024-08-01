package servers

import (
	"app-library/internal/app_config"
	"app-library/internal/databases"
	"app-library/internal/handlers"
	"app-library/internal/repositories"
	"app-library/internal/services"
	"context"
	"log"
	"net/http"
)

// Server main struct implementation.
type Server struct {
	ctx              context.Context
	config           *app_config.AppConfig
	libraryRepo      repositories.Interface
	databaseInstance *databases.Instance
	handler          *handlers.ConnectionHandler
}

// Init server instance.
func Init() *Server {
	s := Server{}

	config := app_config.Get()
	s.config = config
	s.ctx = context.Background()
	s.initializeRepo()

	s.handler = &handlers.ConnectionHandler{
		Ctx:            s.ctx,
		LibraryService: services.NewLibrary(s.libraryRepo),
	}
	return &s
}

func (s *Server) initializeRepo() {
	di := databases.New()
	di.ConnectLibraryDB()
	s.databaseInstance = di
	s.libraryRepo = repositories.NewRepository(s.databaseInstance.LibraryDB)
}

// Run server
func (s *Server) Run() {
	r := router(s.handler)
	log.Println("server starts on", s.config.ENV.PORT)
	if err := http.ListenAndServe(s.config.ENV.PORT, r); err != nil {
		log.Panic("Error serving:", err.Error())
	}
}

func (s *Server) Close() {
	log.Println("Closing server")
	if s.databaseInstance != nil {
		s.databaseInstance.CloseLibraryDB()
	}
}
