package controller

import (
	"github.com/aryaputra777/rest/config"
	"github.com/aryaputra777/rest/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Userlogin struct {
	ID       uint32 `json:"ID"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Phone    string `json:"phone"`
	Role     string `json:"id_role"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *fiber.Ctx) error {
	var db = config.Db

	userlogin := new(Userlogin)

	if err := c.BodyParser(userlogin); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	str := userlogin.Password
	err := db.QueryRow("SELECT  a.id, a.username, a.name, a.email, a.status, a.phone, a.id_role, a.password FROM users as a WHERE username=$1", userlogin.Username).Scan(&userlogin.ID, &userlogin.Username, &userlogin.Name, &userlogin.Email, &userlogin.Status, &userlogin.Phone, &userlogin.Role, &userlogin.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	match := CheckPasswordHash(str, userlogin.Password)
	if match == false {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "User Tidak Valid",
		})
	}

	token, err := utils.GenerateNewAccessToken(userlogin.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	userlogin.Token = token
	return c.JSON(fiber.Map{
		"Id":    userlogin.ID,
		"name":  userlogin.Name,
		"token": userlogin.Token,
	})

}

func Logout(c *fiber.Ctx) error {
	var db = config.Db
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	_, err := db.Query("INSERT INTO forkum.users (username, name, email,status,phone, password, id_role)VALUES ($1, $2, $3, $4, $5, $6, $7)", u.Username, u.Name, u.Email, u.Status, u.Phone, u.Password, u.Role)
	if err != nil {
		return err
	}

	return c.JSON(u)
}
