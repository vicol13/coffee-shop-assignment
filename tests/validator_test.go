package tests

import (
	"aaha/pkg" 
	"testing"
	"time"
)

// 
// Testing Basic Quota with different parametere
//
func TestBasicQuotaWithoutLimit(t *testing.T){
	//given 
	rules := pkg.InitRules()
	validator := pkg.QuotaValidator{RulesContainer: &rules}

	tmg := time.Now()
	order1 := pkg.Order{Product: pkg.Americano ,Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }
	order2 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-5*time.Hour)) }
	orderHistory := pkg.OrderHistory{Orders: []pkg.Order{order1,order2}}

	//when
	americanoError := validator.Handle(pkg.Basic,orderHistory,pkg.Americano)
	espressoError := validator.Handle(pkg.Basic,orderHistory,pkg.Espresso )
	cappuccinoError := validator.Handle(pkg.Basic,orderHistory,pkg.Cappuccino)


	//then
	if americanoError != nil {
		t.Errorf("[americanoError] is expected to be nil with given context")
	}

	if espressoError != nil {
		t.Errorf("[espressoError] is expected to be nil with given context")
	}

	if cappuccinoError != nil {
		t.Errorf("[cappuccinoError] is expected to be nil with given context")
	}

}


func TestBasicQuotaWithCompleteLimitation(t *testing.T){
	//given 
	rules := pkg.InitRules()
	validator := pkg.QuotaValidator{RulesContainer: &rules}

	tmg := time.Now()
	order1 := pkg.Order{Product: pkg.Americano ,Timestamp: tmg.Add(time.Duration(-1*time.Hour)) }
	order2 := pkg.Order{Product: pkg.Americano ,Timestamp: tmg.Add(time.Duration(-2*time.Hour)) }
	order3 := pkg.Order{Product: pkg.Americano ,Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }
	
	order4 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-4*time.Hour)) }
	order5 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-5*time.Hour)) }
	order6 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-6*time.Hour)) }

	order7 := pkg.Order{Product: pkg.Cappuccino ,Timestamp: tmg.Add(time.Duration(-7*time.Hour)) }
	
	orderHistory := pkg.OrderHistory{Orders: []pkg.Order{order1,order2,order3,order4,order5,order6,order7}}

	//when
	americanoError := validator.Handle(pkg.Basic,orderHistory,pkg.Americano )
	espressoError := validator.Handle(pkg.Basic,orderHistory,pkg.Espresso)
	cappuccinoError := validator.Handle(pkg.Basic,orderHistory,pkg.Cappuccino)

	//then
	if americanoError == nil {
		t.Errorf("[americanoError] is expected to be NOT nil with given context")
	}

	if espressoError == nil {
		t.Errorf("[espressoError] is expected to be NOT nil with given context")
	}

	if cappuccinoError == nil {
		t.Errorf("[cappuccinoError] is expected to be NOT nil with given context")
	}

}



func TestBasicQuotaWithPartialLimitation(t *testing.T){
	//given 
	rules := pkg.InitRules()
	validator := pkg.QuotaValidator{RulesContainer: &rules}

	tmg := time.Now()
	order1 := pkg.Order{Product: pkg.Americano,Timestamp: tmg.Add(time.Duration(-1*time.Hour)) }
	order3 := pkg.Order{Product: pkg.Americano,Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }
	
	order4 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-4*time.Hour)) }
	order5 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-5*time.Hour)) }
	order6 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-6*time.Hour)) }

	
	orderHistory := pkg.OrderHistory{Orders: []pkg.Order{order1,order3,order4,order5,order6}}

	//then
	americanoError := validator.Handle(pkg.Basic,orderHistory,pkg.Americano)
	espressoError := validator.Handle(pkg.Basic,orderHistory,pkg.Espresso)
	cappuccinoError := validator.Handle(pkg.Basic,orderHistory,pkg.Cappuccino)


	//when
	if americanoError != nil {
		t.Errorf("[americanoError] is expected to be nil with given context")
	}

	if espressoError == nil {
		t.Errorf("[espressoError] is expected to be NOT nil with given context")
	}

	if cappuccinoError != nil {
		t.Errorf("[cappuccinoError] is expected to be NOT nil with given context")
	}

}


func TestBasicQuotaWithPartialLimitationAndOldOrders(t *testing.T){
	//given 
	rules := pkg.InitRules()
	validator := pkg.QuotaValidator{RulesContainer: &rules}

	tmg := time.Now()
	order1 := pkg.Order{Product: pkg.Americano,Timestamp: tmg.Add(time.Duration(-1*time.Hour)) }
	order2 := pkg.Order{Product: pkg.Americano,Timestamp: tmg.Add(time.Duration(-3*time.Hour)) }
	order3 := pkg.Order{Product: pkg.Americano,Timestamp: tmg.Add(time.Duration(-25*time.Hour)) }
	
	order4 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-4*time.Hour)) }
	order5 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-5*time.Hour)) }
	order6 := pkg.Order{Product: pkg.Espresso,Timestamp: tmg.Add(time.Duration(-6*time.Hour)) }

	order7 := pkg.Order{Product: pkg.Cappuccino,Timestamp: tmg.Add(time.Duration(-26*time.Hour)) }

	orderHistory := pkg.OrderHistory{Orders: []pkg.Order{order1,order2,order3,order4,order5,order6,order7}}

	//then
	americanoError := validator.Handle(pkg.Basic,orderHistory,pkg.Americano)
	espressoError := validator.Handle(pkg.Basic,orderHistory,pkg.Espresso)
	cappuccinoError := validator.Handle(pkg.Basic,orderHistory,pkg.Cappuccino)


	//when
	if americanoError != nil {
		t.Errorf("[americanoError] is expected to be nil with given context")
	}

	if espressoError == nil {
		t.Errorf("[espressoError] is expected to be NOT nil with given context")
	}

	if cappuccinoError != nil {
		t.Errorf("[cappuccinoError] is expected to be  nil with given context")
	}

}



//
//	Tests for next quotas are ommited as they are similar to this one
//


