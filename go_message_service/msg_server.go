//go:build go1.22

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mikaeljonsson/message-service/api"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	log.Println("Loaded swagger spec")
	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	msgServer := api.NewMsgServer()

	r := http.NewServeMux()
	fmt.Println("Registering handlers")
	// We now register our petStore above as the handler for the interface
	api.HandlerFromMux(msgServer, r)

	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	h := middleware.OapiRequestValidator(swagger)(r)

	// Attach authentication
	// ctx := context.Background()
	//authFuncs := map[string]func(ctx context.Context, token string) (context.Context, error){
	//	"BearerAuth": api.BearerAuth,
	//	"APIKeyAuth": api.APIKeyAuth,
	//}

	// Register middleware
	//r.Use(middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
	//	SecurityHandler: authFuncs,
	//}))

	// TiDB connection details
	dsn := "root@tcp(127.0.0.1:4000)/test" // TODO: Update this with your TiDB connection details
	db, err := sql.Open("mysql", dsn)
	log.Println("Connecting to TiDB")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to TiDB:", err)
	}
	fmt.Println("Connected to TiDB successfully!")
	msgServer.Db = db

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}
	fmt.Printf("Listening on %s\n", s.Addr)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}

/*
CreateTime  *time.Time `json:"create_time,omitempty"`
	Id          *int       `json:"id,omitempty"`
	IsFetched   *bool      `json:"is_fetched,omitempty"`
	MessageBody *string    `json:"message_body,omitempty"`
	Recipient   string     `json:"recipient"`
	Url         *string    `json:"url,omitempty"`

	CREATE TABLE Message (
		id SERIAL PRIMARY KEY,
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		recipient VARCHAR(200) NOT NULL,
		message_body TEXT,
		is_fetched BOOLEAN NOT NULL DEFAULT FALSE
	);
*/

/*
// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album

    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }
    return albums, nil
}
*/
