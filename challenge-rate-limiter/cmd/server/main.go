package main

import (
	"challenge-rate-limiter/configs"
	"challenge-rate-limiter/internal/infra/storage_adapters"
	"challenge-rate-limiter/internal/infra/webserver"
	"challenge-rate-limiter/internal/infra/webserver/middlewares"
	"net/http"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	webserver := webserver.NewWebServer(configs.WebServerPort)
	storage_adapter, err := storage_adapters.InitRedisAdapter(configs.RedisAddr)
	// storage_adapter, err := storage_adapters.InitMemoryAdapter()
	if err != nil {
		panic(err)
	}

	customTokens := populateCustomTokens()
	rateLimiterConfig := &middlewares.RateLimiterConfig{
		LimitByIP: &middlewares.RateLimiterRateConfig{
			MaxRequestsPerSecond:  configs.LimitByIPMaxRPS,
			BlockTimeMilliseconds: configs.LimitByIPBlockTimeMs,
		},
		LimitByToken: &middlewares.RateLimiterRateConfig{
			MaxRequestsPerSecond:  configs.LimitByTokenMaxRPS,
			BlockTimeMilliseconds: configs.LimitByTokenBlockTimeMs,
		},
		StorageAdapter: storage_adapter,
		CustomTokens:   &customTokens,
	}

	rateLimiter := middlewares.NewRateLimiter(rateLimiterConfig)
	webserver.Use(rateLimiter)

	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	}
	webserver.AddHandler("/", rootHandler, "GET")
	webserver.Start()
}

func populateCustomTokens() map[string]*middlewares.RateLimiterRateConfig {
	return map[string]*middlewares.RateLimiterRateConfig{
		"ABC": {
			MaxRequestsPerSecond:  20,
			BlockTimeMilliseconds: 3000,
		},
		"DEF": {
			MaxRequestsPerSecond:  20,
			BlockTimeMilliseconds: 3000,
		},
	}
}
