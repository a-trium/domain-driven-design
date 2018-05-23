package db

type Property struct {
	Host     string `default:"localhost"`
	Port     string `default:"3306"`
	UserName string `default:"root"`
	Password string `default:"root"`
	Database string `default:"application"`
}