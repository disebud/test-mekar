package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/disebud/api-users-test/main/master/models"
	"github.com/disebud/api-users-test/main/master/usecases"
	"github.com/disebud/api-users-test/main/master/utils"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUseCase usecases.UserUseCase
}

func UserController(r *mux.Router, service usecases.UserUseCase) {
	UserHandler := UserHandler{service}
	r.HandleFunc("/users", UserHandler.ListUsers).Methods(http.MethodGet)
	r.HandleFunc("/jobs", UserHandler.ListJobs).Methods(http.MethodGet)
	r.HandleFunc("/educations", UserHandler.ListEducations).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", UserHandler.GetUserId).Methods(http.MethodGet)
	r.HandleFunc("/user", UserHandler.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", UserHandler.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", UserHandler.DeleteUserId).Methods(http.MethodDelete)
}

func (e UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars, "GET LIST USER")

	// log.Println(infinfoType)
	currentDate := time.Now()
	Users, err := e.UserUseCase.GetAllUsers()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.Response{}
	response.Date = "Today / " + currentDate.Format("02-January-2006")
	response.Status = http.StatusOK
	response.Message = "List Users"
	response.Result = Users
	byteOfUsers, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfUsers))
}

func (e UserHandler) ListJobs(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// fmt.Println(vars, "ok")

	currentDate := time.Now()
	Jobs, err := e.UserUseCase.GetJob()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.GenerateResponse("Today / "+currentDate.Format("02-January-2006"), http.StatusOK, "List Jobs", Jobs)

	byteOfUsers, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfUsers))
}

func (e UserHandler) ListEducations(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// fmt.Println(vars, "ok")

	currentDate := time.Now()
	Educations, err := e.UserUseCase.GetEducation()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.GenerateResponse("Today / "+currentDate.Format("02-January-2006"), http.StatusOK, "List Educations", Educations)

	byteOfUsers, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(byteOfUsers))
}

func (e UserHandler) GetUserId(w http.ResponseWriter, r *http.Request) {
	currentDate := time.Now()
	vars := mux.Vars(r)
	Employees, err := e.UserUseCase.GetUserById(vars["id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	response := utils.Response{}
	response.Status = http.StatusAccepted
	response.Message = " Data User id User = " + vars["id"]
	response.Date = "Today / " + currentDate.Format("02-January-2006")
	response.Result = Employees
	byteOfUsers, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops try again"))
	}
	w.Header().Set("content-type", "application/json")
	w.Write(byteOfUsers)
}

func (e UserHandler) DeleteUserId(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	_, err := e.UserUseCase.DeleteUserById(param["id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Delete Data successful , User ID = " + param["id"]))
	log.Println("Delete Data successful , User ID = " + param["id"])
}

func (e UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var re models.User
	err := json.NewDecoder(r.Body).Decode(&re)
	if err != nil {
		log.Println(err)
	}
	log.Println(re)
	err = e.UserUseCase.CreateUser(re)
	if err != nil {
		log.Println(err)
	}
	log.Println("Insert successful")
	w.Write([]byte("Insert successful"))

}

func (e UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	currentDate := time.Now()
	date := "Today / " + currentDate.Format("02-January-2006")
	var User models.User
	param := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&User)
	err = e.UserUseCase.UpdateUser(param["id"], User)
	if err != nil {
		w.Write([]byte("Data not found !!"))
		log.Println(err)
		return
	}
	byteOfInfo, _ := json.Marshal(utils.GenerateResponse(date, http.StatusOK, "Success Update", User))
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfInfo)
	w.Write([]byte("\n" + "Update Data successful , NIK = " + param["id"]))
	log.Println("Update Data successful , NIK = " + param["id"])
}
