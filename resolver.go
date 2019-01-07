package backend_go

import (
	"context"
	"github.com/graphql-stack/backend-go/middleware"
	"github.com/graphql-stack/backend-go/service"

	"github.com/graphql-stack/backend-go/model"
	"github.com/graphql-stack/backend-go/types"
)

type Resolver struct{}

func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) Post(ctx context.Context, obj *model.Comment) (*model.Post, error) {
	panic("not implemented")
}
func (r *commentResolver) Author(ctx context.Context, obj *model.Comment) (*model.User, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Register(ctx context.Context, registerInput types.RegisterInput) (model.User, error) {
	u, err := service.Register(registerInput)
	if err != nil {
		return model.User{}, err
	}
	return *u, err
}
func (r *mutationResolver) Login(ctx context.Context, loginInput types.LoginInput) (model.Token, error) {
	tk, err := service.Login(loginInput)
	if err != nil {
		return model.Token{}, err
	}
	return *tk, err
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *model.Post) (*model.User, error) {
	return middleware.GetUserLoader(ctx).Load(obj.AuthorID)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (model.User, error) {
	u := middleware.GetCurrentUser(ctx)
	if u == nil {
		return model.User{}, types.ErrInvalidToken
	}
	return *u, nil
}

func (r *queryResolver) Posts(ctx context.Context, pageParams *model.PageParms) (model.PostsList, error) {
	limit := 10
	offset := 0

	if pageParams != nil && pageParams.Limit != nil {
		limit = *pageParams.Limit
	}
	if pageParams != nil && pageParams.Offset != nil {
		offset = *pageParams.Offset
	}
	resp, err := service.GetPosts(limit, offset)
	if err != nil {
		return model.PostsList{}, err
	}

	return *resp, nil
}
func (r *queryResolver) Post(ctx context.Context, id string) (model.Post, error) {
	panic("not implemented")
}
