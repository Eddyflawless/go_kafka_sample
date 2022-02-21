package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"database/sql"
	"time"
	"gopkg.in/robfig/cron.v2"
	_ "github.com/go-sql-driver/mysql"
)


type Book struct {

	Name string `json:"name"`
	Room_id string `json:"room_id" binding:"required,gte=100,lte=500"`
	Phone string `json:"phone"`
	// start_date time.Time `json:"start_date"`
	// end_date time.Time `json:"end_date"`

}

type Room struct {

	id string
	wing string
	capacity int

}

func RunCron(){

	c := cron.New()
	c.AddFunc("@every 1m", CronTargetFunction)
	c.Start()
}

func CronTargetFunction(){
	fmt.Println("print something")
}

func getBookingHandler(c *gin.Context){

}

func addBookingHandler(c *gin.Context){

}

func middlware01() gin.HandlerFunc {

	return func(c *gin.Context){

		log.Println(c.GetHeader("User-Agent"))

		c.Next()
	}
}

func main(){

	RunCron()

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Some error occrured. Err: %s", err)
	}

	db_name := os.Getenv("MYSQL_DB_DATABASE")
	db_url := fmt.Sprintf("%s:%s", os.Getenv("MYSQL_DB_HOST"), os.Getenv("MYSQL_DB_PORT") )
	username := os.Getenv("MYSQL_DB_USERNAME")
	password := os.Getenv("MYSQL_DB_PASSWORD")

	connection_str := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, db_url, db_name )

	db, err := sql.Open("mysql", connection_str)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("-------------Successfully connected to MySql database----------------")

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	router := gin.Default()

	router.Use(middlware01())

	router.Use(cors.Default())

	v1  := router.Group("/v1")

	v1.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "CORS works!"})
    })

	v1.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1.GET("/booking", func(c *gin.Context) {

		book := Book{ "new booking", "012", "233501198766" }

		for i := 0; i < 1000; i++ {

			fmt.Printf("booking %v", i)
		}

		fmt.Printf("%v \n", book)

		c.JSON(http.StatusOK, book)
	})

	v1.POST("/booking", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":82") // listen and serve on 0.0.0.0:82


}