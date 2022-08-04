package handler_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/TamerB/products-import-service/api/handler"
	mockdb "github.com/TamerB/products-import-service/db/mock"
	db "github.com/TamerB/products-import-service/db/sqlc"
	"github.com/go-playground/assert"

	"github.com/golang/mock/gomock"
)

func TestHandleImportSuccess(t *testing.T) {
	t.Run("Response 200 if stock is found and request succeeds", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)

		params := db.UpdateStockTxParams{
			ProductSKU:  RandomString(12),
			ProductName: RandomString(20),
			CountryCode: RandomString(2),
			Quantity:    RandomInt64(),
		}

		store.EXPECT().CreateOrUpdateStockTx(context.Background(), params).Return(nil)

		handler := handler.ImportRows{
			Rows: []handler.ImportRow{
				{
					Sku:         params.ProductSKU,
					Name:        params.ProductName,
					Country:     params.CountryCode,
					StockChange: params.Quantity,
				},
			},
		}

		errs := handler.HandleImport(store)

		assert.Equal(t, 0, len(errs))
	})
}

func TestHandleImportWillBeLessThanZero(t *testing.T) {
	t.Run("Response 400 if stock will be less than zero if updated", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		store := mockdb.NewMockStore(ctrl)

		params := db.UpdateStockTxParams{
			ProductSKU:  RandomString(12),
			ProductName: RandomString(20),
			CountryCode: RandomString(2),
			Quantity:    RandomInt64(),
		}

		store.EXPECT().CreateOrUpdateStockTx(context.Background(), params).Return(fmt.Errorf("quantity value would make quantity in minus: %d", params.Quantity))

		handler := handler.ImportRows{
			Rows: []handler.ImportRow{
				{
					Sku:         params.ProductSKU,
					Name:        params.ProductName,
					Country:     params.CountryCode,
					StockChange: params.Quantity,
				},
			},
		}

		errs := handler.HandleImport(store)
		assert.Equal(t, 1, len(errs))
	})
}
