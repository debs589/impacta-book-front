package main

import (
	"fmt"
	"impacta-book/config"
	"impacta-book/src/router"
	"impacta-book/src/utils"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()

	utils.LoadTemplates()

	r := router.Generate()

	fmt.Println("Running application")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
