/*
Slurm Rest API

Testing SlurmApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package slurmrestapi

import (
	"context"
	"testing"

	openapiclient "github.com/heromicro/slurmrestapi.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_slurmrestapi_SlurmApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SlurmApiService SlurmV0039CancelJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039CancelJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039DeleteNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039DeleteNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039Diag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039Diag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039GetJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039GetJobs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039GetNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039GetNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039GetNodes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039GetNodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039GetPartition", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var partitionName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039GetPartition(context.Background(), partitionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039GetPartitions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039GetPartitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039GetReservation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var reservationName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039GetReservation(context.Background(), reservationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039GetReservations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039GetReservations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039Ping", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039Ping(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039SlurmctldGetLicenses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039SlurmctldGetLicenses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039SubmitJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039SubmitJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039UpdateJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039UpdateJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmV0039UpdateNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmV0039UpdateNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039AddClusters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039AddClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039AddWckeys", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039AddWckeys(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039DeleteAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039DeleteAccount(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039DeleteAssociation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039DeleteAssociation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039DeleteAssociations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039DeleteAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039DeleteCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039DeleteCluster(context.Background(), clusterName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039DeleteQos", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var qosName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039DeleteQos(context.Background(), qosName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039DeleteUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039DeleteUser(context.Background(), userName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039DeleteWckey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var wckey string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039DeleteWckey(context.Background(), wckey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039Diag", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039Diag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetAccount", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetAccount(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetAccounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetAccounts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetAssociation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetAssociation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetAssociations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetCluster", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clusterName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetCluster(context.Background(), clusterName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetClusters", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetJob", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetJobs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetQos", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetQos(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetSingleQos", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var qosName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetSingleQos(context.Background(), qosName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetTres", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetTres(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userName string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetUser(context.Background(), userName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetWckey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var wckey string

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetWckey(context.Background(), wckey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039GetWckeys", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039GetWckeys(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039SetConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039SetConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039UpdateAccounts", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039UpdateAccounts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039UpdateAssociations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039UpdateAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039UpdateQos", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039UpdateQos(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039UpdateTres", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039UpdateTres(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmApiService SlurmdbV0039UpdateUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SlurmApi.SlurmdbV0039UpdateUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
