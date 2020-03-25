package cassandra

import (
	"github.com/gocql/gocql"
	"os"
)

const (
	cassandra_username = "cassandra_username"
	cassandra_password = "cassandra_password"
)


var (
	session *gocql.Session
	username = os.Getenv(cassandra_username)
	password = os.Getenv(cassandra_password)
)

func init() {
	//Todo move to environment variable
	cluster := gocql.NewCluster("192.168.99.100")
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: username,
		Password: password,
	}
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}