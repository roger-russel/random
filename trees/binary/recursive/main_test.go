package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeImpl_Add(t *testing.T) {
	type fields struct {
		nodes *BinaryTreeNodes
	}
	type args struct {
		index int
		data  any
	}
	tests := []struct {
		name   string
		fields fields
		args   []args
		want   *BinaryTreeNodes
	}{
		{
			name: "simple case",
			fields: fields{
				nodes: nil,
			},
			args: []args{
				{
					index: 1,
					data:  "foo",
				},
			},
			want: &BinaryTreeNodes{
				TreeRight: nil,
				TreeLeft:  nil,
				Index:     1,
				Data:      "foo",
			},
		},
		{
			name: "to the right case",
			fields: fields{
				nodes: nil,
			},
			args: []args{
				{
					index: 10,
					data:  "foo",
				},
				{
					index: 13,
					data:  "boo",
				},
			},
			want: &BinaryTreeNodes{
				TreeRight: &BinaryTreeNodes{
					TreeRight: nil,
					TreeLeft:  nil,
					Index:     13,
					Data:      "boo",
				},
				TreeLeft: nil,
				Index:    10,
				Data:     "foo",
			},
		},
		{
			name: "to the left case",
			fields: fields{
				nodes: nil,
			},
			args: []args{
				{
					index: 10,
					data:  "foo",
				},
				{
					index: 9,
					data:  "boo",
				},
			},
			want: &BinaryTreeNodes{
				TreeRight: nil,
				TreeLeft: &BinaryTreeNodes{
					TreeRight: nil,
					TreeLeft:  nil,
					Index:     9,
					Data:      "boo",
				},
				Index: 10,
				Data:  "foo",
			},
		},
		{
			name: "a bit more complex case",
			fields: fields{
				nodes: nil,
			},
			args: []args{
				{
					index: 10,
					data:  "foo",
				},
				{
					index: 5,
					data:  "boo",
				},
				{
					index: 7,
					data:  "xoo",
				},
			},
			want: &BinaryTreeNodes{
				TreeRight: nil,
				TreeLeft: &BinaryTreeNodes{
					TreeRight: &BinaryTreeNodes{
						TreeRight: nil,
						TreeLeft:  nil,
						Index:     7,
						Data:      "xoo",
					},
					TreeLeft: nil,
					Index:    5,
					Data:     "boo",
				},
				Index: 10,
				Data:  "foo",
			},
		},
		{
			name: "a complex case",
			fields: fields{
				nodes: nil,
			},
			args: []args{
				{
					index: 10,
					data:  "10",
				},
				{
					index: 5,
					data:  "5",
				},
				{
					index: 7,
					data:  "7",
				},
				{
					index: 6,
					data:  "6",
				},
				{
					index: 8,
					data:  "8",
				},
				{
					index: 13,
					data:  "13",
				},
			},
			want: &BinaryTreeNodes{
				TreeRight: &BinaryTreeNodes{
					Index: 13,
					Data:  "13",
				},
				TreeLeft: &BinaryTreeNodes{
					TreeRight: &BinaryTreeNodes{
						TreeRight: &BinaryTreeNodes{
							Index: 8,
							Data:  "8",
						},
						TreeLeft: &BinaryTreeNodes{
							Index: 6,
							Data:  "6",
						},
						Index: 7,
						Data:  "7",
					},
					Index: 5,
					Data:  "5",
				},
				Index: 10,
				Data:  "10",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BinaryTreeImpl{
				nodes: tt.fields.nodes,
			}
			for _, args := range tt.args {
				bt.Add(args.index, args.data)
			}
			assert.Empty(t, cmp.Diff(tt.want, bt.nodes))
		})
	}
}

func TestBinaryTreeImpl_Search(t *testing.T) {
	type fields struct {
		nodes *BinaryTreeNodes
	}
	type args struct {
		index int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData any
		wantErr  error
	}{
		{
			name: "simple case",
			args: args{
				index: 10,
			},
			wantData: "foo",
			wantErr: nil,
			fields: fields{
				nodes: &BinaryTreeNodes{
					TreeRight: nil,
					TreeLeft: nil,
					Index: 10,
					Data:  "foo",
				},
			},
		},
		{
			name: "simple error case",
			args: args{
				index: 10,
			},
			wantData: nil,
			wantErr: ErrIndexNotFound,
			fields: fields{
				nodes: &BinaryTreeNodes{
					TreeRight: nil,
					TreeLeft: nil,
					Index: 9,
					Data:  "foo",
				},
			},
		},
		{
			name: "complex case",
			args: args{
				index: 11,
			},
			wantData: "10",
			wantErr: nil,
			fields: fields{
				nodes: &BinaryTreeNodes{
					TreeRight: &BinaryTreeNodes{
						TreeRight: nil,
						TreeLeft:  &BinaryTreeNodes{
							TreeRight: &BinaryTreeNodes{
								TreeRight: nil,
								TreeLeft: nil,
								Index: 10,
								Data:  "10",
							},
							TreeLeft: nil,
							Index: 9,
							Data:  "foo",
						},
						Index: 14,
						Data:  "foo",
					},
					TreeLeft: nil,
					Index: 8,
					Data:  "foo",
				},
			},
		},			
		{
			name: "complex error case",
			args: args{
				index: 10,
			},
			wantData: nil,
			wantErr: ErrIndexNotFound,
			fields: fields{
				nodes: &BinaryTreeNodes{
					TreeRight: &BinaryTreeNodes{
						TreeRight: nil,
						TreeLeft:  &BinaryTreeNodes{
							TreeRight: nil,
							TreeLeft: nil,
							Index: 9,
							Data:  "foo",
						},
						Index: 14,
						Data:  "foo",
					},
					TreeLeft: nil,
					Index: 8,
					Data:  "foo",
				},
			},
		},		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := &BinaryTreeImpl{
				nodes: tt.fields.nodes,
			}
			gotData, err := bt.Search(tt.args.index)
			if tt.wantErr != nil {
				assert.True(t, errors.Is(err, tt.wantErr), fmt.Sprintf("error got %v, expected %v", err.Error(), tt.wantErr.Error()))
			} else {
				assert.Nil(t, err)
			}
			
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("BinaryTreeImpl.Search() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}
