package framework

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var schema = `
	CREATE TABLE IF NOT EXISTS chat (
		id int,
		message VARCHAR
	);

	CREATE TABLE IF NOT EXISTS person (
		id int,
		firstName VARCHAR,
		lastName VARCHAR
	)
`

// SQLConfig : Configuration object for sql
type SQLConfig struct {
	DriverName     string
	DataSourceName string
}

// NewDBContext : Creates a db connection
func NewDBContext(c *SQLConfig) *sqlx.DB {
	db, err := sqlx.Connect(c.DriverName, c.DataSourceName)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	db.MustExec(schema)

	return db
}

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
