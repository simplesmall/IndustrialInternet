package main

import (
	"IndustrialInternet/config"
	"IndustrialInternet/routers"
)

// @title IndustrialInternet项目API
// @version 1.0
// @description This is a  server of IndustrialInternet.
// @termsOfService https://www.topgoer.com

// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email example@some.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8090
// @BasePath /
func main() {
	defer config.CloseDB()
	routers.InitServer()
}