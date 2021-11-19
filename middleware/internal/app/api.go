package app

import (
	"car-service/internal/app/car_service/endpoint"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"log"

	"github.com/go-chi/chi"
)

// ValidBearer is a hardcoded bearer token for demonstration purposes.
const ValidBearer = "123456"


type Api struct {
	grpcClient endpoint.CarServiceClient
	ctx        context.Context
}

func NewApi(grpcClient endpoint.CarServiceClient) *Api {
	return &Api{grpcClient: grpcClient, ctx: context.Background()}
}

func (a *Api) jsonResponse(w http.ResponseWriter, data interface{}, c int) {
	dj, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

func (a *Api) parseRequest(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	return nil
}

// RequireAuthentication is an example middleware handler that checks for a
// hardcoded bearer token. This can be used to verify session cookies, JWTs
// and more.
func (a *Api) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Make sure an Authorization header was provided
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		// This is where token validation would be done. For this boilerplate,
		// we just check and make sure the token matches a hardcoded string
		if token != ValidBearer {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		// Assuming that passed, we can execute the authenticated handler
		next.ServeHTTP(w, r)
	})
}

// NewRouter returns an HTTP handler that implements the routes for the API
func (a *Api) NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(a.RequireAuthentication)

	// Register the API routes
	r.Get("/user/{id}", a.GetUser)
	r.Post("/user/", a.AddUser)
	r.Post("/garage/get", a.GetGarage) //TODO: прееделать на get
	r.Post("/garage/create", a.CreateGarage)
	r.Post("/car/search", a.CarSearch)
	r.Post("/garage/add", a.AddToGarage)


	return r
}
