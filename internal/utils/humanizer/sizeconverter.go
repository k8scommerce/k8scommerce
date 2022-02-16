package humanizer

import (
	"fmt"

	"github.com/dustin/go-humanize"
)

// SizeToHuman handles decimal places
func SizeToHuman(size int64) string {
	const unit = 1000
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(size)/float64(div), "kMGTPE"[exp])
}

// converts B, MB, GB, TB, etc to int64
func HumanToSize(size string) int64 {
	sUint64, _ := humanize.ParseBytes(size)
	return int64(sUint64)
}
