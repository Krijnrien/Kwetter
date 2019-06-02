package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/ksuid"
	"github.com/krijnrien/Kwetter/db"
	"github.com/krijnrien/Kwetter/event"
	"github.com/krijnrien/Kwetter/schema"
	"github.com/krijnrien/Kwetter/util"
)

func createKweetHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json:"id"`
	}

	ctx := r.Context()

	// Read parameters
	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	// Create Kweet
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create Kweet")
		return
	}
	Kweet := schema.Kweet{
		ID:        id.String(),
		Body:      body,
		CreatedAt: createdAt,
	}
	if err := db.InsertKweet(ctx, Kweet); err != nil {
		log.Println(err)
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create Kweet")
		return
	}

	// Publish event
	if err := event.PublishKweetCreated(Kweet); err != nil {
		log.Println(err)
	}

	// Return new Kweet
	util.ResponseOk(w, response{ID: Kweet.ID})
}
