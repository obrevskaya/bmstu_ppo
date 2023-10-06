/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DeleteWineResponse struct {
	Deleted bool `json:"deleted"`
}

// AssertDeleteWineResponseRequired checks if the required fields are not zero-ed
func AssertDeleteWineResponseRequired(obj DeleteWineResponse) error {
	elements := map[string]interface{}{
		"deleted": obj.Deleted,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertDeleteWineResponseConstraints checks if the values respects the defined constraints
func AssertDeleteWineResponseConstraints(obj DeleteWineResponse) error {
	return nil
}
