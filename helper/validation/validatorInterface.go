package validation

import m "trains/models"

type AppDataValidator interface {
	Validate(appData m.AppData) bool
}

type FileValidator interface {
	Validate(rawData []string) (bool, error)
}

var appDataValidationRules = []AppDataValidator{
	StartStationValidator{},
	EndStationValidator{},
	UniqueCoordinatesForStation{},
}

func ValidateWithAllRules(appData m.AppData) bool {
	for _, v := range appDataValidationRules {
		ok := v.Validate(appData)
		if !ok {
			return false
		}
	}

	return true
}
