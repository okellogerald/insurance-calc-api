package plans

var Plans [6]Plan

func Init() {
	var sumAssuredDependentTerms [7]SAOptions
	sumAssuredDependentTerms[0] = SAOptions{SumAssured: 5000000, Terms: []uint{12, 15}}
	sumAssuredDependentTerms[1] = SAOptions{SumAssured: 7500000, Terms: []uint{12, 15}}
	sumAssuredDependentTerms[2] = SAOptions{SumAssured: 10000000, Terms: []uint{10, 12, 15}}
	sumAssuredDependentTerms[3] = SAOptions{SumAssured: 12500000, Terms: []uint{10, 12, 15}}
	sumAssuredDependentTerms[4] = SAOptions{SumAssured: 15000000, Terms: []uint{10, 12, 15}}
	sumAssuredDependentTerms[5] = SAOptions{SumAssured: 17500000, Terms: []uint{10, 12, 15}}
	sumAssuredDependentTerms[6] = SAOptions{SumAssured: 20000000, Terms: []uint{10, 12, 15}}

	Plans[0] = Plan{ID: 0, Name: "Life Plus Plan - With Cashback", AvailableTerms: []uint{10, 12, 15}}
	Plans[1] = Plan{ID: 1, Name: "Life Plus Plan - No Cashback", AvailableTerms: []uint{10, 12, 15}}
	Plans[2] = Plan{ID: 2, Name: "Life Plan - With Cashback", SumAssuredDependentTerms: sumAssuredDependentTerms[:]}
	Plans[3] = Plan{ID: 3, Name: "Life Plan - No Cashback", SumAssuredDependentTerms: sumAssuredDependentTerms[:]}
	Plans[4] = Plan{ID: 4, Name: "Education Plan - With Cashback", SumAssuredDependentTerms: sumAssuredDependentTerms[:]}
	Plans[5] = Plan{ID: 5, Name: "Education Plan - No Cashback", SumAssuredDependentTerms: sumAssuredDependentTerms[:]}
}
