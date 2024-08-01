package handlers

import (
	"app-library/internal/helpers"
	"app-library/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type borrowBookResponse struct {
	Message string `json:"message"`
}

// BorrowBook godoc
// @Summary      Borrow books
// @Description  Member borrow books
// @Tags         members
// @Accept       json
// @Produce      json
// @Success      200  {object}  borrowBookResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/borrow [post] {object} models.Borrow
func (c *ConnectionHandler) BorrowBook(w http.ResponseWriter, r *http.Request) {
	res := new(borrowBookResponse)

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.Message = fmt.Sprintf("error parse form: %v", err)
		helpers.WriteJSONResponse(w, http.StatusBadRequest, res)
		return
	}

	req := new(models.Borrow)
	err = json.Unmarshal(body, &req)
	if err != nil {
		res.Message = fmt.Sprintf("error decode request: %v", err)
		helpers.WriteJSONResponse(w, http.StatusBadRequest, res)
		return
	}

	if err := r.ParseForm(); err != nil {
		res.Message = fmt.Sprintf("error parse form: %v", err)
		helpers.WriteJSONResponse(w, http.StatusBadRequest, res)
		return
	}

	if err := c.LibraryService.BorrowBook(c.Ctx, req); err != nil {
		res.Message = fmt.Sprintf("failed borrow book: %v", err)
		helpers.WriteJSONResponse(w, http.StatusBadRequest, res)
		return
	}

	res.Message = "success"
	helpers.WriteJSONResponse(w, http.StatusOK, res)
}
