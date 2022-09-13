package pkg

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type CoffeeService struct {
	Redis     *RedisRepository
	Validator *QuotaValidator
}

func (cfs *CoffeeService) ProcessOrder(quota string, user_id string, newOrder string) (string, error) {
	if err := cfs.validateUserAndQuota(quota, user_id); err != nil {
		return "", err
	}

	//validate user history
	history, repoError := cfs.Redis.FindOrEmpty(user_id)
	if repoError != nil {
		return "", repoError
	}

	
	err := cfs.Validator.Handle(quota, history, newOrder)
	if err != nil {
		return "", err
	}

	order := Order{Product: newOrder, Timestamp: time.Now()}
	cfs.Redis.SaveOrder(user_id, order)
	// log.Debug(fmt.Sprintf("Added new order with [%s] for user with id [%s]",newOrder,user_id))
	return fmt.Sprintf("Enjoy your %s ", newOrder), nil
}

// this function will check how right are inserted values into request
func (cfs *CoffeeService) validateUserAndQuota(quota string, user_id string) error {
	//assuming that instead of this map we have a call to a relational database
	//where we do keep quotas <> user relation
	var db = map[string]map[string]bool{
		Basic:  {"1": true, "2": true, "3": true},
		Lover:  {"4": true, "5": true, "6": true},
		Maniac: {"7": true, "8": true, "9": true},
	}

	value, ok := db[quota]

	if ok {
		_, ok2 := value[user_id]
		if ok2 {
			return nil
		} else {
			log.Info(fmt.Sprintf("User with id [%s] is not in the db", user_id))
			return errors.New(fmt.Sprintf("No such user with id:  [%s]", user_id))
		}
	} else {
		log.Info(fmt.Sprintf("Quota with name [%s] is not in the db", quota))
		return errors.New(fmt.Sprintf("No such quota with name:  %s", quota))
	}
}
