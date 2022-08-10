package options_util

type Options struct {
	MySqlUrl   string
	MongoDBUrl string
	Port       string // port to listening
	APIPrefix  string

	// jwtSecretKey for user token generation and validation
	JwtSecretKey string
}