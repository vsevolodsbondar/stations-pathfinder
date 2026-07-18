package validation

import m "trains/models"

type AppDataValidator interface {
	Validate(appData m.AppData) (bool, []error)
}

type FileValidator interface {
	Validate(rawData []string) (bool, []error)
}

var appDataValidationRules = []AppDataValidator{
	StartStationValidator{},
	EndStationValidator{},
	UniqueCoordinatesForStation{},
}

func ValidateWithAllRules(appData m.AppData) (bool, []error) {
	valid := true
	errs := []error{}
	for _, v := range appDataValidationRules {
		ok, err := v.Validate(appData)
		if !ok {
			valid = false
			errs = append(errs, err...)
		}
	}

	return valid, errs
}
