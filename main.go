package main

import (
	"fmt"
	"net/http"

	"github.com/kimyoutora/x-games/protos"
)

func main() {
	fmt.Println("X-Games server starting...")

	// TODO: remove me after testing
	sport := &protos.Sport{
		Id:        1,
		Name:      "Tennis",
		DayOfWeek: protos.Sport_TUESDAY,
	}
	fmt.Println("Sport: ", sport)

	http.ListenAndServe(":8080", http.FileServer(http.Dir("frontend/public")))
}
