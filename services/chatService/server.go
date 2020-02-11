package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Devcon4/finch/services/chatService/framework"
	"github.com/Devcon4/finch/services/chatService/modules/chatmodule"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router := framework.NewRouter(&framework.RouterConfig{
		Prefix:  "/api",
		Version: 1,
	})

	db := framework.NewDBContext(&framework.GORMConfig{
		DriverName: "postgres",
		DataSource: "host=localhost port=4261 user=chat dbname=chat password=FinchDev sslmode=disable",
	})
	// , &personmodule.Person{}
	db.AutoMigrate(&chatmodule.Chat{})

	chatService := chatmodule.NewChatService(db, router)
	chatHandler := chatmodule.NewChatHandler(router, chatService)
	chatHandler.Register()

	// personService := personmodule.NewPersonService(db, router)
	// personHandler := personmodule.NewPersonHandler(router, personService)
	// personHandler.Register()

	router.Use(framework.LogRequestMiddleware)

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("🚀 Server running on ", server.Addr, "!")
	log.Fatal(server.ListenAndServe())
	defer db.Close()
}
