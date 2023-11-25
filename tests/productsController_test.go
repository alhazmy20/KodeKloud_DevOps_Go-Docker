package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	productsController "github.com/alhazmy20/Go-WebAPI/controllers"
	database "github.com/alhazmy20/Go-WebAPI/database"
	migrations "github.com/alhazmy20/Go-WebAPI/migrations"
	model "github.com/alhazmy20/Go-WebAPI/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var router = mux.NewRouter()

func MainTest(m *testing.M) {
	os.Exit(m.Run())
}

func TestIndex(t *testing.T) {
	database.EnvFilePath = "../.env"
	migrations.Drop_Products_Table()
	migrations.Create_Products_Table()
	migrations.Insert_Product()
	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()
	router.HandleFunc("/products", productsController.Index)
	router.ServeHTTP(responseRecorder, req)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

}

func TestStore(t *testing.T) {
	database.EnvFilePath = "../.env"
	migrations.Drop_Products_Table()
	migrations.Create_Products_Table()
	product := model.Product{
		Name:     "Test Product",
		Price:    10.99,
		Quantity: 100,
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products", productsController.Store)
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestShow(t *testing.T) {
	database.EnvFilePath = "../.env"
	migrations.Drop_Products_Table()
	migrations.Create_Products_Table()
	migrations.Insert_Product()
	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", productsController.Show)
	router.ServeHTTP(rr, req)
	eq := assert.Equal(t, http.StatusOK, rr.Code)
	if eq == false {
		t.Errorf("Test Show fails")
	}
}
func TestUpdate(t *testing.T) {
	database.EnvFilePath = "../.env"
	migrations.Drop_Products_Table()
	migrations.Create_Products_Table()
	migrations.Insert_Product()
	product := model.Product{
		Name:     "Updated Product",
		Price:    19.99,
		Quantity: 50,
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("PUT", "/products/1", bytes.NewBuffer(productJSON))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", productsController.Update)
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
func TestDestroy(t *testing.T) {
	database.EnvFilePath = "../.env"
	migrations.Drop_Products_Table()
	migrations.Create_Products_Table()
	migrations.Insert_Product()
	req, err := http.NewRequest("DELETE", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{id}", productsController.Destroy)
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
