package service

import (
	"douyin/db"
	"douyin/response"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

type CommentResponse struct {
	response.Response
	ID         int64         `json:"id"`
	User       response.User `json:"user"`
	Content    string        `json:"content"`
	CreateDate string        `json:"create_date"`
}

func (cs *CommentService) PublishComment(video_id int64, user_id int64, comment_text string) (*CommentResponse, error) {
	commentDao, err := db.SaveComment(video_id, user_id, comment_text)
	if err != nil {
		return nil, err
	}
	user := db.GetUserById(user_id)
	var timeFormat = "01-02"
	return &CommentResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "评论成功",
		},
		User:       user,
		Content:    comment_text,
		CreateDate: commentDao.CreateDate.Format(timeFormat),
		ID:         commentDao.ID,
	}, nil
}

func (cs *CommentService) DelComment(video_id int64, comment_id int64) (*response.Response, error) {
	err := db.DelComment(video_id, comment_id)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		StatusCode: 200,
		StatusMsg:  "删除评论成功",
	}, nil
}
