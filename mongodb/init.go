package mongodb

import (
	"fmt"
	"sync"

	"golibrary/mongodb/api"
)

var mgoConnMap = make(map[string]*api.MongoSession)

func Mongo(key string) (*api.MongoSession, error) {
	if conn, ok := mgoConnMap[key]; ok {
		return conn, nil
	} else {
		return nil, fmt.Errorf("the db %v is not exist ", key)
	}
}

var once sync.Once

func initMongoDB(){

}

func Init() {
	once.Do(initMongoDB)
}
