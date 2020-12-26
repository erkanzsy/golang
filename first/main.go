package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start work. Go http://localhost:8080/test ")

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"hello guys")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
