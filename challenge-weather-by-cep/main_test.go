package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestHandleRequest_ValidCEP(t *testing.T) {
	app := fiber.New()
	app.Get("/:cep", handleRequest)

	// Simular um CEP válido
	req := httptest.NewRequest(http.MethodGet, "/04709110", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)

	// Verificar se os campos de temperatura estão presentes
	_, tempCExists := body["temp_C"]
	_, tempFExists := body["temp_F"]
	_, tempKExists := body["temp_K"]
	assert.True(t, tempCExists)
	assert.True(t, tempFExists)
	assert.True(t, tempKExists)
}

func TestHandleRequest_InvalidCEP(t *testing.T) {
	app := fiber.New()
	app.Get("/:cep", handleRequest)

	// Simular um CEP inválido
	req := httptest.NewRequest(http.MethodGet, "/123", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	var body map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)

	assert.Equal(t, "invalid zipcode", body["error"])
}

func TestHandleRequest_CEPNotFound(t *testing.T) {
	app := fiber.New()
	app.Get("/:cep", handleRequest)

	// Simular um CEP que não existe
	req := httptest.NewRequest(http.MethodGet, "/99999999", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)

	var body map[string]interface{}
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)

	assert.Equal(t, "can not find zipcode", body["error"])
}

func TestFetchData_ValidURL(t *testing.T) {
	// URL simulada
	url := "https://viacep.com.br/ws/04709110/json/"
	resp, err := fetchData(nil, url)

	assert.NoError(t, err)
	assert.NotEmpty(t, resp)

	var data ViaCepResponse
	err = json.Unmarshal(resp, &data)
	assert.NoError(t, err)

	assert.Equal(t, "04709110", data.Cep)
}

func TestRemoveAccents(t *testing.T) {
	input := "São Paulo"
	expected := "Sao Paulo"

	result := removeAccents(input)
	assert.Equal(t, expected, result)
}
