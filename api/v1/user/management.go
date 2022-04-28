package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"path"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"server/utils"
	"strconv"
	"time"
)

type ManagementApi struct {
}

// Login 登录
func (m *ManagementApi) Login(ctx *gin.Context) {
	var userInfo request.LoginUserInfo
	if err := ctx.ShouldBind(&userInfo); err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
		return
	}
	if userInfo.Way <= 0 {
		global.FailWithMsg(ctx, "登录方式错误")
		return
	}
	// 登录
	token, err := userLogic.Login(ctx, &userInfo)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "登录成功", map[string]string{"token": "Bearer " + token})
}

// RegisterUser 注册用户
func (m *ManagementApi) RegisterUser(ctx *gin.Context) {
	var userInfo request.RegisterUserInfo
	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		global.GLOBAL_LOG.Error("提交的信息有误", zap.Error(err))
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
		return
	}
	// 注册用户
	err = userLogic.RegisterUser(ctx, &userInfo)
	if err != nil {
		global.GLOBAL_LOG.Error("注册失败", zap.Error(err))
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithMsg(ctx, "注册成功")
}

// RegisterUsers 上传文件批量注册用户
func (m *ManagementApi) RegisterUsers(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		global.GLOBAL_LOG.Error("解析表单错误", zap.Error(err))
		global.FailWithMsg(ctx, "解析表单错误")
		return
	}
	uuid := time.Now().Unix()
	fileName := string(uuid) + "_" + file.Filename
	dst := path.Join("./upload", fileName)
	if ok, _ := utils.PathExists("upload"); !ok { // 判断是否有upload文件夹
		_ = os.Mkdir("upload", os.ModePerm)
	}
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		global.FailWithMsg(ctx, file.Filename+"上传失败")
	}
	err = userLogic.Register2(ctx, dst)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithMsg(ctx, "录入成功")
}

// GetUserInfo 获取用户相信信息
func (m *ManagementApi) GetUserInfo(ctx *gin.Context) {
	claims, err := utils.GetClaims(ctx)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithDetails(ctx, "", response.UserInfo{
		UserName:   claims.UserName,
		UserId:     claims.UserId,
		Email:      claims.Email,
		Phone:      claims.Phone,
		Role:       global.RoleMap[claims.Role],
		Department: claims.Department,
	})
}

// SetRole 给用户设置角色
func (m *ManagementApi) SetRole(ctx *gin.Context) {
	var info *request.UserRole
	err := ctx.ShouldBind(&info)
	if err != nil {
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
		return
	}
	if info.UserId == "" {
		global.FailWithMsg(ctx, "用户不存在")
		return
	}
	err = userLogic.SetRole(ctx, info)
	if err != nil {
		global.FailWithMsg(ctx, err.Error())
		return
	}
	global.OkWithMsg(ctx, "用户角色设置成功, 请重新登录之后生效")
}

// GetAllUser 获取所有用户信息
func (m *ManagementApi) GetAllUser(ctx *gin.Context) {
	pageNum, _ := ctx.GetQuery("pageNum")
	pageSize, _ := ctx.GetQuery("pageSize")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	var users []*response.UserVO
	var total int64
	if err := global.GLOBAL_DB.Table("sys_user").Count(&total).Limit(size).Offset((num - 1) * size).Find(&users).Error; err != nil {
		global.GLOBAL_LOG.Error("获取所有用户信息失败", zap.Error(err))
		global.FailWithMsg(ctx, "获取失败")
		return
	}
	for _, user := range users {
		user.Role = global.RoleMap[user.Role]
	}
	global.OkWithDetails(ctx, "", map[string]interface{}{
		"users": users,
		"total": total,
	})
}

// UpdateUser 更新用户信息
func (m *ManagementApi) UpdateUser(ctx *gin.Context) {
	var userForm *response.UserVO
	err := ctx.ShouldBind(&userForm)
	if err != nil {
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
		return
	}
	if err = global.GLOBAL_DB.Table("sys_user").Where("user_id = ?", userForm.UserId).Updates(map[string]interface{}{
		"user_name":   userForm.UserName,
		"phone":       userForm.Phone,
		"email":       userForm.Email,
		"department":  userForm.Department,
		"role":        global.Role[userForm.Role],
		"update_time": time.Now(),
	}).Error; err != nil {
		global.GLOBAL_LOG.Error("更新用户信息失败", zap.String("user_id", userForm.UserId), zap.Error(err))
		global.FailWithMsg(ctx, "更新信息失败")
		return
	}
	global.OkWithMsg(ctx, "更改成功")
}

// UpdatePass 修改密码
func (m *ManagementApi) UpdatePass(ctx *gin.Context) {
	var info *request.PasswordInfo
	err := ctx.ShouldBind(&info)
	if err != nil {
		global.FailWithMsg(ctx, "提交的信息不合规, 请仔细检查后再次提交")
		return
	}
	if err = global.GLOBAL_DB.Table("sys_user").Where("user_id", info.UserId).Updates(map[string]interface{}{
		"password":    utils.Encrypt(info.Password),
		"update_time": time.Now(),
	}).Error; err != nil {
		global.GLOBAL_LOG.Error("修改密码失败", zap.String("user_id", info.UserId), zap.Error(err))
		global.FailWithMsg(ctx, "修改密码失败")
		return
	}
	global.OkWithMsg(ctx, "修改成功")
}

func (m *ManagementApi) GetUser(ctx *gin.Context) {
	userId, _ := ctx.GetQuery("userId")
	if userId == "" {
		global.FailWithMsg(ctx, "用户不存在")
		return
	}
	var user *response.UserVO
	if err := global.GLOBAL_DB.Table("sys_user").Where("user_id", userId).Find(&user).Error; err != nil {
		global.GLOBAL_LOG.Error("获取用户信息失败", zap.String("user_id", userId), zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.FailWithMsg(ctx, "用户不存在")
			return
		}
		global.FailWithMsg(ctx, "获取用户信息失败")
		return
	}
	user.Role = global.RoleMap[user.Role]
	global.OkWithDetails(ctx, "", user)
}

func (m *ManagementApi) DelUser(ctx *gin.Context) {
	userId, _ := ctx.GetQuery("userId")
	var user *model.SysUser
	if err := global.GLOBAL_DB.Table("sys_user").Where("user_id", userId).Delete(&user).Error; err != nil {
		global.GLOBAL_LOG.Error("删除用户失败", zap.String("user_id", userId), zap.Error(err))
		global.FailWithMsg(ctx, "删除失败")
		return
	}
	global.OkWithMsg(ctx, "删除成功")
}
