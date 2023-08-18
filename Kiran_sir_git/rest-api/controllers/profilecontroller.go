package controllers

import (
	"net/http"
	"strings"
	"rest-api/interfaces"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileService interfaces.IProfile
}

func InitProfileController(profileService interfaces.IProfile) ProfileController {
	return ProfileController{profileService}
}

func  (pc *ProfileController) CreateProfile(ctx *gin.Context) {
	var profile *models.Profile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := pc.ProfileService.CreateProfile(profile)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}
func (pc *ProfileController)  EditProfile(context *gin.Context)        {}
func (pc *ProfileController)  GetProfileById(context *gin.Context)     {}
func (pc *ProfileController)  GetProfileBySearch(context *gin.Context) {}
func (pc *ProfileController) DeleteProfile(context *gin.Context)      {}
