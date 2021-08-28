package model

import "strconv"

type AppConfig struct {
	Key   AppConfigKey `json:"key"`
	Value string       `json:"value"`
}

type AppConfigKey string

func (ac AppConfig) MustGetInt() int {
	value, _ := strconv.Atoi(ac.Value)

	return value
}

func (ac AppConfig) MustGetFloat() float64 {
	value, _ := strconv.ParseFloat(ac.Value, 64)

	return value
}

func (ac AppConfig) MustGetBool() bool {
	value, _ := strconv.ParseBool(ac.Value)

	return value
}
