package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"instant_messaging_system/db"
	"instant_messaging_system/define"
	"instant_messaging_system/dto"
	"instant_messaging_system/models"
	"instant_messaging_system/utils"
	"net/http"
	"time"
)

func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if !utils.VerifyEmail(email) {
		c.JSON(http.StatusOK, define.NewResponseResult(400, "邮箱地址格式不对", nil))
		return
	}

	count := db.GetUserCountByEmail(email)
	if count > 0 {
		c.JSON(http.StatusOK, define.NewResponseResult(400, "邮箱已被注册", nil))
		return
	}

	code := utils.GetCode()
	err := utils.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "邮件发送失败", nil))
		return
	}

	// 把验证码保存到数据库
	err = db.RedisClient.Set(context.Background(), define.CodePrefix+email, code, time.Second*time.Duration(define.CodeExpireTime)).Err()
	if err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "系统错误:"+err.Error(), nil))
	}

	c.JSON(http.StatusOK, define.NewResponseResult(200, "验证码发送成功", nil))
}

func Register(c *gin.Context) {
	code := c.PostForm("code")
	email := c.PostForm("email")
	password := c.PostForm("password")
	nickname := utils.DefaultPostForm(c, "nickname", email)
	avatar := utils.DefaultPostForm(c, "avatar", "http://localhost:8080/auth/static/img/01.png")
	sex := utils.DefaultPostForm(c, "sex", "0")
	if code == "" || !utils.VerifyEmail(email) || password == "" {
		c.JSON(http.StatusOK, define.NewResponseResult(400, "参数不正确", nil))
		return
	}

	// 验证码是否正确
	result, err := db.RedisClient.Get(context.Background(), define.CodePrefix+email).Result()
	// 把code从redis中移除
	defer db.RedisClient.Del(context.Background(), define.CodePrefix+email)
	if err != nil || result != code {
		c.JSON(http.StatusOK, define.NewResponseResult(400, "验证码不正确或已过期", nil))
		return
	}

	// 将用户注册信息保存起来
	user := &models.User{
		Email:    email,
		Password: utils.GetMd5(password),
		Avatar:   avatar,
		Nickname: nickname,
		Sex:      sex[0] - '0',
	}
	if err = db.InsertUser(user); err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "系统错误:"+err.Error(), nil))
		return
	}

	// 生成token+设置过期时间+存于redis并返回
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "系统错误:"+err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, define.NewResponseResult(200, "注册成功", gin.H{"token": token}))
}

func Login(c *gin.Context) {
	// 获取并校验参数
	email := c.PostForm("email")
	pwd := c.PostForm("password")
	if !utils.VerifyEmail(email) || len(pwd) == 0 {
		c.JSON(http.StatusOK, define.NewResponseResult(401, "参数有误", nil))
		return
	}

	// 查询数据库
	id, err := db.GetUserIdByEmailPwd(email, utils.GetMd5(pwd))
	if err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(401, "账号或密码错误", nil))
	}

	// 生成token+设置过期时间+存于redis并返回
	token, err := utils.GenerateToken(id, email)
	if err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "系统错误:"+err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, define.NewResponseResult(200, "登录成功", gin.H{"token": token}))
}

func PersonalMsg(c *gin.Context) {
	uid := c.GetUint("user_id")
	user, err := db.GetUserById(uid)
	if err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "系统错误,获取失败", nil))
	}
	c.JSON(http.StatusOK, define.NewResponseResult(200, "成功", gin.H{"user_msg": dto.NewUserDto(user)}))
}

func UserQuery(c *gin.Context) {
	email := c.Query("email")
	if !utils.VerifyEmail(email) {
		c.JSON(http.StatusOK, define.NewResponseResult(401, "参数有误, 获取失败", nil))
		return
	}
	user, err := db.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusOK, define.NewResponseResult(500, "系统错误,获取失败", nil))
	}
	// 避免用户隐式信息泄露
	user.ID = 0
	c.JSON(http.StatusOK, define.NewResponseResult(200, "成功", gin.H{"user_msg": dto.NewUserDto(user)}))
}
