package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	InitUserRuoter(r)
	InitRelationRouter(r)
	InitVideoRouter(r)
	InitCommentRouter(r)
	InitFavotiteRouter(r)
}
