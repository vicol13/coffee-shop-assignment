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
func (qv *QuotaValidator) Handle(quota string, history OrderHistory, newOrder string) error {
	//todo: move to map number of products increase
	switch quota {
	case Basic:
		return qv.validate(history, newOrder, qv.RulesContainer.GetProductRules(Basic, newOrder))
	case Maniac:
		return qv.validate(history, newOrder, qv.RulesContainer.GetProductRules(Maniac, newOrder))
	case Lover:
		return qv.validate(history, newOrder, qv.RulesContainer.GetProductRules(Lover, newOrder))
	}

	return errors.New("no such quota: " + quota)
}

// generic function to validate the quota constraints for certain item
// @history -> previous aquired by user
// @newOrder -> desired new item
// @limit -> how many items is user is allowed to get
// @timeWindowHrs -> in which time window user is allowed to get @limit
func (qv *QuotaValidator) validate(history OrderHistory, newOrder string, rules ProductRules) error {
	end := time.Now()
	start := end.Add(-(time.Duration(rules.TimeWindowHrs) * time.Hour))

	counter := 0
	minTime := end //represent timestamp of oldest request

	//iterate over elements (history) of filter @newOrder type
	//within given time window and increment counter
	for _, element := range history.Orders {
		if element.Product == newOrder && InTimeSpan(start, end, element.Timestamp) {
			minTime = Min(minTime, element.Timestamp)
			if counter == rules.Limit-1 {
				timeToWait := GetWaitTime(minTime, end, rules.TimeWindowHrs)
				return errors.New(fmt.Sprintf("excedeed limit of available drinks, please wait : [%s]", timeToWait))
			}
			counter = counter + 1
		}
	}
	return nil
}
