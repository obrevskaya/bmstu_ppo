/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AddElemResponse struct {
	Added bool `json:"added"`
}

// AssertAddElemResponseRequired checks if the required fields are not zero-ed
func AssertAddElemResponseRequired(obj AddElemResponse) error {
	elements := map[string]interface{}{
		"added": obj.Added,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertAddElemResponseConstraints checks if the values respects the defined constraints
func AssertAddElemResponseConstraints(obj AddElemResponse) error {
	return nil
}
