package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// struct which represent mainly interaction with redis
type RedisRepository struct {
	Client *redis.Client
	Ctx    context.Context
}

// should update existing history with order or create a new one
func (r *RedisRepository) SaveOrder(userId string, order Order) error {
	val, err := r.Client.Get(r.Ctx, (userId)).Result()
	if err != nil {
		//create new user with user history
		history := OrderHistory{Orders: []Order{order}}
		r.saveHistory(userId, &history)
	} else {
		// update existing user history
		var history OrderHistory
		json.Unmarshal([]byte(val), &history)
		history.add(order)
		r.saveHistory(userId, &history)
	}
	return nil
}

// should add new history to repository (ACHTUNG it will replace the old one)
// and also it will clean orders which are older than 24 hours
func (r *RedisRepository) saveHistory(userId string, history *OrderHistory) error {
	newHistory := CleanWithinWindow(history, 24)

	data, jsonError := json.Marshal(newHistory)
	if jsonError != nil {
		log.Error(fmt.Sprintf("Erorr while saving history for user with [%s]", userId))
		log.Error(jsonError)
		return jsonError
	}
	err2 := r.Client.Set(r.Ctx, userId, string(data), time.Duration(24*time.Hour)).Err()
	if err2 != nil && jsonError != nil {
		log.Error(fmt.Sprintf("Erorr while saving history for user with [%s]", userId))
		log.Error(jsonError)
		return err2
	}

	log.Trace(fmt.Sprintf("Cleaning and saving new history for user with id [%s]", userId))
	return nil
}

// should return history of existing user otherwise nill
func (r *RedisRepository) FindOrEmpty(userId string) (OrderHistory, error) {
	var history OrderHistory
	val, redisError := r.Client.Get(r.Ctx, userId).Result()
	
	if redis.Nil == redisError {
		//in case we don't have any records for users 
		//return empty history
		return OrderHistory{Orders: []Order{}}, nil
	}
	if redisError  == nil{
		return OrderHistory{}, redisError 
	}


	err := json.Unmarshal([]byte(val), &history)
	if err != nil {
		return OrderHistory{}, err
	}
	return history, nil
}
