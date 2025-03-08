package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"impacta-book/src/models"
	"impacta-book/src/response"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	responseApi, err := http.Post("http://localhost:5000/login/", "application/json", bytes.NewBuffer(user))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}
	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.TreatStatusCodeError(w, responseApi)
		return
	}

	var authenticationData models.AuthenticationData
	if err = json.NewDecoder(responseApi.Body).Decode(&authenticationData); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}

	fmt.Println("response status:", responseApi.StatusCode, "response Body: ", responseApi.Body)
}
