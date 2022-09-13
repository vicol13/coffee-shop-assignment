package pkg

import (
	"encoding/json"
	"time"
)

// function which clean history within a given window in Hours
// if we buy coffe each 23 hours user history  won't be evicted from
// redis but we would have stored data about orders which are not in last 24 hours
func CleanWithinWindow(history *OrderHistory, window int) OrderHistory {
	end := time.Now()
	start := end.Add(-(time.Duration(window) * time.Hour))

	filtered := []Order{}

	for _, order := range history.Orders {
		if InTimeSpan(start, end, order.Timestamp) {
			filtered = append(filtered, order)
		}
	}

	return OrderHistory{Orders: filtered}
}

func Min(v1, v2 time.Time) time.Time {
	if v1.Before(v2) {
		return v1
	} else {
		return v2
	}
}

// function which checks if a certain time is in range
func InTimeSpan(start, end, check time.Time) bool {
	if start.Before(end) {
		return !check.Before(start) && !check.After(end)
	}
	if start.Equal(end) {
		return check.Equal(start)
	}
	return !start.After(check) || !end.Before(check)
}

// Compute amount of which user have wait untill next request, function works in the next
// assuming that we get @firstRequest at 00:01 Aug 1 and we receive secondRequest at 09:01 Aug 1 and
// and @timeWindow of 24hrs with one request (quota policy timming)
// we take @firstRequest add @timeWindow to it in order to check when we are allowed to make next
// request (@nextRequestTime) and then we extract difference from @nextRequestTime and @currentRequestTime
// and we get wait time for request
func GetWaitTime(lastRequest, currentRequest time.Time, timeWindow int) time.Duration {
	nextRequestTime := lastRequest.Add(time.Duration(timeWindow) * time.Hour)
	return nextRequestTime.Sub(currentRequest)
}

// converts error to json message
func jsonMessage(er error) string {
	wrapper := ErrorHttpWrapper{Message: er.Error(), Timestamp: time.Now()}
	data, _ := json.Marshal(wrapper)
	return string(data)
}
