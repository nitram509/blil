package main

import (
    "flag"
    "fmt"
    "time"
    "github.com/boombuler/led"
    "image/color"
    "math/rand"
    "os"
)

var RED = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
var GREEN = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}

var flagSetColor string

func init() {
    flag.StringVar(&flagSetColor, "set-color", "", "Set color for device. The format must be \"#rrggbb\" or \"random\" where rr, gg, bb is any hexadecimal number from 00 to FF.")
}

func main() {

    flag.Parse()

    if (len(flagSetColor) < 1) {
        os.Exit(0)
    }

    for devInfo := range led.Devices() {
        dev, err := devInfo.Open()
        if err != nil {
            fmt.Println(err)
            continue
        }
        defer dev.Close()
        if (flagSetColor == "random") {
            rand.Seed(time.Now().UnixNano())
            dev.SetColor(color.RGBA{uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 0xFF})
        } else {
            dev.SetColor(GREEN)
        }

        time.Sleep(2 * time.Second)
    }

}
