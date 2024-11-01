package main

import "minhajuddinkhan/runng/migrations"

func main() {
	migrations.MigrateTableCreate()
	migrations.MigratePlayerCreate()
}
