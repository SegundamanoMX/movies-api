package utils

import (
	"github.com/magiconair/properties"
	"github.mpi-internal.com/guillermo-dlsg/movies-api/pkg/constants"
)

var PropertiesFiles = []string{"./pkg/resources/config.properties"}

var Props, _ = properties.LoadFiles(PropertiesFiles, properties.UTF8, true)

type ResourceManager struct {
}

func (res ResourceManager) GetProperty(propertyKey string) string {
	var propertyValue string = constants.EMPTY
	var existProperty bool

	propertyValue, existProperty = Props.Get(propertyKey)

	if !existProperty {
		return constants.EMPTY
	} else {
		return propertyValue
	}
}
