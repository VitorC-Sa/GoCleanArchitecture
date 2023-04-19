package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/VitorC-Sa/GolangCleanArchitecture/app/entity"
	"github.com/VitorC-Sa/GolangCleanArchitecture/app/usecase"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(uc usecase.UserUseCase) UserHandler {
	return UserHandler{UserUseCase: uc}
}

func (uh UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		handleHttpResponse(w, http.StatusBadRequest, err)
		return
	}

	if err = uh.UserUseCase.AddUser(&user); err != nil {
		handleHttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleHttpResponse(w, http.StatusCreated, user)
}

func (uh UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	getUser := func() (*entity.User, error) {
		for _, searchBy := range []string{"id", "email"} {
			value := r.Header.Get(searchBy)

			if value != "" {
				return uh.UserUseCase.GetUser(searchBy, value)
			}
		}
		return nil, errors.New("User not found!")
	}

	user, err := getUser()
	if err != nil {
		handleHttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleHttpResponse(w, http.StatusOK, user)
}

func (uh UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handleHttpResponse(w, http.StatusBadRequest, err)
		return
	}

	if err := uh.UserUseCase.UpdateUser(&user); err != nil {
		handleHttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleHttpResponse(w, http.StatusOK, user)
}

func (uh UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	deleteUser := func() error {
		for _, deleteBy := range []string{"id", "email"} {
			value := r.Header.Get(deleteBy)

			if value != "" {
				return uh.UserUseCase.DeleteUser(deleteBy, value)
			}
		}
		return errors.New("User not found!")
	}

	if err := deleteUser(); err != nil {
		handleHttpResponse(w, http.StatusInternalServerError, err)
		return
	}

	handleHttpResponse(w, http.StatusNoContent, nil)
}

func handleHttpResponse(w http.ResponseWriter, status int, response interface{}) {
	w.WriteHeader(status)
	switch response.(type) {
	case error:
		json.NewEncoder(w).Encode(map[string]error{
			"Error": response.(error),
		})
	case nil:
		// Status 204 - No Content (no response body)
	default:
		json.NewEncoder(w).Encode(response)
	}
}
