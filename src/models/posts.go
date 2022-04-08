package models

import (
	"errors"
	"strings"
	"time"
)

type Posts struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (post *Posts) Prepare() error {
	post.format()

	if err := post.validate(); err != nil {
		return err
	}

	return nil
}

func (post *Posts) validate() error {
	if post.Title == "" {
		return errors.New("title is mandatory and cannot be empty")
	}

	if post.Content == "" {
		return errors.New("content is mandatory and cannot be empty")
	}

	return nil
}

func (post *Posts) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
