package main

import (
	"fmt"
	"repository-pattern/repository"
	"repository-pattern/service"
)

func main() {
	// Initialize repository
	userRepo := repository.NewUserRepositoryMemory()

	// Initialize service with repository
	userService := service.NewUserService(userRepo)

	// Use the service
	user, err := userService.RegisterUser("John", "Doe", "john@example.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Registered user: %+v\n", user)

	// Retrieve the user
	retrievedUser, err := userService.GetUser(user.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Retrieved user: %+v\n", retrievedUser)

	// Get all users
	allUsers, err := userService.GetAllUsers()
	if err != nil {
		panic(err)
	}

	fmt.Println("All users:")
	for _, u := range allUsers {
		fmt.Printf("- %+v\n", u)
	}
}
