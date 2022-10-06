// Package android handles Android OS or Android application information.
package android

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewAPI(t *testing.T) {
	type args struct {
		lv uint64
	}
	tests := []struct {
		name string
		args args
		want *API
	}{
		{
			name: "API Lv.1",
			args: args{
				lv: 1,
			},
			want: &API{
				Level:    1,
				CodeName: "(no code name)",
				Version:  "1.0",
			},
		},
		{
			name: "API Lv.2",
			args: args{
				lv: 2,
			},
			want: &API{
				Level:    2,
				CodeName: "(no code name)",
				Version:  "1.1",
			},
		},
		{
			name: "API Lv.3",
			args: args{
				lv: 3,
			},
			want: &API{
				Level:    3,
				CodeName: "Cupcake",
				Version:  "1.5",
			},
		},
		{
			name: "API Lv.4",
			args: args{
				lv: 4,
			},
			want: &API{
				Level:    4,
				CodeName: "Donut",
				Version:  "1.6",
			},
		},
		{
			name: "API Lv.5",
			args: args{
				lv: 5,
			},
			want: &API{
				Level:    5,
				CodeName: "Eclair",
				Version:  "2.0 - 2.1",
			},
		},
		{
			name: "API Lv.6",
			args: args{
				lv: 6,
			},
			want: &API{
				Level:    6,
				CodeName: "Eclair",
				Version:  "2.0 - 2.1",
			},
		},
		{
			name: "API Lv.7",
			args: args{
				lv: 7,
			},
			want: &API{
				Level:    7,
				CodeName: "Eclair",
				Version:  "2.0 - 2.1",
			},
		},
		{
			name: "API Lv.8",
			args: args{
				lv: 8,
			},
			want: &API{
				Level:    8,
				CodeName: "Froyo",
				Version:  "2.2 - 2.2.3",
			},
		},
		{
			name: "API Lv.9",
			args: args{
				lv: 9,
			},
			want: &API{
				Level:    9,
				CodeName: "Gingerbread",
				Version:  "2.3 - 2.3.7",
			},
		},
		{
			name: "API Lv.10",
			args: args{
				lv: 10,
			},
			want: &API{
				Level:    10,
				CodeName: "Gingerbread",
				Version:  "2.3 - 2.3.7",
			},
		},
		{
			name: "API Lv.11",
			args: args{
				lv: 11,
			},
			want: &API{
				Level:    11,
				CodeName: "Honeycomb",
				Version:  "3.0 - 3.2.6",
			},
		},
		{
			name: "API Lv.12",
			args: args{
				lv: 12,
			},
			want: &API{
				Level:    12,
				CodeName: "Honeycomb",
				Version:  "3.0 - 3.2.6",
			},
		},
		{
			name: "API Lv.13",
			args: args{
				lv: 13,
			},
			want: &API{
				Level:    13,
				CodeName: "Honeycomb",
				Version:  "3.0 - 3.2.6",
			},
		},
		{
			name: "API Lv.14",
			args: args{
				lv: 14,
			},
			want: &API{
				Level:    14,
				CodeName: "Ice Cream Sandwich",
				Version:  "4.0 - 4.0.4",
			},
		},
		{
			name: "API Lv.15",
			args: args{
				lv: 15,
			},
			want: &API{
				Level:    15,
				CodeName: "Ice Cream Sandwich",
				Version:  "4.0 - 4.0.4",
			},
		},
		{
			name: "API Lv.16",
			args: args{
				lv: 16,
			},
			want: &API{
				Level:    16,
				CodeName: "Jelly Bean",
				Version:  "4.1 - 4.3.1",
			},
		},
		{
			name: "API Lv.17",
			args: args{
				lv: 17,
			},
			want: &API{
				Level:    17,
				CodeName: "Jelly Bean",
				Version:  "4.1 - 4.3.1",
			},
		},
		{
			name: "API Lv.18",
			args: args{
				lv: 18,
			},
			want: &API{
				Level:    18,
				CodeName: "Jelly Bean",
				Version:  "4.1 - 4.3.1",
			},
		},
		{
			name: "API Lv.19",
			args: args{
				lv: 19,
			},
			want: &API{
				Level:    19,
				CodeName: "KitKat/4.4W",
				Version:  "4.4 - 4.4.4",
			},
		},
		{
			name: "API Lv.20",
			args: args{
				lv: 20,
			},
			want: &API{
				Level:    20,
				CodeName: "KitKat/4.4W",
				Version:  "4.4 - 4.4.4",
			},
		},
		{
			name: "API Lv.21",
			args: args{
				lv: 21,
			},
			want: &API{
				Level:    21,
				CodeName: "Lollipop",
				Version:  "5.0 - 5.1.1",
			},
		},
		{
			name: "API Lv.22",
			args: args{
				lv: 22,
			},
			want: &API{
				Level:    22,
				CodeName: "Lollipop",
				Version:  "5.0 - 5.1.1",
			},
		},
		{
			name: "API Lv.23",
			args: args{
				lv: 23,
			},
			want: &API{
				Level:    23,
				CodeName: "Marshmallow",
				Version:  "6.0 - 6.0.1",
			},
		},
		{
			name: "API Lv.24",
			args: args{
				lv: 24,
			},
			want: &API{
				Level:    24,
				CodeName: "Nougat",
				Version:  "7.0 - 7.1.2",
			},
		},
		{
			name: "API Lv.25",
			args: args{
				lv: 25,
			},
			want: &API{
				Level:    25,
				CodeName: "Nougat",
				Version:  "7.0 - 7.1.2",
			},
		},
		{
			name: "API Lv.26",
			args: args{
				lv: 26,
			},
			want: &API{
				Level:    26,
				CodeName: "Oreo",
				Version:  "8.0 - 8.1",
			},
		},
		{
			name: "API Lv.27",
			args: args{
				lv: 27,
			},
			want: &API{
				Level:    27,
				CodeName: "Oreo",
				Version:  "8.0 - 8.1",
			},
		},
		{
			name: "API Lv.28",
			args: args{
				lv: 28,
			},
			want: &API{
				Level:    28,
				CodeName: "Pie",
				Version:  "9",
			},
		},
		{
			name: "API Lv.29",
			args: args{
				lv: 29,
			},
			want: &API{
				Level:    29,
				CodeName: "Q",
				Version:  "10",
			},
		},
		{
			name: "API Lv.30",
			args: args{
				lv: 30,
			},
			want: &API{
				Level:    30,
				CodeName: "R",
				Version:  "11",
			},
		},
		{
			name: "API Lv.31",
			args: args{
				lv: 31,
			},
			want: &API{
				Level:    31,
				CodeName: "S",
				Version:  "12",
			},
		},
		{
			name: "API Lv.32",
			args: args{
				lv: 32,
			},
			want: &API{
				Level:    32,
				CodeName: "Sv2",
				Version:  "12L",
			},
		},
		{
			name: "API Lv.33",
			args: args{
				lv: 33,
			},
			want: &API{
				Level:    33,
				CodeName: "Tiramisu",
				Version:  "13",
			},
		},
		{
			name: "No applicable API level",
			args: args{
				lv: 0,
			},
			want: &API{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAPI(tt.args.lv)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
