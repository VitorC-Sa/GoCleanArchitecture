package main

import (
	"fmt"
	"net/http"

	dHttp "github.com/VitorC-Sa/GolangCleanArchitecture/app/delivery/http"
	"github.com/VitorC-Sa/GolangCleanArchitecture/app/repository"
	"github.com/VitorC-Sa/GolangCleanArchitecture/app/usecase"
	"github.com/gorilla/mux"
	// "github.com/VitorC-Sa/GolangCleanArchitecture/database"
	// "github.com/gorilla/mux"
)

func main() {
	// db := database.NewPostgresDB()
	// defer db.Close()

	// userRepository := repository.NewUserRepositoryBun(db.BunDB)
	userRepository := repository.NewUserRepositoryMock()
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := dHttp.NewUserHandler(userUseCase)

	// HTTP Handlers
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.AddUser).Methods("POST")
	r.HandleFunc("/users/{searchBy}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users", userHandler.DeleteUser).Methods("DELETE")

	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe(":8080", r)
}
