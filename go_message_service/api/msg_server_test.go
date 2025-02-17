package api

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	createTable := `
	CREATE TABLE Message (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		create_time DATETIME,
		recipient TEXT,
		message_body TEXT,
		is_fetched BOOLEAN
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestRootRetrieve(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc((&MsgServer{}).RootRetrieve)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello, world!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestMessagesCreate(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	server := &MsgServer{Db: db}

	form := "recipient=test&message_body=Hello"
	req, err := http.NewRequest("POST", "/messages/", bytes.NewBufferString(form))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.MessagesCreate)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() == "" {
		t.Errorf("handler returned unexpected body: got empty string")
	}
}

func TestMessagesRetrieve(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	server := &MsgServer{Db: db}

	// Insert a test message
	time := time.Now()
	_, err = addMsg(server, Message{Recipient: "test", CreateTime: &time, MessageBody: strPtr("Hello")})
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/messages/1/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.MessagesRetrieve(w, r, 1)
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func strPtr(s string) *string {
	return &s
}
