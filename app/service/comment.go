package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	global2 "loopy-manager/app/global"
	"loopy-manager/app/model"
	"loopy-manager/pkg/utils"
	"strconv"
	"time"
)

func AddComment(commentForm model.AddCommentForm) utils.Response {
	if err := global2.CommentTableMaster.Transaction(func(tx *gorm.DB) error {
		comment := model.Comment{
			CreateTime: time.Now(),
			MomentId:   commentForm.MomentId,
			UserId:     commentForm.UserId,
			Content:    commentForm.Content,
		}
		if commentForm.ParentId != 0 {
			comment.ParentId = &commentForm.ParentId
		}
		fmt.Println(comment)
		if err := tx.Debug().Create(&comment).Error; err != nil {
			return fmt.Errorf("创建评论失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", "1")
}

func AddMoment(addMomentForm model.AddMomentForm) utils.Response {
	if err := global2.MomentTableMaster.Transaction(func(tx *gorm.DB) error {
		moment := model.Moment{
			Id:         global2.MomentSnowFlake.Generate().Int64(),
			CreateTime: time.Now(),
			UserId:     addMomentForm.UserId,
			Content:    addMomentForm.Content,
		}
		if err := tx.Debug().Create(&moment).Error; err != nil {
			return fmt.Errorf("创建动态失败:%w", err)
		}
		return nil
	}); err != nil {
		return utils.ErrorMess("事务失败", err.Error())
	}
	return utils.SuccessMess("插入成功", "1")
}

func GetComment(momentId string) utils.Response {
	var commentTrees []model.CommentTree
	mid, _ := strconv.Atoi(momentId)
	commentTrees = GetMomentComment(mid)
	return utils.SuccessMess("成功", commentTrees)
}

func GetMomentComment(mid int) []model.CommentTree {
	var commentTrees []model.CommentTree
	var comments []model.Comment
	tx := global2.CommentTableSlave.ChooseSlave()
	if err := tx.Debug().Where("moment_id = ? AND parent_id IS NULL", mid).Order("id desc").Find(&comments).Error; err != nil {
		logrus.Error("查询失败", err)
		return nil
	}
	txUser := global2.UserTableSlave.ChooseSlave()
	for _, comment := range comments {
		var user model.User
		commentId := int(comment.Id)
		uid := comment.UserId
		txUser.Where("id = ?", uid).First(&user)
		commentTree := model.CommentTree{
			CommentId: commentId,
			Content:   comment.Content,
			Author:    gin.H{"name": user.Name},
			CreatedAt: comment.CreateTime.Format("2006-01-02 15:04"),
			Children:  []*model.CommentTree{},
		}
		GetMomentCommentChild(commentId, &commentTree)
		commentTrees = append(commentTrees, commentTree)
	}
	return commentTrees
}

func GetMomentCommentChild(pid int, commentTree *model.CommentTree) {
	var comments []model.Comment
	tx := global2.CommentTableSlave.ChooseSlave()
	tx.Where("parent_id = ?", pid).Find(&comments)
	// 查询二级及以下的多级评论
	txUser := global2.UserTableSlave.ChooseSlave()
	for i, _ := range comments {
		var user model.User
		cid := int(comments[i].Id)
		uid := comments[i].UserId
		txUser.Where("id = ?", uid).First(&user)
		child := model.CommentTree{
			CommentId: cid,
			Content:   comments[i].Content,
			Author:    gin.H{"user": user.Name},
			CreatedAt: comments[i].CreateTime.Format("2006-01-02 15:04"),
			Children:  []*model.CommentTree{},
		}
		commentTree.Children = append(commentTree.Children, &child)
		GetMomentCommentChild(cid, &child)
	}
}
