package main

import (
	"github.com/Kong/go-pdk"
	"github.com/google/uuid"
)

type Config struct {
}

func New() interface{} {
	return &Config{}
}

// Custom Plugin Logic Space
func (conf Config) Access(kong *pdk.PDK) {
	correlationId, err := kong.Request.GetHeader("x-correlation-id")
	if correlationId == "" || err != nil {
		newId := uuid.New()
		kong.ServiceRequest.SetHeader("x-correlation-id", newId.String())
	}
}
