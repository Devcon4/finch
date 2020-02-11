package framework

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GetEnvOrDefault ...
func GetEnvOrDefault(key string, fallback string) string {
	v, f := os.LookupEnv(key)

	if !f {
		v = fallback
	}
	return v
}

// GORMConfig : Config to create a gorm connection
type GORMConfig struct {
	DriverName string
	DataSource string
}

// NewDBContext : Create a gorm db connection
func NewDBContext(c *GORMConfig) *gorm.DB {
	db, err := gorm.Open(c.DriverName, c.DataSource)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("dataSource: ", c.DataSource)
		panic("failed to connect database")
	}

	// defer db.Close()

	return db
}

// sqlx setup
// // SQLConfig : Configuration object for sql
// type SQLConfig struct {
// 	DriverName     string
// 	DataSourceName string
// }

// // NewDBContext : Creates a db connection
// func NewDBContext(c *SQLConfig) *sqlx.DB {
// 	db, err := sqlx.Connect(c.DriverName, c.DataSourceName)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil
// 	}

// 	db.MustExec(schema)

// 	return db
// }

// RouterConfig : Config object for the router
type RouterConfig struct {
	Prefix  string
	Version int
}

// NewRouter : Creates a mux router
func NewRouter(c *RouterConfig) *mux.Router {
	return mux.NewRouter().PathPrefix(c.Prefix + "/V" + fmt.Sprint(c.Version)).Subrouter()
}

// JSONHandler : Handler helper to return json response
func JSONHandler(w http.ResponseWriter, o interface{}, e error) {
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(o)
	// fmt.Println("Json: ", string(json))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// LogRequestMiddleware : Logs each request
func LogRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
