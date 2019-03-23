package testapi

type Config struct {
	User		string `toml:"user"`
	Password	string `toml:"password"`
	DBName		string `toml:"dbname"`
	SSLMode		string `toml:"sslmode"`
}