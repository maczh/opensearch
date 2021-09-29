package opensearch

import "fmt"

type QueryClause []string

func NewQueryClause() *QueryClause {
	qc := make(QueryClause, 0)
	return &qc
}

func (qc *QueryClause) LeftParenthesis() *QueryClause {
	*qc = append(*qc, "(")
	return qc
}

func (qc *QueryClause) RightParenthesis() *QueryClause {
	*qc = append(*qc, ")")
	return qc
}

func (qc *QueryClause) FirstQuery(field, keyword string) *QueryClause {
	if field == "" {
		field = "default"
	}
	kv := fmt.Sprintf("%s:'%s'", field, keyword)
	*qc = append(*qc, kv)
	return qc
}

func (qc *QueryClause) And(field, keyword string) *QueryClause {
	if field == "" {
		field = "default"
	}
	kv := fmt.Sprintf(" AND %s:'%s'", field, keyword)
	*qc = append(*qc, kv)
	return qc
}

func (qc *QueryClause) Or(field, keyword string) *QueryClause {
	if field == "" {
		field = "default"
	}
	kv := fmt.Sprintf(" OR %s:'%s'", field, keyword)
	*qc = append(*qc, kv)
	return qc
}

func (qc *QueryClause) AndNot(field, keyword string) *QueryClause {
	if field == "" {
		field = "default"
	}
	kv := fmt.Sprintf(" ANDNOT %s:'%s'", field, keyword)
	*qc = append(*qc, kv)
	return qc
}

func (qc *QueryClause) Rank(field, keyword string) *QueryClause {
	if field == "" {
		field = "default"
	}
	kv := fmt.Sprintf(" RANK %s:'%s'", field, keyword)
	*qc = append(*qc, kv)
	return qc
}

func (qc *QueryClause) String() string {
	clause := ""
	for _, str := range *qc {
		clause = clause + str
	}
	return clause
}
