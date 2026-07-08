package models

type Train struct {
	Route           []Station
	StartingStation Station
	EndingStation   Station
}
