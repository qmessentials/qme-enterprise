package lib

const PostgresDefaultPort = 5432

type PostgresConfig struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

func NewPostgresConfig() (*PostgresConfig, error) {
	port, err := getEnvInt("POSTGRES_PORT", PostgresDefaultPort)
	if err != nil {
		return nil, err
	}
	user, err := getEnvStringRequired("POSTGRES_USER")
	if err != nil {
		return nil, err
	}
	result := &PostgresConfig{
		DBHost:     getEnvString("POSTGRES_HOST", "localhost"),
		DBPort:     port,
		DBName:     getEnvString("POSTGRES_DATABASE", "auth"),
		DBUser:     user,
		DBPassword: getEnvString("POSTGRES_PASSWORD", ""),
	}
	return result, nil
}
