package main

import (
	"fmt"
	"time"
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Title string `toml:"title"`
	Owner struct {
		Name         string    `toml:"name"`
		Organization string    `toml:"organization"`
		Bio          string    `toml:"bio"`
		Dob          time.Time `toml:"dob"`
	} `toml:"owner"`
	Database struct {
		Server        string `toml:"server"`
		Ports         []int  `toml:"ports"`
		ConnectionMax int    `toml:"connection_max"`
		Enabled       bool   `toml:"enabled"`
	} `toml:"database"`
	Servers struct {
		Alpha struct {
			IP string `toml:"ip"`
			Dc string `toml:"dc"`
		} `toml:"alpha"`
		Beta struct {
			IP string `toml:"ip"`
			Dc string `toml:"dc"`
		} `toml:"beta"`
	} `toml:"servers"`
	Clients struct {
		//data = [ ["gamma", "delta"], [1, 2] ] 出错
		Data  [][]interface{} `toml:"data"`
		Hosts []string        `toml:"hosts"`
	} `toml:"clients"`
}

func main() {
	var config tomlConfig
	if _,err := toml.DecodeFile("example.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Owner: %s (%s, %s), Born: %s\n",
		config.Owner.Name, config.Owner.Organization, config.Owner.Bio, config.Owner.Organization)
	fmt.Printf("Database: %s %v (Max conn. %d), Enabled? %v\n",
		config.Database.Server, config.Database.Ports, config.Database.ConnectionMax,
		config.Database.Enabled)
	fmt.Printf("Client data: %v\n", config.Clients.Data)
	fmt.Printf("Client hosts: %v\n", config.Clients.Hosts)
}
