package main

import (
	"net/http"

	"github.com/nikkbh/social-app/internal/store"
)

// GetPostFeed godoc
//
//	@Summary		Fetches post feed
//	@Description	Fetches post feed for logged in user
//	@Tags			feed
//	@Accept			json
//	@Produce		json
//
//	@Param			since	query		string	false	"Since"
//	@Param			until	query		string	false	"Until"
//	@Param			limit	query		int		false	"Limit"
//	@Param			offset	query		int		false	"Offset"
//	@Param			sort	query		int		false	"Sort"
//	@Param			tags	query		string	false	"Tags"
//	@Param			Search	query		string	false	"Search"
//
//	@Success		200		{object}	[]store.PostWithMetadata
//	@Failure		500		{object}	error
//	@Failure		404		{object}	error	"User not found"
//	@Security		ApiKeyAuth
//	@Router			/users/feed [get]
func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	fq := store.PaginatedFeedQuery{
		Limit:  20,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	user := getUserFromCtx(r)
	feed, err := app.store.Posts.GetUserFeed(ctx, user.ID, fq)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}
