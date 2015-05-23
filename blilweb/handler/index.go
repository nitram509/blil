package handler

import (
	"encoding/json"
	"github.com/nitram509/blil/blilweb/info"
	"github.com/nitram509/blil/blilweb/led"
	"net/http"
)

type link struct {
	Href  string `json:"href"`
	Title string `json:"title"`
}

type links struct {
	Self link `json:"self"`
}

type embedded struct {
	Led []led.LedInfo `json:"led"`
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
		Led: led.DetectAllLeds().Leds,
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
