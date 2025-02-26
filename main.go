package main

import (
	"LMS/services"
	"LMS/controllers"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Library Managment system")
	service := services.NewLibraryService()
	controller := controllers.NewLibraryServiceControl(service)
	controller.Library_controller()
	
}