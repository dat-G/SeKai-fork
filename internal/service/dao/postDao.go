package dao

import (
	"SeKai/internal/logger"
	"SeKai/internal/model"
	"SeKai/internal/model/param"
	"SeKai/internal/util"
	"errors"
	"gorm.io/gorm"
)

func NewPost(newPostParam *param.NewPostParam, userId uint) (uint, error) {
	post := model.Post{
		Title:         newPostParam.Title,
		Author:        userId,
		Content:       newPostParam.Content,
		PostStatus:    newPostParam.PostStatus,
		CommentStatus: newPostParam.CommentStatus,
	}
	if err := util.Datebase.Create(&post).Error; err != nil {
		logger.ServerLogger.Error(err.Error())
		return 0, errors.New("系统错误")
	} else {
		return post.ID, nil
	}
}

func GetPost(id uint) (post model.Post, err error) {
	if err = util.Datebase.Where(&model.Post{Model: gorm.Model{ID: id}}).First(&post).Error; err != nil {
		logger.ServerLogger.Debug(err)
		return model.Post{}, errors.New("找不到文章")
	}
	return post, nil
}
