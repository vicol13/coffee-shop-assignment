package pkg

import (
	"errors"
	"fmt"
	"time"

)


type QuotaValidator struct {
	RulesContainer *RulesContainer
}

// represents the generic function to handle new requests to validate the quota
// @quota -> represent the of current user
// @history -> represent the previous reqeusts of the user 
// @newOrder -> represent the new product which user wants to get (e.g. americano, espresso)
func (qv *QuotaValidator) Handle(quota string,history OrderHistory, newOrder string ) error {
	//todo: move to map number of products increase
	switch quota {
		case Basic: return qv.onBasic(history,newOrder)
		case Maniac: return qv.onManiac(history,newOrder)
		case Lover: return qv.onLover(history,newOrder)
	}
	
	return errors.New("no such quota: " + quota)
}

//
// concrete handler for quotas
//
func (qv *QuotaValidator) onBasic(history OrderHistory, newOrder string) error {
	config := qv.RulesContainer.GetProductRules(Basic,newOrder)
	return qv.validate(history, newOrder, config.Limit, config.TimeWindowHrs)
}

func (qv *QuotaValidator) onLover(history OrderHistory,newOrder string) error {
	config := qv.RulesContainer.GetProductRules(Lover,newOrder)
	return qv.validate(history, newOrder, config.Limit, config.TimeWindowHrs)
}

func (qv *QuotaValidator) onManiac(history OrderHistory, newOrder string) error {
	config := qv.RulesContainer.GetProductRules(Maniac,newOrder)
	return qv.validate(history, newOrder,config.Limit, config.TimeWindowHrs)
}


// generic function to validate the quota constraints for certain item
// @history -> previous aquired by user
// @newOrder -> desired new item
// @limit -> how many items is user is allowed to get 
// @timeWindowHrs -> in which time window user is allowed to get @limit
func (qv *QuotaValidator) validate(history OrderHistory, newOrder string, limit int, timeWindowHrs int) error {
	end := time.Now()
	start := end.Add(-(time.Duration(timeWindowHrs) * time.Hour))

	counter :=0
	minTime := end //represent timestamp of oldest request

	//iterate over elements (history) of filter @newOrder type
	//within given time window and increment counter
	for _,element := range history.Orders{
		if element.Product == newOrder && InTimeSpan(start,end,element.Timestamp){
			minTime = Min(minTime,element.Timestamp)
			if(counter == limit -1){
				timeToWait := GetWaitTime(minTime,end,timeWindowHrs)
				return errors.New(fmt.Sprintf("excedeed limit of available drinks, please wait : [%s]",timeToWait))
			}
			counter = counter + 1
		}	
    }
	return nil
}




