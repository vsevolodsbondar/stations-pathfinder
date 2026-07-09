package models

type Station struct {
	Name        string
	Connections []Station
	X_axis      int
	Y_axis      int
}
