package main

import (
    "fmt"
    "time"
    "github.com/boombuler/led"
    "image/color"
    "math/rand"
    "os"
    "sort"
    "strings"
    "regexp"
    "encoding/hex"
    "gopkg.in/alecthomas/kingpin.v1"
)

var (
    VERSION = "0.0.1"
    flagSetColor = kingpin.Flag("set-color", "Set color for device. The format must be \"#rrggbb\" or \"random\" or an HTML color name, like \"green\"").String()
    flagListColors = kingpin.Flag("list-colors", "List all available HTML color names").Bool()
)

func printListColorNames() {
    var colorNames []string
    for k := range colors {
        colorNames = append(colorNames, k)
    }
    sort.Strings(colorNames)
    for k := range colorNames {
        if (k > 0) {
            fmt.Print(",")
        }
        fmt.Print(colorNames[k])
    }
}

func getFlagColor() color.Color {
    *flagSetColor = strings.ToLower(*flagSetColor)
    if (*flagSetColor == "random") {
        rand.Seed(time.Now().UnixNano())
        return color.RGBA{uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 0xFF}
    }
    if (colors[*flagSetColor] != nil) {
        return colors[*flagSetColor]
    }
    validHexCode := regexp.MustCompile(`^#?([a-f0-9]{6})$`)
    if (validHexCode.MatchString(*flagSetColor)) {
        hexStr := validHexCode.FindStringSubmatch(*flagSetColor)
        bytes, err := hex.DecodeString(hexStr[1])
        if (err != nil) {
            fmt.Printf("invalid color code '%s'. use '#rrggbb' instead", hexStr[1])
            return nil
        }
        return color.RGBA{uint8(bytes[0]), uint8(bytes[1]), uint8(bytes[2]), 0xFF}
    }
    return nil
}

func main() {

    kingpin.Version(VERSION)
    kingpin.Parse()

    if (len(os.Args) <= 1) {
        kingpin.Usage()
        os.Exit(0)
    }

    if (flagListColors != nil && *flagListColors) {
        printListColorNames()
        os.Exit(0)
    }

    if (flagSetColor == nil) {
        os.Exit(0)
    }

    for devInfo := range led.Devices() {
        dev, err := devInfo.Open()
        if err != nil {
            fmt.Println(err)
            continue
        }
        defer dev.Close()

        col := getFlagColor()
        if (col != nil) {
            dev.SetColor(col)
            time.Sleep(2 * time.Second)
        }

    }

}
