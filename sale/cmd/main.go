package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"fmt"
	"log"
)

type Database struct {
	Connection *sqlx.DB
}

func main() {
	connection, err := sqlx.Connect("mysql", "sealteam:sckshuhari@(store-database:3306)/toy")
	if err != nil {
		log.Fatalln("cannot connect to database", err)
	}
	api := Database{
		Connection: connection,
	}
	route := gin.Default()
	route.GET("/api/v1/health", api.healthCheckHandle)
	log.Fatal(route.Run(":3030"))

}

func (api Database) healthCheckHandle(context *gin.Context) {
	user, err := GetUserNameFromDB(api.Connection)
	if err != nil {
		context.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(200, gin.H{
		"message": user,
	})
}

func GetUserNameFromDB(connection *sqlx.DB) (User, error) {
	user := User{}
	err := connection.Get(&user, "SELECT id,name FROM user WHERE ID=1")
	if err != nil {
		fmt.Printf("Get user name from tearup get error : %s", err.Error())
		return User{}, err
	}
	return user, err
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
