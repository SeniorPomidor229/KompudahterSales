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

var ProductCollection *mongo.Collection = configs.GetCollection(configs.DB, "Products")
var CategoryCollection *mongo.Collection = configs.GetCollection(configs.DB, "Categories")

func CreateProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	auth, _ := JwtFromHeader(c, fiber.HeaderAuthorization)
	var product models.Product
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


	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": err.Error()})
	}

	if validateErr := validate.Struct(&product); validateErr != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": validateErr.Error()})
	}

	var cat models.Category

	mongoErr := CategoryCollection.FindOne(ctx, bson.M{"_id": product.Category.ID}).Decode(&cat)
	if mongoErr != nil {
	    if mongoErr == mongo.ErrNoDocuments {
	        _, insertErr := CategoryCollection.InsertOne(ctx, product.Category)
	        if insertErr != nil {
	            return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": insertErr.Error()})
	        }
	    }
	    return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": mongoErr.Error()})
	}

	newProduct := models.Product {
		ID: primitive.NewObjectID(),
		Name: product.Name,
		PhotoUrl: product.PhotoUrl,
		Description: product.Description,
		Price: product.Price,
		Category: product.Category,
	}

	_, inserr := ProductCollection.InsertOne(ctx, newProduct); if inserr != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":inserr.Error()})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"data":"ok"})
}

func CreateCategory(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	auth, _ := JwtFromHeader(c, fiber.HeaderAuthorization)
	var cat models.Category
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

	newCat := models.Category {
		ID: primitive.NewObjectID(),
		Name: cat.Name,
		Description: cat.Description,
	}

	_, ererr := CategoryCollection.InsertOne(ctx, newCat); if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{"error":ererr.Error()})
	}

	return c.Status(http.StatusOK).JSON(map[string]string{"data":"ok"})
}

func GetAllProduct(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	categoryId := c.Query("categoryId")
	defer cancel()

	filter := bson.M{}
    if categoryId != "" {
        objID, err := primitive.ObjectIDFromHex(categoryId)
        if err != nil {
            return c.Status(http.StatusBadRequest).JSON(map[string]string{"error": "Invalid categoryId"})
        }
        filter["category._id"] = objID
    }

	cursor, err := ProductCollection.Find(ctx, filter)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
    }
    defer cursor.Close(ctx)

    var products []models.Product
    if err := cursor.All(ctx, &products); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
    }

    return c.JSON(products)
}

func GetAllCategories(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

    cursor, err := CategoryCollection.Find(ctx, bson.M{})
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
    }
    defer cursor.Close(ctx)

    var categories []models.Category
    if err := cursor.All(ctx, &categories); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
    }

    return c.JSON(categories)
}

func GetProductById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	Id := c.Query("Id")
	defer cancel()
}