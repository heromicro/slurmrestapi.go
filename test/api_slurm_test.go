/*
Slurm Rest API

Testing SlurmAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package slurmrestapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_slurmrestapi_SlurmAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SlurmAPIService SlurmV0038CancelJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		httpRes, err := apiClient.SlurmAPI.SlurmV0038CancelJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038Diag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038Diag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetNodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetNodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetPartition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partitionName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetPartition(context.Background(), partitionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetPartitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetPartitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetReservation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reservationName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetReservation(context.Background(), reservationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038GetReservations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038GetReservations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038Ping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038Ping(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038SlurmctldGetLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038SlurmctldGetLicenses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038SubmitJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0038SubmitJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0038UpdateJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		httpRes, err := apiClient.SlurmAPI.SlurmV0038UpdateJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039CancelJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039CancelJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039DeleteNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039DeleteNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039Diag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039Diag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039GetJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039GetJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039GetNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039GetNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039GetNodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039GetNodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039GetPartition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partitionName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039GetPartition(context.Background(), partitionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039GetPartitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039GetPartitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039GetReservation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reservationName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039GetReservation(context.Background(), reservationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039GetReservations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039GetReservations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039Ping", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039Ping(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039SlurmctldGetLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039SlurmctldGetLicenses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039SubmitJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039SubmitJob(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039UpdateJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039UpdateJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0039UpdateNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0039UpdateNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040DeleteJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040DeleteJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040DeleteNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040DeleteNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetDiag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetDiag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetLicenses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetLicenses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetNodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetNodes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetPartition", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var partitionName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetPartition(context.Background(), partitionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetPartitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetPartitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetPing", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetPing(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetReconfigure", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetReconfigure(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetReservation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var reservationName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetReservation(context.Background(), reservationName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetReservations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetReservations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040GetShares", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040GetShares(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040PostJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040PostJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040PostJobSubmit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040PostJobSubmit(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmV0040PostNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nodeName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmV0040PostNode(context.Background(), nodeName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038AddClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038AddClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038AddWckeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038AddWckeys(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038DeleteAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038DeleteAccount(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038DeleteAssociation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038DeleteAssociation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038DeleteAssociations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038DeleteAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038DeleteCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clusterName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038DeleteCluster(context.Background(), clusterName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038DeleteQos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var qosName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038DeleteQos(context.Background(), qosName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038DeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038DeleteUser(context.Background(), userName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038DeleteWckey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var wckey string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038DeleteWckey(context.Background(), wckey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038Diag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038Diag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetAccount(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetAccounts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetAccounts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetAssociation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetAssociation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetAssociations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clusterName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetCluster(context.Background(), clusterName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetQos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetQos(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetSingleQos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var qosName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetSingleQos(context.Background(), qosName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetTres", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetTres(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetUser(context.Background(), userName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetWckey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var wckey string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetWckey(context.Background(), wckey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038GetWckeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038GetWckeys(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038SetConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038SetConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038UpdateAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038UpdateAccount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038UpdateAssociations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038UpdateAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038UpdateQos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038UpdateQos(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038UpdateTres", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038UpdateTres(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0038UpdateUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0038UpdateUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039AddClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039AddClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039AddWckeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039AddWckeys(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039DeleteAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039DeleteAccount(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039DeleteAssociation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039DeleteAssociation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039DeleteAssociations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039DeleteAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039DeleteCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clusterName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039DeleteCluster(context.Background(), clusterName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039DeleteQos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var qosName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039DeleteQos(context.Background(), qosName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039DeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039DeleteUser(context.Background(), userName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039DeleteWckey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var wckey string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039DeleteWckey(context.Background(), wckey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039Diag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039Diag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetAccount(context.Background(), accountName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetAccounts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetAccounts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetAssociation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetAssociation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetAssociations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clusterName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetCluster(context.Background(), clusterName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetClusters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetJob(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetJobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetJobs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetQos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetQos(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetSingleQos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var qosName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetSingleQos(context.Background(), qosName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetTres", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetTres(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userName string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetUser(context.Background(), userName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetWckey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var wckey string

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetWckey(context.Background(), wckey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039GetWckeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039GetWckeys(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039SetConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039SetConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039UpdateAccounts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039UpdateAccounts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039UpdateAssociations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039UpdateAssociations(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039UpdateQos", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039UpdateQos(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039UpdateTres", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039UpdateTres(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SlurmAPIService SlurmdbV0039UpdateUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SlurmAPI.SlurmdbV0039UpdateUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
