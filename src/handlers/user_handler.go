package handlers

import (
	"bytes"
	"encoding/json"
	"impacta-book/src/response"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nickname": r.FormValue("nickname"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := "http://localhost:5000/user/"
	responseApi, err := http.Post(url, "application/json", bytes.NewBuffer(user))
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
