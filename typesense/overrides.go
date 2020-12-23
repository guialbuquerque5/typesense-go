package typesense

import (
	"context"

	"github.com/v-byte-cpu/typesense-go/typesense/api"
)

// OverridesInterface is a type for Search Overrides API operations
type OverridesInterface interface {
	Upsert(overrideID string, overrideSchema *api.SearchOverrideSchema) (*api.SearchOverride, error)
	Retrieve() ([]*api.SearchOverride, error)
}

// overrides is internal implementation of OverridesInterface
type overrides struct {
	apiClient      api.ClientWithResponsesInterface
	collectionName string
}

func (o *overrides) Upsert(overrideID string, overrideSchema *api.SearchOverrideSchema) (*api.SearchOverride, error) {
	response, err := o.apiClient.UpsertSearchOverrideWithResponse(context.Background(),
		o.collectionName, overrideID, api.UpsertSearchOverrideJSONRequestBody(*overrideSchema))
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &httpError{status: response.StatusCode(), body: response.Body}
	}
	return response.JSON200, nil
}

func (o *overrides) Retrieve() ([]*api.SearchOverride, error) {
	response, err := o.apiClient.GetSearchOverridesWithResponse(context.Background(), o.collectionName)
	if err != nil {
		return nil, err
	}
	if response.JSON200 == nil {
		return nil, &httpError{status: response.StatusCode(), body: response.Body}
	}
	return response.JSON200.Overrides, nil
}
