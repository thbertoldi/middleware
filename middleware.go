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

func startAsyncServices() {
	go mqtt.Run()
}

func main() {
	verP := flag.Bool("version", false, "show version")
	devP := flag.Bool("dev", false, "run dev mode")
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
	app.Get("/api/config", func(c *fiber.Ctx) error {
		err := c.Status(200).JSON(conf)
		return err
	})
	app.Post("/api/config", func(c *fiber.Ctx) error {
		if err := c.BodyParser(&conf); err != nil {
			return err
		}
		fmt.Println(conf)
		return nil
	})
	app.Get("*", func(c *fiber.Ctx) error {
		err := c.SendFile("web/public/index.html")
		return err
	})
	if !*devP{
		startAsyncServices()
	}
	fmt.Println(app.Listen(":" + strconv.Itoa(conf.HTTP.Port)))
}
