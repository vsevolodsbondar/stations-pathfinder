package cli

import (
	"fmt"
	val "trains/helper/validation"
	io "trains/io"
	m "trains/models"
)

func DataConfiguration(conf m.FlagConfig) (m.AppData, error) {
	appData := m.AppData{}

	stations, err := (io.HandleInitialInputFile(conf.NetworkMapPath))
	if err != nil {
		return appData, err
	}

	var stationsSlice []*m.Station

	for _, v := range stations {
		stationsSlice = append(stationsSlice, v)
	}

	appData.NetworkMap = stationsSlice
	appData.StartingStation = stations[conf.StartingStation]
	appData.EndingStation = stations[conf.EndingStation]
	appData.TrainNumb = conf.TrainNumb

	ok := val.ValidateWithAllRules(appData)
	if !ok {
		return m.AppData{}, fmt.Errorf("Some issue with input data.")
	}

	return appData, nil
}
