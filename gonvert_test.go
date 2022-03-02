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
	tests := []struct {
		name    string
		value   interface{}
		want    []interface{}
		wantErr bool
	}{
		{
			name:    "nil",
			value:   nil,
			want:    []interface{}{},
			wantErr: false,
		},
		{
			name:    "int",
			value:   1,
			want:    []interface{}{1},
			wantErr: false,
		},
		{
			name:    "bool",
			value:   true,
			want:    []interface{}{true},
			wantErr: false,
		},
		{
			name:    "float",
			value:   69.42,
			want:    []interface{}{69.42},
			wantErr: false,
		},
		{
			name:    "string",
			value:   "foobar",
			want:    []interface{}{"foobar"},
			wantErr: false,
		},
		{
			name:    "interface_slice",
			value:   []interface{}{1, true, 69, 42.69, "foobar"},
			want:    []interface{}{1, true, 69, 42.69, "foobar"},
			wantErr: false,
		},
		{
			name:    "int_slice",
			value:   []int64{1, 2, 3, 4, 5},
			want:    []interface{}{int64(1), int64(2), int64(3), int64(4), int64(5)},
			wantErr: false,
		},
		{
			name:    "string_slice",
			value:   []string{"lorem", "ipsum", "dolor", "sit", "amet"},
			want:    []interface{}{"lorem", "ipsum", "dolor", "sit", "amet"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToSlice(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %+v (%T), want %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    []string
		wantErr bool
	}{
		{
			name:    "nil",
			value:   nil,
			want:    []string{},
			wantErr: false,
		},
		{
			name:    "int",
			value:   1,
			want:    []string{"1"},
			wantErr: false,
		},
		{
			name:    "bool",
			value:   true,
			want:    []string{"true"},
			wantErr: false,
		},
		{
			name:    "float",
			value:   69.42,
			want:    []string{"69.42"},
			wantErr: false,
		},
		{
			name:    "string",
			value:   "foobar",
			want:    []string{"foobar"},
			wantErr: false,
		},
		{
			name:    "interface_slice",
			value:   []interface{}{1, true, 69, 42.69, "foobar"},
			want:    []string{"1", "true", "69", "42.69", "foobar"},
			wantErr: false,
		},
		{
			name:    "int_slice",
			value:   []int64{1, 2, 3, 4, 5},
			want:    []string{"1", "2", "3", "4", "5"},
			wantErr: false,
		},
		{
			name:    "string_slice",
			value:   []string{"lorem", "ipsum", "dolor", "sit", "amet"},
			want:    []string{"lorem", "ipsum", "dolor", "sit", "amet"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToStringSlice(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToStringSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringSlice() = %+v (%T), want %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestToIntSlice(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    []int64
		wantErr bool
	}{
		{
			name:    "nil",
			value:   nil,
			want:    []int64{},
			wantErr: false,
		},
		{
			name:    "int",
			value:   1,
			want:    []int64{1},
			wantErr: false,
		},
		{
			name:    "bool",
			value:   true,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "float",
			value:   69.42,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "string",
			value:   "foobar",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "interface_slice",
			value:   []interface{}{1, true, 69, 42.69, "foobar"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "int_slice",
			value:   []int64{1, 2, 3, 4, 5},
			want:    []int64{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "int_slice",
			value:   []float64{1, 2, 3, 4, 5},
			want:    []int64{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "string_slice",
			value:   []string{"lorem", "ipsum", "dolor", "sit", "amet"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToIntSlice(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToIntSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSlice() = %+v (%T), want %+v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}

func TestToMapString(t *testing.T) {
	tests := []struct {
		name    string
		value   interface{}
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "nil",
			value:   nil,
			want:    map[string]interface{}{},
			wantErr: false,
		},
		{
			name: "map[string]interface{}",
			value: map[string]interface{}{
				"foo":   "bar",
				"lorem": 1,
				"ipsum": 69.42,
				"dolor": true,
			},
			want: map[string]interface{}{
				"foo":   "bar",
				"lorem": 1,
				"ipsum": 69.42,
				"dolor": true,
			},
			wantErr: false,
		},
		{
			name: "map[interface{}]interface{}",
			value: map[interface{}]interface{}{
				"foo": "bar",
				69:    42,
				69.42: 69.42,
				false: true,
			},
			want: map[string]interface{}{
				"foo":   "bar",
				"69":    42,
				"69.42": 69.42,
				"false": true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMapString(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMapString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapString() = %v (%T), want %v (%T)", got, got, tt.want, tt.want)
			}
		})
	}
}
