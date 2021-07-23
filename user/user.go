package user

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	err      error
	DSN      string
	user     string
	pwd      string
	database string
)

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LasName   string `json:"lastname"`
	Email     string `json:"email"`
}

func init() {
	godotenv.Load()
	user = strings.Trim(os.Getenv("MYSQL_USR"), " ")
	pwd = strings.Trim(os.Getenv("MYSQL_PWD"), " ")
	database = strings.Trim(os.Getenv("MYSQL_DATABASE"), " ")
	DSN = user + ":" + pwd + "@tcp(127.0.0.1:3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func InitialMigration() {
	fmt.Printf("DSN: %s", DSN)
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Cannot connect to Database")
		// panic("Cannot connect to Database")
	}

	DB.AutoMigrate(&User{})
	fmt.Println("Cannot connect to Database")
}

func GetUsers(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)
	return c.JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.Find(&user, id)
	return c.JSON(&user)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Create(&user)
	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User is not available")
	}

	DB.Delete(&user)
	return c.Status(202).SendString("User deleted")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	DB.First(&user, id)
	if user.Email == "" {
		return c.Status(500).SendString("User is not available")
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&user)
	return c.JSON(&user)
}
