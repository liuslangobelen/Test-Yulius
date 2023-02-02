package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"tog.test/no6/model"
	"tog.test/no6/router"
)

func TestGetPing(t *testing.T) {

	r := router.SetupRouter()
	// r.GET("/api/show-cart", controller.ShowCart)
	req, _ := http.NewRequest("GET", "/api/ping", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var shoppingCart string
	json.Unmarshal(w.Body.Bytes(), &shoppingCart)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, shoppingCart)
}

func TestGetAllCart(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.SetupRouter()
	req, _ := http.NewRequest("GET", "/api/show-cart", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var shoppingCart []string
	json.Unmarshal(w.Body.Bytes(), &shoppingCart)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, shoppingCart)
}

func TestNewCompanyHandler(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.SetupRouter()
	code := "code-test" //uuid.New()
	shoppingCart := model.ShoppingCart{
		CodeProduct: code,
		NameProduct: "Demo Product",
		UserId:      1,
		Quantity:    2,
	}
	jsonValue, _ := json.Marshal(shoppingCart)
	req, _ := http.NewRequest("POST", "/api/add-cart", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDeleteCompany(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := router.SetupRouter()
	code := "code-test" //uuid.New()

	req, _ := http.NewRequest("DELETE", "/api/delete-cart/"+code, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
