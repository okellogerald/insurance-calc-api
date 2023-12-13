package plans

type Plan struct {
	ID                       uint        `json:"id"`
	Name                     string      `json:"name"`
	SumAssuredDependentTerms []SAOptions `json:"sum_assured_dependent_terms"`
	AvailableTerms           []uint      `json:"available_terms"`
}

type SAOptions struct {
	SumAssured uint64 `json:"sum_assured"`
	Terms      []uint `json:"terms"`
}

func (p Plan) DependentOnSumAssured() bool {
	return len(p.AvailableTerms) == 0 && len(p.SumAssuredDependentTerms) > 0
}
