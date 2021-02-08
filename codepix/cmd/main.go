package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/thauska/imersao-fullstack-fullcycle/codepix/application/grpc"
	"github.com/thauska/imersao-fullstack-fullcycle/codepix/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}