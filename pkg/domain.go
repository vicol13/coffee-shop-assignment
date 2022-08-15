package pkg

import "time"

type ProductRules struct {
	Limit         int
	TimeWindowHrs int
}
type ErrorHttpWrapper struct {
	Message   string
	Timestamp time.Time
}

type HttpResponse struct {
	Message string
}

type Order struct {
	Product   string
	Timestamp time.Time
}

type OrderHistory struct {
	Orders []Order
}

func (r *OrderHistory) add(order Order) {
	r.Orders = append(r.Orders, order)
}
