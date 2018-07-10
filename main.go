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
        GenerateSampleEnvfile(filepath)
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
    {
      Name: "add",
      Usage: "add new key-value line in envfile",
      Action: func(c *cli.Context) error {
        filepath := c.String("envfile")
        if len(c.Args()) == 2 {
          key := c.Args().Get(0)
          value := c.Args().Get(1)
          AddNewEnvToFile(filepath, key, value)
          AddNewEnvToFile(filepath + ".sample", key, value)
          return nil
        } else {
          return nil
        }
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

  app.Run(os.Args)
}

func GenerateSampleEnvfile(filepath string) {
  texts := ReadEnvFile(filepath)
  WriteEnvFile(filepath + ".sample", BuildValueRemovedLines(texts))
  fmt.Printf("generated sample file: %v => %v\n", filepath, filepath + ".sample")
}

func AddNewEnvToFile(filepath string, key string, value string) {
  AppendLineToFile(filepath, key + "=" + value)
}

