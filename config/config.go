package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration.
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"
	DbUser = "db.user"
	DbPass = "db.pass"

	ServerHost = "server.host"
	ServerPort = "server.port"

	GCPProjectID = "gcp.project.id"
	GCPCredFile  = "gcp.cred.file"

	DataFactory = "data.factory"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbPort, "DB_PORT")
	_ = viper.BindEnv(DbUser, "DB_USER")
	_ = viper.BindEnv(DbPass, "DB_PASS")

	// env var for server
	_ = viper.BindEnv(ServerHost, "SERVER_HOST")
	_ = viper.BindEnv(ServerPort, "SERVER_PORT")

	// env var for gcp
	_ = viper.BindEnv(GCPProjectID, "GCP_PROJECT_ID")
	_ = viper.BindEnv(GCPCredFile, "GCP_CRED_FILE")

	_ = viper.BindEnv(DataFactory, "DATA_FACTORY")

	// defaults
	viper.SetDefault(DbName, "inventory")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")

	viper.SetDefault(ServerHost, "0.0.0.0")
	viper.SetDefault(ServerPort, "8080")

	viper.SetDefault(GCPProjectID, "inventory-backend-a8ec5")
	viper.SetDefault(GCPCredFile, "../cred.json")

	viper.SetDefault(DataFactory, "mongo")
}
