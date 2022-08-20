package main

import (
	"log"
	"os"

	"github.com/bodagovsky/metro-project/database"
)

func main() {
	logger := log.New(os.Stdout, "prefix", log.LstdFlags)
	_ = database.New()
	logger.Printf("successfully initiated storage")
}