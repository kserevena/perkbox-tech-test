package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kserevena/perkbox-tech-test/modal"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_PostCoupon(t *testing.T) {

	// Init database and clear it afterwards
	storage := initStorage()
	defer clearDb(t, storage)

	testCreationTime := time.Date(2020, 9, 13, 0, 0, 0, 0, time.UTC)
	testExpiryTime := time.Date(2021, 9, 13, 0, 0, 0, 0, time.UTC)

	// Define input data
	inputCoupon := modal.Coupon{
		Name:      "testCoupon1",
		Brand:     "testBrand1",
		Value:     10,
		CreatedAt: testCreationTime.String(),
		Expiry:    testExpiryTime.String(),
	}

	// Start server:
	router := mux.NewRouter()
	router = initRoutes(router, storage)

	server := httptest.NewServer(router)
	defer server.Close()

	// POST Coupon to endpoint
	//recorder := httptest.NewRecorder()
	postBody := new(bytes.Buffer)
	assert.NoError(t, json.NewEncoder(postBody).Encode(inputCoupon))
	//request := httptest.NewRequest(http.MethodPost, server.URL +"/coupon", postBody)

	response, err := server.Client().Post(server.URL+"/coupon", "application/json", postBody)
	assert.NoError(t, err)

	var responseBody modal.Coupon
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	assert.NoError(t, err)

	assert.Equal(t, inputCoupon, responseBody)
}

func clearDb(t *testing.T, db *sql.DB) {
	statement := "drop table coupon"
	_, err := db.Exec(statement)
	assert.NoError(t, err)
}
