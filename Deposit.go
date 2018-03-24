package main

import (
	"fmt"
	"github.com/howeyc/gopass"
	"eth-deposit/accounts"
	"os"
	"eth-deposit/logger"
	"runtime"
	"github.com/julienschmidt/httprouter"
	"eth-deposit/getethaddress"
	"log"
	"net/http"
	"eth-deposit/config"
)

const (
	VERSION = "0.01"
)

func init() {
	config.Load()
}

func main() {

	fmt.Printf("Database password: ")
	pass, _ := gopass.GetPasswd()

	if len(pass) > 1 {
		accounts.KEY = string(pass)
	}

	logger.Log.Printf("Server v%s pid=%d started with processes: %d", VERSION, os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))

	router := httprouter.New()
	router.GET("/account/:id", getethaddress.GetAddress)

	log.Fatal(http.ListenAndServe(":9010", router))
}
