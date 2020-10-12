package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/luhanm/bexs-backend-exam/routes/handler"
	"github.com/luhanm/bexs-backend-exam/routes/scale"
	"github.com/luhanm/bexs-backend-exam/routes/util"
)

func main() {

	if len(os.Args) < 1 {
		log.Fatalf("Arquivo CSV não foi informado.")
	}
	scale.LoadScalesFile(os.Args[1])

	go func() {
		router := mux.NewRouter()
		router.HandleFunc("/route", handler.HttpInsertRoute).Methods("POST")
		router.HandleFunc("/route/{route}", handler.HttpGetRoute).Methods("GET")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()
	fmt.Println("Listening :8080 ")

	var command string
	for {
		fmt.Print("please enter the route: ")
		fmt.Scan(&command)
		command = strings.ToUpper(command)
		if command == "EXIT" {
			break
		} else {

			origin, destination, err := util.SplitRoute(command)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			bestRoute, err := util.FindCheapestWay(origin, destination)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			fmt.Printf("best route: %s > $%d", bestRoute.CompleteWay, bestRoute.Cost)
			fmt.Println("")
		}
	}
}
