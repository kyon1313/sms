package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"practice/database"
	"practice/model"
	"practice/route"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	accountSid string
	authToken  string
	fromPhone  string
	toPhone    string
	client     *twilio.RestClient
)

func SendMessage(msg string) {

	params := openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	response, err := client.Api.CreateMessage(&params)
	if err != nil {
		fmt.Printf("error creating and sending message: %s\n", err.Error())
		return
	}
	fmt.Printf("Message SID: %s\n", *response.Sid)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error loading .env: %s\n", err.Error())
		os.Exit(1)
	}

	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	fromPhone = os.Getenv("FROM_PHONE")
	toPhone = os.Getenv("TO_PHONE")

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

}

func Routes(app *fiber.App) {
	app.Post("/cars", route.AddCars)
	app.Get("/cars", route.GetCars)

	//user
	app.Get("/users", route.GetAllUsers)
	app.Get("/users/:id", route.GetUser)
	app.Post("/user", route.AddUser)
	app.Delete("/user/:id", route.DeleteUser)
	app.Put("/user/:id", route.UpdateUser)

	//transaction
	app.Post("/transac", route.AddTransaction)
	app.Post("/message", Mess)

	//code
	app.Post("/code", Confirm)
}

func main() {
	app := fiber.New()
	database.Migration()
	Routes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}

type Messaging struct {
	User       model.User `json:"user"`
	TheMessage string     `json:"the_message"`
}

func MessageSerializer(user model.User, message model.Message) Messaging {

	return Messaging{

		User:       user,
		TheMessage: message.TheMessage,
	}

}

func Mess(c *fiber.Ctx) error {
	var message model.Message
	if err := c.BodyParser(&message); err != nil {
		return c.SendString("Fuck you cant post a message")
	}
	var user model.User
	database.DB.Find(&user, "id=?", message.UserId)
	if user.ID == 0 {
		return c.SendString("This user not exist")
	}

	response := MessageSerializer(user, message)
	rand.Seed(time.Now().Unix())
	var random = []int{}
	for len(random) != 6 {
		var rannumbers = rand.Intn(9)
		random = append(random, rannumbers)
	}

	var rancomNumber string
	for _, v := range random {
		r := strconv.Itoa(v)
		rancomNumber = rancomNumber + r
	}
	i, err := strconv.Atoi(rancomNumber)

	if err != nil {
		fmt.Print("cant convert")
	} else {
		message.TheCode = i
	}

	database.DB.Create(&message)
	msg := fmt.Sprintf(message.TheMessage, message.TheCode, user.Name)
	SendMessage(msg)

	return c.JSON(&fiber.Map{
		"message": "message successfully added",
		"data":    response,
		"success": true,
	})
}

type Confirmation struct {
	User int `json:"user"`
	Code int `json:"code"`
}

func Confirm(c *fiber.Ctx) error {
	var confirm Confirmation

	var message model.Message
	if err := c.BodyParser(&confirm); err != nil {
		return c.SendString("Parsing problem idiot")
	}

	database.DB.Find(&message, "user_id=?", confirm.User)
	if message.UserId != confirm.User {
		return c.SendString("user not exist")
	} else {
		if message.TheCode != confirm.Code {
			return c.SendString("Code wrong")
		}
	}
	var user model.User
	database.DB.Find(&user, "id=?", message.UserId)

	return c.JSON(&fiber.Map{
		"message": "Code Correct",
		"date":    user,
	})

}
