package pkg

// this object should be responsible for
// loading file from external resources /xml/json/yaml
type RulesContainer struct {
	//nested map with limits and timing for each quouta
	ruleMap map[string]map[string]ProductRules
}

// function which will return product rules under given quota
func (rc *RulesContainer) GetProductRules(quota string, product string) ProductRules {
	return rc.ruleMap[quota][product]
}

func InitRules() RulesContainer {

	mm := map[string]map[string]ProductRules{
		Basic: {
			Espresso:   ProductRules{Limit: 3, TimeWindowHrs: 24},
			Americano:  ProductRules{Limit: 3, TimeWindowHrs: 24},
			Cappuccino: ProductRules{Limit: 1, TimeWindowHrs: 24},
		},
		Lover: {
			Espresso:   ProductRules{Limit: 5, TimeWindowHrs: 24},
			Americano:  ProductRules{Limit: 5, TimeWindowHrs: 24},
			Cappuccino: ProductRules{Limit: 5, TimeWindowHrs: 24},
		},
		Maniac: {
			Espresso:   ProductRules{Limit: 5, TimeWindowHrs: 24},
			Americano:  ProductRules{Limit: 3, TimeWindowHrs: 24},
			Cappuccino: ProductRules{Limit: 1, TimeWindowHrs: 24},
		},
	}

	return RulesContainer{ruleMap: mm}
}
