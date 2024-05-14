package lilymetrics

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp(dir_name string) {

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	migration_path := os.Getenv("PWD") + "/migrations/" + dir_name
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", "postgres://"+dbHost+":5432/postgres?sslmode=disable&user="+dbUser+"&password="+dbPassword)
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{MigrationsTable: dir_name + "_version"})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///"+migration_path,
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migration" + migration_path)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Error running migrations:", err)
	}

}
func isDuplicateKeyViolation(err error) bool {
	// Check if the error is related to a duplicate key violation
	// This assumes you're using the pq driver which returns a specific error for duplicate key violations
	// Adjust this function based on the error type returned by your PostgreSQL driver
	if pqErr, ok := err.(*pq.Error); ok {
		return pqErr.Code == "23505" // PostgreSQL error code for unique_violation
	}
	return false
}
func CopyDir(src string, dst string) error {
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	fmt.Println("Creating directory: ", dst)
	os.MkdirAll(dst, os.ModePerm)

	// Check if the directory exists
	_, err = os.Stat(dst)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalf("Failed to create directory: %v", err)
		} else {
			log.Fatalf("Error checking directory: %v", err)
		}
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		fileInfo, err := os.Stat(srcPath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		case os.ModeSymlink, os.ModeNamedPipe, os.ModeSocket, os.ModeDevice:
			if err := os.Link(srcPath, dstPath); err != nil {
				return err
			}
		default:
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}
func CopyFile(src string, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}
