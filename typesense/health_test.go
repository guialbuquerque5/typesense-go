package typesense

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/guialbuquerque5/typesense-go/typesense/api"
	"github.com/guialbuquerque5/typesense-go/typesense/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestHealthStatus(t *testing.T) {
	tests := []struct {
		ok bool
	}{
		{
			ok: true,
		},
		{
			ok: false,
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)

		mockAPIClient.EXPECT().
			HealthWithResponse(gomock.Not(gomock.Nil())).
			Return(&api.HealthResponse{
				JSON200: &api.HealthStatus{Ok: tt.ok},
			}, nil).
			Times(1)

		client := NewClient(WithAPIClient(mockAPIClient))
		result, err := client.Health(context.Background(), 2*time.Second)
		assert.NoError(t, err)
		assert.Conditionf(t, func() bool {
			return result == tt.ok
		}, "health status expected to be %v", tt.ok)
	}
}

func TestHealthStatusOnApiClientErrorReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)

	mockAPIClient.EXPECT().
		HealthWithResponse(gomock.Not(gomock.Nil())).
		Return(nil, errors.New("failed request")).
		Times(1)

	client := NewClient(WithAPIClient(mockAPIClient))
	result, err := client.Health(context.Background(), 2*time.Second)
	assert.Error(t, err)
	assert.False(t, result)
}

func TestHealthStatusOnHttpStatusErrorCodeReturnsError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockAPIClient := mocks.NewMockAPIClientInterface(ctrl)

	mockAPIClient.EXPECT().
		HealthWithResponse(gomock.Not(gomock.Nil())).
		Return(&api.HealthResponse{
			HTTPResponse: &http.Response{
				StatusCode: 500,
			},
			Body: []byte("Internal Server error"),
		}, nil).
		Times(1)

	client := NewClient(WithAPIClient(mockAPIClient))
	result, err := client.Health(context.Background(), 2*time.Second)
	assert.Error(t, err)
	assert.False(t, result)
}
