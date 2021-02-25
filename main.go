package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {

	// EINVAL Error Code: #define EINVAL          22      /* Invalid argument */
	const EINVAL = 22

	app := cli.NewApp()

	var quantity int
	const missQuantityVal = 0

	app.Flags = []cli.Flag{
		&cli.IntFlag{
			Name:        "quantity",
			Aliases:     []string{"q"},
			Value:       missQuantityVal,
			Usage:       "quantity in Mb the application will occupy in ram",
			Destination: &quantity,
		},
	}

	app.Action = func(c *cli.Context) error {
		if quantity <= 0 {
			return cli.NewExitError("quantity must be more than 0Mb", EINVAL)
		}
		RunMemFill(quantity)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
