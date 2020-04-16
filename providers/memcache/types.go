package memcache

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/fasthttp/session/v2"
)

// Config configuration of provider
type Config struct {
	// ServerList memcache server list
	ServerList []string

	// MaxIdleConns specifies the maximum number of idle connections that will
	// be maintained per address. If less than one, DefaultMaxIdleConns will be
	// used.
	//
	// Consider your expected traffic rates and latency carefully. This should
	// be set to a number higher than your peak parallel requests.
	MaxIdleConns int

	// KeyPrefix sessionID as memcache key prefix
	KeyPrefix string

	// SerializeFunc session value serialize func
	SerializeFunc func(src session.Dict) ([]byte, error)

	// UnSerializeFunc session value unSerialize func
	UnSerializeFunc func(dst *session.Dict, src []byte) error
}

// Provider backend manager
type Provider struct {
	config Config
	db     *memcache.Client
}