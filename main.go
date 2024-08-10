package main

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/cmd"
	_ "github.com/cesc1802/onboarding-and-volunteer-service/docs"
)

// @title Onboarding and Volunteer Service API
// @version 1.0
// @description This is a volunteer service API
// @securityDefinitions.apiKey bearerToken
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
