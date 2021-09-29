package opensearch

import "fmt"

type ConfigClause struct {
	Start      int    `json:"start"`
	Hit        int    `json:"hit"`
	Format     string `json:"format"`
	RerankSize int    `json:"rerank_size"`
}

func NewDefaultConfigClause() *ConfigClause {
	defaultConfigClause := &ConfigClause{
		Start:      0,
		Hit:        10,
		Format:     "json",
		RerankSize: 200,
	}
	return defaultConfigClause
}

func NewConfigClause(start, hit int, format string, rerankSize int) *ConfigClause {
	defaultConfigClause := &ConfigClause{
		Start:      start,
		Hit:        hit,
		Format:     format,
		RerankSize: rerankSize,
	}
	return defaultConfigClause
}

func (c *ConfigClause) SetStart(start int) *ConfigClause {
	c.Start = start
	return c
}

func (c *ConfigClause) SetHit(hit int) *ConfigClause {
	c.Hit = hit
	return c
}

func (c *ConfigClause) SetFormat(format string) *ConfigClause {
	c.Format = format
	return c
}

func (c *ConfigClause) SetRerankSize(rerankSize int) *ConfigClause {
	c.RerankSize = rerankSize
	return c
}

func (c *ConfigClause) String() string {
	clause := fmt.Sprintf("start:%d,hit:%d,format:%s", c.Start, c.Hit, c.Format)
	if c.RerankSize > 0 {
		clause = fmt.Sprintf("%s,rerank_size:%d", clause, c.RerankSize)
	}
	return clause
}
