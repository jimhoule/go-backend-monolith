package ports

type AppRepositoryPort interface {
	Diagnose() error
}