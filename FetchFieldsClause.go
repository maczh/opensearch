package opensearch

type FetchFieldsClause []string

func NewFetchFieldsClause() *FetchFieldsClause {
	ff := make(FetchFieldsClause, 0)
	return &ff
}

func (ff *FetchFieldsClause) AddFieldName(field string) *FetchFieldsClause {
	if len(*ff) == 0 {
		*ff = append(*ff, field)
	} else {
		*ff = append(*ff, ";"+field)
	}
	return ff
}

func (ff *FetchFieldsClause) String() string {
	clause := ""
	for _, str := range *ff {
		clause = clause + str
	}
	return clause
}
