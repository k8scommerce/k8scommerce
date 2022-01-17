package logic

import (
	"context"
	"strconv"
	"strings"

	"k8scommerce/services/api/client/internal/types"

	"github.com/bojanz/currency"
)

func convertOutgoingPrices(ctx context.Context, p *types.Price) {
	if p.Currency == "" {
		p.Currency = "USD"
	}

	locale := ctx.Value(types.Locale).(string)
	if locale == "" {
		locale = "en"
	}

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

func httpResponse(statusCode int64, message ...string) types.ResponseStatus {
	return types.ResponseStatus{
		StatusCode:    statusCode,
		StatusMessage: strings.Join(message, " "),
	}
}
