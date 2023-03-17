package config

import "time"

const WindowWidth = 468
const WindowHeight = 468
const OpenWeatherMapApiKey = "ea42e6f5e54a1766d53c3ab30605f7ea"
const MaxClouds = 32
const MaxRaindrops = 512
const HTTPTimeout = time.Minute * 5
const WindSpeedModifier = 0.1
const UpdateWeatherInterval = -5 * time.Second
const DefaultWeatherId = 800
