package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/kserevena/perkbox-tech-test/modal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Define tests data
var testCreationTime = time.Date(2020, 9, 13, 0, 0, 0, 0, time.UTC)
var testExpiryTime = time.Date(2021, 9, 13, 0, 0, 0, 0, time.UTC)

var inputCoupon = modal.Coupon{
	Name:      "testCoupon1",
	Brand:     "testBrand1",
	Value:     10,
	CreatedAt: testCreationTime.String(),
	Expiry:    testExpiryTime.String(),
}

// Tests below

func Test_PostCoupon(t *testing.T) {

	// Init database and clear it afterwards
	storage := initStorage()
	defer clearDb(t, storage)

	// Start server:
	router := mux.NewRouter()
	router = initRoutes(router, storage)

	server := httptest.NewServer(router)
	defer server.Close()

	// POST Coupon to endpoint
	postBody := new(bytes.Buffer)
	assert.NoError(t, json.NewEncoder(postBody).Encode(inputCoupon))

	response, err := server.Client().Post(server.URL+"/coupon", "application/json", postBody)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseBody modal.Coupon
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Equal(t, inputCoupon, responseBody)
}

func Test_GetCoupons(t *testing.T) {

	// Init database and clear it afterwards
	storage := initStorage()
	defer clearDb(t, storage)

	// Start server:
	router := mux.NewRouter()
	router = initRoutes(router, storage)

	server := httptest.NewServer(router)
	defer server.Close()

	// Populate DB with test coupons

	err := inputCoupon.CreateCoupon(storage)
	assert.NoError(t, err)

	// Call GET endpoint
	response, err := server.Client().Get(server.URL + "/coupon")

	// Verify response
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	var responseBody []modal.Coupon
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	assert.NoError(t, err)

	// TODO: work out why date fields are not returned correctly by DB query
	assert.Equal(t, []modal.Coupon{inputCoupon}, responseBody)

}

func clearDb(t *testing.T, db *sql.DB) {
	statement := "drop table coupon"
	_, err := db.Exec(statement)
	assert.NoError(t, err)
}
