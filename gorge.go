package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gorge"
	app.Usage = "README generation for new projects"
	app.Version = "0.0.1"
	app.Action = func(c *cli.Context) {
		println("Generating a README for: ")
		get_info()
	}
	app.Run(os.Args)
}

func get_info() {
	dir_name, _ := os.Getwd()
	var i int
	ext := ""

	println(dir_name)
	println("Select a format for your README: ")
	println("  (1) Markdown\n  (2) Text")
	_, unrecognizable := fmt.Scanf("%d", &i)
	if unrecognizable != nil {
		panic(unrecognizable)
	}
	if i == 1 {
		ext = "markdown"
	} else if i == 2 {
		ext = "txt"
	} else {
		println("Not recognized. Quitting.")
		os.Exit(1)
	}
	fmt.Printf("Creating README.%s\n", ext)
	/* fi, err := os.Create("README.md") */
	/* if err != nil { */
	/* 	panic(err) */
	/* } */

}
