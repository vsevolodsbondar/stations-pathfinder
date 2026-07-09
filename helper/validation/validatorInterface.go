package validation

import m "trains/models"

type AppDataValidator interface {
	Validate(appData m.AppData) bool
}

type FileValidator interface {
	Validate(rawData []string) (bool, error)
}
