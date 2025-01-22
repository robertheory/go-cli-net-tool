package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate function to generate the CLI app
func Generate() *cli.App {
	
	app := &cli.App{
		Name: "cli-net-tool",
		Usage: "A simple CLI network tool for IP and NS lookup",
	}

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "amazom.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "IP lookup",
			Flags: flags,
			Action: ipLookup,
		},
		{
			Name: "ns",
			Usage: "NS lookup",
			Flags: flags,
			Action: nsLookup,
		},
	}

	return app
}

func ipLookup(c *cli.Context) {
	host := c.String("host")
	
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func nsLookup(c *cli.Context) {
	host := c.String("host")

	nss, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ns := range nss {
		fmt.Println(ns.Host)
	}
}