package main

import "parishioner_management/cmd/deployment/scripts"

func main() {
	scripts.LoadEnv()
	scripts.CreateAdminAccount()
}
