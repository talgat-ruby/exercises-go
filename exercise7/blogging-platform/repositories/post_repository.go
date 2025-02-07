// repositories/post_repository.go
package repositories

import (
	"errors"
	"strings"

	"alinurmyrzakhanov/database"
	"alinurmyrzakhanov/models"
)

func CreatePost(post *models.Post) error {
	post.TagsRaw = strings.Join(post.Tags, ",") // массив тегов => строка
	if err := database.DB.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := database.DB.First(&post, id).Error; err != nil {
		return nil, err
	}
	if post.TagsRaw != "" {
		post.Tags = strings.Split(post.TagsRaw, ",")
	} else {
		post.Tags = []string{}
	}

	return &post, nil
}

func UpdatePost(post *models.Post) error {
	post.TagsRaw = strings.Join(post.Tags, ",")
	if err := database.DB.Save(post).Error; err != nil {
		return err
	}
	return nil
}

func DeletePost(id uint) error {
	result := database.DB.Delete(&models.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("post not found")
	}
	return nil
}
func GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := database.DB.Find(&posts).Error; err != nil {
		return nil, err
	}

	for i := range posts {
		if posts[i].TagsRaw != "" {
			posts[i].Tags = strings.Split(posts[i].TagsRaw, ",")
		} else {
			posts[i].Tags = []string{}
		}
	}
	return posts, nil
}

func SearchPostsByTerm(term string) ([]models.Post, error) {
	var posts []models.Post
	likeTerm := "%" + term + "%"

	if err := database.DB.Where(
		"title LIKE ? OR content LIKE ? OR category LIKE ?",
		likeTerm, likeTerm, likeTerm).
		Find(&posts).Error; err != nil {
		return nil, err
	}

	for i := range posts {
		if posts[i].TagsRaw != "" {
			posts[i].Tags = strings.Split(posts[i].TagsRaw, ",")
		} else {
			posts[i].Tags = []string{}
		}
	}

	return posts, nil
}
