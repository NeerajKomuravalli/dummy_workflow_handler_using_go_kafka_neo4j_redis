package neo4jmanager

import (
	log "github.com/sirupsen/logrus"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jClient struct {
	Driver neo4j.Driver
}

func NewNeo4jClient(uri, username, password string) *Neo4jClient {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		log.Panic(err)
	}
	return &Neo4jClient{
		driver,
	}
}

func (neo4jClient *Neo4jClient) CloseConnection() {
	neo4jClient.Driver.Close()
}
