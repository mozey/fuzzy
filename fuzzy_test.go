package fuzzy_test

import (
	"encoding/json"
	"testing"

	"github.com/mozey/fuzzy"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	type Data struct {
		String fuzzy.String `json:"string"`
	}
	d := Data{}

	// null
	b := []byte(`{"string": null}`)
	err := json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "", string(d.String), "value must match")

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "", string(d.String), "value must match")

	// string
	b = []byte(`{"string": "123"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "123", string(d.String), "value must match")

	// int
	b = []byte(`{"string": 123}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "123", string(d.String), "value must match")

	b = []byte(`{"string": 0}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "0", string(d.String), "value must match")

	b = []byte(`{"string": -123}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "-123", string(d.String), "value must match")

	// float
	b = []byte(`{"string": 123.456}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "123.456", string(d.String), "value must match")

	b = []byte(`{"string": -123.456}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "-123.456", string(d.String), "value must match")

	// bool
	b = []byte(`{"string": true}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "true", string(d.String), "value must match")

	b = []byte(`{"string": false}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, "false", string(d.String), "value must match")
}

func TestInt(t *testing.T) {
	type Data struct {
		Int fuzzy.Int `json:"int"`
	}
	d := Data{}

	// null
	b := []byte(`{"int": null}`)
	err := json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, int64(0), int64(d.Int), "value must match")

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, int64(0), int64(d.Int), "value must match")

	// string
	b = []byte(`{"int": "123"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, int64(123), int64(d.Int), "value must match")

	b = []byte(`{"int": "-123"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, int64(-123), int64(d.Int), "value must match")

	b = []byte(`{"int": "abc"}`)
	err = json.Unmarshal(b, &d)
	require.Error(t, err)
	// TODO Add context to err message
	// TODO Check err message

	// int
	b = []byte(`{"int": -123}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, int64(-123), int64(d.Int), "value must match")

	// float
	b = []byte(`{"int": -123.456}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, int64(-123), int64(d.Int), "value must match")

	// bool
	b = []byte(`{"int": true}`)
	err = json.Unmarshal(b, &d)
	require.Error(t, err)
	// TODO Add context to err message
	// TODO Check err message
}

func TestFloat(t *testing.T) {
	type Data struct {
		Float fuzzy.Float `json:"float"`
	}
	d := Data{}

	// null
	b := []byte(`{"float": null}`)
	err := json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, float64(0), float64(d.Float), "value must match")

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, float64(0), float64(d.Float), "value must match")

	// string
	b = []byte(`{"float": "1.618"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, float64(1.618), float64(d.Float), "value must match")

	b = []byte(`{"float": "-1.618"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, float64(-1.618), float64(d.Float), "value must match")

	b = []byte(`{"float": "abc"}`)
	err = json.Unmarshal(b, &d)
	require.Error(t, err)
	// TODO Add context to err message
	// TODO Check err message

	// int
	b = []byte(`{"float": -1}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, float64(-1), float64(d.Float), "value must match")

	// float
	b = []byte(`{"float": -1.618}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, float64(-1.618), float64(d.Float), "value must match")

	// bool
	b = []byte(`{"float": true}`)
	err = json.Unmarshal(b, &d)
	require.Error(t, err)
	// TODO Add context to err message
	// TODO Check err message
}

func TestBool(t *testing.T) {
	type Data struct {
		Bool fuzzy.Bool `json:"bool"`
	}
	d := Data{}

	// null
	b := []byte(`{"bool": null}`)
	err := json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, bool(d.Bool), "value must match")

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, bool(d.Bool), "value must match")

	// string
	b = []byte(`{"bool": "false"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, bool(d.Bool), "value must match")

	b = []byte(`{"bool": "0"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, bool(d.Bool), "value must match")

	b = []byte(`{"bool": ""}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, bool(d.Bool), "value must match")

	b = []byte(`{"bool": "true"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, bool(d.Bool), "value must match")

	b = []byte(`{"bool": "1"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, bool(d.Bool), "value must match")

	b = []byte(`{"bool": "abc"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, bool(d.Bool), "value must match")

	// int
	b = []byte(`{"bool": 0}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, bool(d.Bool), "value must match")

	b = []byte(`{"bool": 1}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, bool(d.Bool), "value must match")

	b = []byte(`{"bool": -1}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, bool(d.Bool), "value must match")

	// float
	b = []byte(`{"bool": 1.23}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, bool(d.Bool), "value must match")

	b = []byte(`{"bool": -1.23}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, bool(d.Bool), "value must match")

	// bool
	b = []byte(`{"bool": false}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, bool(d.Bool), "value must match")

	b = []byte(`{"bool": true}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, bool(d.Bool), "value must match")
}

func TestMarshalToJSON(t *testing.T) {
	type Data struct {
		String fuzzy.String `json:"string"`
		Int    fuzzy.Int    `json:"int"`
		Bool   fuzzy.Bool   `json:"bool"`
		Float  fuzzy.Float  `json:"float"`
	}

	// Valid values
	d := Data{}
	d.String = "foo"
	d.Int = fuzzy.Int(123)
	d.Bool = true
	d.Float = fuzzy.Float(1.618)
	b, err := json.Marshal(d)
	require.NoError(t, err)
	require.Equal(t, `{"string":"foo","int":123,"bool":true,"float":1.618}`, string(b))

	// Empty value is empty
	d = Data{}
	b, err = json.Marshal(d)
	require.NoError(t, err)
	require.Equal(t, `{"string":"","int":0,"bool":false,"float":0}`, string(b))

	// Missing properties are marshaled as empty.
	// Note that Marshal does not output the same as input to Unmarshal
	d = Data{}
	err = json.Unmarshal([]byte("{}"), &d)
	require.NoError(t, err)
	b, err = json.Marshal(d)
	require.NoError(t, err)
	require.Equal(t, `{"string":"","int":0,"bool":false,"float":0}`, string(b))

	// omitempty
	type Data2 struct {
		String fuzzy.String `json:"string,omitempty"`
		Int    fuzzy.Int    `json:"int,omitempty"`
		Bool   fuzzy.Bool   `json:"bool,omitempty"`
		Float  fuzzy.Float  `json:"float,omitempty"`
	}
	d2 := Data2{}
	err = json.Unmarshal([]byte("{}"), &d2)
	require.NoError(t, err)
	b, err = json.Marshal(d2)
	require.NoError(t, err)
	require.Equal(t, `{}`, string(b))
}
