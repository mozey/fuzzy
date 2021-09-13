package fuzzy_test

import (
	"encoding/json"
	"testing"

	"github.com/guregu/null"
	"github.com/mozey/fuzzy"
	"github.com/stretchr/testify/require"
)

func TestNullString(t *testing.T) {
	type Data struct {
		String fuzzy.NullString `json:"string"`
	}
	d := Data{}

	// null
	b := []byte(`{"string": null}`)
	err := json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, d.String.Valid, "string must not be valid")
	require.Equal(t, "", d.String.String, "value must match")

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, d.String.Valid, "string must not be valid")
	require.Equal(t, "", d.String.String, "value must match")

	// string
	b = []byte(`{"string": "123"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.String.Valid, "string must be valid")
	require.Equal(t, "123", d.String.String, "value must match")

	// int
	b = []byte(`{"string": 123}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.String.Valid, "string must be valid")
	require.Equal(t, "123", d.String.String, "value must match")

	b = []byte(`{"string": 0}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.String.Valid, "string must be valid")
	require.Equal(t, "0", d.String.String, "value must match")

	b = []byte(`{"string": -123}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.String.Valid, "string must be valid")
	require.Equal(t, "-123", d.String.String, "value must match")

	// float
	b = []byte(`{"string": 123.456}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.String.Valid, "string must be valid")
	require.Equal(t, "123.456", d.String.String, "value must match")

	b = []byte(`{"string": -123.456}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.String.Valid, "string must be valid")
	require.Equal(t, "-123.456", d.String.String, "value must match")

	// bool
	b = []byte(`{"string": true}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.String.Valid, "string must be valid")
	require.Equal(t, "true", d.String.String, "value must match")

	b = []byte(`{"string": false}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.String.Valid, "string must be valid")
	require.Equal(t, "false", d.String.String, "value must match")
}

func TestNullInt(t *testing.T) {
	type Data struct {
		Int fuzzy.NullInt `json:"int"`
	}
	d := Data{}

	// null
	b := []byte(`{"int": null}`)
	err := json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, d.Int.Valid, "int must not be valid")
	require.Equal(t, int64(0), d.Int.Int64, "value must match")

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, d.Int.Valid, "int must not be valid")
	require.Equal(t, int64(0), d.Int.Int64, "value must match")

	// string
	b = []byte(`{"int": "123"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Int.Valid, "int must be valid")
	require.Equal(t, int64(123), d.Int.Int64, "value must match")

	b = []byte(`{"int": "-123"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Int.Valid, "int must be valid")
	require.Equal(t, int64(-123), d.Int.Int64, "value must match")

	b = []byte(`{"int": "abc"}`)
	err = json.Unmarshal(b, &d)
	require.Error(t, err)
	// TODO Add context to err message
	// TODO Check err message

	// int
	b = []byte(`{"int": -123}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Int.Valid, "int must be valid")
	require.Equal(t, int64(-123), d.Int.Int64, "value must match")

	// float
	b = []byte(`{"int": -123.456}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Int.Valid, "int must be valid")
	require.Equal(t, int64(-123), d.Int.Int64, "value must match")

	// bool
	b = []byte(`{"int": true}`)
	err = json.Unmarshal(b, &d)
	require.Error(t, err)
	// TODO Add context to err message
	// TODO Check err message
}

func TestNullFloat(t *testing.T) {
	type Data struct {
		Float fuzzy.NullFloat `json:"float"`
	}
	d := Data{}

	// null
	b := []byte(`{"float": null}`)
	err := json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, d.Float.Valid, "must not be valid")
	require.Equal(t, float64(0), d.Float.Float64, "value must match")

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, d.Float.Valid, "must not be valid")
	require.Equal(t, float64(0), d.Float.Float64, "value must match")

	// string
	b = []byte(`{"float": "1.618"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Float.Valid, "must be valid")
	require.Equal(t, float64(1.618), d.Float.Float64, "value must match")

	b = []byte(`{"float": "-1.618"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Float.Valid, "must be valid")
	require.Equal(t, float64(-1.618), d.Float.Float64, "value must match")

	b = []byte(`{"float": "abc"}`)
	err = json.Unmarshal(b, &d)
	require.Error(t, err)
	// TODO Add context to err message
	// TODO Check err message

	// int
	b = []byte(`{"float": -1}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Float.Valid, "must be valid")
	require.Equal(t, float64(-1), d.Float.Float64, "value must match")

	// float
	b = []byte(`{"float": -1.618}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Float.Valid, "int must be valid")
	require.Equal(t, float64(-1.618), d.Float.Float64, "value must match")

	// bool
	b = []byte(`{"float": true}`)
	err = json.Unmarshal(b, &d)
	require.Error(t, err)
	// TODO Add context to err message
	// TODO Check err message
}

