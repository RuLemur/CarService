package app

import (
	"car-service/middleware/internal/datastruct"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

// AddUserResponse ...
type AddUserResponse struct {
	Id int64 `json:"id"`
}

// CreateGarageResponse ...
type CreateGarageResponse struct {
	Id int64 `json:"id"`
}

// GetUserResponse ...
type GetUserResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	GarageID  int64     `json:"garage_id"`
	UpdatedAt time.Time `json:"Updated_at"`
}

type GetGarageResponse struct {
	Cars     []int64 `json:"cars"`
	GarageID int64   `json:"garage_id"`
}

type CarSearchResponse struct {
	Cars []*datastruct.CarModel `json:"car"`
}

type AddToGarageRequest struct {
	GarageId int64 `json:"garage_id"`
	ModelId  int64 `json:"model_id"`
}

type AddToGarageResponse struct {
	CarId int64 `json:"id"`
}

type Error struct {
	Message string
}

// GetUser returns a basic "Hello World!" message
func (a *Api) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response := Error{
			Message: fmt.Sprintf("Fail to parse userID"),
		}
		a.jsonResponse(w, response, http.StatusBadRequest)
		return
	}
	getUserResponse, err := a.GetUserGRPC(userID)
	if err != nil {
		response := Error{
			Message: fmt.Sprintf("Fail server %d!", userID),
		}
		a.jsonResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := GetUserResponse{
		ID:        getUserResponse.Id,
		Username:  getUserResponse.Username,
		GarageID:  getUserResponse.GarageId,
		UpdatedAt: getUserResponse.UpdatedAt,
	}
	a.jsonResponse(w, response, http.StatusOK)
}

// AddUser returns
func (a *Api) AddUser(w http.ResponseWriter, r *http.Request) {

	var user datastruct.User
	err := a.parseRequest(r, &user)
	if err != nil {
		errResponse := Error{
			Message: fmt.Sprintf("Fail to parse request!"),
		}
		a.jsonResponse(w, errResponse, http.StatusBadRequest)
	}
	addUserResponse, err := a.AddUserGRPC(user)
	if err != nil {
		errResponse := Error{
			Message: fmt.Sprint("Fail server!"),
		}
		a.jsonResponse(w, errResponse, http.StatusInternalServerError)
		return
	}
	response := AddUserResponse{
		Id: addUserResponse.Id,
	}
	a.jsonResponse(w, response, http.StatusOK)
}

// CreateGarage returns
func (a *Api) CreateGarage(w http.ResponseWriter, r *http.Request) {
	createGarageResponse, err := a.CreateGarageGRPC()
	if err != nil {
		response := Error{
			Message: fmt.Sprint("Fail server!"),
		}
		a.jsonResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := CreateGarageResponse{
		Id: createGarageResponse.Id,
	}
	a.jsonResponse(w, response, http.StatusOK)
}

// GetGarage returns
func (a *Api) GetGarage(w http.ResponseWriter, r *http.Request) {
	var garage datastruct.Garage
	err := a.parseRequest(r, &garage)
	if err != nil {
		response := Error{
			Message: fmt.Sprint("Fail server!"),
		}
		a.jsonResponse(w, response, http.StatusInternalServerError)
	}

	addUserResponse, err := a.GetGarageGRPC(garage.GarageID)
	if err != nil {
		response := Error{
			Message: fmt.Sprint("Fail server!"),
		}
		a.jsonResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := GetGarageResponse{
		Cars:     addUserResponse.Cars,
		GarageID: addUserResponse.GarageId,
	}
	a.jsonResponse(w, response, http.StatusOK)
}

// CarSearch returns
func (a *Api) CarSearch(w http.ResponseWriter, r *http.Request) {
	var model datastruct.CarModel
	err := a.parseRequest(r, &model)
	if err != nil {
		response := Error{
			Message: fmt.Sprint("Fail server!"),
		}
		a.jsonResponse(w, response, http.StatusInternalServerError)
	}

	carSearchResponse, err := a.CarSearchGRPC(model)
	if err != nil {
		response := Error{
			Message: fmt.Sprint("Fail server!"),
		}
		a.jsonResponse(w, response, http.StatusInternalServerError)
		return
	}
	var cars []*datastruct.CarModel
	for _, car := range carSearchResponse.Car {
		cars = append(cars, &datastruct.CarModel{
			ID:         car.Id,
			Brand:      car.Brand,
			Model:      car.Model,
			Equipment:  car.Equipment,
			EngineType: car.EngineType,
		})
	}
	response := CarSearchResponse{
		Cars: cars,
	}
	a.jsonResponse(w, response, http.StatusOK)
}

// AddToGarage returns
func (a *Api) AddToGarage(w http.ResponseWriter, r *http.Request) {
	var garage AddToGarageRequest
	err := a.parseRequest(r, &garage)
	if err != nil {
		response := Error{
			Message: fmt.Sprint("Fail server!"),
		}
		a.jsonResponse(w, response, http.StatusInternalServerError)
	}

	addToGarageResponse, err := a.AddToGarageGRPC(garage.GarageId, garage.ModelId)
	if err != nil {
		response := Error{
			Message: fmt.Sprint("Fail server!"),
		}
		a.jsonResponse(w, response, http.StatusInternalServerError)
		return
	}
	response := AddToGarageResponse{
		CarId: addToGarageResponse.CarId,
	}
	a.jsonResponse(w, response, http.StatusOK)
}
