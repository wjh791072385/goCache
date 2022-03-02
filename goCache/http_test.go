package goCache

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

var dbHttp = map[string]string{
	"hello": "1132",
	"mike":  "world",
}

func TestHttp(t *testing.T) {
	//本机创建group
	NewGroup("scores", 2<<10, GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDB] search key", key)
			if v, ok := dbHttp[key]; ok {
				return []byte(v), nil
			}
			return nil, fmt.Errorf("%s not exist", key)
		}))

	addr := "localhost:8888"
	//peers实现了serveHTTP接口
	peers := NewHTTPPool(addr)
	http.ListenAndServe(addr, peers)

}
