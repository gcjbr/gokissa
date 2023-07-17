package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/souljacker/gokissa/initializers"
)

func init() {
	initializers.ConnectDB()
	initializers.SyncDb()

}

func main() {}
