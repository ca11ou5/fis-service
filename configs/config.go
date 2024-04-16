package configs

type Config struct {
	Port     string `env:"PORT"`
	LogLevel string `env:"LOG_LEVEL"`

	GraylogUrl string `env:"GRAYLOG_URL"`
	MongoUrl   string `env:"MONGO_URL"`

	DadataBaseUrl    string `env:"DADATA_BASE_URL"`
	DadataAddressUrl string `env:"DADATA_ADDRESS_URL"`

	OpenApiBaseUrl    string `env:"OPENAPI_BASE_URL"`
	OpenApiScoringUrl string `env:"OPENAPI_SCORING_URL"`
	OpenApiSecretKey  string `env:"OPENAPI_SECRET_KEY"`
	OpenApiClientID   string `env:"OPENAPI_CLIENT_ID"`
}
