//go:build go1.22

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=cfg.yaml ../../petstore-expanded.yaml

package api

import (
	"net/http"
	"sync"
)

type MsgServer struct {
	Lock sync.Mutex
}

// Make sure we conform to ServerInterface

var _ ServerInterface = (*MsgServer)(nil)

func NewMsgServer() *MsgServer {
	return &MsgServer{
		Lock: sync.Mutex{},
	}
}

// (GET /)
func (p *MsgServer) RootRetrieve(w http.ResponseWriter, r *http.Request) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (GET /messages/)
func (p *MsgServer) MessagesList(w http.ResponseWriter, r *http.Request, params MessagesListParams) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages/)
func (p *MsgServer) MessagesCreate(w http.ResponseWriter, r *http.Request) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages/bulk-delete)
func (p *MsgServer) MessagesBulkDeleteCreate(w http.ResponseWriter, r *http.Request) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages/bulk-delete{format})
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

// (POST /messages/fetch-new{format})
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
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
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

// (DELETE /messages/{id}{format})
func (p *MsgServer) MessagesFormattedDestroy(w http.ResponseWriter, r *http.Request, id int, format MessagesFormattedDestroyParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (GET /messages/{id}{format})
func (p *MsgServer) MessagesFormattedRetrieve(w http.ResponseWriter, r *http.Request, id int, format MessagesFormattedRetrieveParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (PATCH /messages/{id}{format})
func (p *MsgServer) MessagesFormattedPartialUpdate(w http.ResponseWriter, r *http.Request, id int, format MessagesFormattedPartialUpdateParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (PUT /messages/{id}{format})
func (p *MsgServer) MessagesFormattedUpdate(w http.ResponseWriter, r *http.Request, id int, format MessagesFormattedUpdateParamsFormat) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (GET /messages{format})
func (p *MsgServer) MessagesFormattedList(w http.ResponseWriter, r *http.Request, format MessagesFormattedListParamsFormat, params MessagesFormattedListParams) {
	_ = p
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, world!"))
}

// (POST /messages{format})
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
