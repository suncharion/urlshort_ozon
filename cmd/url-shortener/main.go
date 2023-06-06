package main

import (
	"flag"
	"log"
	"net/http"
	"ozon/internal/service"
	"ozon/internal/storage"
	storagememory "ozon/internal/storageMemory"
	storagepostgres "ozon/internal/storagePostgres"
)

var storageType = flag.String("storage", "inmemory", "stroage type - inmemory (default) or postgres")

func main() {
	log.Println("Starting url shortener service")

	flag.Parse()

	if storageType == nil {
		log.Println("No storage type specified, default is inmemory")
		*storageType = "inmemory"
	}

	var storage storage.Storage
	var err error

	if *storageType == "inmemory" {
		storage, err = storagememory.NewStorageMemory()
		if err != nil {
			log.Fatal(err)
		}
	} else if *storageType == "postgres" {
		storage, err = storagepostgres.NewStoragePostgres()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("Unkown storage type")
	}

	log.Println("Storage inited")

	serv, err := service.NewService(storage)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Starting to serve http")

	http.HandleFunc("/", serv.HttpHandlerGet)
	http.HandleFunc("/save", serv.HttpHandlerPost)
	http.ListenAndServe(":80", nil)

}
