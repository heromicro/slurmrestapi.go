/*
Slurm Rest API

Testing OpenapiAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package slurmrestapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/heromicro/slurmrestapi.go"
)

func Test_slurmrestapi_OpenapiAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OpenapiAPIService OpenapiGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OpenapiAPI.OpenapiGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenapiAPIService OpenapiJsonGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OpenapiAPI.OpenapiJsonGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenapiAPIService OpenapiV3Get", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OpenapiAPI.OpenapiV3Get(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OpenapiAPIService OpenapiYamlGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OpenapiAPI.OpenapiYamlGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
