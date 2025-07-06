package config

type Config struct {
	Server Server
	Database Database
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	URL string
}