package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
	"path"
	"server/global"
	"server/model/request"
	"server/model/response"
	"server/utils"
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
	global.OkWithData(ctx, response.UserInfo{
		UserName:   claims.UserName,
		UserId:     claims.UserId,
		Email:      claims.Email,
		Phone:      claims.Phone,
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
