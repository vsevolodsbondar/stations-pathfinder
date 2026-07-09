package validation

import (
	"testing"
	m "trains/models"
)

func mockAppData() m.AppData {
	station1 := m.Station{
		Name:   "waterloo",
		X_axis: 1,
		Y_axis: 1,
	}

	station2 := m.Station{
		Name:   "euston",
		X_axis: 1,
		Y_axis: 2,
	}

	station3 := m.Station{
		Name:   "victoria",
		X_axis: 2,
		Y_axis: 2,
	}

	station1.Connections = []m.Station{station2, station3}
	station2.Connections = []m.Station{station1, station3}
	station3.Connections = []m.Station{station1, station2}

	return m.AppData{
		NetworkMap: []m.Station{
			station1,
			station2,
			station3,
		},
		StartingStation: station1,
		EndingStation:   station3,
		TrainNumb:       2,
	}
}

func TestShouldFailWhenStartingStationNotPresent(t *testing.T) {

}
