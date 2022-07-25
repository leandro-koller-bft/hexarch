package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leandro-koller-bft/hexarch/adapters/cli"
	"github.com/leandro-koller-bft/hexarch/app"
	mock_app "github.com/leandro-koller-bft/hexarch/app/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := "abc"
	name := "Product Test"
	price := 25.99
	status := app.ENABLED

	mock := mock_app.NewMockIProduct(ctrl)
	mock.EXPECT().GetID().Return(id).AnyTimes()
	mock.EXPECT().GetName().Return(name).AnyTimes()
	mock.EXPECT().GetPrice().Return(price).AnyTimes()
	mock.EXPECT().GetStatus().Return(status).AnyTimes()

	service := mock_app.NewMockIProductService(ctrl)
	service.EXPECT().Create(name, price).Return(mock, nil).AnyTimes()
	service.EXPECT().Get(id).Return(mock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(mock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(mock, nil).AnyTimes()

	resultExpected := fmt.Sprintf(
		"Product ID #%s named '%s' has been created with price %f, status set to %s.",
		id,
		name,
		price,
		status)
	result, err := cli.Run(service, "create", "", name, price)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product '%s' has been enabled.", name)
	result, err = cli.Run(service, "enable", id, "", price)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product '%s' has been disabled.", name)
	result, err = cli.Run(service, "disable", id, "", price)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s.",
		id,
		name,
		price,
		status)
	result, err = cli.Run(service, "get", id, "", price)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
