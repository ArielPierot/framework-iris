package main

import (
	"log"

	"github.com/arielpierot/iris-framework/handler"
	"github.com/arielpierot/iris-framework/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/basicauth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=desafio_db port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.Produto{})
	if err != nil {
		log.Fatal(err)
	}

	handler := handler.NewHandlerEstoque(db)

	app := iris.Default()

	opts := basicauth.Options{
		Allow: basicauth.AllowUsersFile("users.yml", basicauth.BCRYPT),
		Realm: basicauth.DefaultRealm,
	}

	auth := basicauth.New(opts)

	app.Use(auth)

	app.Get("/produto", handler.List)
	app.Post("/produto", handler.Create)
	app.Get("/produto/{codigo}", handler.Fetch)
	app.Put("/produto/{codigo}", handler.Update)
	app.Delete("/produto/{codigo}", handler.Delete)

	err = app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
