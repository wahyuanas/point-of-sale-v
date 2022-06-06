package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	_ "github.com/go-sql-driver/mysql"

	_accountRepo "github.com/wahyuanas/point-of-sale-v/account/repository"

	_accountUsecase "github.com/wahyuanas/point-of-sale-v/account/usecase"

	_accountHandler "github.com/wahyuanas/point-of-sale-v/account/api/handler"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		log.Println("succes getting the env values")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbDriver := os.Getenv("DB_DRIVER")
	appPort := os.Getenv("APP_PORT")

	ctxTimeOut := os.Getenv("CONTEXT_TIME_OUT")
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(dbDriver, dsn)

	if err != nil {
		log.Printf("Cannot connect to %s database", dbDriver)
		log.Fatal(err)
	} else {
		log.Printf("We are connected to the %s database", dbDriver)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	db := bun.NewDB(dbConn, mysqldialect.New())

	app := fiber.New()
	accountRepo := _accountRepo.NewAccountRepository(db)
	ctm, _ := strconv.Atoi(ctxTimeOut)
	timeoutContext := time.Duration(ctm) * time.Second
	accountUsecase := _accountUsecase.NewAccountUsecase(accountRepo, timeoutContext)
	_accountHandler.NewAccountHandler(app, accountUsecase)
	log.Fatal(app.Listen(":" + appPort))
}
