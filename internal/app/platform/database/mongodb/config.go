package mongodb

import "os"

func GetMongoDBURI() string {
	return os.Getenv("MONGODB_URI")
}
