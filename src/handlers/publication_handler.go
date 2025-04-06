package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"impacta-book/src/config"
	"impacta-book/src/request"
	"impacta-book/src/response"
	"net/http"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	publication, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/", config.APIURL)
	responseApi, err := request.RequestWithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(publication))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}
	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.TreatStatusCodeError(w, responseApi)
		return
	}

	response.JSON(w, responseApi.StatusCode, nil)

}
