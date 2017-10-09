package main

import (
	"flag"
	"fmt"
)

var (
	version    = flag.Int("version", 1, "Current version of this software")
	configFile string
)

func init() {
	flag.StringVar(&configFile, "config", "config.json", "Default config file path")
}

func main() {
	flag.Parse()
	fmt.Printf("Version: %d\n", *version)
	fmt.Printf("Config: %s\n", configFile)
}