func TestNullBool(t *testing.T) {
	type Data struct {
		Bool fuzzy.NullBool `json:"bool"`
	}
	d := Data{}

	// null
	b := []byte(`{"bool": null}`)
	err := json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, d.Bool.Valid, "bool must not be valid")
	require.Equal(t, false, d.Bool.Bool, "value must match")

	b = []byte(`{}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, false, d.Bool.Valid, "bool must not be valid")
	require.Equal(t, false, d.Bool.Bool, "value must match")

	// string
	b = []byte(`{"bool": "false"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, false, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": "0"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, false, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": ""}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, false, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": "true"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, true, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": "1"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, true, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": "abc"}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, true, d.Bool.Bool, "value must match")

	// int
	b = []byte(`{"bool": 0}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, false, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": 1}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, true, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": -1}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, true, d.Bool.Bool, "value must match")

	// float
	b = []byte(`{"bool": 1.23}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, true, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": -1.23}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, true, d.Bool.Bool, "value must match")

	// bool
	b = []byte(`{"bool": false}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, false, d.Bool.Bool, "value must match")

	b = []byte(`{"bool": true}`)
	err = json.Unmarshal(b, &d)
	require.NoError(t, err)
	require.Equal(t, true, d.Bool.Valid, "bool must be valid")
	require.Equal(t, true, d.Bool.Bool, "value must match")
}

func TestNullMarshalToJSON(t *testing.T) {
	type Data struct {
		String fuzzy.NullString `json:"string"`
		Int    fuzzy.NullInt    `json:"int"`
		Bool   fuzzy.NullBool   `json:"bool"`
		Float  fuzzy.NullFloat  `json:"float"`
	}

	// Valid values
	d := Data{}
	d.String = fuzzy.NullString(null.StringFrom("foo"))
	d.Int = fuzzy.NullInt(null.IntFrom(123))
	d.Bool = fuzzy.NullBool(null.BoolFrom(true))
	d.Float = fuzzy.NullFloat(null.FloatFrom(1.618))
	b, err := json.Marshal(d)
	require.NoError(t, err)
	require.Equal(t, `{"string":"foo","int":123,"bool":true,"float":1.618}`, string(b))

	// Null values
	d = Data{}
	d.String = fuzzy.NullString(null.String{})
	d.Int = fuzzy.NullInt(null.Int{})
	d.Bool = fuzzy.NullBool(null.Bool{})
	d.Float = fuzzy.NullFloat(null.Float{})
	b, err = json.Marshal(d)
	require.NoError(t, err)
	require.Equal(t, `{"string":null,"int":null,"bool":null,"float":null}`, string(b))

	// Empty value is null
	d = Data{}
	b, err = json.Marshal(d)
	require.NoError(t, err)
	require.Equal(t, `{"string":null,"int":null,"bool":null,"float":null}`, string(b))

	// Missing properties are marshaled as null.
	// Note that Marshal does not output the same as input to Unmarshal
	d = Data{}
	err = json.Unmarshal([]byte("{}"), &d)
	require.NoError(t, err)
	b, err = json.Marshal(d)
	require.NoError(t, err)
	require.Equal(t, `{"string":null,"int":null,"bool":null,"float":null}`, string(b))

	// omitempty does not omit null
	// https://github.com/guregu/null#bugs
	// https://github.com/golang/go/issues/11939
	// However, I'm not sure that should be considered a bug.
	type Data2 struct {
		String fuzzy.NullString `json:"string,omitempty"`
		Int    fuzzy.NullInt    `json:"int,omitempty"`
		Bool   fuzzy.NullBool   `json:"bool,omitempty"`
		Float  fuzzy.NullFloat  `json:"float,omitempty"`
	}
	d2 := Data2{}
	err = json.Unmarshal([]byte("{}"), &d2)
	require.NoError(t, err)
	b, err = json.Marshal(d2)
	require.NoError(t, err)
	require.Equal(t, `{"string":null,"int":null,"bool":null,"float":null}`, string(b))
}
