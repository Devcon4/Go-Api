package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Devcon4/Go-Api/modules/chatmodule"
	"github.com/Devcon4/Go-Api/modules/personmodule"
	"github.com/Devcon4/Go-Api/tools/framework"

	_ "github.com/lib/pq"
)

func main() {
	router := framework.NewRouter(&framework.RouterConfig{
		Prefix:  "/api",
		Version: 1,
	})

	db := framework.NewDBContext(&framework.SQLConfig{
		DriverName:     "postgres",
		DataSourceName: "user=dev dbname=Go password=GoDev sslmode=disable",
	})

	chatService := chatmodule.NewChatService(db, router)
	chatHandler := chatmodule.NewChatHandler(router, chatService)
	chatHandler.Register()

	personService := personmodule.NewPersonService(db, router)
	personHandler := personmodule.NewPersonHandler(router, personService)
	personHandler.Register()

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:4255",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("🚀 Server running on ", server.Addr, "!")
	log.Fatal(server.ListenAndServe())
}
