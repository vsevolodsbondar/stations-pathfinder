package models

type AppData struct {
	NetworkMap      []Station
	StartingStation Station
	EndingStation   Station
	TrainNumb       int
}

type AppDataPointer struct {
	NetworkMap      []PointerStation
	StartingStation PointerStation
	EndingStation   PointerStation
	TrainNumb       int
}
