package handler

import (
	"encoding/json"
	"github.com/nitram509/blil/blilweb/info"
	"github.com/nitram509/blil/blilweb/led"
	"net/http"
	"strconv"
)

type link struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

type selfLink struct {
	Self link `json:"self"`
}

type ledInfoResource struct {
	led.LedInfo            //embedded
	Links       []selfLink `json:"_links"`
}

type embedded struct {
	Led []ledInfoResource `json:"leds"`
}

type indexResource struct {
	Version  string   `json:"version"`
	Name     string   `json:"name"`
	Embedded embedded `json:"_embedded"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	embeddedleds := embedded{
		Led: make([]ledInfoResource, 0),
	}

	leds := led.DetectAllLeds().Leds
	for l := range leds {
		led := leds[l]
		lir := ledInfoResource{
			Links: make([]selfLink, 0),
		}
		lir.Links = append(lir.Links, selfLink{Self: link{
			Href:  "http://" + r.Host + "/led/" + strconv.Itoa(led.Number),
			Title: "Set or get color on this LED",
		}})
		lir.Number = led.Number
		lir.Path = led.Path
		lir.Type = led.Type
		embeddedleds.Led = append(embeddedleds.Led, lir)
	}

	index := &indexResource{
		Version:  info.VERSION,
		Name:     info.NAME,
		Embedded: embeddedleds,
	}

	if err := json.NewEncoder(w).Encode(index); err != nil {
		panic(err)
	}
}
