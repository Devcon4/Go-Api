package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Devcon4/Go-Api/modules/chatmodule"
	"github.com/Devcon4/Go-Api/modules/personmodule"
	"github.com/Devcon4/Go-Api/tools/framework"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func buildDataSource() string {
	host := framework.GetEnvOrDefault("DB_HOST", "localhost")
	port := framework.GetEnvOrDefault("DB_PORT", "4261")
	user := framework.GetEnvOrDefault("DB_USER", "dev")
	dbname := framework.GetEnvOrDefault("DB_DBNAME", "chat")
	password := framework.GetEnvOrDefault("DB_PASSWORD", "FinchDev")

	return fmt.Sprint("host=", host, " port=", port, " user=", user, " dbname=", dbname, " password=", password, " sslmode=disable")
}

func buildServerHost() string {
	host := framework.GetEnvOrDefault("SERVICE_HOST", "localhost")
	port := framework.GetEnvOrDefault("SERVICE_PORT", "80")

	return fmt.Sprint(host, ":", port)
}

func main() {
	router := framework.NewRouter(&framework.RouterConfig{
		Prefix:  "/api",
		Version: 1,
	})

	db := framework.NewDBContext(&framework.GORMConfig{
		DriverName: "postgres",
		DataSource: buildDataSource(),
	})

	db.AutoMigrate(&chatmodule.Chat{}, &personmodule.Person{})

	chatService := chatmodule.NewChatService(db, router)
	chatHandler := chatmodule.NewChatHandler(router, chatService)
	chatHandler.Register()

	personService := personmodule.NewPersonService(db, router)
	personHandler := personmodule.NewPersonHandler(router, personService)
	personHandler.Register()

	router.Use(framework.LogRequestMiddleware)

	server := &http.Server{
		Handler:      router,
		Addr:         buildServerHost(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("🚀 Server running on ", server.Addr, "!")
	log.Fatal(server.ListenAndServe())
	defer db.Close()
}
