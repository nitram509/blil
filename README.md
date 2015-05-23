# BliL - Blinking Light

1. ```blil``` A command line client, written in GO, works on Windows and Mac OS X
2. ```blilweb``` A web server, which makes the devices available via REST API (HTTP + JSON)

#### License

The MIT License (MIT)


## blil - command line

Usage:

```bash
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


## blilweb


### get all available devices

GET http://localhost:8080/

```json
{
    "version": "0.0.1",
    "name": "BliL - Blinking Light",
    "_embedded": {
        "leds": [
            {
                "number": 0,
                "type": "BlinkStick",
                "path": "USB_20a0_41e5_14100000",
                "_links": [
                    {
                        "self": {
                            "href": "http://localhost:8080/led/0",
                            "title": "Set or get color on this LED"
                        }
                    }
                ]
            }
        ]
    }
}
```

### set a color

POST http://localhost:8080/led/0/green

```json
{
    "number": 0,
    "color": "008000"
}
```

## Supported devices

* [blink(1)](http://blink1.thingm.com/)
* [LinkM / BlinkM](http://thingm.com/products/linkm/)
* [BlinkStick](http://www.blinkstick.com/)
* [Blync](http://www.blynclight.com/)
* [Busylight UC](http://www.busylight.com/busylight-uc.html)
* [Busylight Lync](http://www.busylight.com/busylight-lync.html)
* [DreamCheeky USBMailNotifier](http://www.dreamcheeky.com/webmail-notifier)

_powered by_ https://github.com/boombuler/led

