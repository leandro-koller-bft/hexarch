package app_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/leandro-koller-bft/hexarch/app"
	mock_app "github.com/leandro-koller-bft/hexarch/app/mocks"
	"github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	persistence := mock_app.NewMockIProductPersistence(ctrl)
	// any time persistence.Get is called, it will return a product mock
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("123")
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	persistence := mock_app.NewMockIProductPersistence(ctrl)
	// any time persistence.Save is called, it will return a product mock
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)

	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_app.NewMockIProduct(ctrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_app.NewMockIProductPersistence(ctrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := app.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)

	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)

	require.Nil(t, err)
	require.Equal(t, product, result)
}
