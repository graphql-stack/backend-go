package backend_go

import (
	"context"

	"github.com/graphql-stack/backend-go/model"
)

type Resolver struct{}

func (r *Resolver) Comment() CommentResolver {
	return &commentResolver{r}
}
func (r *Resolver) Post() PostResolver {
	return &postResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type commentResolver struct{ *Resolver }

func (r *commentResolver) Post(ctx context.Context, obj *model.Comment) (model.Post, error) {
	panic("not implemented")
}
func (r *commentResolver) Author(ctx context.Context, obj *model.Comment) (model.User, error) {
	panic("not implemented")
}

type postResolver struct{ *Resolver }

func (r *postResolver) Author(ctx context.Context, obj *model.Post) (model.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Me(ctx context.Context) (model.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Posts(ctx context.Context) (model.PostsList, error) {
	panic("not implemented")
}
func (r *queryResolver) Post(ctx context.Context, id string) (model.Post, error) {
	panic("not implemented")
}
