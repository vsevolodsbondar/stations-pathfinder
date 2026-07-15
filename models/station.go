package models

type Station struct {
	Name        string
	Connections []Station //should be []*Station
	X_axis      int
	Y_axis      int
}

type PointerStation struct {
	Name        string
	Connections []*PointerStation
	X_axis      int
	Y_axis      int
}
