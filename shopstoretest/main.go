package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//var SecretKey string = "ItIsCompletelySecret"

func main() {

	db, err := sql.Open("mysql",
		"shopstore:shopstoret0lk2o20@tcp(localhost:1234)/shopstore_db")

	dberr := db.Ping()
	fmt.Println("dberr :", dberr)

	fmt.Println("error : ", err)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	fmt.Println("\n\nerror ", err2)
	//_, err := mysqlRepo.DB.Query("INSERT INTO users (userID) VALUES(23)")

	//fmt.Println("error :", err, "\n\n\n", version)
	//uService := userservice.Service{Repository: mysqlRepo}

	/*req := userservice.RegisterRequest{
		Name:        "hamed",
		Password:    "123444",
		PhoneNumber: "888",
	}

	res, err := uService.Register(req)
	if err != nil {

		fmt.Println("cant save in database", err)
	}

	fmt.Println("\n\n", res)
	//req := userservice.RegisterRequest{}

	//dbCfg := cfg.NewDatabaseCfg(hamedcfg.DatabaseUser, hamedcfg.DatabasePassword, hamedcfg.DatabaseHost,
	//hamedcfg.DatabasePort, hamedcfg.DatabaseProtocol, hamedcfg.DatabaseName)

	//mysqlRepo := mysql.New(dbCfg)

	//uService := userservice.Service{Repository: mysqlRepo}
	//fmt.Println("fuck u")

	httpserver.Serve()
	/*e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("hi here")
		type Message struct {
			message string
		}

		myMessage := Message{message: "good boy"}

		return c.JSON(http.StatusOK, myMessage)
	})
	e.Logger.Fatal(e.Start(":5555"))

	e.GET("/health-check", httpserver.HeathCheck)
	fmt.Println("my name is hamed")

	e.POST("/users/register", func(c echo.Context) error {
		fmt.Println("here")

		bErr := c.Bind(&req)
		if bErr != nil {

			return bErr
		}
		_, err := uService.Register(req)
		if err != nil {

			return fmt.Errorf("cant register user %w", err)
		}

		return c.String(http.StatusOK, "res")
	})*/

}
