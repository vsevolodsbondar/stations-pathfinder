package schedulingutils

import (
	m "trains/models"
)

func TrainMaker(startingStation string, trains int) []m.Train {
	trainsSlice := make([]m.Train, 0, trains)

	for i := 1; i <= trains; i++ {
		train := m.Train{
			ID:          i,
			CurrStation: startingStation,
			Turn:        0,
			Finished:    false,
			CanMove:     false,
		}

		trainsSlice = append(trainsSlice, train)
	}

	return trainsSlice
}
