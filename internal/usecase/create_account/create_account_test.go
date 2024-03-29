package create_account

import (
	"testing"

	"github.com/elieudomaia/ms-wallet-app/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateAccountExecute(t *testing.T) {
	client, _ := entity.NewClient("Elieudo Maia", "any@mail.com")

	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Get", client.ID).Return(client, nil)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)
	input := &CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)
	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
