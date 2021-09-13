package fuzzy

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// String can be used to decode any JSON value to string
type String string

// MarshalJSON method with value receiver for String
// Method must not have a pointer receiver!
// See https://stackoverflow.com/a/21394657/639133
func (fs String) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, fs)), nil
}

// UnmarshalJSON for String
func (fs *String) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", uint64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fs = ""
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		*fs = String(s)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fs = String(bArr[:])
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fs = String(bArr[:])
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*fs = String(bArr[:])
		return
	}

	return
}

// Int can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type Int int64

// MarshalJSON method for Int
func (fi Int) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(fi), 10)), nil
}

// UnmarshalJSON method for Int
func (fi *Int) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fi = Int(0)
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		i, err2 := strconv.ParseInt(s, 10, 64)
		if err2 != nil {
			// Value is null if int could not be parsed from the string
			//*fi = Int(0) // This is not a good idea...
			return err2
		}
		*fi = Int(i)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fi = Int(i)
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fi = Int(int64(f))
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		//*fi = Int(0) // This is not a good idea...
		return errors.WithStack(fmt.Errorf("value is a bool"))
	}

	return
}

// Float can be used to decode any JSON value to int64.
// Strings that are not valid representation of a number will error.
// Boolean values will error
type Float float64

// MarshalJSON method for Float
func (fi Float) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatFloat(float64(fi), 'f', -1, 64)), nil
}

// UnmarshalJSON method for Float
func (fi *Float) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fi = Float(0)
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		i, err2 := strconv.ParseFloat(s, 10)
		if err2 != nil {
			// Value is null if int could not be parsed from the string
			//*fi = Float(0) // This is not a good idea...
			return err2
		}
		*fi = Float(i)
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		*fi = Float(i)
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		*fi = Float(f)
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		//*fi = Float(0) // This is not a good idea...
		return errors.WithStack(fmt.Errorf("value is a bool"))
	}

	return
}

// Bool can be used to decode any JSON value to bool.
// Empty strings as well as "false" and "0" evaluate to false,
// all other strings are true.
// Numbers equal to 0 will evaluate to false,
// all other numbers are true.
type Bool bool

// MarshalJSON method for Bool
func (fb Bool) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatBool(bool(fb))), nil
}

// UnmarshalJSON method for Bool
func (fb *Bool) UnmarshalJSON(bArr []byte) (err error) {
	s, i, f, b :=
		"", int64(0), float64(0), false

	// Value is null
	if string(bArr) == "null" {
		*fb = false
		return
	}

	// Value is a...
	// string
	if err = json.Unmarshal(bArr, &s); err == nil {
		compare := strings.ToLower(strings.TrimSpace(s))
		if compare == "false" || compare == "0" || compare == "" {
			*fb = false
			return
		}
		*fb = true
		return
	}

	// int
	if err = json.Unmarshal(bArr, &i); err == nil {
		if i == 0 {
			*fb = false
			return
		}
		*fb = true
		return
	}

	// float
	if err = json.Unmarshal(bArr, &f); err == nil {
		if f == 0 {
			*fb = false
			return
		}
		*fb = true
		return
	}

	// bool
	if err = json.Unmarshal(bArr, &b); err == nil {
		*fb = Bool(b)
		return
	}

	return
}
