package service

import (
	"errors"
	"github.com/graphql-stack/backend-go/db"
	"github.com/graphql-stack/backend-go/model"
	"github.com/graphql-stack/backend-go/utils"
)

var ErrExpired = errors.New("TOKEN_EXPIRED")

func RefreshToken(token model.Token) {
	db.ORM.Model(&token).Update("token", utils.GenerateToken())
}

func GetOrCreateToken(user model.User) (*model.Token, error) {
	var token model.Token
	err := db.ORM.Where(model.Token{UserID: user.ID}).Attrs(model.Token{Token: utils.GenerateToken()}).FirstOrCreate(&token).Error
	if token.IsExpired() {
		RefreshToken(token)
	}
	return &token, err
}

func GetUserByToken(tk string) (*model.User, error) {
	var token model.Token
	err := db.ORM.Preload("User").Where("token = ?", tk).First(&token).Error
	if err != nil {
		return nil, err
	}
	if token.IsExpired() {
		return nil, ErrExpired
	}
	return token.User, err
}
