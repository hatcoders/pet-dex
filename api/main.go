package main

import (
	"fmt"
	"log"
	"net/http"
	"pet-dex-backend/v2/api/controllers"
	petcontroller "pet-dex-backend/v2/api/controllers/pet"
	"pet-dex-backend/v2/api/routes"
	"pet-dex-backend/v2/infra/config"
	"pet-dex-backend/v2/infra/db"
	"pet-dex-backend/v2/usecase"
)

func main() {
	env, err := config.LoadEnv(".")
	if err != nil {
		panic(err)
	}

	database := config.InitConfigs()
	config.RunMigrations(database)
	dbPetRepo := db.NewPetRepository(database)

	exampleUseCase := usecase.NewExampleUseCase(dbPetRepo)
	petUseCase := usecase.NewPetUseCase(dbPetRepo)

	petUsecase := usecase.NewPetUseCase(dbPetRepo)
	petController := controllers.NewPetController(petUsecase)
	exampleController := petcontroller.NewExampleController(exampleUseCase)
	findPetController := petcontroller.NewFindPetController(petUseCase)
	petController := controllers.NewPetController(petUseCase)

	contrllers := routes.Controllers{
		FindPetController:      findPetController,
		ListUserPetsController: listUserPetsController,
		ExampleController:      exampleController,
		PetController:          petController,
	}
	router := routes.InitializeRouter(contrllers)

	fmt.Printf("running on port %v \n", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}
