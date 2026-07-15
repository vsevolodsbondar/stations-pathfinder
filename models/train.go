package models

type Train struct {
	ID              int
	Route           []Station //index will be representation of the turn and Station will be current position of the train on some station
	StartingStation Station   //may be not necessary
	EndingStation   Station   //may be not necessary
}
