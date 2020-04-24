package pkg

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Params ...
type Params struct {
	Values map[string]interface{}
}

// ErrConflict ...
type errConflict struct {
	keys []string
}

func (e *errConflict) Error() string {
	return fmt.Sprintf("parameters %s and %s conflict", e.keys[0], e.keys[1])
}

// IsConflictErr ...
func IsConflictErr(err error) (bool, []string) {
	if e, ok := err.(*errConflict); ok {
		return true, e.keys
	}
	return false, nil
}

func descKeyFrom(params map[string]interface{}) string {
	for key, val := range params {
		switch t := val.(type) {
		case map[string]interface{}:
			return fmt.Sprintf("%s.%s", key, descKeyFrom(t))
		default:
			return key
		}
	}
	return ""
}

func setVal(
	params map[string]interface{},
	key string,
	val string) error {
	keys := strings.Split(key, ".")

	last := keys[len(keys)-1]
	rest := keys[:len(keys)-1]

	for _, sk := range rest {
		switch t := params[sk].(type) {
		case nil:
			nm := map[string]interface{}{}
			params[sk] = nm
			params = nm
		case map[string]interface{}:
			params = t
		default:
			return &errConflict{
				keys: []string{
					key,
					strings.Join(rest, "."),
				},
			}
		}
	}

	switch t := params[last].(type) {
	case nil:
		params[last] = val
	case string:
		params[last] = []string{t, val}
	case []string:
		params[last] = append(t, val)
	case map[string]interface{}:
		return &errConflict{
			keys: []string{
				key,
				fmt.Sprintf("%s.%s", key, descKeyFrom(t)),
			},
		}
	default:
		return fmt.Errorf("unexpected error on key: %s", key)
	}

	return nil
}

// Set ...
func (p *Params) Set(value string) error {
	if p.Values == nil {
		p.Values = map[string]interface{}{}
	}

	ix := strings.Index(value, "=")
	if ix < 0 {
		return setVal(p.Values, value, "")
	}
	return setVal(p.Values, value[:ix], value[ix+1:])
}

func (p *Params) String() string {
	if v := p.Values; v != nil {
		b, err := json.Marshal(v)
		if err == nil {
			return string(b)
		}
	}
	return ""
}
