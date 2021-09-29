package opensearch

const (
	URI_DATA       = "/v3/openapi/apps/$app_name/$table_name/actions/bulk"
	URI_SEARCH     = "/v3/openapi/apps/$app_name/search"
	CMD_DOC_ADD    = "add"
	CMD_DOC_DEL    = "delete"
	CMD_DOC_UPDATE = "update"
)

type Config struct {
	OS_ACCESS_KEY string
	OS_SECRET_KEY string
	OS_HOST       string
	OS_APPNAME    string
}
