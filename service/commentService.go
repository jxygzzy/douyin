package service

import (
	"douyin/db"
	"douyin/response"
	"sort"
	"sync"
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
	user, _ := db.GetUserById(user_id, user_id)
	var timeFormat = "01-02"
	return &CommentResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "评论成功",
		},
		User:       *user,
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

type CommentListResponse struct {
	response.Response
	CommentList *[]response.Comment `json:"comment_list"`
}

func (cs *CommentService) CommentList(video_id int64, user_id int64) (*CommentListResponse, error) {
	commentDaos, err := db.CommentList(video_id)
	if err != nil {
		return nil, err
	}
	wg := sync.WaitGroup{}
	commentList := make([]response.Comment, 0, len(*commentDaos))
	for i, n := 0, len(*commentDaos); i < n; i++ {
		wg.Add(1)
		go func(commentDao db.CommentDao) {
			defer wg.Done()
			comment := response.Comment{}
			comment.Id = commentDao.ID
			comment.Content = commentDao.Content
			comment.CreateDate = commentDao.CreateDate.Format("01-02")
			user, _ := db.GetUserById(user_id, commentDao.UserId)
			comment.User = *user
			commentList = append(commentList, comment)
		}((*commentDaos)[i])
	}
	wg.Wait()
	sort.Slice(commentList, func(i, j int) bool {
		return commentList[i].Id > commentList[j].Id
	})
	return &CommentListResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "获取评论列表成功",
		},
		CommentList: &commentList,
	}, nil
}
