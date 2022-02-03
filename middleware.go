package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	"codigos.ufsc.br/g.manoel/pi_das_2021_2/config"
	"codigos.ufsc.br/g.manoel/pi_das_2021_2/device"
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
		newConf := new(config.Internal)
		if err := json.Unmarshal(c.Body(), &newConf); err != nil {
			fmt.Println(err.Error())
			return err
		}
		conf = *newConf
		err := config.Write(c.Body())
		if err != nil {
			fmt.Println(err.Error())
		}
		return err
	})
	app.Get("/api/device", func(c *fiber.Ctx) error {
		result := device.GetAll()
		err := c.Status(200).JSON(result)
		return err
	})
	app.Get("*", func(c *fiber.Ctx) error {
		err := c.SendFile("web/public/index.html")
		return err
	})
	if !*devP{
		startAsyncServices()
	}
	_ = device.GetOrCreate("MF", "000")
	fmt.Println(app.Listen(":" + strconv.Itoa(conf.HTTP.Port)))
}
