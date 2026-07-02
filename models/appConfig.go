package models

type AppConfig struct {
	NetworkMap      []Station
	StartingStation Station
	EndingStation   Station
	TrainNumb       int
}
