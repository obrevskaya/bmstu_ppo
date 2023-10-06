/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type GetByOrderRequest struct {
	Id string `json:"id"`
}

// AssertGetByOrderRequestRequired checks if the required fields are not zero-ed
func AssertGetByOrderRequestRequired(obj GetByOrderRequest) error {
	elements := map[string]interface{}{
		"id": obj.Id,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertGetByOrderRequestConstraints checks if the values respects the defined constraints
func AssertGetByOrderRequestConstraints(obj GetByOrderRequest) error {
	return nil
}
