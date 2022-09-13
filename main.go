package main

import (
	"aaha/pkg"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetLevel(log.TraceLevel)
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rulesContainer := pkg.InitRules()
	quotaValidator := pkg.QuotaValidator{RulesContainer: &rulesContainer}
	redis := pkg.RedisRepository{Client: client, Ctx: ctx}
	service := pkg.CoffeeService{Redis: &redis, Validator: &quotaValidator}
	webserviceHandler := pkg.CoffeeRequestHandler{Coffeeservice: &service}

	router := httprouter.New()
	router.GET("/coffee/:name", webserviceHandler.GetCoffe)
	log.Fatal(http.ListenAndServe(":8000", router))
}
