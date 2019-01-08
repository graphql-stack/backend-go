package service

import (
	"github.com/graphql-stack/backend-go/db"
	"github.com/graphql-stack/backend-go/model"
	"github.com/graphql-stack/backend-go/types"
	"github.com/graphql-stack/backend-go/validator"
	"github.com/zcong1993/libgo/mysql"
)

func GetPosts(limit, offset int) (*model.PostsList, error) {
	var posts []model.Post
	totalCount, err := mysql.PaginationQuery(db.ORM.Model(new(model.Post)), limit, offset, &posts)
	if err != nil {
		return nil, err
	}

	return &model.PostsList{TotalCount: totalCount, Posts: posts}, nil
}

func GetPostByID(id string) (model.Post, error) {
	var post model.Post
	err := db.ORM.Where("id = ?", id).First(&post).Error
	return post, err
}

func GetComments(post *model.Post) ([]*model.Comment, error) {
	var comments []*model.Comment
	err := db.ORM.Model(post).Related(&comments).Error
	return comments, err
}

func CreatePost(input types.PostInput, user *model.User) (*model.Post, error) {
	err := validator.Validate.Struct(&input)
	if err != nil {
		return nil, types.ErrValidateError
	}

	post := &model.Post{
		Title:   input.Title,
		Content: input.Content,
		Author:  user,
	}

	err = db.ORM.Create(post).Error

	if err != nil {
		return nil, err
	}

	return post, nil
}

func CreateComment(input types.CommentInput, user *model.User) (*model.Comment, error) {
	err := validator.Validate.Struct(&input)
	if err != nil {
		return nil, types.ErrValidateError
	}

	post, err := GetPostByID(input.PostID)
	if err != nil {
		return nil, err
	}

	comment := &model.Comment{
		Content: input.Content,
		Post:    &post,
		Author:  user,
	}

	err = db.ORM.Create(comment).Error

	if err != nil {
		return nil, err
	}

	return comment, nil
}
