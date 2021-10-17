package main

import (
	"fmt"
	"github.com/darencorp/gotodo/sql"
	_ "github.com/darencorp/gotodo/todo"
	"github.com/joho/godotenv"
	"net/http"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	sql.Init()
	godotenv.Load()
	fmt.Println("Listen on 0.0.0.0:80")
	http.ListenAndServe("0.0.0.0:80", nil)
}
