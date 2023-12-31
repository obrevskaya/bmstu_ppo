/*
 * API for ppo project
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type CreateElems struct {
	Elems []Elem `json:"elems"`
}

// AssertCreateElemsRequired checks if the required fields are not zero-ed
func AssertCreateElemsRequired(obj CreateElems) error {
	elements := map[string]interface{}{
		"elems": obj.Elems,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Elems {
		if err := AssertElemRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCreateElemsConstraints checks if the values respects the defined constraints
func AssertCreateElemsConstraints(obj CreateElems) error {
	return nil
}
