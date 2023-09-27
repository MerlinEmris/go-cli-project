package main

import (
	"fmt"

	"github.com/MerlinEmris/go-cli-project/internal/network"
)

func main() {
	fmt.Println("Entrypoint")
	network.Ping("127.0.0.1")
}
