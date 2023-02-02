package sscp

import (
	"bytes"
	"reflect"
	"sync"
	"testing"
)

func Test_bufferPool_Get(t *testing.T) {
	type fields struct {
		pool sync.Pool
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *bytes.Buffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &bufferPool{
				pool: tt.fields.pool,
			}
			if got := p.Get(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bufferPool.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bufferPool_Put(t *testing.T) {
	type fields struct {
		pool sync.Pool
	}
	type args struct {
		b *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &bufferPool{
				pool: tt.fields.pool,
			}
			p.Put(tt.args.b)
		})
	}
}

func Test_newBufferPool(t *testing.T) {
	tests := []struct {
		name string
		want *bufferPool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newBufferPool(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBufferPool() = %v, want %v", got, tt.want)
			}
		})
	}
}
