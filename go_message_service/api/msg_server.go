//go:build go1.22

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=cfg.yaml ../openapi.yaml

package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type MsgServer struct {
	Lock sync.Mutex
	Db   *sql.DB
}

// Make sure we conform to ServerInterface

var _ ServerInterface = (*MsgServer)(nil)

func NewMsgServer() *MsgServer {
	return &MsgServer{
		Lock: sync.Mutex{},
		Db:   nil,
	}
}

// (GET /)
func (p *MsgServer) RootRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("RootRetrieve")
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
	log.Println("RootRetrieve end")
}

// (GET /messages/)
func (p *MsgServer) MessagesList(w http.ResponseWriter, r *http.Request, params MessagesListParams) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages/)
func (p *MsgServer) MessagesCreate(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("message_body")
	time := time.Now()
	id, err := addMsg(p, Message{Recipient: r.FormValue("recipient"),
		CreateTime:  &time,
		MessageBody: &body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("Error adding message: %v", err)))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(strconv.FormatInt(id, 10)))
}

// (POST /messages/bulk-delete)
func (p *MsgServer) MessagesBulkDeleteCreate(w http.ResponseWriter, r *http.Request) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages/bulk-delete/{format})
func (p *MsgServer) MessagesBulkDeleteFormattedCreate(w http.ResponseWriter, r *http.Request, format MessagesBulkDeleteFormattedCreateParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages/fetch-new)
func (p *MsgServer) MessagesFetchNewCreate(w http.ResponseWriter, r *http.Request) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages/fetch-new/{format})
func (p *MsgServer) MessagesFetchNewFormattedCreate(w http.ResponseWriter, r *http.Request, format MessagesFetchNewFormattedCreateParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (DELETE /messages/{id}/)
func (p *MsgServer) MessagesDestroy(w http.ResponseWriter, r *http.Request, id int) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (GET /messages/{id}/)
func (p *MsgServer) MessagesRetrieve(w http.ResponseWriter, r *http.Request, id int) {
	msg, err := msgByID(p, int64(id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(fmt.Sprintf("Message %d not found", id)))
		return
	}
	w.WriteHeader(http.StatusOK)
	if msg.MessageBody != nil {
		_, _ = w.Write([]byte(*msg.MessageBody))
	} else { // can this happen?
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Message body not found"))
	}
}

// (PATCH /messages/{id}/)
func (p *MsgServer) MessagesPartialUpdate(w http.ResponseWriter, r *http.Request, id int) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (PUT /messages/{id}/)
func (p *MsgServer) MessagesUpdate(w http.ResponseWriter, r *http.Request, id int) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (DELETE /messages/{id}/{format})
func (p *MsgServer) MessagesFormattedDestroy(w http.ResponseWriter, r *http.Request, id int, format MessagesFormattedDestroyParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (GET /messages/{id}/{format})
func (p *MsgServer) MessagesFormattedRetrieve(w http.ResponseWriter, r *http.Request, id int, format MessagesFormattedRetrieveParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (PATCH /messages/{id}/{format})
func (p *MsgServer) MessagesFormattedPartialUpdate(w http.ResponseWriter, r *http.Request, id int, format MessagesFormattedPartialUpdateParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (PUT /messages/{id}/{format})
func (p *MsgServer) MessagesFormattedUpdate(w http.ResponseWriter, r *http.Request, id int, format MessagesFormattedUpdateParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (GET /messages/{format})
func (p *MsgServer) MessagesFormattedList(w http.ResponseWriter, r *http.Request, format MessagesFormattedListParamsFormat, params MessagesFormattedListParams) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages/{format})
func (p *MsgServer) MessagesFormattedCreate(w http.ResponseWriter, r *http.Request, format MessagesFormattedCreateParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (GET /{format})
func (p *MsgServer) RootFormattedRetrieve(w http.ResponseWriter, r *http.Request, format RootFormattedRetrieveParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// Authenticator for Bearer Token
func BearerAuth(ctx context.Context, token string) (context.Context, error) {
	//if token != "valid_token" {
	//	return ctx, errors.New("invalid token")
	//}
	return ctx, nil
}

func APIKeyAuth(ctx context.Context, key string) (context.Context, error) {
	//if key != "my-secret-key" {
	//	return ctx, errors.New("invalid API key")
	//}
	return ctx, nil
}

// addMsg adds the specified message to the database,
// returning the message ID of the new entry
func addMsg(p *MsgServer, msg Message) (int64, error) {
	result, err := p.Db.Exec("INSERT INTO Message (create_time, recipient, message_body) VALUES (?, ?, ?)",
		msg.CreateTime, msg.Recipient, msg.MessageBody)
	if err != nil {
		return 0, fmt.Errorf("addMsg: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addMsg: %v", err)
	}
	return id, nil
}

// msgByID queries for the message with the specified ID.
func msgByID(p *MsgServer, id int64) (Message, error) {
	// An album to hold data from the returned row.
	var msg Message

	row := p.Db.QueryRow("SELECT * FROM Message WHERE id = ?", id)
	if err := row.Scan(&msg.Id, &msg.CreateTime, &msg.Recipient, &msg.MessageBody, &msg.IsFetched); err != nil {
		if err == sql.ErrNoRows {
			return msg, fmt.Errorf("msgById %d: no such message", id)
		}
		return msg, fmt.Errorf("msgById %d: %v", id, err)
	}
	return msg, nil
}

// sendPetStoreError wraps sending of an error in the Error format, and
// handling the failure to marshal that.
/*
func sendMsgServerError(w http.ResponseWriter, code int, message string) {
	petErr := Error{
		Code:    int32(code),
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(petErr)
}
*/

// FindPets implements all the handlers in the ServerInterface
/*
func (p *MsgServer) FindMessages(w http.ResponseWriter, r *http.Request) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	var result []Message
    messages := []Message{}
	for _, message := range messages {
		// TODO: Implement filtering
		result = append(result, message)
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(result)
}

func (p *MsgServer) AddMessage(w http.ResponseWriter, r *http.Request) {
	// We expect a NewPet object in the request body.
	var newPet NewPet
	if err := json.NewDecoder(r.Body).Decode(&newPet); err != nil {
		sendPetStoreError(w, http.StatusBadRequest, "Invalid format for NewPet")
		return
	}

	// We now have a pet, let's add it to our "database".

	// We're always asynchronous, so lock unsafe operations below
	p.Lock.Lock()
	defer p.Lock.Unlock()

	// We handle pets, not NewPets, which have an additional ID field
	var pet Message
	pet.Name = newPet.Name
	pet.Tag = newPet.Tag
	pet.Id = p.NextId
	p.NextId++

	// Insert into map
	p.Pets[pet.Id] = pet

	// Now, we have to return the NewPet
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(pet)
}

func (p *PetStore) FindPetByID(w http.ResponseWriter, r *http.Request, id int64) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	pet, found := p.Pets[id]
	if !found {
		sendPetStoreError(w, http.StatusNotFound, fmt.Sprintf("Could not find pet with ID %d", id))
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(pet)
}

func (p *PetStore) DeletePet(w http.ResponseWriter, r *http.Request, id int64) {
	p.Lock.Lock()
	defer p.Lock.Unlock()

	_, found := p.Pets[id]
	if !found {
		sendPetStoreError(w, http.StatusNotFound, fmt.Sprintf("Could not find pet with ID %d", id))
		return
	}
	delete(p.Pets, id)

	w.WriteHeader(http.StatusNoContent)
}
*/
