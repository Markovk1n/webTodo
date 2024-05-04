package models

type AppConfigs struct {
	Postgres Postgres
	Token    string
}
type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}
