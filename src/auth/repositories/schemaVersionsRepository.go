package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type SchemaVersionsRepository struct {
	conn *pgx.Conn
}

func NewSchemaVersionRepository(conn *pgx.Conn) *SchemaVersionsRepository {
	return &SchemaVersionsRepository{conn: conn}
}

func (repo *SchemaVersionsRepository) GetScalar() (string, error) {
	sql := `select version_number from schema_version`
	var version_number string
	if err := repo.conn.QueryRow(context.Background(), sql).Scan(&version_number); err != nil {
		return "", err
	}
	return version_number, nil
}
