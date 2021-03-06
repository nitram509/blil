package main

import (
	"fmt"
	"github.com/boombuler/led"
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
	"sort"
	"strconv"
)

const (
	VERSION           = "0.2.0"
	DEFAULT_NO_NUMBER = -1
	DEFAULT_NO_PATH   = "<<no-path>>"
)

var (
	flagSetColor    = kingpin.Flag("set-color", "Set color for device. The format must be \"#rrggbb\", \"random\", \"off\" or an CSS3 color keyword, e.g. \"green\"").String()
	flagListColors  = kingpin.Flag("list-colors", "List all available CSS3 color keywords, as defined in http://www.w3.org/TR/css3-color/").Bool()
	flagListDevices = kingpin.Flag("list-devices", "List all connected devices").Short('l').Bool()
	flagNumber      = kingpin.Flag("number", "Select device by number, starts with 0, default: action is applied to all").Short('n').Default(strconv.Itoa(DEFAULT_NO_NUMBER)).Int()
	flagPath        = kingpin.Flag("path", "Select device by path (the path is platform dependant)").Short('p').Default(DEFAULT_NO_PATH).String()
)

func printListColorNames() {
	var colorNames []string
	for k := range Colors {
		colorNames = append(colorNames, k)
	}
	sort.Strings(colorNames)
	for k := range colorNames {
		if k > 0 {
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

func allDevicesSelected() bool {
	return DEFAULT_NO_NUMBER == *flagNumber && DEFAULT_NO_PATH == *flagPath
}

func selectedByNumber(number int) bool {
	return *flagNumber == number
}

func selectedByPath(devInfo led.DeviceInfo) bool {
	return *flagPath == devInfo.GetPath()
}

func main() {

	kingpin.Version(VERSION)
	kingpin.Parse()

	if len(os.Args) <= 1 {
		kingpin.Usage()
		os.Exit(0)
	}

	if flagListColors != nil && *flagListColors {
		printListColorNames()
		os.Exit(0)
	}

	if flagListDevices != nil && *flagListDevices {
		printListDevices()
		os.Exit(0)
	}

	if flagSetColor == nil {
		os.Exit(0)
	}

	var number int = 0
	for devInfo := range led.Devices() {
		if allDevicesSelected() || selectedByNumber(number) || selectedByPath(devInfo) {
			col := MapColor(*flagSetColor)
			if col != nil {
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
		}
		number++
	}

}
