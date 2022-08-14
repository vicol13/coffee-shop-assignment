package pkg

import (
	"errors"
	"fmt"
	"time"

)


type QuotaValidator struct {
	RulesContainer *RulesContainer
}

func (qv *QuotaValidator) Handle(quota string,history OrderHistory, newOrder string ) error {
	//todo: move to map number of products increase
	switch quota {
		case BASIC: return qv.onBasic(history,newOrder)
		case MANIAC: return qv.onManiac(history,newOrder)
		case LOVER: return qv.onLover(history,newOrder)
	}
	
	return errors.New("no such quota: " + quota)
}

func (qv *QuotaValidator) onBasic(history OrderHistory, newOrder string) error {
	rules := qv.RulesContainer.getRule(BASIC)
	timeWindow :=qv.RulesContainer.getTiming(BASIC)
	return qv.validate(history, newOrder, rules ,timeWindow)
}

func (qv *QuotaValidator) onLover(history OrderHistory,newOrder string) error {
	rules := qv.RulesContainer.getRule(LOVER)
	timeWindow :=qv.RulesContainer.getTiming(LOVER)
	return qv.validate(history, newOrder, rules ,timeWindow)
}

func (qv *QuotaValidator) onManiac(history OrderHistory, newOrder string) error {
	rules := qv.RulesContainer.getRule(MANIAC)
	timeWindow :=qv.RulesContainer.getTiming(MANIAC)
	return qv.validate(history, newOrder, rules ,timeWindow)
}

//generic function to validate quotas
func (qv *QuotaValidator) validate(history OrderHistory, newOrder string,rules map[string]int, timeWindowHrs int) error {
	end := time.Now()
	start := end.Add(-(time.Duration(timeWindowHrs) * time.Hour))

	counter :=0
	minTime := end //represent timestamp of oldest request

	//iterate over element of @newProduct type
	//within given time window and increment counter
	for _,element := range history.Orders{
		if element.Product == newOrder && InTimeSpan(start,end,element.Timestamp){
			minTime = Min(minTime,element.Timestamp)
			if(counter == rules[element.Product]-1){
				timeToWait := GetWaitTime(minTime,end,timeWindowHrs)
				return errors.New(fmt.Sprintf("excedeed limit of available drinks, please wait : [%s]",timeToWait))
			}
			counter = counter + 1
		}	
    }
	return nil
}




