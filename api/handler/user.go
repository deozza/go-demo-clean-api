package handler

import (
	"api-test/api/presenter"
	"api-test/useCases/user"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func post(w http.ResponseWriter, r *http.Request, service user.UseCase) {
	w.Header().Set("Content-Type", "application/json")

	var input struct{
		Email    string `json:"email" validate:"required,email,max=100"`
		Password string `json:"password" validate:"required,max=255"`
		Username string `json:"username" validate:"required,max=100"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error adding user"))
		return
	}

	validate := validator.New()
	if err = validate.Struct(input); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		errorAsJson, _ := json.Marshal(ValidationErrorInJson(err))
		w.Write(errorAsJson)
		return
	}

	id, err := service.CreateUser(input.Email, input.Password, input.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error adding user"))
		return
	}

	response := &presenter.User{
		ID: id,
		Email: input.Email,
		Username: input.Username,
	}

	if err = json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error adding user"))
		return
	}
}

func getList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET users"}`))
}

func getSpecific(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	userId := -1
	var err error
	if val, ok := params["userId"]; ok {
		userId, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": "%d"}`, userId)))
}

func BuildUserRouter(r *mux.Router, service user.UseCase){
	userApi := r.PathPrefix("/users").Subrouter()

	userApi.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		post(w, r, service)
	}).Methods(http.MethodPost)
	userApi.HandleFunc("", getList).Methods(http.MethodGet)
	userApi.HandleFunc("/{userId}", getSpecific).Methods(http.MethodGet)
}