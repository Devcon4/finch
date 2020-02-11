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

func buildDataSource() string {
	host := framework.GetEnvOrDefault("DB_HOST", "localhost")
	port := framework.GetEnvOrDefault("DB_PORT", "4261")
	user := framework.GetEnvOrDefault("DB_USER", "dev")
	dbname := framework.GetEnvOrDefault("DB_DBNAME", "chat")
	password := framework.GetEnvOrDefault("DB_PASSWORD", "FinchDev")

	return fmt.Sprint("host=", host, " port=", port, " user=", user, " dbname=", dbname, " password=", password, " sslmode=false")
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
