package main

import (
	"fmt"
	"log"
	"os"
	c "trains/cli"
	io "trains/io"
)

func main() {
	conf, err := c.FlagHandling()
	if err != nil {
		log.Fatal(err)
	}
	stations, err := (io.HandleInitialInputFile(conf.NetworkMapPath))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, station := range stations {
		fmt.Print(station.Name + ": ")
		for _, con := range station.Connections {
			fmt.Println(con.Name)
		}
		println()
	}
	fmt.Println(conf.NetworkMapPath)
	fmt.Println(conf.StartingStation)
	fmt.Println(conf.EndingStation)
	fmt.Println(conf.TrainNumb)
}
