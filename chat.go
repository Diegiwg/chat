package main

import "github.com/Diegiwg/cli"

func main() {
	app := cli.NewApp()

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
