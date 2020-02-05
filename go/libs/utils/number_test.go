package utils_test

import (
	"github.com/shopspring/decimal"
	"github.com/velo-protocol/DRSv2/go/libs/utils"
	"testing"
)

func TestIsDecimalValid(t *testing.T) {
	type args struct {
		s decimal.Decimal
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Success", args{decimal.NewFromFloat(100000000000)}, true},
		{"Success", args{decimal.NewFromFloat(10000000000)}, true},
		{"Success", args{decimal.NewFromFloat(1000000000)}, true},
		{"Success", args{decimal.NewFromFloat(100000000)}, true},
		{"Success", args{decimal.NewFromFloat(10000000)}, true},
		{"Success", args{decimal.NewFromFloat(1000000)}, true},
		{"Success", args{decimal.NewFromFloat(100000)}, true},
		{"Success", args{decimal.NewFromFloat(10000)}, true},
		{"Success", args{decimal.NewFromFloat(1000)}, true},
		{"Success", args{decimal.NewFromFloat(100)}, true},
		{"Success", args{decimal.NewFromFloat(100.1)}, true},
		{"Success", args{decimal.NewFromFloat(100.12)}, true},
		{"Success", args{decimal.NewFromFloat(100.123)}, true},
		{"Success", args{decimal.NewFromFloat(100.1234)}, true},
		{"Success", args{decimal.NewFromFloat(100.12345)}, true},
		{"Success", args{decimal.NewFromFloat(100.123456)}, true},
		{"Success", args{decimal.NewFromFloat(100.1234567)}, true},
		{"Failed", args{decimal.NewFromFloat(100.12345678)}, false},
		{"Failed", args{decimal.NewFromFloat(100.123456789)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsDecimalValid(tt.args.s); got != tt.want {
				t.Errorf("IsDecimalValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
