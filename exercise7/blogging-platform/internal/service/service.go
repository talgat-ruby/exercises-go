package service

import (
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/db"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

var Id int

func ServiceGetNumberPost(id int) (error, models.Blog) {
	var OneBlog models.Blog

	return nil, OneBlog
}

func CreationPost(body models.Blog) error {

	if err := db.NewPostCreator(body); err != nil {
		return err
	}
	return nil
}
func ServiceGetPosts() ([]models.Blog, error) {
	return db.DBGetPosts()
}
