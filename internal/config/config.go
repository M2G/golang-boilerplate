package config

import "github.com/urfave/cli/v3"

type Config struct {
	Port string
}

func Load(cmd *cli.Command) Config {
	return Config{
		Port: cmd.String("port"),
	}
}
