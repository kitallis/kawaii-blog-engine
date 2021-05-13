package jwt

import (
	"fmt"
	"log"
)

//var DenyList *memory.Storage
//
//func InitDenyList() {
//	DenyList = memory.New(memory.Config{GCInterval: 10 * time.Hour})
//	fmt.Println("👍🏽 Initialized JWT Deny List")
//}

//var DenyList *redis.Storage
//
//func InitDenyList() {
//	DenyList = redis.New()
//	fmt.Println("👍🏽 Initialized JWT Deny List")
//}

var DenyList []string

func InitDenyList() {
	fmt.Println("👍🏽 Initialized JWT Deny List")
}

func Set(key string) {
	log.Printf("Revoking... %s", key)
	DenyList = append(DenyList, key)
}

func Get(key string) (bool, error) {
	log.Printf("Getting... %s", key)
	log.Println(DenyList)
	return contains(DenyList, key), nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}
