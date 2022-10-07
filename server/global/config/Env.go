package config

var AppInfo struct {
	Name    string `bson:"name"`
	Version string `bson:"version"`
	Port    int    `bson:"Port"`
}
