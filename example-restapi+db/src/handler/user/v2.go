package user

import (
	"example-restapi+db/src/model"
	"example-restapi+db/src/service/sql"
	"example-restapi+db/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandlerV2 struct{}

func (User *UserHandlerV2) GetUserList(ctx *gin.Context) {
	var user []model.UserProfile
	inst := sql.GetConnection()
	util.PrintInfo("DB Inst : %+v", inst)

	if err := inst.Sql.Debug().Table("user").Limit(10).Offset(280).Find(&user).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server Error"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

func (User *UserHandlerV2) GetUser(ctx *gin.Context) {
	idParam, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not acceptable user id"})
		return
	}
	userId := uint(idParam)

	var user model.UserProfile
	inst := sql.GetConnection()
	util.PrintInfo("DB Inst : %+v", inst)

	if err := inst.Sql.Debug().Table("user").Find(&user, userId).Error; err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, user)
}

func (User *UserHandlerV2) PostUser(ctx *gin.Context) {
	newUser := new(model.UserProfile)
	if err := ctx.BindJSON(newUser); err != nil {
		util.PrintError("%+v", err)
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Not acceptable user data"})
		return
	}

	inst := sql.GetConnection()
	util.PrintInfo("DB Inst : %+v", inst)
	util.PrintInfo("%+v", newUser)

	if err := inst.Sql.Debug().Table("user").Create(&newUser).Error; err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server Error"})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, *newUser)
}
