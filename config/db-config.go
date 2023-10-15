package config

var info string

func init() {
	info = "Successfully connected to database!"
}

func GetInfo() string {
	return info
}
