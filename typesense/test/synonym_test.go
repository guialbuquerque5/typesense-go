//go:build integration
// +build integration

package test

import (
	"context"
	"testing"

	"github.com/guialbuquerque5/typesense-go/typesense/api/pointer"
	"github.com/stretchr/testify/require"
)

func TestSearchSynonymRetrieve(t *testing.T) {
	collectionName := createNewCollection(t, "products")
	synonymID := newUUIDName("customize-apple")
	expectedResult := newSearchSynonym(synonymID)

	body := newSearchSynonymSchema()
	_, err := typesenseClient.Collection(collectionName).Synonyms().Upsert(context.Background(), synonymID, body)
	require.NoError(t, err)

	result, err := typesenseClient.Collection(collectionName).Synonym(synonymID).Retrieve(context.Background())

	require.NoError(t, err)
	expectedResult.Root = pointer.String("")
	require.Equal(t, expectedResult, result)
}

func TestSearchSynonymDelete(t *testing.T) {
	collectionName := createNewCollection(t, "products")
	synonymID := newUUIDName("customize-apple")

	body := newSearchSynonymSchema()
	_, err := typesenseClient.Collection(collectionName).Synonyms().Upsert(context.Background(), synonymID, body)
	require.NoError(t, err)

	result, err := typesenseClient.Collection(collectionName).Synonym(synonymID).Delete(context.Background())

	require.NoError(t, err)
	require.Equal(t, synonymID, result.Id)

	_, err = typesenseClient.Collection(collectionName).Synonym(synonymID).Retrieve(context.Background())
	require.Error(t, err)
}
