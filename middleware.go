package main

import (
	"flag"
	"fmt"
	"strconv"

	"codigos.ufsc.br/g.manoel/pi_das_2021_2/internal"
	"codigos.ufsc.br/g.manoel/pi_das_2021_2/protocol/mqtt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	path := flag.String("path", ".", "path to configuration file")
	flag.Parse()
	conf, err := internal.LoadConfig(*path)
	if err != nil {
		panic("Failed to read configuration file")
	}
	app := fiber.New()
	app.Static("/", "./web/public")
	go mqtt.Run(&conf)
	fmt.Println(app.Listen(":" + strconv.Itoa(conf.ServerPort)))
}
