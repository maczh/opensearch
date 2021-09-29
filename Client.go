package opensearch

import (
	"github.com/levigross/grequests"
	"github.com/sadlil/gologger"
	"github.com/xxjwxc/public/tools"
	"net/url"
	"strings"
	"time"
)

var logger = gologger.GetLogger()

type OpenSearchClient struct {
	Conf Config `json:"conf,omitempty"`
}

func NewOpenSearchClient(cf Config) *OpenSearchClient {
	client := new(OpenSearchClient)
	client.Conf = cf
	return client
}

func (client *OpenSearchClient) AddDocument(tableName string, data interface{}) (*OpenSearchCommandResponse, error) {
	uri := strings.ReplaceAll(strings.ReplaceAll(URI_DATA, "$app_name", client.Conf.OS_APPNAME), "$table_name", tableName)
	cmd := OpenSearchCommand{
		Cmd:       CMD_DOC_ADD,
		Fields:    data,
		Timestamp: time.Now().UnixNano(),
	}
	cmds := []OpenSearchCommand{cmd}
	return client.post(uri, cmds)
}

func (client *OpenSearchClient) AddDocuments(tableName string, data []interface{}) (*OpenSearchCommandResponse, error) {
	uri := strings.ReplaceAll(strings.ReplaceAll(URI_DATA, "$app_name", client.Conf.OS_APPNAME), "$table_name", tableName)
	cmds := make([]OpenSearchCommand, 0)
	for _, d := range data {
		cmd := OpenSearchCommand{
			Cmd:       CMD_DOC_ADD,
			Fields:    d,
			Timestamp: time.Now().UnixNano(),
		}
		cmds = append(cmds, cmd)
	}
	return client.post(uri, cmds)
}

func (client *OpenSearchClient) DeleteDocument(tableName, id string) (*OpenSearchCommandResponse, error) {
	uri := strings.ReplaceAll(strings.ReplaceAll(URI_DATA, "$app_name", client.Conf.OS_APPNAME), "$table_name", tableName)
	cmd := OpenSearchCommand{
		Cmd:       CMD_DOC_DEL,
		Fields:    map[string]string{"id": id},
		Timestamp: time.Now().UnixNano(),
	}
	cmds := []OpenSearchCommand{cmd}
	return client.post(uri, cmds)
}

func (client *OpenSearchClient) Search(query *QueryClause, config *ConfigClause, sort *SortClause, fetchFields *FetchFieldsClause) (*OpenSearchResponse, error) {
	searchQuery := NewSearchQuery(query, config, sort, fetchFields)
	uri := strings.ReplaceAll(URI_SEARCH, "$app_name", client.Conf.OS_APPNAME)
	return client.get(uri, searchQuery.ToMap())
}

func (client *OpenSearchClient) post(uri string, commands []OpenSearchCommand) (*OpenSearchCommandResponse, error) {
	body := tools.JSONDecode(commands)
	logger.Debug("阿里云OpenSearch请求参数:" + body)
	resp, err := grequests.Post(client.Conf.OS_HOST+uri, &grequests.RequestOptions{
		Headers: getSignedHeader(client.Conf, "POST", uri, body),
		JSON:    commands,
	})
	if err != nil {
		logger.Error("阿里云OpenSearch连接异常:" + err.Error())
		return nil, err
	}
	logger.Debug("OpenSearch返回结果:" + resp.String())
	var openSearchRsp OpenSearchCommandResponse
	err = resp.JSON(&openSearchRsp)
	if err != nil {
		logger.Error("结果转换错误:" + err.Error())
		return nil, err
	} else {
		return &openSearchRsp, nil
	}
}

func (client *OpenSearchClient) get(uri string, params map[string]string) (*OpenSearchResponse, error) {
	body := ""
	p := url.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	querystr := strings.ReplaceAll(p.Encode(), "+", "%20")
	logger.Debug("阿里云OpenSearch请求参数:" + querystr)
	resp, err := grequests.Get(client.Conf.OS_HOST+uri+"?"+querystr, &grequests.RequestOptions{
		//Params:  params,
		Headers: getSignedHeader(client.Conf, "GET", uri+"?"+querystr, body),
	})
	if err != nil {
		logger.Error("阿里云OpenSearch连接异常:" + err.Error())
		return nil, err
	}
	logger.Debug("OpenSearch返回结果:" + resp.String())
	var openSearchRsp OpenSearchResponse
	err = resp.JSON(&openSearchRsp)
	if err != nil {
		logger.Error("结果转换错误:" + err.Error())
		return nil, err
	} else {
		return &openSearchRsp, nil
	}
}
