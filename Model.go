package opensearch

type OpenSearchResponse struct {
	Errors []struct {
		Code    int    `json:"code" bson:"code"`
		Message string `json:"message" bson:"message"`
	} `json:"errors" bson:"errors"`
	OpsRequestMisc string `json:"ops_request_misc" bson:"ops_request_misc"`
	Qp             []struct {
		AppName             string `json:"app_name" bson:"app_name"`
		QueryCorrectionInfo []struct {
			CorrectedQuery  string `json:"corrected_query" bson:"corrected_query"`
			CorrectionLevel int    `json:"correction_level" bson:"correction_level"`
			Index           string `json:"index" bson:"index"`
			OriginalQuery   string `json:"original_query" bson:"original_query"`
			ProcessorName   string `json:"processor_name" bson:"processor_name"`
		} `json:"query_correction_info" bson:"query_correction_info"`
	} `json:"qp" bson:"qp"`
	RequestID string `json:"request_id" bson:"request_id"`
	Result    struct {
		ComputeCost []struct {
			IndexName string  `json:"index_name" bson:"index_name"`
			Value     float64 `json:"value" bson:"value"`
		} `json:"compute_cost" bson:"compute_cost"`
		Facet      []interface{} `json:"facet" bson:"facet"`
		Items      []interface{} `json:"items" bson:"items"`
		Num        int           `json:"num" bson:"num"`
		Searchtime float64       `json:"searchtime" bson:"searchtime"`
		Total      int           `json:"total" bson:"total"`
		Viewtotal  int           `json:"viewtotal" bson:"viewtotal"`
	} `json:"result" bson:"result"`
	Status string `json:"status" bson:"status"`
	Tracer string `json:"tracer" bson:"tracer"`
}

type OpenSearchCommand struct {
	Cmd       string      `json:"cmd" bson:"cmd"`
	Fields    interface{} `json:"fields" bson:"fields"`
	Timestamp int64       `json:"timestamp" bson:"timestamp"`
}

type OpenSearchCommandResponse struct {
	Errors []struct {
		Code    int    `json:"code" bson:"code"`
		Message string `json:"message" bson:"message"`
		Params  struct {
			FriendlyMessage string `json:"friendly_message" bson:"friendly_message"`
		} `json:"params" bson:"params"`
	} `json:"errors" bson:"errors"`
	RequestID string `json:"request_id" bson:"request_id"`
	Result    bool   `json:"result" bson:"result"`
	Status    string `json:"status" bson:"status"`
}
