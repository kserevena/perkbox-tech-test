package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rychanfox/couponService/modal"

	"github.com/gorilla/mux"
)

func NewIndex(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/coupon", func(res http.ResponseWriter, req *http.Request) {
		var c modal.Coupon
		coupons, err := c.GetCoupons(db)
		if err != nil {
			log.Printf("error: %s", err.Error())
			ErrorResponse(res, http.StatusBadRequest, "Database error")
			return
		}
		json.NewEncoder(res).Encode(coupons)
	}).Methods("GET")
}

func NewCreate(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/coupon", func(res http.ResponseWriter, req *http.Request) {
		var c modal.Coupon
		decoder := json.NewDecoder(req.Body)
		if err := decoder.Decode(&c); err != nil {
			log.Printf("error: %s", err.Error())
			ErrorResponse(res, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer req.Body.Close()

		if err := c.CreateCoupon(db); err != nil {
			log.Printf("error: %s", err.Error())
			ErrorResponse(res, http.StatusBadRequest, "Database issue")
			return
		}

		Respone(res, http.StatusOK, c)
	})
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	response, _ := json.Marshal(map[string]string{"error": message})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func Respone(w http.ResponseWriter, code int, message interface{}) {
	response, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
