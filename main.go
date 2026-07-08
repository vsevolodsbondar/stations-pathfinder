package main

import (
	"fmt"
	"log"
	c "trains/cli"
)

func main() {
	conf, err := c.FlagHandling()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf.NetworkMapPath)
	fmt.Println(conf.StartingStation)
	fmt.Println(conf.EndingStation)
	fmt.Println(conf.TrainNumb)
}
