package pkg


// this object should be responsible for
// loading file from external resources /xml/json/yaml
type RulesContainer struct {
	ruleMap map[string]map[string]int
	timing map[string]int
}

func (rc *RulesContainer) getRule(rule string) map[string]int{
	return rc.ruleMap[rule]
}

func (rc *RulesContainer) getTiming(rule string) int {
	return rc.timing[rule]
}


func InitRules() (RulesContainer) {
	mm := map[string]map[string]int {
		BASIC: {
			ESPRESSO:3,
			AMERICANO:3,
			CAPPUCCINO :1,
		},
		LOVER: {
			ESPRESSO:5,
			AMERICANO:5,
			CAPPUCCINO :5,
		},
		MANIAC:{
			ESPRESSO:5,
			AMERICANO:5,
			CAPPUCCINO :5,
		},
	}
	
	tt := map[string]int{
		BASIC:24,
		LOVER:24,
		MANIAC:1,
	}
	return RulesContainer{ruleMap: mm,timing:  tt}
}