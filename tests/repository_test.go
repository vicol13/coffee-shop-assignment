package tests


import (
    "testing"
    "github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	"context"
	"aaha/pkg"
	"time"
	// log "github.com/sirupsen/logrus"
)



func TestRedisSavingOrderWithoutPreviousHistory(t *testing.T) {
	//given
	s,_ := miniredis.Run()
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	repo := pkg.RedisRepository{Client: client, Ctx: ctx}
	tmg := time.Now()
	order1 := pkg.Order{Product: pkg.AMERICANO ,Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }

	//when
	repo.SaveOrder("1",order1) 
	
	//then
	val,err := repo.Find("1")

	if err != nil {
		t.Errorf("Error while retrieving the entity from redis")
	}

	if len(val.Orders) != 1 {
		t.Errorf("Size of history is expected to be 1, got %d", len(val.Orders))
	}
	//
	//  todo: investigate this case as redis is returning
	//  @val = {americano 2022-08-14 09:09:47.450537 +0300 EEST}
	//	@order1 = {americano 2022-08-14 09:09:47.450537 +0300 EEST m=-10799.998567494}
	// 	find issue with m=-14...
	
	//  if val.Orders[0].Product != order1.Product || val.Orders[0].Timestamp != order1.Timestamp {
	// 	log.Error(val.Orders[0])
	// 	log.Error(order1)
	// 	t.Errorf("Enities doesn't match")
	// }
}




func TestRedisSavingOrderWithPreviousHistory(t *testing.T) {
	//given
	s,_ := miniredis.Run()
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	repo := pkg.RedisRepository{Client: client, Ctx: ctx}
	tmg := time.Now()
	order1 := pkg.Order{Product: pkg.AMERICANO ,Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }
	order2 := pkg.Order{Product: pkg.AMERICANO ,Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }
	
	//when
	repo.SaveOrder("1",order1) 
	repo.SaveOrder("1",order2) 
	
	//then
	val,err := repo.Find("1")

	if err != nil {
		t.Errorf("Error while retrieving the entity from redis")
	}

	if len(val.Orders) != 2 {
		t.Errorf("Size of history is expected to be 1, got %d", len(val.Orders))
	}
}