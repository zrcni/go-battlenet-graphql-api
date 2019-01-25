package api

type ElasticSearchMountHit struct {
	Index  string                 `mapstructure:"_index"`
	Type   string                 `mapstructure:"_type"`
	ID     string                 `mapstructure:"_id"`
	Source map[string]interface{} `mapstructure:"_source"`
}
