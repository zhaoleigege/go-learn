package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "app"
	app.Usage = `使用方法
                 直接执行
				`
	app.Action = func(c *cli.Context) error {
		fmt.Println("程序启动")
		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
