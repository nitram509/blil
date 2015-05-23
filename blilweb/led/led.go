package led

import (
	"fmt"
	"github.com/boombuler/led"
	"image/color"
)

type LedInfo struct {
	Number int    `json:"number"`
	Type   string `json:"type"`
	Path   string `json:"path"`
}

type Led struct {
	Number int    `json:"number"`
	Color  string `json:"color"`
}

type LedCollectionResource struct {
	Leds []LedInfo `json:"leds"`
}

func DetectAllLeds() LedCollectionResource {
	var i int = 0
	leds := []LedInfo{}
	for devInfo := range led.Devices() {
		led := &LedInfo{i, devInfo.GetType().String(), devInfo.GetPath()}
		leds = append(leds, *led)
		i++
	}
	return LedCollectionResource{Leds: leds}
}

func SetLedColor(nr int, col color.Color) {
	var number int = 0
	for devInfo := range led.Devices() {
		if number == nr {
			dev, err := devInfo.Open()
			if err != nil {
				fmt.Println(err)
				continue
			}
			dev.SetKeepActive(true)
			dev.SetColor(col)
			defer func() {
				dev.Close()
			}()

		}
		number++
	}
}
