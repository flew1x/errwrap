package errwrap

import (
	"sync"
)

type config struct {
	domain string
}

var (
	globalCfg     = &config{domain: "service"} // default
	globalCfgLock sync.RWMutex
)

// Option function-config
type Option func(*config)

// WithDomain set domain
func WithDomain(domain string) Option {
	return func(c *config) {
		c.domain = domain
	}
}

// Configure set config
func Configure(opts ...Option) {
	globalCfgLock.Lock()
	defer globalCfgLock.Unlock()

	for _, opt := range opts {
		opt(globalCfg)
	}
}

// internal getter
func getDomain() string {
	globalCfgLock.RLock()
	defer globalCfgLock.RUnlock()

	return globalCfg.domain
}
