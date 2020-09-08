package user

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hesen/blog/pkg/res"
	"github.com/hesen/blog/pkg/util"

	"strconv"
	// "log"
	"net/http"
)

// GetUsers 获取用户列表
func GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	users := UserService.findUsers(page, limit)

	c.JSON(http.StatusOK, users)
}

// Register 注册用户
func Register(c *gin.Context) {
	user := NewUserModel()
	var err error

	if err = c.ShouldBindWith(&user, binding.FormPost); err != nil {
		c.JSON(http.StatusOK, res.Error(res.MissingRequiredParameterCode))
		return
	}
	user.IP = c.ClientIP()

	user.Password = util.Md5(user.Password)

	isExist := NewUserModel()
	err = user.findUser(&isExist)

	if err != nil {
		c.JSON(http.StatusOK, res.Error(res.SystemErrorCode))
		return
	}

	if isExist.Phone != "" {
		c.JSON(http.StatusOK, res.Error(res.UserAlreadyExistsCode))
		return
	}

	err = user.create()

	if err != nil {
		c.JSON(http.StatusOK, res.Error(res.SystemErrorCode))
		return
	}

	c.JSON(http.StatusOK, res.Success(gin.H{ "phone": user.Phone, "id": user.ID }))
}

// Login 登录
func Login(c *gin.Context) {
	user := NewUserModel()
	var err error

	user.Phone = c.PostForm("phone")
	user.Password = c.PostForm("password")

	if user.Phone == "" || user.Password == "" {
		c.JSON(http.StatusOK, res.Error(res.MissingRequiredParameterCode))
		return
	}

	isExist := NewUserModel()
	err = user.findUserWithFullInfo(&isExist)

	if err != nil {
		c.JSON(http.StatusOK, res.Error(res.SystemErrorCode))
		return
	}

	if isExist.Phone == "" {
		c.JSON(http.StatusOK, res.Error(res.UserNotExistsCode))
		return
	}
	user.Password = util.Md5(user.Password)
	if isExist.Password != user.Password {
		c.JSON(http.StatusOK, res.Error(res.WrongPasswordCode))
		return
	}

	result := gin.H{
		"phone": isExist.Phone,
		"email": isExist.Email,
		"age": isExist.Age,
		"address": isExist.Address,
	}

	result["token"], err = util.BuildJwt(util.JwtClaims{
		Phone: isExist.Phone, Email: isExist.Email, Name: isExist.Name,
	}, 2)

	if err != nil {
		c.JSON(http.StatusOK, res.Error(res.SystemErrorCode))
		return
	}

	result["token"] = "Bearer " + result["token"].(string)

	c.JSON(http.StatusOK, res.Success(result))
}
