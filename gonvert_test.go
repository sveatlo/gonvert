package gonvert

import (
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    string
		wantErr bool
	}{
		{
			name:  "nil",
			value: nil,
			want:  "",
		},
		{
			name:  "string",
			value: "foo",
			want:  "foo",
		},
		{
			name:  "int",
			value: 123,
			want:  "123",
		},
		{
			name:  "float",
			value: 1.23,
			want:  "1.23",
		},
		{
			name:  "bool",
			value: true,
			want:  "true",
		},
		// TODO: add more tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToString(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    int64
		wantErr bool
	}{
		{
			name:    "nil",
			value:   nil,
			wantErr: true,
		},
		{
			name:  "int",
			value: uint64(1),
			want:  1,
		},
		{
			name:  "float w/o truncation",
			value: float64(1),
			want:  1,
		},
		{
			name:    "float w truncation",
			value:   1.23,
			wantErr: true,
		},
		{
			name:  "string - valid int",
			value: "123",
			want:  123,
		},
		{
			name:    "string - invalid float",
			value:   "1.23",
			wantErr: true,
		},
		// TODO: add more tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    float64
		wantErr bool
	}{
		{
			name:    "nil",
			value:   nil,
			wantErr: true,
		},
		{
			name:  "int",
			value: int32(123),
			want:  123,
		},
		{
			name:  "string - int",
			value: "123",
			want:  123,
		},
		{
			name:  "string - float",
			value: "1.23",
			want:  1.23,
		},
		{
			name:  "float",
			value: float64(1.23), // float32 loses precision and makes the number not equal
			want:  1.23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToFloat(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    bool
		wantErr bool
	}{
		{
			name:    "nil",
			value:   nil,
			wantErr: true,
		},
		{
			name:  "int-true",
			value: 1,
			want:  true,
		},
		{
			name:  "int-false",
			value: 0,
			want:  false,
		},
		{
			name:  "yes",
			value: "yes",
			want:  true,
		},
		{
			name:  "no",
			value: "no",
			want:  false,
		},
		{
			name:  "YES",
			value: "YES",
			want:  true,
		},
		{
			name:  "string-true",
			value: "true",
			want:  true,
		},
		{
			name:  "string-false",
			value: "false",
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToBool(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToSlice(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSlice(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMapString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMapString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}
