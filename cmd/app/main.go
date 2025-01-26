package main

import (
	"fmt"
	"github.com/iwerqfx/url-shortener/internal/config"
)

func main() {
	cfg := config.Get()

	fmt.Println(cfg)
}
