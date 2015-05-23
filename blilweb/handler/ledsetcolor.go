package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nitram509/blil/blilweb/led"
	"github.com/nitram509/blil/shared"
	"strconv"
)

func LedSetColor(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ledNr, err := strconv.Atoi(vars["ledNr"])
	if err != nil {
		ledNr = -1
	}
	col := shared.MapColor(vars["color"])

	if ledNr < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if col == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	led.SetLedColor(ledNr, col)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	r, g, b, _ := col.RGBA()
	led := &led.Led{Number: ledNr, Color: fmt.Sprintf("%.2x%.2x%.2x", uint8(r), uint8(g), uint8(b))}
	if err := json.NewEncoder(w).Encode(led); err != nil {
		panic(err)
	}
}
