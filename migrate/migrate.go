package migrate

import (
	"log"

	"github.com/CalvinLe08/todo-app/initializers"
	"github.com/CalvinLe08/todo-app/models"
)

func init() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.TodoItem{})
}



