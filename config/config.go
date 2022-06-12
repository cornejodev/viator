package config

// Config server RESTful API.
type Config struct {
	Port     string // Port address for server
	Database        // Database to be setup
}

// Database config.
type Database struct {
	Engine   string // Engine eg.: "mysql" or "postgres".
	User     string // User of database, eg.: "root".
	Password string // Password of User database
	Name     string // Name of SQL database.
}

func New() (*Config, error) {
	cfg := Config{
		Port: ":8080",
		Database: Database{
			Engine:   "postgres",
			User:     "johndoe",
			Password: "johndoe",
			Name:     "viator",
		},
	}

	return &cfg, nil
}
