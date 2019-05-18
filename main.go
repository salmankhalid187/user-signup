package main

import (
	"flag"
	"os"

	log "github.com/daesu/payments/logging"
	"github.com/go-openapi/loads"
	"github.com/salmankhalid187/user-signup/cmd"
	"github.com/salmankhalid187/user-signup/currenttime"
	"github.com/salmankhalid187/user-signup/gen/restapi"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations"
	"github.com/salmankhalid187/user-signup/usersignup"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

const dbUserName = "sa"
const dbPassword = "reallyStrongPwd123"
const serverName = "localhost"
const dataBaseName = "TutorialDB"

func startApiServer(db *gorm.DB) {

	host, err := os.Hostname()
	if err != nil {
		log.Fatal("unable to get Hostname", err)
	}
	log.WithFields(logrus.Fields{
		"Host": host,
	}).Info("Service Startup")

	var portFlag = flag.Int("port", 8080, "Port to listen for web requests on")

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal("Invalid swagger file for initializing", err)
	}

	api := operations.NewUserSignUpApisAPI(swaggerSpec)

	repo := usersignup.NewRepository(db)
	signUpService := usersignup.New(repo)
	usersignup.Configure(api, signUpService)

	currentTimeService := currenttime.New()
	currenttime.Configure(api, currentTimeService)

	if err := cmd.Start(api, *portFlag); err != nil {
		log.Fatal("Failed to start", err)
	}
}

func connectDb() (db *gorm.DB) {

	db, err := gorm.Open("mssql", "sqlserver://"+dbUserName+":"+dbPassword+"@"+serverName+":1433?database="+dataBaseName)
	if err != nil {
		log.Println("Error while connecting to data base")
		return nil
	}

	return db

}

func main() {
	db := connectDb()
	defer db.Close()
	startApiServer(db)
}
