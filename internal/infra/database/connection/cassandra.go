package connection

import (
	"time"

	"github.com/gocql/gocql"
	configs "github.com/sandronister/cassandra-go/config"
)

func NewCassandraConnection(conf *configs.Conf) (*gocql.Session, error) {
	cluster := gocql.NewCluster(conf.Ips...)
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 5
	session, err := cluster.CreateSession()

	if err != nil {
		return nil, err
	}

	_, err = session.Query("Select * from system.local").Iter().SliceMap()

	if err != nil {
		return nil, err
	}

	return session, nil
}
