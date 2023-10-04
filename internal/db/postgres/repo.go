package postgres

type PostgresRepo struct {
	AnalyzerRepo
	TypeRepo
	EventRepo
	GasRepo
	SensorRepo
	UserRepo
}

func NewPostgresRepo() *PostgresRepo {
	return &PostgresRepo{}
}
