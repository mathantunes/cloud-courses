package maxheap

import (
	"testing"
)

func TestMaxHeap_Insert(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		h    *MaxHeap
		args args
	}{
		{
			name: "First",
			h: &MaxHeap{
				array: []int{20, 10, 11},
			},
			args: args{21},
		},
		{
			name: "Second",
			h: &MaxHeap{
				array: []int{20, 10, 11, 5, 3, 2, 4},
			},
			args: args{7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Insert(tt.args.key)
			t.Log(tt.h)
		})
	}
}

func TestMaxHeap_Extract(t *testing.T) {
	tests := []struct {
		name string
		h    *MaxHeap
		want int
	}{
		{
			name: "First",
			h: &MaxHeap{
				array: []int{20, 10, 11, 5, 3, 2, 4},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Extract(); got != tt.want {
				t.Errorf("MaxHeap.Extract() = %v, want %v", got, tt.want)
			}
			t.Log(tt.h)
		})
	}
}
