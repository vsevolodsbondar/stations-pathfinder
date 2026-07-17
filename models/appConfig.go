package models

type AppData struct {
	NetworkMap      []*Station
	StartingStation *Station
	EndingStation   *Station
	TrainNumb       int
}

func (a AppData) FindStationByName(name string) *Station {
	for _, v := range a.NetworkMap {
		if v.Name == name {
			return v
		}
	}

	return nil
}
