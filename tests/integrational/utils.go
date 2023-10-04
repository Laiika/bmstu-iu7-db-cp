package integrational

import (
	"context"
	"db_cp_6_sem/pkg/client/postgresql"
)

func TruncateTables(client postgresql.Client, ctx context.Context) error {
	q := `
		TRUNCATE gas_analyzers, sensors, gases, events, analyzer_types, types_gases
	`
	_, err := client.Exec(ctx, q)
	if err != nil {
		return err
	}

	return nil
}
