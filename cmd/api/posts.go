package main

import (
	"github.com/go-chi/chi/v5"
	"go-learning/internal/store"
	"net/http"
	"strconv"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		UserID:  3,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postId")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := r.Context()

	post, err := app.store.Posts.GetById(ctx, id)

	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
