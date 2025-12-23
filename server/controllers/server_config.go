package controllers

import "github.com/LuisanaMTDev/spaced_learning/server/database/gosql_queries"

// This struct will contain the server config, as its name says, and the controllers.
type ServerConfig struct {
	DBQueries *gosql_queries.Queries
	Platform  string
}

func NewServerConfig(dbQueries *gosql_queries.Queries, runningPlatform string) *ServerConfig {
	return &ServerConfig{
		DBQueries: dbQueries,
		Platform:  runningPlatform,
	}
}
