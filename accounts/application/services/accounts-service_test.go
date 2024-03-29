package services

import (
	"app/accounts/application/payloads"
	"app/accounts/domain/factories"
	"app/accounts/domain/models"
	"app/accounts/persistence/fake/repositories"
	"app/crypto"
	"app/uuid"
	"testing"
)

func getTestContext() (*AccountsService, func(), func() (*models.Account, error)) {
	accountsService := &AccountsService{
		AccountsFactory: factories.AccountsFactory{
			UuidService:   uuid.GetService(),
			CryptoService: crypto.GetService(),
		},
		AccountsRepository: &repositories.FakeAccountsRepository{},
	}

	createAccount := func() (*models.Account, error) {
		return accountsService.Create(payloads.CreateAccountPayload{
			FirstName: "Dummy first name",
			LastName: "Dummy last name",
			Email: "dummy@dummy.com",
			Password: "1234",
			PlanId: "dummyPlanId",
		})
	}

	return accountsService, repositories.ResetFakeAccountsRepository, createAccount
}

func TestCreateAccountService(t *testing.T) {
	_, reset, createAccount := getTestContext()
	defer reset()

	_, err := createAccount()
	if err != nil {
		t.Errorf("Expected to create an Account but got %v", err)
		return
	}
}

func TestFindAllAccountsService(t *testing.T) {
	accountsService, reset, createAccount := getTestContext()
	defer reset()
	newAccount, _ := createAccount()

	accounts, err := accountsService.FindAll()
	if err != nil {
		t.Errorf("Expected to find all Accounts but got %v", err)
		return
	}

	if len(accounts) != 1 {
		t.Errorf("Expected slice of Accounts with a length of 1 but got %d", len(accounts))
		return
	}

	if accounts[0] != newAccount {
		t.Errorf("Expected first Account of slice to be equal to New Account but got %v", accounts[0])
		return
	}
}

func TestFindAccountByIdService(t *testing.T) {
	accountsService, reset, createAccount := getTestContext()
	defer reset()
	newAccount, _ := createAccount()

	account, err := accountsService.FindById(newAccount.Id)
	if err != nil {
		t.Errorf("Expected to find an Account by id but got %v", err)
		return
	}

	if newAccount != account {
		t.Errorf("Expected Account to be equal to New Account but got %v", account)
		return
	}
}

func TestFindAccountByEmailService(t *testing.T) {
	accountsService, reset, createAccount := getTestContext()
	defer reset()
	newAccount, _ := createAccount()

	account, err := accountsService.FindByEmail(newAccount.Email)
	if err != nil {
		t.Errorf("Expected to find an Account by email but got %v", err)
		return
	}

	if newAccount != account {
		t.Errorf("Expected Account to be equal to New Account but got %v", account)
		return
	}
}