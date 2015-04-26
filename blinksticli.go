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
    flagSetColor = kingpin.Flag("set-color", "Set color for device. The format must be \"#rrggbb\", \"random\", \"off\" or an CSS3 color keyword, e.g. \"green\"").String()
    flagListColors = kingpin.Flag("list-colors", "List all available CSS3 color keywords, as defined in http://www.w3.org/TR/css3-color/").Bool()
    flagListDevices = kingpin.Flag("list-devices", "List all connected devices").Short('l').Bool()
    flagNumber = kingpin.Flag("number", "Select device by number, starts with 0, default: action is applied to all").Short('n').Int()
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

func printListDevices() {
    var i int = 0
    fmt.Printf("%s\t%s\t%s\n", "Number", "Type", "Path")
    for devInfo := range led.Devices() {
        fmt.Printf("%d\t%s\t%s\n", i, devInfo.GetType(), devInfo.GetPath())
        i++
    }
}

func getFlagColor() color.Color {
    col := strings.ToLower(*flagSetColor)
    if (col == "random") {
        rand.Seed(time.Now().UnixNano())
        return color.RGBA{uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 0xFF}
    }
    if (col == "off") {
        return color.Black
    }
    if (colors[col] != nil) {
        return colors[col]
    }
    validHexCode := regexp.MustCompile(`^#?([a-f0-9]{6})$`)
    if (validHexCode.MatchString(col)) {
        hexStr := validHexCode.FindStringSubmatch(col)
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

    if (flagListDevices != nil && *flagListDevices) {
        printListDevices()
        os.Exit(0)
    }

    if (flagSetColor == nil) {
        os.Exit(0)
    }

    var number int = 0
    for devInfo := range led.Devices() {
        if (flagNumber == nil || *flagNumber == number) {
            col := getFlagColor()
            if (col != nil) {
                dev, err := devInfo.Open()
                dev.SetKeepActive(true)
                if err != nil {
                    fmt.Println(err)
                    continue
                }
                dev.SetColor(col)
                defer func() {
                    dev.Close()
                }()
            }
        }
        number++;
    }

}
