package compactarrays

import (
	"encoding/json"
)

// CompactMarshalIndent marshals a slice of structs and compacts []int fields.
func CompactMarshalIndent[T any](data []T, intFields []string, prefix, indent string) ([]byte, error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var objs []map[string]interface{}
	if err := json.Unmarshal(raw, &objs); err != nil {
		return nil, err
	}

	for _, obj := range objs {
		for _, field := range intFields {
			if val, ok := obj[field]; ok {
				if arr, ok := val.([]interface{}); ok {
					var compact []int
					for _, v := range arr {
						if f, ok := v.(float64); ok {
							compact = append(compact, int(f))
						}
					}
					obj[field] = compact
				}
			}
		}
	}

	return json.MarshalIndent(objs, prefix, indent)
}
