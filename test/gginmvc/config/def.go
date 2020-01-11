package config

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Title string `toml:"title"`
	Db    struct {
		Name         string    `toml:"name"`
		Organization string    `toml:"organization"`
		Bio          string    `toml:"bio"`
		Dob          time.Time `toml:"dob"`
	} `toml:"db"`
	Rou struct {
		Roufunc []string `toml:"roufunc"`
	} `toml:"rou"`
}

func Tconf() TomlConfig {
	var config TomlConfig
	TomlFilePath := `./tomls/def.toml`
	if _, err := toml.DecodeFile(TomlFilePath, &config); err != nil {
		fmt.Println(err)
	}
	return config
}
