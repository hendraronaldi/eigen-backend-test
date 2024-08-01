package handlers

import (
	"app-library/internal/helpers"
	"app-library/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type returnBookResponse struct {
	Message string `json:"message"`
}

// ReturnBook godoc
// @Summary      Return books
// @Description  Member return books
// @Tags         members
// @Accept       json
// @Produce      json
// @Param data body models.Return true "The input return book by member_id, book_ids, date returned_at"
// @Success      200  {object}  returnBookResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/return [post]
func (c *ConnectionHandler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	res := new(returnBookResponse)

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.Message = fmt.Sprintf("error parse form: %v", err)
		helpers.WriteJSONResponse(w, http.StatusBadRequest, res)
		return
	}

	req := new(models.Return)
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

	if err := c.LibraryService.ReturnBook(c.Ctx, req); err != nil {
		res.Message = fmt.Sprintf("failed borrow book: %v", err)
		helpers.WriteJSONResponse(w, http.StatusBadRequest, res)
		return
	}

	res.Message = "success"
	helpers.WriteJSONResponse(w, http.StatusOK, res)
}
