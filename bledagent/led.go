package main

type Led struct {
    Number int    `json:"number"`
    Type   string `json:"type"`
    Path   string `json:"path"`
}

type LedResource struct {
    Leds []Led `json:"leds"`
}