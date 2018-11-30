package datastore

import (
	"errors"
)

type DataStoreDefinition struct {
	Name        string                `json:"name"`
	Type        string                `json:"type"`
	Validations string                `json:"validations"`
	Fields      []DataStoreDefinition `json:"fields"`
}

type DataStore struct {
	Definition DataStoreDefinition
	Data       map[string]interface{}
}

func NewDataStore(definition DataStoreDefinition) DataStore {
	return DataStore{
		Definition: definition,
	}
}

func buildDataFromDefinition(raw_data interface{}, definition DataStoreDefinition) (interface{}, error) {
	switch definition.Type {
	case "int":
		data, ok := raw_data.(float64)

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		return int64(data), nil
	case "float":
		data, ok := raw_data.(float64)

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		return data, nil
	case "string":
		data, ok := raw_data.(string)

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		return data, nil
	case "boolean":
		data, ok := raw_data.(bool)

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		return data, nil
	case "object":
		unfiltered_data, ok := raw_data.(map[string]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		data := make(map[string]interface{})

		for _, field := range definition.Fields {
			unfiltered_fields, exists := unfiltered_data[field.Name]

			if exists {
				var err error
				data[field.Name], err = buildDataFromDefinition(unfiltered_fields, field)

				if err != nil {
					return nil, err
				}
			} else {
				data[field.Name] = defaultValueForType(field.Type)
			}
		}

		return data, nil
	case "[]int":
		data, ok := raw_data.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		int_data := make([]int64, len(data))

		for i := 0; i < len(data); i++ {
			element, ok := data[i].(float64)

			if !ok {
				return nil, errors.New("Type mismatch in data and definition")
			}

			int_data[i] = int64(element)
		}

		return int_data, nil
	case "[]float":
		data, ok := raw_data.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		float_data := make([]float64, len(data))

		for i := 0; i < len(data); i++ {
			element, ok := data[i].(float64)

			if !ok {
				return nil, errors.New("Type mismatch in data and definition")
			}

			float_data[i] = element
		}

		return float_data, nil
	case "[]string":
		data, ok := raw_data.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		string_data := make([]string, len(data))

		for i := 0; i < len(data); i++ {
			element, ok := data[i].(string)

			if !ok {
				return nil, errors.New("Type mismatch in data and definition")
			}

			string_data[i] = element
		}

		return string_data, nil
	case "[]boolean":
		data, ok := raw_data.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		bool_data := make([]bool, len(data))

		for i := 0; i < len(data); i++ {
			element, ok := data[i].(bool)

			if !ok {
				return nil, errors.New("Type mismatch in data and definition")
			}

			bool_data[i] = element
		}

		return bool_data, nil
	case "[]object":
		unfiltered_data, ok := raw_data.([]interface{})

		if !ok {
			return nil, errors.New("Type mismatch in data and definition")
		}

		data := make([]map[string]interface{}, len(unfiltered_data))

		for i := 0; i < len(unfiltered_data); i++ {
			element := make(map[string]interface{})

			for _, field := range definition.Fields {
				unfiltered_data_element, ok := unfiltered_data[i].(map[string]interface{})

				if !ok {
					return nil, errors.New("Type mismatch in data and definition")
				}

				unfiltered_fields, exists := unfiltered_data_element[field.Name]

				if exists {
					var err error
					element[field.Name], err = buildDataFromDefinition(unfiltered_fields, field)

					if err != nil {
						return nil, err
					}
				} else {
					element[field.Name] = defaultValueForType(field.Type)
				}
			}

			data[i] = element
		}

		return data, nil
	default:
		return nil, errors.New("Invalid type in definition")
	}
}

func defaultValueForType(tpe string) interface{} {
	switch tpe {
	case "string":
		return ""
	case "int":
		return 0
	case "float":
		return 0.0
	case "boolean":
		return false
	case "object":
		return nil
	case "[]string":
		return nil
	case "[]int":
		return nil
	case "[]float":
		return nil
	case "[]boolean":
		return nil
	case "[]object":
		return nil
	default:
		return nil
	}
}
