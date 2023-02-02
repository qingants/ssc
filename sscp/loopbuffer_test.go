package sscp

import (
	"reflect"
	"testing"
)

func Test_loopBuffer_Len(t *testing.T) {
	type fields struct {
		buf    []byte
		off    int
		looped bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &loopBuffer{
				buf:    tt.fields.buf,
				off:    tt.fields.off,
				looped: tt.fields.looped,
			}
			if got := b.Len(); got != tt.want {
				t.Errorf("loopBuffer.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loopBuffer_Cap(t *testing.T) {
	type fields struct {
		buf    []byte
		off    int
		looped bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &loopBuffer{
				buf:    tt.fields.buf,
				off:    tt.fields.off,
				looped: tt.fields.looped,
			}
			if got := b.Cap(); got != tt.want {
				t.Errorf("loopBuffer.Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loopBuffer_Write(t *testing.T) {
	type fields struct {
		buf    []byte
		off    int
		looped bool
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &loopBuffer{
				buf:    tt.fields.buf,
				off:    tt.fields.off,
				looped: tt.fields.looped,
			}
			got, err := b.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("loopBuffer.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loopBuffer.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loopBuffer_Read(t *testing.T) {
	type fields struct {
		buf    []byte
		off    int
		looped bool
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &loopBuffer{
				buf:    tt.fields.buf,
				off:    tt.fields.off,
				looped: tt.fields.looped,
			}
			got, err := b.Read(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("loopBuffer.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loopBuffer.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loopBuffer_Reset(t *testing.T) {
	type fields struct {
		buf    []byte
		off    int
		looped bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &loopBuffer{
				buf:    tt.fields.buf,
				off:    tt.fields.off,
				looped: tt.fields.looped,
			}
			b.Reset()
		})
	}
}

func Test_loopBuffer_CopyTo(t *testing.T) {
	type fields struct {
		buf    []byte
		off    int
		looped bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &loopBuffer{
				buf:    tt.fields.buf,
				off:    tt.fields.off,
				looped: tt.fields.looped,
			}
			b.CopyTo()
		})
	}
}
