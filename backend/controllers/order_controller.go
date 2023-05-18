package controllers

import (
	"context"
	"net/http"
	"time"
	
	//"strings"
	//"errors"
	
	"comp/configs"
	"comp/models"
	"comp/utils"
	
	//"auto/middleware"
	
	//"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var OrderCollection *mongo.Collection = configs.GetCollection(configs.DB, "Orders")

func CreateOrder(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	auth, _ := JwtFromHeader(c, fiber.HeaderAuthorization)
	var order models.Order
	defer cancel()

	claims, err := utils.EncodeAccsesToken(auth); if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":err.Error()})
	}

	objId, _ := primitive.ObjectIDFromHex(claims.Id)

	var user models.User

	mongErr := UserCollection.FindOne(ctx, bson.M{"_id":objId}).Decode(&user); if mongErr != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":mongErr.Error()})
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	if validateErr := validate.Struct(&order); validateErr != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": validateErr.Error()})
	}

	newOrder := models.Order {
		ID: primitive.NewObjectID(),
		UserID: objId,
		TotalPrice: order.TotalPrice,
		Products: order.Products,
		CreatedAt: time.Now(),
	}

	_, inserr := ProductCollection.InsertOne(ctx, newOrder); if inserr != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":inserr.Error()})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"data":"ok"})
}

func GetOrders(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	auth, _ := JwtFromHeader(c, fiber.HeaderAuthorization)
	defer cancel()

	claims, err := utils.EncodeAccsesToken(auth); if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":err.Error()})
	}

	objId, _ := primitive.ObjectIDFromHex(claims.Id)

	var user models.User

	mongErr := UserCollection.FindOne(ctx, bson.M{"_id":objId}).Decode(&user); if mongErr != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":mongErr.Error()})
	}

	if (!user.IsAdmin){
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":"permission denaid"})
	}

	cursor, err := OrderCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}
	defer cursor.Close(ctx)

	var orders []models.Order
	if err := cursor.All(ctx, &orders); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	}

	return c.JSON(orders)
}