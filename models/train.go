package models

import "fmt"

type Train struct {
	ID          int
	Route       Route //index will be representation of the turn and Station will be current position of the train on some station
	CurrStation string
	Turn        int
	Finished    bool
	CanMove     bool
}

func (t Train) PrintTrain() {
	formatted := fmt.Sprintf("Train ID: %d, Route ID: %d, Currently on: %s", t.ID, t.Route.ID, t.CurrStation)
	fmt.Println(formatted)
}

func (t *Train) Move() (bool, string) {
	t.Turn++
	t.CurrStation = t.Route.Route[t.Turn]
	formatted := fmt.Sprintf("T%d-%s ", t.ID, t.CurrStation)

	t.CanMove = false
	if t.CurrStation == t.Route.Route[len(t.Route.Route)-1] {
		t.Finished = true
		return false, formatted
	}

	return true, formatted
}
