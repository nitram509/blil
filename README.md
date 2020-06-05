# BliL - Blinking Light

A command line client, which can control a blinkstick and/or compatible devices,
written in GO, works on Windows and Mac OS X

#### Build status

[![Build Status](https://travis-ci.org/nitram509/blil.svg?branch=master)](https://travis-ci.org/nitram509/blil)


#### License

The MIT License (MIT)


## blil - command line

Usage:

```
usage: blil [<flags>]

Flags:
  --help              Show help.
  --set-color=SET-COLOR  
                      Set color for device. The format must be "#rrggbb", "random", "off" or an CSS3 color keyword, e.g. "green"
  --list-colors       List all available CSS3 color keywords, as defined in http://www.w3.org/TR/css3-color/
  -l, --list-devices  List all connected devices
  -n, --number=-1     Select device by number, starts with 0, default: action is applied to all
  -p, --path="<<no-path>>"  
                      Select device by path (the path is platform dependant)
  --version           Show application version.
```


## installation

There are no pre-build binaries available at the moment, so you have to download the sources and build from scratch.
This is quite easy, since it just requires a Go runtime, to be downloaded from https://golang.org/dl/

### requirements

* Go 1.14 or newer

```shell script
go get github.com/nitram509/blil/cmd/blil
go install github.com/nitram509/blil/cmd/blil
~/go/bin/blil --version
## 0.2.0
```

You might want to copy the binary ```~/go/bin/blil``` into ```/usr/local/bin```
or add the folder to your PATH variable.


## Supported devices

* [blink(1)](http://blink1.thingm.com/)
* [LinkM / BlinkM](http://thingm.com/products/linkm/)
* [BlinkStick](http://www.blinkstick.com/)
* [Blync](http://www.blynclight.com/)
* [Busylight UC](http://www.busylight.com/busylight-uc.html)
* [Busylight Lync](http://www.busylight.com/busylight-lync.html)
* [DreamCheeky USBMailNotifier](http://www.dreamcheeky.com/webmail-notifier)

_powered by_ https://github.com/boombuler/led

