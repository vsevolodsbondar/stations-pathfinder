package main

import (
	"fmt"
	"log"
	c "trains/cli"
	"trains/io"
)

func main() {
	conf, err := c.FlagHandling()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(io.HandleInitialInputFile("stations.map"))
	fmt.Println(conf.NetworkMapPath)
	fmt.Println(conf.StartingStation)
	fmt.Println(conf.EndingStation)
	fmt.Println(conf.TrainNumb)
}
