package main

import (
	. "config"
	. "models/client"
	. "models/session_token"
)

func main() {
	db := DatabaseConnection()
	db.AutoMigrate(&Client{}, &SessionToken{})

	defer db.Close()
}
