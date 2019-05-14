package main

import (
	"flag"
	"os"

	log "github.com/daesu/payments/logging"
	"github.com/go-openapi/loads"
	"github.com/salmankhalid187/user-signup/cmd"
	"github.com/salmankhalid187/user-signup/gen/restapi"
	"github.com/salmankhalid187/user-signup/gen/restapi/operations"
	"github.com/salmankhalid187/user-signup/usersignup"

	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func start(db *gorm.DB) {

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

	// Health setup
	repo := usersignup.NewRepository(db)
	signUpService := usersignup.New(repo)
	usersignup.Configure(api, signUpService)

	if err := cmd.Start(api, *portFlag); err != nil {
		log.Fatal("Failed to start", err)
	}
}

func main() {
	db, err := gorm.Open("mssql", "sqlserver://sa:reallyStrongPwd123@localhost:1433?database=TutorialDB")
	if err != nil {
		log.Println("Error while connecting to data base")
	}
	defer db.Close()
	start(db)
}
