package logic

import (
	"context"
	"fmt"
	"strconv"

	"k8scommerce/services/api/client/internal/types"

	"github.com/bojanz/currency"
)

func convertOutgoingPrices(ctx context.Context, p *types.Price) {
	if p.Currency == "" {
		p.Currency = "USD"
	}

	clientLocale := ctx.Value(types.ClientLocale).(string)
	fmt.Println("INCOMING LOCALE:", clientLocale)

	if clientLocale == "" {
		clientLocale = "en"
	}

	locale := currency.NewLocale(clientLocale)
	formatter := currency.NewFormatter(locale)

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
