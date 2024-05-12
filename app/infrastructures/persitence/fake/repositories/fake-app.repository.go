package repositories

type FakeAppRepository struct{}

func (*FakeAppRepository) Diagnose() error {
	return nil
}