package main

import (
	"fmt"
	chttp "go_cyfra_homework_2/internal/service/http"
	"go_cyfra_homework_2/internal/service/store"
	"net/http"
)

func main() {
	store.New()
	server := chttp.New()

	fmt.Println("Server started at http://localhost:3000")
	err := http.ListenAndServe(":3000", server.BuildRoutes())

	if err != nil {
		panic(err)
	}

}
