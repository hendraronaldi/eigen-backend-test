package handlers

import (
	"app-library/internal/helpers"
	"app-library/internal/models"
	"fmt"
	"net/http"
)

type getAllBooksResponse struct {
	Message string         `json:"message"`
	Data    []*models.Book `json:"data"`
}

// GetAllBooks godoc
// @Summary      List books
// @Description  Get all books
// @Tags         books
// @Produce      json
// @Success      200  {object}  getAllBooksResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/books [get]
func (c *ConnectionHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	res := new(getAllBooksResponse)

	datas, err := c.LibraryService.GetAllBooks(c.Ctx)
	if err != nil {
		res.Message = fmt.Sprintf("error get all books: %v", err)
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, res)
		return
	}
	res.Data = datas

	res.Message = "success"
	helpers.WriteJSONResponse(w, http.StatusOK, res)
}
