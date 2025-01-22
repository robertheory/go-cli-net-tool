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

	app.Commands = []cli.Command{
		{
			Name: "ip",
			Usage: "IP lookup",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "google.com",
				},
			},
			Action: ipLookup,
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