package helpers

import (
	"context"
	"k8scommerce/services/api/admin/internal/types"
	"strconv"

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

	if p.Amount != 0 {
		amount, _ := currency.NewAmountFromInt64(int64(p.Amount), p.Currency)
		floatAmount, _ := strconv.ParseFloat(amount.Number(), 64)
		p.Amount = floatAmount
		p.DisplayAmount = formatter.Format(amount)
	}

	if p.CompareAtAmount != 0 {
		amount, _ := currency.NewAmountFromInt64(int64(p.CompareAtAmount), p.Currency)
		floatAmount, _ := strconv.ParseFloat(amount.Number(), 64)
		p.CompareAtAmount = floatAmount
		p.DisplayCompareAtAmount = formatter.Format(amount)
	}
}