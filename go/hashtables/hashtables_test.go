package hashtables

import (
	"testing"
)

func TestHashTable_Insert(t *testing.T) {
	type args struct {
		keys []string
	}
	tests := []struct {
		name string
		h    *HashTable
		args args
	}{
		{
			name: "Insert",
			h:    New(),
			args: args{keys: []string{"Key", "Yke", "Kye"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertKeys(tt.h, tt.args.keys)
			tt.h.Print()
		})
	}
}

func insertKeys(h *HashTable, keys []string) {
	for _, key := range keys {
		h.Insert(key)
	}
}

func TestHashTable_Search(t *testing.T) {

	h := New()
	insertKeys(h, []string{"Search1", "Search2", "arch1", "earch1", "ch1", "h1", "3"})

	type args struct {
		key string
	}
	tests := []struct {
		name string
		h    *HashTable
		args args
		want bool
	}{
		{
			name: "Search",
			h:    h,
			args: args{key: "arch1"},
			want: true,
		},
		{
			name: "Search unexisting",
			h:    h,
			args: args{key: "not there"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Search(tt.args.key); got != tt.want {
				t.Errorf("HashTable.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashTable_Delete(t *testing.T) {
	h := New()
	insertKeys(h, []string{"Search1", "Search2", "Search2", "earch1", "ch1", "h1", "3"})

	type args struct {
		key string
	}
	tests := []struct {
		name string
		h    *HashTable
		args args
	}{
		{
			name: "Delete OK",
			h:    h,
			args: args{key: "Search2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Delete(tt.args.key)
			if ok := tt.h.Search(tt.args.key); ok {
				t.Errorf("Wanted %v to be deleted but hashtable is currently", tt.args.key)
			}
			tt.h.Print()
		})
	}
}
