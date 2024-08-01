package handlers

import (
	"app-library/internal/services"
	"context"
)

type ConnectionHandler struct {
	Ctx            context.Context
	LibraryService services.LibraryInterface
}
