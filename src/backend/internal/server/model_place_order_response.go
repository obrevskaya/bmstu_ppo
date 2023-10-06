/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PlaceOrderResponse struct {
	Placed bool `json:"placed"`
}

// AssertPlaceOrderResponseRequired checks if the required fields are not zero-ed
func AssertPlaceOrderResponseRequired(obj PlaceOrderResponse) error {
	elements := map[string]interface{}{
		"placed": obj.Placed,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPlaceOrderResponseConstraints checks if the values respects the defined constraints
func AssertPlaceOrderResponseConstraints(obj PlaceOrderResponse) error {
	return nil
}