package main

import (
	"flag"
	"fmt"
	"strconv"

	"codigos.ufsc.br/g.manoel/pi_das_2021_2/config"
	"codigos.ufsc.br/g.manoel/pi_das_2021_2/protocol/mqtt"
	"github.com/gofiber/fiber/v2"
)

var version string // set by the compiler!

func main() {
	verP := flag.Bool("version", false, "show version")
	flag.Parse()
	if *verP {
		fmt.Println(version)
		return
	}
	conf, err := config.LoadInternal()
	if err != nil {
		panic("Failed to read configuration file")
	}
	app := fiber.New()
	app.Static("/", "./web/public")
	app.Get("/api", func(c *fiber.Ctx) error {
		err := c.SendString("hi there")
		return err
	})
	app.Get("*", func(c *fiber.Ctx) error {
		err := c.SendFile("web/public/index.html")
		return err
	})
	go mqtt.Run()
	fmt.Println(app.Listen(":" + strconv.Itoa(conf.HTTP.Port)))
}
