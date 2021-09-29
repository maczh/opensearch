package opensearch

type SortClause []string

func NewSortClause() *SortClause {
	sc := make(SortClause, 0)
	return &sc
}

func (sc *SortClause) Asc(field string) *SortClause {
	if len(*sc) == 0 {
		*sc = append(*sc, "+"+field)
	} else {
		*sc = append(*sc, ";+"+field)
	}
	return sc
}

func (sc *SortClause) Dsc(field string) *SortClause {
	if len(*sc) == 0 {
		*sc = append(*sc, "-"+field)
	} else {
		*sc = append(*sc, ";-"+field)
	}
	return sc
}

func (sc *SortClause) String() string {
	clause := ""
	for _, str := range *sc {
		clause = clause + str
	}
	return clause
}
