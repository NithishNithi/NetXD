package services

import (
	"context"
	"errors"
	"rest-api/interfaces"
	"rest-api/models"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProfileService struct {
	ProfileCollection *mongo.Collection
	ctx               context.Context
}

func NewProfileServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.IProfile {
	return &ProfileService{ collection , ctx}
}

func (p *ProfileService) CreateProfile(user *models.Profile) (*models.DBResponse, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = false
	user.Role = "user"

	//hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = user.Password
	res, err := p.ProfileCollection.InsertOne(p.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	// Create a unique index for the email field
	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := p.ProfileCollection.Indexes().CreateOne(p.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = p.ProfileCollection.FindOne(p.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
func (p *ProfileService) EditProfile(profile *models.Profile) {}
func (p *ProfileService) DeleteProfile(profileId int)         {}
func (p *ProfileService) GetProfileById(profileId int)        {}
func (p *ProfileService) GetProfileBySearch()                 {}
