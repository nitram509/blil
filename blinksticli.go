package main

import (
    "fmt"
    "github.com/boombuler/led"
    "image/color"
    "time"
)

var RED = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
var GREEN = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}

func main() {
    for devInfo := range led.Devices() {
        dev, err := devInfo.Open()
        if err != nil {
            fmt.Println(err)
            continue
        }
        defer dev.Close()
        dev.SetColor(GREEN)

        time.Sleep(2 * time.Second)
    }

}
