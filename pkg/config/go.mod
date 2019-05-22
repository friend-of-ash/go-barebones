module config

go 1.12

require (
	github.com/joho/godotenv v1.3.0
	github.com/kelseyhightower/envconfig v1.3.0

	pkg/log v1.0.0
)

replace pkg/log => ../log
