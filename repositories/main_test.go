package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true",
		dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setupTestData() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/setupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/cleanupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		fmt.Println("connectDB", err)
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup", err)
		return err
	}
	return nil
}

func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		// testing.M は Fatal系のメソッドを持たないため、os.Exit関数でプログラムを終了させる
		// os.Exit関数：引数として与えられた終了コードでプログラムを終了させる
		os.Exit(1)
	}

	m.Run()

	teardown()
}
