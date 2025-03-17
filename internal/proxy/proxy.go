package proxy

import (
	"caching-proxy/internal/cache"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type ProxyObject struct {
	Cache  map[string]*cache.CacheObject
	Mutex  sync.RWMutex
	Origin string
}

func GetNewProxy(origin string) *ProxyObject {
	return &ProxyObject{
		Cache:  make(map[string]*cache.CacheObject),
		Origin: origin,
	}
}

func (p *ProxyObject) isCached(key string) bool {
	_, ok := p.Cache[key]
	return ok
}

func (p *ProxyObject) cacheValue(key string, value *cache.CacheObject) {
	p.Cache[key] = value
}

func (p *ProxyObject) getValue(key string) *cache.CacheObject {
	val := p.Cache[key]
	return val
}

func (p *ProxyObject) RequestHandler(writer http.ResponseWriter, req *http.Request) {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", p.Origin+req.URL.String(), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	cacheKey := req.URL.String()

	if p.isCached(cacheKey) {
		fmt.Println("Cache hit :)")
		writer.Header().Set("X-Cache", "HIT")
	} else {
		fmt.Println("Cache miss :(")
		writer.Header().Set("X-Cache", "MISS")
		p.cacheValue(cacheKey, &cache.CacheObject{Content: body, CreatedAt: time.Now()})
	}
	cacheResult := p.getValue(cacheKey)

	_, err = writer.Write(cacheResult.Content)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
