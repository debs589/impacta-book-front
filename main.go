package main

import (
	"fmt"
	"impacta-book/src/config"
	"impacta-book/src/cookies"
	"impacta-book/src/router"
	"impacta-book/src/utils"
	"log"
	"net/http"
)

func main() {
	config.LoadEnv()
	cookies.Configure()

	utils.LoadTemplates()

	r := router.Generate()

	fmt.Println("Running application")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
