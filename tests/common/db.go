package common

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func CheckEnv(key string) string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Error().Err(err).Msg("failed to load .env")
	// } else {
	// 	log.Info().Msg("Successfully loaded .env file")
	// }

	val := os.Getenv(key)
	if val == "" {
		fmt.Printf("%v cannot be empty: ", key)
	}
	return val
}

func InitTestDB(t *testing.T) *pgxpool.Pool {

	dsn := CheckEnv("POSTGRES_TEST_DSN")
	if dsn == "" {
		t.Fatalf("POSTGRES_TEST_DSN is not set")
	}

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse dsn")
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	require.NoError(t, err)

	return pool
}
