package handlers

import (
	"app-library/internal/helpers"
	"app-library/internal/models"
	"fmt"
	"net/http"
)

type getAllMembersResponse struct {
	Message string           `json:"message"`
	Data    []*models.Member `json:"data"`
}

func (c *ConnectionHandler) GetAllMembers(w http.ResponseWriter, r *http.Request) {
	res := new(getAllMembersResponse)

	datas, err := c.LibraryService.GetAllMembers(c.Ctx)
	if err != nil {
		res.Message = fmt.Sprintf("error get all books: %v", err)
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, res)
		return
	}
	res.Data = datas

	res.Message = "success"
	helpers.WriteJSONResponse(w, http.StatusOK, res)
}
