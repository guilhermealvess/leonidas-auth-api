package rest

import (
	"api-auth/src/adapter/repository"
	"api-auth/src/entity"
	"api-auth/src/usecase"
	"fmt"
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

func (a *AccountController) ActivationAccount(w http.ResponseWriter, r *http.Request) {
	activationKey := r.URL.Query().Get("key")

	processAccount := usecase.NewProcessAccount(a.AccountRepository, a.ProjectRepository)
	output := processAccount.ActivateAccount(activationKey)

	if output.Success {
		http.Redirect(w, r, output.Url, http.StatusPermanentRedirect)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, output.Error)
	return
}
