package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conn, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=true",
		getEnvOrDefault("MYSQL_USERNAME", "root"),
		getEnvOrDefault("MYSQL_PASSWORD", "password"),
		getEnvOrDefault("MYSQL_HOST", "127.0.0.1"),
		getEnvOrDefault("MYSQL_PORT", "3306"),
	))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	dbs := []string{
		"booq_test",
	}

	for _, name := range dbs {
		if _, err = conn.Exec("DROP DATABASE IF EXISTS " + name); err != nil {
			panic(err)
		}
		if _, err = conn.Exec("CREATE DATABASE `" + name + "` CHARACTER SET = utf8mb4"); err != nil {
			panic(err)
		}
		log.Println("Database `" + name + "` was created")
	}
}

func getEnvOrDefault(env string, def string) string {
	s := os.Getenv(env)
	if len(s) == 0 {
		return def
	}
	return s
}
