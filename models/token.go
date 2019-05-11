package models

import (
	"time"

	"github.com/aliforever/gista/models/token"
)

type Token struct {
	CarrierName                 string        `json:"carrier_name"`
	CarrierId                   int           `json:"carrier_id"`
	Ttl                         int64         `json:"ttl"`
	Features                    interface{}   `json:"features"`
	RequestTime                 string        `json:"request_time"`
	TokenHash                   string        `json:"token_hash"`
	RewriteRules                []RewriteRule `json:"rewrite_rules"`
	EnabledWalletDefsKeys       interface{}   `json:"enabled_wallet_defs_keys"`
	Deadline                    string        `json:"deadline"`
	ZeroCmsFetchIntervalSeconds int           `json:"zero_cms_fetch_interval_seconds"`
}

func (t *Token) ExpiresAt() int64 {
	ttl := t.Ttl
	if ttl == 0 {
		ttl = token.DefaultTTL
	}
	return time.Now().Unix() + ttl
}
