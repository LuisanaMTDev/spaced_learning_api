package helpers

import "github.com/LuisanaMTDev/spaced_learning/server/database/gosql_queries"

type ServerConfig struct {
	DBQueries *gosql_queries.Queries
	Platform  string
}
