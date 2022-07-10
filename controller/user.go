package controller

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aryaputra777/rest/config"
	"github.com/aryaputra777/rest/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     int    `json:"id_role"`
}

type Users struct {
	Users []User `json:"data"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func Getdatauser(c *fiber.Ctx) error {
	var db = config.Db
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	expires := claims.Expires
	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	_, err = utils.ExtractTokenID(c)

	rows, err := db.Query("SELECT a.id, a.username, a.name, a.email, a.status, a.phone, a.id_role FROM users as a order by id")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	defer rows.Close()
	result := Users{}

	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Status, &user.Phone, &user.Role); err != nil {
			return err
		}

		result.Users = append(result.Users, user)
	}
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"code":  200,
		"data":  result.Users,
	})
}

func Savedatauser(c *fiber.Ctx) error {
	var db = config.Db
	u := new(User)
	validate := utils.NewValidator()

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := validate.Struct(u); err != nil {
		fmt.Println("err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}
	bycrypt, _ := HashPassword(u.Password)

	resp, _ := db.Query("SELECT  a.id, a.username, a.name, a.email, a.status, a.phone, a.id_role FROM users as a WHERE username=$1 or email=$2", u.Username, u.Email)

	if resp == nil {
		err := db.QueryRow("INSERT INTO users (username, name, email,status,phone, password, id_role)VALUES ($1, $2, $3, $4, $5, $6, $7) returning id", u.Username, u.Name, u.Email, u.Status, u.Phone, bycrypt, u.Role).Scan(&u.ID)

		if err != nil {
			return err
		}

		return c.JSON(u)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "user already exist",
		})
	}

}

func Detailuser(c *fiber.Ctx) error {
	var db = config.Db
	var id = c.Params("id")
	user := User{}
	err := db.QueryRow("SELECT  a.id, a.username, a.name, a.email, a.status, a.phone, a.id_role FROM users as a WHERE id=$1", id).Scan(&user.ID, &user.Username, &user.Name, &user.Email, &user.Status, &user.Phone, &user.Role)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(user)
}

func Updateuser(c *fiber.Ctx) error {
	var db = config.Db
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var id = c.Params("id")
	inte, err := strconv.Atoi(id)
	u.ID = inte

	_, err = db.Query("UPDATE users SET username=$1,name=$2,email=$3, status=$4, phone=$5, password=$6, id_role=$7 WHERE id=$8", u.Username, u.Name, u.Email, u.Status, u.Phone, u.Password, u.Role, u.ID)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(u)
}

func Deleteuser(c *fiber.Ctx) error {
	var db = config.Db
	var id = c.Params("id")
	_, err := db.Query("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return c.JSON("Deleted")
}
