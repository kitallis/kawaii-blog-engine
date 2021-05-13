package jwt

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

var DenyList *cache.Cache

func InitDenyList() {
	DenyList = cache.New(60 * time.Minute, 10 * time.Minute)
	fmt.Println("👍🏽 Initialized JWT Deny List")
}