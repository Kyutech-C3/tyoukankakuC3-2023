package controllers

import (
	"log"
	"net/http"

	"github.com/Simo-C3/tyoukankakuC3-2023/api/db"
	"github.com/Simo-C3/tyoukankakuC3-2023/api/models"
	"github.com/Simo-C3/tyoukankakuC3-2023/api/schemas"

	"github.com/gin-gonic/gin"
)

// @Summary Todo一覧を配列で返す
// @Tags ToDo
// @Produce  json
// @Success 200 {object} schemas.ToDos
// @Failure 400 {object} error
// @Router /api/v1/todo [get]
func GetToDo(ctx *gin.Context) {
	db := db.GetDB()

	var todos []models.ToDo

	if err := db.Order("created_at").Find(&todos).Error; err == nil {
		ctx.JSON(http.StatusOK, todos)
		return
	}

	return
}

// @Summary ToDoを作成する
// @Tags ToDo
// @Produce  json
// @Param body body schemas.PostToDo true "request body"
// @Success 200 {object} schemas.ToDo
// @Failure 400 {object} error
// @Router /api/v1/todo [post]
func PostToDo(ctx *gin.Context) {
	var requestBody schemas.PostToDo

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		log.Fatal(err)
		return
	}

	u := models.ToDo{Title: requestBody.Title, Detail: requestBody.Detail, Status: requestBody.Status}
	if err := db.GetDB().Create(&u).Error; err == nil {
		ctx.JSON(http.StatusOK, u)
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		ctx.Abort()
		log.Fatal(err)
		return
	}
}

// @Summary ToDoを編集する
// @Tags ToDo
// @Produce json
// @Param id path string true "id"
// @Param body body schemas.PatchToDo true "request body"
// @Success 200 {object} schemas.ToDo
// @Failure 400 {object} error
// @Router /api/v1/todo/{id} [patch]
func PutToDo(ctx *gin.Context) {
	todoId := ctx.Param("id")

	db := db.GetDB()

	var todo models.ToDo

	if err := db.Where("id = ?", todoId).First(&todo).Error; err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{"msg": err.Error()})
		ctx.Abort()
		log.Fatal(err)
		return
	}

	if err := ctx.BindJSON(&todo); err == nil {
		db.Save(&todo)
		ctx.JSON(http.StatusOK, todo)
		return
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		ctx.Abort()
		log.Fatal(err)
		return
	}
}

// @Summary ToDoを編集する
// @Tags ToDo
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} schemas.DeleteSuccessResponse
// @Failure 400 {object} error
// @Router /api/v1/todo/{id} [delete]
func DeleteToDo(ctx *gin.Context) {
	todoId := ctx.Param("id")

	db := db.GetDB()

	var todo models.ToDo

	if err := db.Where("id = ?", todoId).Delete(&todo).Error; err == nil {
		ctx.JSON(http.StatusOK, schemas.DeleteStatus)
		return
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{"msg": err.Error()})
		ctx.Abort()
		log.Fatal(err)
		return
	}
}
