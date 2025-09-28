package main

import (
	"fmt"
	"shorturl/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
