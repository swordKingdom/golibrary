package api

import (
	"github.com/globalsign/mgo"
)

type MongoSession struct {
	*mgo.Session
	conf *MgoConf
}

func (s *MongoSession) Find() {
	s.DB(s.conf.Database)
}

func NewMongoDBSessionWithConf(conf *MgoConf) (*MongoSession, error) {
	cloneConf := conf.Cype()
	dialInfo := &mgo.DialInfo{
		Username:      cloneConf.Username,
		Password:      cloneConf.Password,
		Addrs:         cloneConf.Addrs,
		Database:      cloneConf.Database,
		Timeout:       cloneConf.Timeout,
		MaxIdleTimeMS: cloneConf.MaxIdleTimeMS,
		PoolLimit:     cloneConf.PoolLimit,
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}
	return &MongoSession{session, cloneConf}, nil
}
