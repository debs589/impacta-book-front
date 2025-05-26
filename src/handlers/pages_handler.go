package handlers

import (
	"encoding/json"
	"fmt"
	"impacta-book/src/config"
	"impacta-book/src/cookies"
	"impacta-book/src/models"
	"impacta-book/src/request"
	"impacta-book/src/response"
	"impacta-book/src/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "login.html", nil)

}

func LoadUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "user-register.html", nil)
}

func LoadMainPage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publication", config.APIURL)
	responseApi, err := request.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}
	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.TreatStatusCodeError(w, responseApi)
		return
	}

	var publications []models.Publication
	if err = json.NewDecoder(responseApi.Body).Decode(&publications); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.Atoi(cookie["id"])

	utils.RenderTemplate(w, "home.html", struct {
		Publications []models.Publication
		UserID       int
	}{
		Publications: publications,
		UserID:       userID,
	})

}

func LoadPageEditPublication(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["publicationId"])
	if err != nil {
		response.JSON(w, http.StatusBadRequest, response.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publication/%d", config.APIURL, id)
	responseApi, err := request.RequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErrorAPI{Error: err.Error()})
		return
	}
	defer responseApi.Body.Close()

	if responseApi.StatusCode >= 400 {
		response.TreatStatusCodeError(w, responseApi)
		return
	}

	var publication models.Publication
	if err = json.NewDecoder(responseApi.Body).Decode(&publication); err != nil {
		response.JSON(w, http.StatusUnprocessableEntity, response.ErrorAPI{Error: err.Error()})
		return
	}

	utils.RenderTemplate(w, "edit-publication.html", publication)
}
