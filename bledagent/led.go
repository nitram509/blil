package main
import "github.com/boombuler/led"

type Led struct {
    Number int    `json:"number"`
    Type   string `json:"type"`
    Path   string `json:"path"`
}

type LedCollectionResource struct {
    Leds []Led `json:"leds"`
}

func detectAllLeds() LedCollectionResource {
    var i int = 0
    leds := []Led{}
    for devInfo := range led.Devices() {
        led := &Led{i, devInfo.GetType().String(), devInfo.GetPath()}
        leds = append(leds, *led)
        i++
    }
    return LedCollectionResource{Leds:leds}
}