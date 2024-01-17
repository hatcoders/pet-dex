package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pet-dex-backend/v2/api/errors"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/usecase"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PetController struct {
	Usecase *usecase.PetUseCase
}

func NewPetController(usecase *usecase.PetUseCase) *PetController {
	return &PetController{
		Usecase: usecase,
	}
}

func (pc *PetController) Update(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	petID := chi.URLParam(r, "petID")

	var petToBeUpdated entity.Pet
	err := json.NewDecoder(r.Body).Decode(&petToBeUpdated)
	defer r.Body.Close()

	if err != nil {
		fmt.Printf("Invalid request: could not decode pet data from request body %s", err.Error())
		err := errors.ErrInvalidBody{
			Description: "The body is invalid",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	newSize := petToBeUpdated.Size

	err = pc.Usecase.Update(petID, userID, &entity.Pet{Size: newSize})

	if err != nil {
		fmt.Printf("Error in usecase: %s", err.Error())

		err := errors.ErrInvalidID{
			Description: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
}

func (cntrl *PetController) ListUserPets(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userID")

	userID, erro := strconv.Atoi(idStr)
	if erro != nil {
		http.Error(w, "Error converting 'userID' to int", http.StatusBadRequest)
		return
	}
	pets, err := cntrl.UseCase.ListUserPets(userID)

	if err != nil {
		w.WriteHeader(400)
		return
	}
	json.NewEncoder(w).Encode(&pets)
	w.WriteHeader(http.StatusOK)
}
