package validation

import (
	"testing"
	m "trains/models"
)

// run all tests in the project: go test ./...

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

var startVal = StartStationValidator{}

func TestShouldPassWithValidDataStart(t *testing.T) {
	data := mockAppData()

	ok := startVal.Validate(data)
	if !ok {
		t.Errorf("Should be good with valid data.")
	}
}

func TestShouldFailWhenStartingStationNotPresent(t *testing.T) {
	data := mockAppData()
	data.StartingStation = m.Station{}

	ok := startVal.Validate(data)
	if ok {
		t.Errorf("Should fail with empty starting station.")
	}
}

func TestShouldFailWhenStartingStationNotInMap(t *testing.T) {
	data := mockAppData()
	data.NetworkMap = make([]m.Station, 0)

	ok := startVal.Validate(data)
	if ok {
		t.Errorf("Should fail with empty stations map.")
	}
}

var endVal = EndStationValidator{}

func TestShouldPassWithValidDataEnd(t *testing.T) {
	data := mockAppData()

	ok := endVal.Validate(data)
	if !ok {
		t.Errorf("Should be good with valid data.")
	}
}

func TestShouldFailWhenEndStationNotPresent(t *testing.T) {
	data := mockAppData()
	data.EndingStation = m.Station{}

	ok := endVal.Validate(data)
	if ok {
		t.Errorf("Should fail with empty starting station.")
	}
}

func TestShouldFailWhenEndStationNotInMap(t *testing.T) {
	data := mockAppData()
	data.NetworkMap = make([]m.Station, 0)

	ok := endVal.Validate(data)
	if ok {
		t.Errorf("Should fail with empty stations map.")
	}
}

var coordVal = UniqueCoordinatesForStation{}

func TestShouldPassWhenCoordinatesWasUnique(t *testing.T) {
	data := mockAppData()

	ok := coordVal.Validate(data)
	if !ok {
		t.Errorf("Should be good with valid data.")
	}
}

func TestShouldFailWhenCoordinatesWasntUnique(t *testing.T) {
	data := mockAppData()

	station3 := m.Station{
		Name:   "victoria",
		X_axis: 2,
		Y_axis: 2,
	}

	data.NetworkMap[0] = station3

	ok := coordVal.Validate(data)
	if ok {
		t.Errorf("Should fail as 1st and 3rd station share same coordinates.")
	}
}

var dupVal = DuplicateConnectionsSliceValidator{}

func TestShouldPassWhenConnectionsWasUnique(t *testing.T) {
	var mockValidSliceData = []string{
		"waterloo-victoria",
		"waterloo-euston",
		"st_pancras-euston",
		"victoria-st_pancras",
	}
	ok, _ := dupVal.Validate(mockValidSliceData)
	if !ok {
		t.Errorf("Should be good with valid data.")
	}
}

func TestShouldFailWhenConnectionsWasntUnique(t *testing.T) {
	var mockInvalidSliceData = []string{
		"waterloo-victoria",
		"waterloo-euston",
		"st_pancras-euston",
		"victoria-st_pancras",
		"victoria-waterloo",
	}
	ok, err := dupVal.Validate(mockInvalidSliceData)
	if ok {
		t.Errorf("Should fail with invalid data.")
	}
	if err == nil {
		t.Errorf("Should give info about connection that wasn't unique.")
	}
}

var stLineVal StationLineValidator

func TestShouldPassWhenStationLineIsCorrect(t *testing.T) {
	ok := stLineVal.Validate("waterloo,2,5")
	if !ok {
		t.Errorf("Should pass with valid station line")
	}
	ok = stLineVal.Validate("waterloo,2,5 #international")
	if !ok {
		t.Errorf("Should pass if line contains comment")
	}
}

func TestShouldFailWhenStationLineNameIsIncorrect(t *testing.T) {
	mockInvalidLines := []string{
		"w-terloo,2,5",
		",2,4",
		"waterloo,,2,5",
	}
	for _, line := range mockInvalidLines {
		ok := stLineVal.Validate(line)
		if ok {
			t.Errorf("Should fail with invalid station name (%s)", line)
		}
	}
}

func TestShouldFailWhenStationLineXAxisIsIncorrect(t *testing.T) {
	mockInvalidLines := []string{
		"waterloo,,5",
		"waterloo,0.5,4",
		"waterloo,abc,5",
	}
	for _, line := range mockInvalidLines {
		ok := stLineVal.Validate(line)
		if ok {
			t.Errorf("Should fail with invalid station name (%s)", line)
		}
	}
}

func TestShouldFailWhenStationLineYAxisIsIncorrect(t *testing.T) {
	mockInvalidLines := []string{
		"waterloo,2,",
		"waterloo,2,0.4",
		"waterloo,2,abc",
	}
	for _, line := range mockInvalidLines {
		ok := stLineVal.Validate(line)
		if ok {
			t.Errorf("Should fail with invalid station name (%s)", line)
		}
	}
}

var conLineVal ConnectionLineValidator

func TestShouldPassWhenConnectionLineIsCorrect(t *testing.T) {
	ok := conLineVal.Validate("waterloo-london")
	if !ok {
		t.Errorf("Should pass with valid station line")
	}
	ok = conLineVal.Validate("waterloo-london #international")
	if !ok {
		t.Errorf("Should pass if line contains comment")
	}
}
