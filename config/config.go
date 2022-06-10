package config

// Config server RESTful API.
type Config struct {
	// Port for address server, if is empty by default will are 80.
	Port string

	Database
}

// Database config.
type Database struct {
	// Engine eg.: "mysql" or "postgres".
	Engine string

	// User of database, eg.: "root".
	User string

	// Password of User database
	Password string

	// Name of SQL database.
	Name string
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
