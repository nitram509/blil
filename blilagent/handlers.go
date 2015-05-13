package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/nitram509/bled/shared"
    "strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func LedIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    leds := detectAllLeds()
    if err := json.NewEncoder(w).Encode(leds); err != nil {
        panic(err)
    }
}

func LedSetColor(w http.ResponseWriter, req *http.Request) {
    vars := mux.Vars(req)
    ledNr, err := strconv.Atoi(vars["ledNr"])
    if (err != nil) {
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

    setLedColor(ledNr, col)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    r, g, b, _ := col.RGBA()
    led := &Led{Number:ledNr, Color: fmt.Sprintf("%.2x%.2x%.2x", uint8(r), uint8(g), uint8(b))}
    if err := json.NewEncoder(w).Encode(led); err != nil {
        panic(err)
    }
}
