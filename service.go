package goeximsubscriberlib

import (
	"os"
	"time"

	baseLib "github.com/exim-id/go-exim-base-lib"
	"github.com/exim-id/go-exim-base-lib/db"
	"github.com/exim-id/go-exim-base-lib/errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func HttpStart(endpoints func(*echo.Echo)) {
	defer errors.NormalError()
	go backGround()
	baseLib.ServerStart(endpoints)
}

func backGround() {
	loop := loopService()
	for loop {
		loop = loopService()
		time.Sleep(time.Minute * 30)
	}
	os.Exit(0)
}

func loopService() bool {
	return db.NewMongoTransaction[bool](false).Transaction(func(d *mongo.Database) bool { return true })
}
