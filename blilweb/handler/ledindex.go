package handler

import (
	"encoding/json"
	"github.com/nitram509/blil/blilweb/led"
	"net/http"
)

func LedIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	leds := led.DetectAllLeds()
	if err := json.NewEncoder(w).Encode(leds); err != nil {
		panic(err)
	}
}
