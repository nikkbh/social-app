package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound     = errors.New("resource not found")
	ErrDeleteFailed = errors.New("delete resource failed")
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetById(context.Context, int64) (*Post, error)
		DeletePostById(context.Context, int64) error
		Update(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]Comment, error)
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts:    &PostStore{db},
		Users:    &UserStore{db},
		Comments: &CommentStore{db},
	}
}
