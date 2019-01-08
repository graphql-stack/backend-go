package main

import (
	"fmt"
	"github.com/graphql-stack/backend-go/db"
	"github.com/graphql-stack/backend-go/model"
)

func main() {
	for i := 0; i < 3; i++ {
		comment := model.CommentFactory.MustCreate().(*model.Comment)
		fmt.Printf("%+v\n", comment)
		db.ORM.Create(comment)
	}
}
