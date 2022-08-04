package handler

import (
	"context"

	db "github.com/TamerB/products-import-service/db/sqlc"
)

type ImportRows struct {
	Rows []ImportRow
}

type ImportRow struct {
	Country     string `csv:"country"` // .csv column headers
	Sku         string `csv:"sku"`
	Name        string `csv:"name"`
	StockChange int64  `csv:"stock_change"`
}

func CreateImportRow(country string, sku string, name string, stockChange int64) ImportRow {
	return ImportRow{
		Country:     country,
		Sku:         sku,
		Name:        name,
		StockChange: stockChange,
	}
}

// HandleImport performs stock update, or product/country/stock if any is missing
func (i *ImportRows) HandleImport(store db.Store) []error {
	errs := []error{}

	for _, row := range i.Rows {
		err := store.CreateOrUpdateStockTx(context.Background(), db.UpdateStockTxParams{
			ProductSKU:  row.Sku,
			ProductName: row.Name,
			CountryCode: row.Country,
			Quantity:    int64(row.StockChange),
		})
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
