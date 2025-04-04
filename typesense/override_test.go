package typesense

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/guialbuquerque5/typesense-go/typesense/api"
	"github.com/guialbuquerque5/typesense-go/typesense/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSearchOverrideRetrieve(t *testing.T) {
	expectedResult := createNewSearchOverride("customize-apple")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)
	mockedResult := createNewSearchOverride("customize-apple")

	mockAPIClient.EXPECT().
		GetSearchOverrideWithResponse(gomock.Not(gomock.Nil()), "companies", "customize-apple").
		Return(&api.GetSearchOverrideResponse{
			JSON200: mockedResult,
		}, nil).
		Times(1)

	client := NewClient(WithAPIClient(mockAPIClient))
	result, err := client.Collection("companies").Override("customize-apple").Retrieve(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestSearchOverrideRetrieveOnApiClientErrorReturnsError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)

	mockAPIClient.EXPECT().
		GetSearchOverrideWithResponse(gomock.Not(gomock.Nil()), "companies", "customize-apple").
		Return(nil, errors.New("failed request")).
		Times(1)

	client := NewClient(WithAPIClient(mockAPIClient))
	_, err := client.Collection("companies").Override("customize-apple").Retrieve(context.Background())
	assert.NotNil(t, err)
}

func TestSearchOverrideRetrieveOnHttpStatusErrorCodeReturnsError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)

	mockAPIClient.EXPECT().
		GetSearchOverrideWithResponse(gomock.Not(gomock.Nil()), "companies", "customize-apple").
		Return(&api.GetSearchOverrideResponse{
			HTTPResponse: &http.Response{
				StatusCode: 500,
			},
			Body: []byte("Internal Server error"),
		}, nil).
		Times(1)

	client := NewClient(WithAPIClient(mockAPIClient))
	_, err := client.Collection("companies").Override("customize-apple").Retrieve(context.Background())
	assert.NotNil(t, err)
}

func TestSearchOverrideDelete(t *testing.T) {
	expectedResult := &api.SearchOverrideDeleteResponse{Id: "customize-apple"}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)
	mockedResult := &api.SearchOverrideDeleteResponse{Id: "customize-apple"}

	mockAPIClient.EXPECT().
		DeleteSearchOverrideWithResponse(gomock.Not(gomock.Nil()), "companies", "customize-apple").
		Return(&api.DeleteSearchOverrideResponse{
			JSON200: mockedResult,
		}, nil).
		Times(1)

	client := NewClient(WithAPIClient(mockAPIClient))
	result, err := client.Collection("companies").Override("customize-apple").Delete(context.Background())

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestSearchOverrideDeleteOnApiClientErrorReturnsError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)

	mockAPIClient.EXPECT().
		DeleteSearchOverrideWithResponse(gomock.Not(gomock.Nil()), "companies", "customize-apple").
		Return(nil, errors.New("failed request")).
		Times(1)

	client := NewClient(WithAPIClient(mockAPIClient))
	_, err := client.Collection("companies").Override("customize-apple").Delete(context.Background())
	assert.NotNil(t, err)
}

func TestSearchOverrideDeleteOnHttpStatusErrorCodeReturnsError(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)

	mockAPIClient.EXPECT().
		DeleteSearchOverrideWithResponse(gomock.Not(gomock.Nil()), "companies", "customize-apple").
		Return(&api.DeleteSearchOverrideResponse{
			HTTPResponse: &http.Response{
				StatusCode: 500,
			},
			Body: []byte("Internal Server error"),
		}, nil).
		Times(1)

	client := NewClient(WithAPIClient(mockAPIClient))
	_, err := client.Collection("companies").Override("customize-apple").Delete(context.Background())
	assert.NotNil(t, err)
}
