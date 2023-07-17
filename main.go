package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/souljacker/gokissa/initializers"
)

func init() {
	initializers.ConnectDB()
	initializers.SyncDb()
	initializers.SetApp()
	initializers.SetRoutes()

}

func main() {

	initializers.App.Listen(":" + os.Getenv("PORT"))
}
