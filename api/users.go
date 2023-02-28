package endpoints

import (
	"Backend_for_Capstone/api/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	UserId    int64  `json:"user_id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
}

func GetAllUsers(c *gin.Context) {
	db := database.Database
	var users []*User
	query, e := db.Query("select userid, firstname, lastname, username from users")
	if e != nil {
		log.Fatalf("Database not found")
	}
	for query.Next() {

		u := new(User)
		if err := query.Scan(&u.UserId, &u.Firstname, &u.Lastname, &u.Username); err != nil {
			log.Fatalf("Errow: %s", err)
		}
		users = append(users, u)
		//log.Printf("id: %d, fname: %s, lname: %s, username: %s", userId, firstname, lastname, username)
	}

	c.JSON(http.StatusOK, users)
}
