package main

import (
	"fmt"
	"path"

	appconfig "github.com/reinhardlinardi/atm-report/internal/config/app"
)

func main() {
	dir := "config"
	file := "config.yaml"

	config, err := appconfig.Parse(path.Join(dir, file))
	if err != nil {
		return
	}

	fmt.Println(config.DB)
	fmt.Println(config.FileCron)
}
