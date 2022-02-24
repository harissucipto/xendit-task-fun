package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/harissucipto/xendit-task/db/sqlc"
)

// create comment
type CreateCommentRequestURI struct {
	OrgName string `uri:"org-name" binding:"required"`
}

type CreateCommentRequestJSON struct {
	Comment string `json:"comment" binding:"required"`
}

func (server *Server) createComment(ctx *gin.Context) {
	var requestURI CreateCommentRequestURI
	var requestJSON CreateCommentRequestJSON

	if err := ctx.ShouldBindUri(&requestURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&requestJSON); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check is valid org-name
	if _, err := server.CheckIsValidOrg(requestURI.OrgName); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	arg := db.CreateCommentParams{
		OrgName: requestURI.OrgName,
		Content: requestJSON.Comment,
	}

	comment, err := server.store.CreateComment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ "comment": comment.Content })
}



// delete comment
type ListCommentsRequest struct {
	OrgName string `uri:"org-name" binding:"required"`
}

type ListCommentsResponse []struct {
	Id 				int64  `json:"id"`
	Comment 	string 	`json:"comment"`
	Org				string	`json:"org"`
	CreatedAt	time.Time	`json:"created_at"`
}


func newCommentResponse(comments []db.Comment) ListCommentsResponse {
	 newComments := make(ListCommentsResponse, len(comments))
	 for i, comment := range comments {
			newComments[i].Id = comment.ID
			newComments[i].Comment = comment.Content
			newComments[i].Org = comment.OrgName
			newComments[i].CreatedAt = comment.CreatedAt
	}
	return newComments
}


func (server *Server) listComments(ctx *gin.Context) {
	var requestURI ListCommentsRequest

	if err := ctx.ShouldBindUri(&requestURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check is valid org-name
	if _, err := server.CheckIsValidOrg(requestURI.OrgName); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	comments, err := server.store.ListComments(ctx, requestURI.OrgName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	respComments := newCommentResponse(comments)
	ctx.JSON(http.StatusOK, respComments)
}


// delete comment
type DeleteCommentRequestURI struct {
	OrgName string `uri:"org-name" binding:"required"`
}

func (server *Server) deleteComment(ctx *gin.Context) {
	var requestURI DeleteCommentRequestURI

	if err := ctx.ShouldBindUri(&requestURI); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// check is valid org-name
	if _, err := server.CheckIsValidOrg(requestURI.OrgName); err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	err := server.store.DeleteComment(ctx, requestURI.OrgName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// sendt ok deleted all comments
	ctx.JSON(http.StatusOK, gin.H{ "success": "deleted all comments"})
}