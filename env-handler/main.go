package main

import (
  "os"
  "fmt"
  "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Name = "env-handler"
  app.Usage = "This app handles env file."
  app.Version = "0.0.1"


  app.Commands = []cli.Command {
    {
      Name: "gen-sample",
      Usage: "generate sample file from existing env file",
      Action: func(c *cli.Context) error {
        filepath := c.String("envfile")
        fmt.Println(filepath)
        texts := ReadEnvFile(filepath)
        WriteEnvFile(filepath + ".sample", texts)
        fmt.Println("generated sample file")
        return nil
      },
      Flags: []cli.Flag {
        cli.StringFlag {
          Name: "envfile, e",
          Value: "./.env",
          Usage: "path to envfile",
        },
      },
    },
  }

  app.Action = func(context *cli.Context) error {
    if len(context.Args()) == 0 {
      fmt.Println("USAGE: env-handler [-option]")
    }

    return nil
  }

  app.Run(os.Args)
}

