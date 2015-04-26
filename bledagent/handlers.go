package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/boombuler/led"
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

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func detectAllLeds() LedResource {
    var i int = 0
    leds := []Led{}
    for devInfo := range led.Devices() {
        led := &Led{i, devInfo.GetType().String(), devInfo.GetPath()}
        leds = append(leds, *led)
        i++
    }
    return LedResource{Leds:leds}
}