package internal

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

var Session *gocql.Session

func init() {
	err := godotenv.Load("cassandra.env")
	if err != nil {
		log.Fatal(err)
	}

	user := os.Getenv("CASUSERNAME")
	pass := os.Getenv("CASPASSWORD")
	connstring := os.Getenv("CONNSTRING")
	cluster := gocql.NewCluster(strings.Split(connstring, ",")...)
	cluster.Keyspace = "restfulapi"
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: user,
		Password: pass,
	}

	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	_ = Session
	fmt.Println("cassandra well initialized")
}
