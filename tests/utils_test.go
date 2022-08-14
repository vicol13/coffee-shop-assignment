package tests

import (
	"aaha/pkg"
	"fmt"
	"testing"
	"time"
)


func TestMin(t *testing.T) {
	// given
	now:=time.Now()
	expected := now.Add(time.Duration(-24*time.Hour))
	
	//when
	newValue := pkg.Min(now, expected)

	//then
	if newValue != expected {
		t.Errorf("expected '%s' but got '%s'", expected.String(), newValue.String())
	}

}


func TestPositiveInSpan(t *testing.T){
	//given
	end := time.Now()
	start := end.Add(time.Duration(-24*time.Hour))
	pivot := end.Add(time.Duration(-6*time.Hour))

	//when 
	response := pkg.InTimeSpan(start,end,pivot)
	
	//then
	if response != true {
		t.Errorf("expected [true] but got [false] as pivot is in range")
	}
}


func TestNegativeeInSpan(t *testing.T){
	//given
	end := time.Now()
	start := end.Add(time.Duration(-24*time.Hour))
	pivot := end.Add(time.Duration(-30*time.Hour))

	//when 
	response := pkg.InTimeSpan(start,end,pivot)
	
	//then
	if response != false {
		t.Errorf("expected [false] but got [true] as pivot not in range")
	}
}



func TestCleanWithinWindowWithNoOldOrders(t *testing.T) {
	//given
	tmg := time.Now()
	order1 := pkg.Order{Product: "americano",Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }
	order2 := pkg.Order{Product: "espresso",Timestamp: tmg.Add(time.Duration(-5*time.Hour)) }
	order3 := pkg.Order{Product: "cappuccino",Timestamp: tmg.Add(time.Duration(-7*time.Hour)) }
	orderHistory := pkg.OrderHistory{Orders: []pkg.Order{order1,order2,order3}}

	//when
	newHistory := pkg.CleanWithinWindow(&orderHistory,10)

	//then
	if len(newHistory.Orders) != len(orderHistory.Orders){
		t.Errorf("Size of [newHistory] doesn't match with [orderHistory], is expected to be same size")
	}

}


func TestCleanWithinWindowWithOldOrders(t *testing.T) {
	//given
	tmg := time.Now()
	order1 := pkg.Order{Product: "americano",Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }
	order2 := pkg.Order{Product: "espresso",Timestamp: tmg.Add(time.Duration(-5*time.Hour)) }
	order3 := pkg.Order{Product: "cappuccino",Timestamp: tmg.Add(time.Duration(-7*time.Hour)) }
	orderHistory := pkg.OrderHistory{Orders: []pkg.Order{order1,order2,order3}}

	//when
	newHistory := pkg.CleanWithinWindow(&orderHistory,6)

	//then
	if len(newHistory.Orders) != 2{
		t.Errorf(fmt.Sprintf("Size of parse [newHistory] is [%d], is expected to be 2",len(newHistory.Orders)))
	}

}


func TestGetWaitTime(t *testing.T){
	//given
	expectedDuration,_ := time.ParseDuration("1h")
	now := time.Now()
	then := now.Add(-23 * time.Hour)
	//when
	waitTime := pkg.GetWaitTime(then,now,24)

	//then
	if waitTime != expectedDuration {
		t.Errorf(fmt.Sprintf("Expcted duration is [%s] but got [%s] ", expectedDuration.String(),waitTime.String()))
	}

}