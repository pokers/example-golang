package main

import (
	"example-restapi+db/src/router"
	"example-restapi+db/src/service/sql"
	"example-restapi+db/src/util"
	"strconv"
)

// Before run you shoud run "$>go get -u golang.org/x/sys" on your terminal
func main() {
	util.PrintInfo("Run Application...")
	defer util.PrintInfo("Exit Application...")

	port, err := strconv.Atoi(util.GetEnvString("DB_RDS_PORT"))
	if err != nil {
		util.PrintError("Invalid Port number")
		panic(err.Error())
	}

	var dbConfig sql.DBConfig = sql.DBConfig{
		Type:     "mysql",
		Username: util.GetEnvString("DB_RDS_USER"),
		Pass:     util.GetEnvString("DB_RDS_PASS"),
		URL:      util.GetEnvString("GO_MYSQL_URL"),
		Port:     uint(port),
		DBName:   util.GetEnvString("DB_RDS_DATABASE"),
	}
	// util.PrintInfo("DB Config : %+v", dbConfig)

	sql.InitDB(&dbConfig)
	defer sql.CloseDB()
	router.RunRouter()
}
