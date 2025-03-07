package main

import (
	"fmt"
	"net/http"

	"github.com/zoldyzdk/golink/internal/router"
)	

func main() {
	r := router.Routes()

	fmt.Println("Starting API at port 3000")
	http.ListenAndServe(":3000", r)
}