package models

type APIResponse struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	Timeseries []Timeseries `json:"timeseries"`
}

type Timeseries struct {
	Time string `json:"time"`
	Data Data   `json:"data"`
}

type Data struct {
	Instant Instant `json:"instant"`
}

type Instant struct {
	Details Details `json:"details"`
}

type Details struct {
	AirTemperature float64 `json:"air_temperature"`
}
