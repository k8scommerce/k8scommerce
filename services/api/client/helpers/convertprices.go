package helpers

import (
	"context"
	"strconv"

	"github.com/k8scommerce/k8scommerce/services/api/client/internal/types"

	"github.com/bojanz/currency"
)

func ConvertOutgoingPrices(ctx context.Context, p *types.Price) {
	if p.Currency == "" {
		p.Currency = "USD"
	}

	// locale := ctx.Value(types.Locale).(string)
	// if locale == "" {
	locale := "en"
	// }

	formatter := currency.NewFormatter(
		currency.NewLocale(locale),
	)

	if p.SalePrice != 0 {
		amount, _ := currency.NewAmountFromInt64(int64(p.SalePrice), p.Currency)
		floatAmount, _ := strconv.ParseFloat(amount.Number(), 64)
		p.SalePrice = floatAmount
		p.FormattedSalePrice = formatter.Format(amount)
	}

	if p.RetailPrice != 0 {
		amount, _ := currency.NewAmountFromInt64(int64(p.RetailPrice), p.Currency)
		floatAmount, _ := strconv.ParseFloat(amount.Number(), 64)
		p.RetailPrice = floatAmount
		p.FormattedRetailPrice = formatter.Format(amount)
	}
}
