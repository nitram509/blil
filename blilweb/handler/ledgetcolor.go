package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitram509/blil/blilweb/led"
	"strconv"
)

func LedGetColor(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ledNr, err := strconv.Atoi(vars["ledNr"])
	if err != nil {
		ledNr = -1
	}

	if ledNr < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	led := &led.Led{Number: ledNr}
	if err := json.NewEncoder(w).Encode(led); err != nil {
		panic(err)
	}
}
