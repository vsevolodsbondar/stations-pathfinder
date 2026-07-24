package cli

import (
	val "trains/helper/validation"
	io "trains/io"
	m "trains/models"
)

func DataConfiguration(conf m.FlagConfig) (m.AppData, []error) {
	appData := m.AppData{}

	stations, errrs := (io.HandleInitialInputFile(conf.NetworkMapPath))
	if errrs != nil {
		return appData, errrs
	}

	var stationsSlice []*m.Station

	for _, v := range stations {
		stationsSlice = append(stationsSlice, v)
	}

	appData.NetworkMap = stationsSlice
	appData.StartingStation = stations[conf.StartingStation]
	appData.EndingStation = stations[conf.EndingStation]
	appData.TrainNumb = conf.TrainNumb

	ok, errs := val.ValidateWithAllRules(appData)
	if !ok {
		return m.AppData{}, errs
	}

	return appData, nil
}
