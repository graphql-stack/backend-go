package resolver

import (
	"context"
	"github.com/graphql-stack/backend-go/middleware"
	"github.com/graphql-stack/backend-go/model"
	"github.com/graphql-stack/backend-go/service"
	"github.com/graphql-stack/backend-go/types"
)

func Register(ctx context.Context, registerInput types.RegisterInput) (model.User, error) {
	u, err := service.Register(registerInput)
	if err != nil {
		return model.User{}, err
	}
	return *u, err
}

func Login(ctx context.Context, loginInput types.LoginInput) (model.Token, error) {
	tk, err := service.Login(loginInput)
	if err != nil {
		return model.Token{}, err
	}
	return *tk, err
}

func PostAuthor(ctx context.Context, obj *model.Post) (*model.User, error) {
	return middleware.GetUserLoader(ctx).Load(obj.AuthorID)
}

func CommentAuthor(ctx context.Context, obj *model.Comment) (*model.User, error) {
	return middleware.GetUserLoader(ctx).Load(obj.AuthorID)
}

func Me(ctx context.Context) (model.User, error) {
	u := middleware.GetCurrentUser(ctx)
	if u == nil {
		return model.User{}, types.ErrInvalidToken
	}
	return *u, nil
}

func Posts(ctx context.Context, limit *int, offset *int) (model.PostsList, error) {
	l := 10
	o := 0
	if limit != nil {
		l = *limit
	}
	if offset != nil {
		o = *offset
	}

	resp, err := service.GetPosts(l, o)
	if err != nil {
		return model.PostsList{}, err
	}

	return *resp, nil
}

func Post(ctx context.Context, id string) (model.Post, error) {
	return service.GetPostByID(id)
}

func CreatePost(ctx context.Context, postInput types.PostInput) (model.Post, error) {
	user := middleware.GetCurrentUser(ctx)
	if user == nil {
		return model.Post{}, types.ErrInvalidToken
	}
	pt, err := service.CreatePost(postInput, user)
	if err != nil {
		return model.Post{}, err
	}

	return *pt, nil
}

func CreateComment(ctx context.Context, commentInput types.CommentInput) (model.Comment, error) {
	user := middleware.GetCurrentUser(ctx)
	if user == nil {
		return model.Comment{}, types.ErrInvalidToken
	}

	cm, err := service.CreateComment(commentInput, user)
	if err != nil {
		return model.Comment{}, err
	}

	return *cm, nil
}
