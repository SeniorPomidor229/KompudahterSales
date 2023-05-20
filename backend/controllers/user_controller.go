package controllers

import (
	"context"

	"net/http"
	"time"
	"strings"
	"errors"

	"comp/configs"
	"comp/models"
	"comp/utils"
	"comp/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var UserCollection *mongo.Collection = configs.GetCollection(configs.DB, "Users")
var validate = validator.New()

func Register(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	if validateErr := validate.Struct(&user); validateErr != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": validateErr.Error()})
	}

	hash, _ := utils.CreateHash(user.Password)

	if (!utils.IsEmailValid(user.Email)) {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":"Invalid Username"})
	}

	newUser := models.User{
		ID: 			primitive.NewObjectID(),
		Email:   		utils.NormalizeEmail(user.Email),
		Password:   	string(hash),
		FirstName:  	user.FirstName,
		LastName:   	user.LastName,
		IsAdmin:    	false,
	}

	_, err := UserCollection.InsertOne(ctx, newUser); if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":err.Error()})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"data":"ok"})
}

func Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.UserDto
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	if validateErr := validate.Struct(&user); validateErr != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": validateErr.Error()})
	}

	var userdb models.User

	err := UserCollection.FindOne(ctx, bson.M{"email": user.Username}).Decode(&userdb); if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	match, err := utils.VerifyHash(user.Password, string(userdb.Password));	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":err.Error()})
	}

	if !match {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":"Invalid passwprd"})
	}

	token, _ := utils.GenerateNewAccessToken(userdb.ID.Hex())
	return c.Status(http.StatusOK).JSON(map[string]string{"accses_token":token})
}

func JwtFromHeader(c *fiber.Ctx, header string) (string, error) {
	auth := c.Get(header)
	l := len(middleware.JWTAuthScheme)
	if len(auth) > l+1 && strings.EqualFold(auth[:l], middleware.JWTAuthScheme) {
		return auth[l+1:], nil
	}
	return "", errors.New("Invalid token")
}