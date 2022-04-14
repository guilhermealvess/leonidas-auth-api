package rest

import (
	"api-auth/src/adapter/repository"
	"api-auth/src/entity"
	"api-auth/src/usecase"
	"encoding/json"
	"net/http"
)

type AccountController struct {
	AccountRepository entity.AccountRepository
	ProjectRepository entity.ProjectRepository
}

func NewAccountController(db repository.DocumentDB, cache repository.Cache) *AccountController {
	accountRepo := repository.NewAccountRepositoryDB(db, cache)
	projectRepo := repository.NewProjectRepositoryDB(db, cache)

	return &AccountController{
		AccountRepository: accountRepo,
		ProjectRepository: projectRepo,
	}
}

func (a *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var payload usecase.AccountDtoInput
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&payload); err != nil {
		respondWithError(w, http.StatusBadRequest, msgErrorPayload)
	}

	processAccount := usecase.NewProcessAccount(a.AccountRepository, a.ProjectRepository)
	output, err := processAccount.ExecuteCreateNewAccount(payload)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusCreated, output)
}
