package utils

import (
	"context"
	"database/sql"
	"sync"
	"testing"

	"github.com/UpRightSofia/lottolodge/src/models/config"
	"github.com/docker/go-connections/nat"
	"github.com/pressly/goose"
	log "github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var dbConfig = config.DBConfig{
	User:     "postgres",
	Password: "postgres",
	DBName:   "testdb",
}

// This function is overly complicated because t.Parallel() doesn't work with wg so we cannot stop the execution of cleanup.
// So we create our own t.Parallel() with cocaine and hookers (and wg)
func WithParallel(wg *sync.WaitGroup, testBlock func()) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		testBlock()
	}()
}

func WithPostgres(t *testing.T, testBlock func(*sql.DB, *sync.WaitGroup)) {
	ctx := context.Background()
	port := "5432/tcp"
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15.3",
		ExposedPorts: []string{port},
		WaitingFor:   wait.ForListeningPort(nat.Port(port)),
		Env: map[string]string{
			"POSTGRES_DB":       dbConfig.DBName,
			"POSTGRES_PASSWORD": dbConfig.Password,
			"POSTGRES_USER":     dbConfig.User,
		},
	}
	psqlC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		if err := psqlC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err.Error())
		}
	}()

	host, err := psqlC.Host(ctx)
	if err != nil {
		t.Error(err)
		return
	}

	actualPort, err := psqlC.MappedPort(ctx, nat.Port(port))
	if err != nil {
		t.Error(err)
		return
	}

	dbConfig.Host = host
	dbConfig.Port = actualPort.Port()

	db, err := sql.Open("pgx", dbConfig.GetDSN())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	// Run migrations
	migrationsDir := "../migrations"
	err = goose.Up(db, migrationsDir)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	//Allow testBlock to attach references to the waitgroup
	var wg sync.WaitGroup

	testBlock(db, &wg)

	//Wait for all testBlocks to finish then allow cleanup
	wg.Wait()
}
