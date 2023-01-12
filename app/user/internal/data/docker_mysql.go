package data

import (
	"database/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ory/dockertest/v3"
	"time"
)

func DockerMysql(img, version string) (string, func()) {
	return innerDockerMysql(img, version)
}

func innerDockerMysql(img, version string) (string, func()) {
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}
	if err := pool.Client.Ping(); err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	resource, err := pool.Run(img, version, []string{"MARIADB_ROOT_PASSWORD=test1234", "MARIADB_ROOT_HOST=%", "MARIADB_DATABASE=galileo"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	conStr := fmt.Sprintf("root:test1234@tcp(localhost:%s)/galileo?parseTime=true", resource.GetPort("3306/tcp"))

	if err := pool.Retry(func() error {
		var err error
		db, err := sql.Open("mysql", conStr)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return conStr, func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}
}
