package responses

import "github.com/aliforever/gista/models"

type FetchQPData struct {
	Response
	RequestStatus       string               `json:"request_status"`
	ExtraInfo           []models.QpExtraInfo `json:"extra_info"`
	QpData              []models.QpData      `json:"qp_data"`
	ClientCacheTtlInSec int                  `json:"client_cache_ttl_in_sec"`
	ErrorMsg            interface{}          `json:"error_msg"`
}
