package opensearch

type SearchQuery struct {
	Query       string `json:"query,omitempty" xml:"query,omitempty" require:"true"`
	Config      string `json:"config,omitempty" xml:"config,omitempty"`
	Sort        string `json:"sort,omitempty" xml:"sort,omitempty"`
	FetchFields string `json:"fetch_fields,omitempty"`
}

func NewSearchQuery(query *QueryClause, config *ConfigClause, sort *SortClause, fetchFields *FetchFieldsClause) *SearchQuery {
	searchQuery := new(SearchQuery)
	if query != nil {
		searchQuery.Query = query.String()
	}
	if config != nil {
		searchQuery.Config = config.String()
	}
	if sort != nil {
		searchQuery.Sort = sort.String()
	}
	if fetchFields != nil {
		searchQuery.FetchFields = fetchFields.String()
	}
	return searchQuery
}

func (s *SearchQuery) SetQuery(v string) *SearchQuery {
	s.Query = v
	return s
}

func (s *SearchQuery) SetConfig(v string) *SearchQuery {
	s.Config = v
	return s
}

func (s *SearchQuery) SetSort(v string) *SearchQuery {
	s.Sort = v
	return s
}

func (s *SearchQuery) ToMap() map[string]string {
	m := make(map[string]string)
	if s.FetchFields != "" {
		m["fetch_fields"] = s.FetchFields
	}
	query := "query=" + s.Query
	if s.Config != "" {
		query = query + "&&config=" + s.Config
	}
	if s.Sort != "" {
		query = query + "&&sort=" + s.Sort
	}
	m["query"] = query
	return m
}
