package migrate

import (
	"github.com/Triatmono/Golang-FP-CRUD-Gin/initializers"
	"github.com/Triatmono/Golang-FP-CRUD-Gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Post{})
}