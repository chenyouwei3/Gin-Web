package controller

import (
	mysqlDB "gin-web/init/mysql"
	"gin-web/internal/manage/models"
	"gin-web/internal/manage/types"
	"gin-web/pkg"
	"gin-web/pkg/extendController"
	"gin-web/pkg/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandlerController struct {
	extendController.BaseController
}

// 查询用户列表
func (u *UserHandlerController) GetList() gin.HandlerFunc {
	return func(c *gin.Context) {
		userReq := &types.UserGetListReq{
			Name:      c.Query("name"),
			Email:     c.Query("email"),
			CurrPage:  c.DefaultQuery("currPage", "1"),
			PageSize:  c.DefaultQuery("pageSize", "10"),
			StartTime: c.Query("startTime"),
			EndTime:   c.Query("endTime"),
		}
		skip, limit, err := pkg.GetPage(userReq.CurrPage, userReq.PageSize)
		if err != nil {
			u.SendCustomResponseByBacked(c, "分页失败", "Paging failed", err)
			return
		}
		//DB操作
		userDB := models.User{
			Name:  userReq.Name,
			Email: userReq.Email,
		}
		resDB, total, err := userDB.GetList(skip, limit, userReq.StartTime, userReq.EndTime)
		if err != nil {
			u.SendServerErrorResponse(c, 5130, err)
			return
		}
		u.SendSuccessResponse(c, types.UserGetListResp{
			Users: resDB,
			Total: total,
		})
	}
}

// 根据用户id查询用户
func (u *UserHandlerController) GetRolesByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Query("id")
		if userId == "" {
			u.SendParameterErrorResponse(c, 4001, nil)
			return
		}
		id, err := strconv.Atoi(userId)
		if err != nil {
			u.SendParameterErrorResponse(c, 4003, err)
			return
		}
		var user models.User
		err = mysqlDB.DB.
			Preload("Roles", func(db *gorm.DB) *gorm.DB {
				return db.Select("role.id", "role.name") // 指定表名.字段更稳妥
			}).
			Where("id = ?", id).
			First(&user).Error
		if err != nil {
			u.SendServerErrorResponse(c, 5130, err)
			return
		}
		u.SendSuccessResponse(c, user.Roles)
	}
}

// 新增用户
func (u *UserHandlerController) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqUser types.UserInsertReq
		if err := c.Bind(&reqUser); err != nil {
			u.SendParameterErrorResponse(c, 4002, err)
			return
		}
		//DB操作
		userDB := models.User{
			Name:      reqUser.User.Name,
			Email:     reqUser.User.Email,
			Account:   reqUser.User.Account,
			Password:  reqUser.User.Password,
			AvatarUrl: reqUser.User.AvatarUrl,
		}
		isExist, err := userDB.IsExist()
		if isExist || err != nil {
			u.SendServerErrorResponse(c, 5101, err)
			return
		}
		if err = userDB.SetPassword(reqUser.User.Password); err != nil {
			u.SendServerErrorResponse(c, 5100, err)
			return
		}
		if err = userDB.Insert(reqUser.AddRoles); err != nil {
			u.SendServerErrorResponse(c, 5100, err)
			return
		}
		u.SendSuccessResponse(c, "success")
	}
}

// 删除用户
func (u *UserHandlerController) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		//参数校验
		var userReq types.RoleDeleteReq
		if err := c.ShouldBind(&userReq); err != nil {
			u.SendParameterErrorResponse(c, 4002, err)
			return
		}
		//DB操作
		userDB := models.User{
			ID: userReq.Id,
		}
		if err := userDB.Delete(); err != nil {
			u.SendServerErrorResponse(c, 5110, err)
			return
		}
		u.SendSuccessResponse(c, "success")
	}
}

// 更新用户
func (u *UserHandlerController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userReq types.UserUpdateReq
		err := c.ShouldBind(&userReq)
		if err != nil {
			u.SendParameterErrorResponse(c, 4002, err)
			return
		}
		userDB := models.User{
			ID:        userReq.User.Id,
			Name:      userReq.User.Name,
			Email:     userReq.User.Email,
			Account:   userReq.User.Account,
			Password:  userReq.User.Password,
			AvatarUrl: userReq.User.AvatarUrl,
		}
		if err = userDB.SetPassword(userDB.Password); err != nil {
			u.SendServerErrorResponse(c, 5100, err)
			return
		}
		//DB操作
		if err := userDB.Update(userReq.AddRoles, userReq.DeletedRoles); err != nil {
			u.SendServerErrorResponse(c, 5120, err)
			return
		}
		u.SendSuccessResponse(c, "success")
	}
}

// 用户登录
func (u *UserHandlerController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqUser types.UserLoginReq
		if err := c.ShouldBind(&reqUser); err != nil {
			u.SendParameterErrorResponse(c, 4002, err)
			return
		}
		var tempUser = &models.User{Account: reqUser.Account, Password: reqUser.Password}
		user, err := tempUser.GetOne()
		if err != nil {
			u.SendServerErrorResponse(c, 5101, err)
			return
		}
		//校验密码
		bol := user.CheckPassword(tempUser.Password)
		if !bol {
			u.SendCustomResponseByBacked(c, "密码错误", "Password error", nil)
			return
		}
		accessToken, refreshToken, err := jwt.GenerateToken(user.Name)
		if err != nil {

			u.SendCustomResponseByBacked(c, "生成token失败", "Failed to generate token", err)
			return
		}
		u.SendSuccessResponse(c, types.UserLoginResp{
			User:         user,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		})
	}
}
