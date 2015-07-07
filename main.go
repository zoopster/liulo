package main

import (
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	r "github.com/dancannon/gorethink"
	"github.com/dwarvesf/liulo/conf"
	_ "github.com/dwarvesf/liulo/routers"
)

const (
	AppVer  = "0.1"
	DB_NAME = "liulo"
)

// InitDB is used for initial RethinkDB
func InitDB() error {

	session, err := r.Connect(r.ConnectOpts{
		Address:  os.Getenv("RETHINKDB_URL"),
		Database: DB_NAME,
		MaxIdle:  10,
		MaxOpen:  10,
		Timeout:  time.Second * 10,
	})

	if err != nil {
		beego.Info(err)
	}

	conf.Session = session

	migrate := beego.AppConfig.String("migrate")
	switch migrate {
	case "drop":
		// TODO: Drop database and recreate all the table
		beego.Warning("Something went wrong?")

		// // Create table `Event`
		// _, err = r.Db(dbName).TableCreate("event").RunWrite(session)

		// // Create table `Session`
		// _, err = r.Db(dbName).TableCreate("session").RunWrite(session)

		// // Create table `Question`
		// _, err = r.Db(dbName).TableCreate("question").RunWrite(session)

		// // Create table `User`
		// _, err = r.Db(dbName).TableCreate("user").RunWrite(session)
		break

	case "alter":
		// TODO: Run migration script
		beego.Info("Prepare to run migration script")

		break

	default:
		beego.Info("Database needs to be safe.")
	}

	return nil
}

func main() {

	InitDB()
	beego.EnableAdmin = true

	// Configure env
	beego.Info(beego.AppName, AppVer)

	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
