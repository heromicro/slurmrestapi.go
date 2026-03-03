# \SlurmdbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmdbV0041DeleteAccount**](SlurmdbAPI.md#SlurmdbV0041DeleteAccount) | **Delete** /slurmdb/v0.0.41/account/{account_name} | Delete account
[**SlurmdbV0041DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0041DeleteAssociation) | **Delete** /slurmdb/v0.0.41/association/ | Delete association
[**SlurmdbV0041DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0041DeleteAssociations) | **Delete** /slurmdb/v0.0.41/associations/ | Delete associations
[**SlurmdbV0041DeleteCluster**](SlurmdbAPI.md#SlurmdbV0041DeleteCluster) | **Delete** /slurmdb/v0.0.41/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0041DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0041DeleteSingleQos) | **Delete** /slurmdb/v0.0.41/qos/{qos} | Delete QOS
[**SlurmdbV0041DeleteUser**](SlurmdbAPI.md#SlurmdbV0041DeleteUser) | **Delete** /slurmdb/v0.0.41/user/{name} | Delete user
[**SlurmdbV0041DeleteWckey**](SlurmdbAPI.md#SlurmdbV0041DeleteWckey) | **Delete** /slurmdb/v0.0.41/wckey/{id} | Delete wckey
[**SlurmdbV0041GetAccount**](SlurmdbAPI.md#SlurmdbV0041GetAccount) | **Get** /slurmdb/v0.0.41/account/{account_name} | Get account info
[**SlurmdbV0041GetAccounts**](SlurmdbAPI.md#SlurmdbV0041GetAccounts) | **Get** /slurmdb/v0.0.41/accounts/ | Get account list
[**SlurmdbV0041GetAssociation**](SlurmdbAPI.md#SlurmdbV0041GetAssociation) | **Get** /slurmdb/v0.0.41/association/ | Get association info
[**SlurmdbV0041GetAssociations**](SlurmdbAPI.md#SlurmdbV0041GetAssociations) | **Get** /slurmdb/v0.0.41/associations/ | Get association list
[**SlurmdbV0041GetCluster**](SlurmdbAPI.md#SlurmdbV0041GetCluster) | **Get** /slurmdb/v0.0.41/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0041GetClusters**](SlurmdbAPI.md#SlurmdbV0041GetClusters) | **Get** /slurmdb/v0.0.41/clusters/ | Get cluster list
[**SlurmdbV0041GetConfig**](SlurmdbAPI.md#SlurmdbV0041GetConfig) | **Get** /slurmdb/v0.0.41/config | Dump all configuration information
[**SlurmdbV0041GetDiag**](SlurmdbAPI.md#SlurmdbV0041GetDiag) | **Get** /slurmdb/v0.0.41/diag/ | Get slurmdb diagnostics
[**SlurmdbV0041GetInstance**](SlurmdbAPI.md#SlurmdbV0041GetInstance) | **Get** /slurmdb/v0.0.41/instance/ | Get instance info
[**SlurmdbV0041GetInstances**](SlurmdbAPI.md#SlurmdbV0041GetInstances) | **Get** /slurmdb/v0.0.41/instances/ | Get instance list
[**SlurmdbV0041GetJob**](SlurmdbAPI.md#SlurmdbV0041GetJob) | **Get** /slurmdb/v0.0.41/job/{job_id} | Get job info
[**SlurmdbV0041GetJobs**](SlurmdbAPI.md#SlurmdbV0041GetJobs) | **Get** /slurmdb/v0.0.41/jobs/ | Get job list
[**SlurmdbV0041GetQos**](SlurmdbAPI.md#SlurmdbV0041GetQos) | **Get** /slurmdb/v0.0.41/qos/ | Get QOS list
[**SlurmdbV0041GetSingleQos**](SlurmdbAPI.md#SlurmdbV0041GetSingleQos) | **Get** /slurmdb/v0.0.41/qos/{qos} | Get QOS info
[**SlurmdbV0041GetTres**](SlurmdbAPI.md#SlurmdbV0041GetTres) | **Get** /slurmdb/v0.0.41/tres/ | Get TRES info
[**SlurmdbV0041GetUser**](SlurmdbAPI.md#SlurmdbV0041GetUser) | **Get** /slurmdb/v0.0.41/user/{name} | Get user info
[**SlurmdbV0041GetUsers**](SlurmdbAPI.md#SlurmdbV0041GetUsers) | **Get** /slurmdb/v0.0.41/users/ | Get user list
[**SlurmdbV0041GetWckey**](SlurmdbAPI.md#SlurmdbV0041GetWckey) | **Get** /slurmdb/v0.0.41/wckey/{id} | Get wckey info
[**SlurmdbV0041GetWckeys**](SlurmdbAPI.md#SlurmdbV0041GetWckeys) | **Get** /slurmdb/v0.0.41/wckeys/ | Get wckey list
[**SlurmdbV0041PostAccounts**](SlurmdbAPI.md#SlurmdbV0041PostAccounts) | **Post** /slurmdb/v0.0.41/accounts/ | Add/update list of accounts
[**SlurmdbV0041PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0041PostAccountsAssociation) | **Post** /slurmdb/v0.0.41/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0041PostAssociations**](SlurmdbAPI.md#SlurmdbV0041PostAssociations) | **Post** /slurmdb/v0.0.41/associations/ | Set associations info
[**SlurmdbV0041PostClusters**](SlurmdbAPI.md#SlurmdbV0041PostClusters) | **Post** /slurmdb/v0.0.41/clusters/ | Get cluster list
[**SlurmdbV0041PostConfig**](SlurmdbAPI.md#SlurmdbV0041PostConfig) | **Post** /slurmdb/v0.0.41/config | Load all configuration information
[**SlurmdbV0041PostQos**](SlurmdbAPI.md#SlurmdbV0041PostQos) | **Post** /slurmdb/v0.0.41/qos/ | Add or update QOSs
[**SlurmdbV0041PostTres**](SlurmdbAPI.md#SlurmdbV0041PostTres) | **Post** /slurmdb/v0.0.41/tres/ | Add TRES
[**SlurmdbV0041PostUsers**](SlurmdbAPI.md#SlurmdbV0041PostUsers) | **Post** /slurmdb/v0.0.41/users/ | Update users
[**SlurmdbV0041PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0041PostUsersAssociation) | **Post** /slurmdb/v0.0.41/users_association/ | Add users with conditional association
[**SlurmdbV0041PostWckeys**](SlurmdbAPI.md#SlurmdbV0041PostWckeys) | **Post** /slurmdb/v0.0.41/wckeys/ | Add or update wckeys
[**SlurmdbV0042DeleteAccount**](SlurmdbAPI.md#SlurmdbV0042DeleteAccount) | **Delete** /slurmdb/v0.0.42/account/{account_name} | Delete account
[**SlurmdbV0042DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0042DeleteAssociation) | **Delete** /slurmdb/v0.0.42/association/ | Delete association
[**SlurmdbV0042DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0042DeleteAssociations) | **Delete** /slurmdb/v0.0.42/associations/ | Delete associations
[**SlurmdbV0042DeleteCluster**](SlurmdbAPI.md#SlurmdbV0042DeleteCluster) | **Delete** /slurmdb/v0.0.42/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0042DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0042DeleteSingleQos) | **Delete** /slurmdb/v0.0.42/qos/{qos} | Delete QOS
[**SlurmdbV0042DeleteUser**](SlurmdbAPI.md#SlurmdbV0042DeleteUser) | **Delete** /slurmdb/v0.0.42/user/{name} | Delete user
[**SlurmdbV0042DeleteWckey**](SlurmdbAPI.md#SlurmdbV0042DeleteWckey) | **Delete** /slurmdb/v0.0.42/wckey/{id} | Delete wckey
[**SlurmdbV0042GetAccount**](SlurmdbAPI.md#SlurmdbV0042GetAccount) | **Get** /slurmdb/v0.0.42/account/{account_name} | Get account info
[**SlurmdbV0042GetAccounts**](SlurmdbAPI.md#SlurmdbV0042GetAccounts) | **Get** /slurmdb/v0.0.42/accounts/ | Get account list
[**SlurmdbV0042GetAssociation**](SlurmdbAPI.md#SlurmdbV0042GetAssociation) | **Get** /slurmdb/v0.0.42/association/ | Get association info
[**SlurmdbV0042GetAssociations**](SlurmdbAPI.md#SlurmdbV0042GetAssociations) | **Get** /slurmdb/v0.0.42/associations/ | Get association list
[**SlurmdbV0042GetCluster**](SlurmdbAPI.md#SlurmdbV0042GetCluster) | **Get** /slurmdb/v0.0.42/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0042GetClusters**](SlurmdbAPI.md#SlurmdbV0042GetClusters) | **Get** /slurmdb/v0.0.42/clusters/ | Get cluster list
[**SlurmdbV0042GetConfig**](SlurmdbAPI.md#SlurmdbV0042GetConfig) | **Get** /slurmdb/v0.0.42/config | Dump all configuration information
[**SlurmdbV0042GetDiag**](SlurmdbAPI.md#SlurmdbV0042GetDiag) | **Get** /slurmdb/v0.0.42/diag/ | Get slurmdb diagnostics
[**SlurmdbV0042GetInstance**](SlurmdbAPI.md#SlurmdbV0042GetInstance) | **Get** /slurmdb/v0.0.42/instance/ | Get instance info
[**SlurmdbV0042GetInstances**](SlurmdbAPI.md#SlurmdbV0042GetInstances) | **Get** /slurmdb/v0.0.42/instances/ | Get instance list
[**SlurmdbV0042GetJob**](SlurmdbAPI.md#SlurmdbV0042GetJob) | **Get** /slurmdb/v0.0.42/job/{job_id} | Get job info
[**SlurmdbV0042GetJobs**](SlurmdbAPI.md#SlurmdbV0042GetJobs) | **Get** /slurmdb/v0.0.42/jobs/ | Get job list
[**SlurmdbV0042GetPing**](SlurmdbAPI.md#SlurmdbV0042GetPing) | **Get** /slurmdb/v0.0.42/ping/ | ping test
[**SlurmdbV0042GetQos**](SlurmdbAPI.md#SlurmdbV0042GetQos) | **Get** /slurmdb/v0.0.42/qos/ | Get QOS list
[**SlurmdbV0042GetSingleQos**](SlurmdbAPI.md#SlurmdbV0042GetSingleQos) | **Get** /slurmdb/v0.0.42/qos/{qos} | Get QOS info
[**SlurmdbV0042GetTres**](SlurmdbAPI.md#SlurmdbV0042GetTres) | **Get** /slurmdb/v0.0.42/tres/ | Get TRES info
[**SlurmdbV0042GetUser**](SlurmdbAPI.md#SlurmdbV0042GetUser) | **Get** /slurmdb/v0.0.42/user/{name} | Get user info
[**SlurmdbV0042GetUsers**](SlurmdbAPI.md#SlurmdbV0042GetUsers) | **Get** /slurmdb/v0.0.42/users/ | Get user list
[**SlurmdbV0042GetWckey**](SlurmdbAPI.md#SlurmdbV0042GetWckey) | **Get** /slurmdb/v0.0.42/wckey/{id} | Get wckey info
[**SlurmdbV0042GetWckeys**](SlurmdbAPI.md#SlurmdbV0042GetWckeys) | **Get** /slurmdb/v0.0.42/wckeys/ | Get wckey list
[**SlurmdbV0042PostAccounts**](SlurmdbAPI.md#SlurmdbV0042PostAccounts) | **Post** /slurmdb/v0.0.42/accounts/ | Add/update list of accounts
[**SlurmdbV0042PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0042PostAccountsAssociation) | **Post** /slurmdb/v0.0.42/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0042PostAssociations**](SlurmdbAPI.md#SlurmdbV0042PostAssociations) | **Post** /slurmdb/v0.0.42/associations/ | Set associations info
[**SlurmdbV0042PostClusters**](SlurmdbAPI.md#SlurmdbV0042PostClusters) | **Post** /slurmdb/v0.0.42/clusters/ | Get cluster list
[**SlurmdbV0042PostConfig**](SlurmdbAPI.md#SlurmdbV0042PostConfig) | **Post** /slurmdb/v0.0.42/config | Load all configuration information
[**SlurmdbV0042PostQos**](SlurmdbAPI.md#SlurmdbV0042PostQos) | **Post** /slurmdb/v0.0.42/qos/ | Add or update QOSs
[**SlurmdbV0042PostTres**](SlurmdbAPI.md#SlurmdbV0042PostTres) | **Post** /slurmdb/v0.0.42/tres/ | Add TRES
[**SlurmdbV0042PostUsers**](SlurmdbAPI.md#SlurmdbV0042PostUsers) | **Post** /slurmdb/v0.0.42/users/ | Update users
[**SlurmdbV0042PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0042PostUsersAssociation) | **Post** /slurmdb/v0.0.42/users_association/ | Add users with conditional association
[**SlurmdbV0042PostWckeys**](SlurmdbAPI.md#SlurmdbV0042PostWckeys) | **Post** /slurmdb/v0.0.42/wckeys/ | Add or update wckeys
[**SlurmdbV0043DeleteAccount**](SlurmdbAPI.md#SlurmdbV0043DeleteAccount) | **Delete** /slurmdb/v0.0.43/account/{account_name} | Delete account
[**SlurmdbV0043DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0043DeleteAssociation) | **Delete** /slurmdb/v0.0.43/association/ | Delete association
[**SlurmdbV0043DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0043DeleteAssociations) | **Delete** /slurmdb/v0.0.43/associations/ | Delete associations
[**SlurmdbV0043DeleteCluster**](SlurmdbAPI.md#SlurmdbV0043DeleteCluster) | **Delete** /slurmdb/v0.0.43/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0043DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0043DeleteSingleQos) | **Delete** /slurmdb/v0.0.43/qos/{qos} | Delete QOS
[**SlurmdbV0043DeleteUser**](SlurmdbAPI.md#SlurmdbV0043DeleteUser) | **Delete** /slurmdb/v0.0.43/user/{name} | Delete user
[**SlurmdbV0043DeleteWckey**](SlurmdbAPI.md#SlurmdbV0043DeleteWckey) | **Delete** /slurmdb/v0.0.43/wckey/{id} | Delete wckey
[**SlurmdbV0043GetAccount**](SlurmdbAPI.md#SlurmdbV0043GetAccount) | **Get** /slurmdb/v0.0.43/account/{account_name} | Get account info
[**SlurmdbV0043GetAccounts**](SlurmdbAPI.md#SlurmdbV0043GetAccounts) | **Get** /slurmdb/v0.0.43/accounts/ | Get account list
[**SlurmdbV0043GetAssociation**](SlurmdbAPI.md#SlurmdbV0043GetAssociation) | **Get** /slurmdb/v0.0.43/association/ | Get association info
[**SlurmdbV0043GetAssociations**](SlurmdbAPI.md#SlurmdbV0043GetAssociations) | **Get** /slurmdb/v0.0.43/associations/ | Get association list
[**SlurmdbV0043GetCluster**](SlurmdbAPI.md#SlurmdbV0043GetCluster) | **Get** /slurmdb/v0.0.43/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0043GetClusters**](SlurmdbAPI.md#SlurmdbV0043GetClusters) | **Get** /slurmdb/v0.0.43/clusters/ | Get cluster list
[**SlurmdbV0043GetConfig**](SlurmdbAPI.md#SlurmdbV0043GetConfig) | **Get** /slurmdb/v0.0.43/config | Dump all configuration information
[**SlurmdbV0043GetDiag**](SlurmdbAPI.md#SlurmdbV0043GetDiag) | **Get** /slurmdb/v0.0.43/diag/ | Get slurmdb diagnostics
[**SlurmdbV0043GetInstance**](SlurmdbAPI.md#SlurmdbV0043GetInstance) | **Get** /slurmdb/v0.0.43/instance/ | Get instance info
[**SlurmdbV0043GetInstances**](SlurmdbAPI.md#SlurmdbV0043GetInstances) | **Get** /slurmdb/v0.0.43/instances/ | Get instance list
[**SlurmdbV0043GetJob**](SlurmdbAPI.md#SlurmdbV0043GetJob) | **Get** /slurmdb/v0.0.43/job/{job_id} | Get job info
[**SlurmdbV0043GetJobs**](SlurmdbAPI.md#SlurmdbV0043GetJobs) | **Get** /slurmdb/v0.0.43/jobs/ | Get job list
[**SlurmdbV0043GetPing**](SlurmdbAPI.md#SlurmdbV0043GetPing) | **Get** /slurmdb/v0.0.43/ping/ | ping test
[**SlurmdbV0043GetQos**](SlurmdbAPI.md#SlurmdbV0043GetQos) | **Get** /slurmdb/v0.0.43/qos/ | Get QOS list
[**SlurmdbV0043GetSingleQos**](SlurmdbAPI.md#SlurmdbV0043GetSingleQos) | **Get** /slurmdb/v0.0.43/qos/{qos} | Get QOS info
[**SlurmdbV0043GetTres**](SlurmdbAPI.md#SlurmdbV0043GetTres) | **Get** /slurmdb/v0.0.43/tres/ | Get TRES info
[**SlurmdbV0043GetUser**](SlurmdbAPI.md#SlurmdbV0043GetUser) | **Get** /slurmdb/v0.0.43/user/{name} | Get user info
[**SlurmdbV0043GetUsers**](SlurmdbAPI.md#SlurmdbV0043GetUsers) | **Get** /slurmdb/v0.0.43/users/ | Get user list
[**SlurmdbV0043GetWckey**](SlurmdbAPI.md#SlurmdbV0043GetWckey) | **Get** /slurmdb/v0.0.43/wckey/{id} | Get wckey info
[**SlurmdbV0043GetWckeys**](SlurmdbAPI.md#SlurmdbV0043GetWckeys) | **Get** /slurmdb/v0.0.43/wckeys/ | Get wckey list
[**SlurmdbV0043PostAccounts**](SlurmdbAPI.md#SlurmdbV0043PostAccounts) | **Post** /slurmdb/v0.0.43/accounts/ | Add/update list of accounts
[**SlurmdbV0043PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0043PostAccountsAssociation) | **Post** /slurmdb/v0.0.43/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0043PostAssociations**](SlurmdbAPI.md#SlurmdbV0043PostAssociations) | **Post** /slurmdb/v0.0.43/associations/ | Set associations info
[**SlurmdbV0043PostClusters**](SlurmdbAPI.md#SlurmdbV0043PostClusters) | **Post** /slurmdb/v0.0.43/clusters/ | Get cluster list
[**SlurmdbV0043PostConfig**](SlurmdbAPI.md#SlurmdbV0043PostConfig) | **Post** /slurmdb/v0.0.43/config | Load all configuration information
[**SlurmdbV0043PostQos**](SlurmdbAPI.md#SlurmdbV0043PostQos) | **Post** /slurmdb/v0.0.43/qos/ | Add or update QOSs
[**SlurmdbV0043PostTres**](SlurmdbAPI.md#SlurmdbV0043PostTres) | **Post** /slurmdb/v0.0.43/tres/ | Add TRES
[**SlurmdbV0043PostUsers**](SlurmdbAPI.md#SlurmdbV0043PostUsers) | **Post** /slurmdb/v0.0.43/users/ | Update users
[**SlurmdbV0043PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0043PostUsersAssociation) | **Post** /slurmdb/v0.0.43/users_association/ | Add users with conditional association
[**SlurmdbV0043PostWckeys**](SlurmdbAPI.md#SlurmdbV0043PostWckeys) | **Post** /slurmdb/v0.0.43/wckeys/ | Add or update wckeys
[**SlurmdbV0044DeleteAccount**](SlurmdbAPI.md#SlurmdbV0044DeleteAccount) | **Delete** /slurmdb/v0.0.44/account/{account_name} | Delete account
[**SlurmdbV0044DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0044DeleteAssociation) | **Delete** /slurmdb/v0.0.44/association/ | Delete association
[**SlurmdbV0044DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0044DeleteAssociations) | **Delete** /slurmdb/v0.0.44/associations/ | Delete associations
[**SlurmdbV0044DeleteCluster**](SlurmdbAPI.md#SlurmdbV0044DeleteCluster) | **Delete** /slurmdb/v0.0.44/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0044DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0044DeleteSingleQos) | **Delete** /slurmdb/v0.0.44/qos/{qos} | Delete QOS
[**SlurmdbV0044DeleteUser**](SlurmdbAPI.md#SlurmdbV0044DeleteUser) | **Delete** /slurmdb/v0.0.44/user/{name} | Delete user
[**SlurmdbV0044DeleteWckey**](SlurmdbAPI.md#SlurmdbV0044DeleteWckey) | **Delete** /slurmdb/v0.0.44/wckey/{id} | Delete wckey
[**SlurmdbV0044GetAccount**](SlurmdbAPI.md#SlurmdbV0044GetAccount) | **Get** /slurmdb/v0.0.44/account/{account_name} | Get account info
[**SlurmdbV0044GetAccounts**](SlurmdbAPI.md#SlurmdbV0044GetAccounts) | **Get** /slurmdb/v0.0.44/accounts/ | Get account list
[**SlurmdbV0044GetAssociation**](SlurmdbAPI.md#SlurmdbV0044GetAssociation) | **Get** /slurmdb/v0.0.44/association/ | Get association info
[**SlurmdbV0044GetAssociations**](SlurmdbAPI.md#SlurmdbV0044GetAssociations) | **Get** /slurmdb/v0.0.44/associations/ | Get association list
[**SlurmdbV0044GetCluster**](SlurmdbAPI.md#SlurmdbV0044GetCluster) | **Get** /slurmdb/v0.0.44/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0044GetClusters**](SlurmdbAPI.md#SlurmdbV0044GetClusters) | **Get** /slurmdb/v0.0.44/clusters/ | Get cluster list
[**SlurmdbV0044GetConfig**](SlurmdbAPI.md#SlurmdbV0044GetConfig) | **Get** /slurmdb/v0.0.44/config | Dump all configuration information
[**SlurmdbV0044GetDiag**](SlurmdbAPI.md#SlurmdbV0044GetDiag) | **Get** /slurmdb/v0.0.44/diag/ | Get slurmdb diagnostics
[**SlurmdbV0044GetInstance**](SlurmdbAPI.md#SlurmdbV0044GetInstance) | **Get** /slurmdb/v0.0.44/instance/ | Get instance info
[**SlurmdbV0044GetInstances**](SlurmdbAPI.md#SlurmdbV0044GetInstances) | **Get** /slurmdb/v0.0.44/instances/ | Get instance list
[**SlurmdbV0044GetJob**](SlurmdbAPI.md#SlurmdbV0044GetJob) | **Get** /slurmdb/v0.0.44/job/{job_id} | Get job info
[**SlurmdbV0044GetJobs**](SlurmdbAPI.md#SlurmdbV0044GetJobs) | **Get** /slurmdb/v0.0.44/jobs/ | Get job list
[**SlurmdbV0044GetPing**](SlurmdbAPI.md#SlurmdbV0044GetPing) | **Get** /slurmdb/v0.0.44/ping/ | ping test
[**SlurmdbV0044GetQos**](SlurmdbAPI.md#SlurmdbV0044GetQos) | **Get** /slurmdb/v0.0.44/qos/ | Get QOS list
[**SlurmdbV0044GetSingleQos**](SlurmdbAPI.md#SlurmdbV0044GetSingleQos) | **Get** /slurmdb/v0.0.44/qos/{qos} | Get QOS info
[**SlurmdbV0044GetTres**](SlurmdbAPI.md#SlurmdbV0044GetTres) | **Get** /slurmdb/v0.0.44/tres/ | Get TRES info
[**SlurmdbV0044GetUser**](SlurmdbAPI.md#SlurmdbV0044GetUser) | **Get** /slurmdb/v0.0.44/user/{name} | Get user info
[**SlurmdbV0044GetUsers**](SlurmdbAPI.md#SlurmdbV0044GetUsers) | **Get** /slurmdb/v0.0.44/users/ | Get user list
[**SlurmdbV0044GetWckey**](SlurmdbAPI.md#SlurmdbV0044GetWckey) | **Get** /slurmdb/v0.0.44/wckey/{id} | Get wckey info
[**SlurmdbV0044GetWckeys**](SlurmdbAPI.md#SlurmdbV0044GetWckeys) | **Get** /slurmdb/v0.0.44/wckeys/ | Get wckey list
[**SlurmdbV0044PostAccounts**](SlurmdbAPI.md#SlurmdbV0044PostAccounts) | **Post** /slurmdb/v0.0.44/accounts/ | Add/update list of accounts
[**SlurmdbV0044PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0044PostAccountsAssociation) | **Post** /slurmdb/v0.0.44/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0044PostAssociations**](SlurmdbAPI.md#SlurmdbV0044PostAssociations) | **Post** /slurmdb/v0.0.44/associations/ | Set associations info
[**SlurmdbV0044PostClusters**](SlurmdbAPI.md#SlurmdbV0044PostClusters) | **Post** /slurmdb/v0.0.44/clusters/ | Get cluster list
[**SlurmdbV0044PostConfig**](SlurmdbAPI.md#SlurmdbV0044PostConfig) | **Post** /slurmdb/v0.0.44/config | Load all configuration information
[**SlurmdbV0044PostJob**](SlurmdbAPI.md#SlurmdbV0044PostJob) | **Post** /slurmdb/v0.0.44/job/{job_id} | Update job
[**SlurmdbV0044PostJobs**](SlurmdbAPI.md#SlurmdbV0044PostJobs) | **Post** /slurmdb/v0.0.44/jobs/ | Update jobs
[**SlurmdbV0044PostQos**](SlurmdbAPI.md#SlurmdbV0044PostQos) | **Post** /slurmdb/v0.0.44/qos/ | Add or update QOSs
[**SlurmdbV0044PostTres**](SlurmdbAPI.md#SlurmdbV0044PostTres) | **Post** /slurmdb/v0.0.44/tres/ | Add TRES
[**SlurmdbV0044PostUsers**](SlurmdbAPI.md#SlurmdbV0044PostUsers) | **Post** /slurmdb/v0.0.44/users/ | Update users
[**SlurmdbV0044PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0044PostUsersAssociation) | **Post** /slurmdb/v0.0.44/users_association/ | Add users with conditional association
[**SlurmdbV0044PostWckeys**](SlurmdbAPI.md#SlurmdbV0044PostWckeys) | **Post** /slurmdb/v0.0.44/wckeys/ | Add or update wckeys



## SlurmdbV0041DeleteAccount

> SlurmdbV0041DeleteAccount200Response SlurmdbV0041DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAccount`: SlurmdbV0041DeleteAccount200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlurmdbV0041DeleteAccount200Response**](SlurmdbV0041DeleteAccount200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteAssociation

> V0041OpenapiAssocsRemovedResp SlurmdbV0041DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAssociation`: V0041OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsRemovedResp**](V0041OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteAssociations

> V0041OpenapiAssocsRemovedResp SlurmdbV0041DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteAssociations`: V0041OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsRemovedResp**](V0041OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteCluster

> SlurmdbV0041DeleteCluster200Response SlurmdbV0041DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteCluster`: SlurmdbV0041DeleteCluster200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**SlurmdbV0041DeleteCluster200Response**](SlurmdbV0041DeleteCluster200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteSingleQos

> SlurmdbV0041DeleteSingleQos200Response SlurmdbV0041DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteSingleQos`: SlurmdbV0041DeleteSingleQos200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlurmdbV0041DeleteSingleQos200Response**](SlurmdbV0041DeleteSingleQos200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteUser

> V0041OpenapiResp SlurmdbV0041DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteUser`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041DeleteWckey

> SlurmdbV0041DeleteWckey200Response SlurmdbV0041DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041DeleteWckey`: SlurmdbV0041DeleteWckey200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SlurmdbV0041DeleteWckey200Response**](SlurmdbV0041DeleteWckey200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAccount

> V0041OpenapiAccountsResp SlurmdbV0041GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAccount`: V0041OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAccounts

> V0041OpenapiAccountsResp SlurmdbV0041GetAccounts(ctx).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	dELETED := "dELETED_example" // string | include deleted associations (optional)
	withAssociations := "withAssociations_example" // string | query includes associations (optional)
	withCoordinators := "withCoordinators_example" // string | query includes coordinators (optional)
	noUsersAreCoords := "noUsersAreCoords_example" // string | remove users as coordinators (optional)
	usersAreCoords := "usersAreCoords_example" // string | users are coordinators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAccounts(context.Background()).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAccounts`: V0041OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **dELETED** | **string** | include deleted associations | 
 **withAssociations** | **string** | query includes associations | 
 **withCoordinators** | **string** | query includes coordinators | 
 **noUsersAreCoords** | **string** | remove users as coordinators | 
 **usersAreCoords** | **string** | users are coordinators | 

### Return type

[**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAssociation

> V0041OpenapiAssocsResp SlurmdbV0041GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAssociation`: V0041OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetAssociations

> V0041OpenapiAssocsResp SlurmdbV0041GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted associations (optional)
	withRawQos := "withRawQos_example" // string | Include a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | Include sub acct information (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | Exclude parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | Exclude limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetAssociations`: V0041OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | Filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted associations | 
 **withRawQos** | **string** | Include a raw qos or delta_qos | 
 **withSubAccts** | **string** | Include sub acct information | 
 **withoutParentInfo** | **string** | Exclude parent id/name | 
 **withoutParentLimits** | **string** | Exclude limits from parents | 

### Return type

[**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetCluster

> V0041OpenapiClustersResp SlurmdbV0041GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetCluster`: V0041OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetClusters

> V0041OpenapiClustersResp SlurmdbV0041GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetClusters`: V0041OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 

### Return type

[**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetConfig

> V0041OpenapiSlurmdbdConfigResp SlurmdbV0041GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetConfig`: V0041OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetConfigRequest struct via the builder pattern


### Return type

[**V0041OpenapiSlurmdbdConfigResp**](V0041OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetDiag

> SlurmdbV0041GetDiag200Response SlurmdbV0041GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetDiag`: SlurmdbV0041GetDiag200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetDiagRequest struct via the builder pattern


### Return type

[**SlurmdbV0041GetDiag200Response**](SlurmdbV0041GetDiag200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetInstance

> V0041OpenapiInstancesResp SlurmdbV0041GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetInstance`: V0041OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0041OpenapiInstancesResp**](V0041OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetInstances

> V0041OpenapiInstancesResp SlurmdbV0041GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetInstances`: V0041OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0041OpenapiInstancesResp**](V0041OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetJob

> V0041OpenapiSlurmdbdJobsResp SlurmdbV0041GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetJob`: V0041OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiSlurmdbdJobsResp**](V0041OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetJobs

> V0041OpenapiSlurmdbdJobsResp SlurmdbV0041GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV wckey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetJobs`: V0041OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV wckey list | 

### Return type

[**V0041OpenapiSlurmdbdJobsResp**](V0041OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetQos

> V0041OpenapiSlurmdbdQosResp SlurmdbV0041GetQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetQos`: V0041OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **withDeleted** | **string** | Include deleted QOS | 

### Return type

[**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetSingleQos

> V0041OpenapiSlurmdbdQosResp SlurmdbV0041GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetSingleQos`: V0041OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetTres

> V0041OpenapiTresResp SlurmdbV0041GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetTres`: V0041OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetTresRequest struct via the builder pattern


### Return type

[**V0041OpenapiTresResp**](V0041OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetUser

> V0041OpenapiUsersResp SlurmdbV0041GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetUser`: V0041OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include wckeys | 

### Return type

[**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetUsers

> V0041OpenapiUsersResp SlurmdbV0041GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default wckey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With wckeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetUsers`: V0041OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default wckey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With wckeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetWckey

> V0041OpenapiWckeyResp SlurmdbV0041GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetWckey`: V0041OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041GetWckeys

> V0041OpenapiWckeyResp SlurmdbV0041GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041GetWckeys`: V0041OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted wckeys | 

### Return type

[**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAccounts

> V0041OpenapiResp SlurmdbV0041PostAccounts(ctx).V0041OpenapiAccountsResp(v0041OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041OpenapiAccountsResp := *openapiclient.NewV0041OpenapiAccountsResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespAccountsInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespAccountsInner("Description_example", "Name_example", "Organization_example")}) // V0041OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAccounts(context.Background()).V0041OpenapiAccountsResp(v0041OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAccounts`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiAccountsResp** | [**V0041OpenapiAccountsResp**](V0041OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAccountsAssociation

> SlurmdbV0041PostAccountsAssociation200Response SlurmdbV0041PostAccountsAssociation(ctx).SlurmdbV0041PostAccountsAssociationRequest(slurmdbV0041PostAccountsAssociationRequest).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	slurmdbV0041PostAccountsAssociationRequest := *openapiclient.NewSlurmdbV0041PostAccountsAssociationRequest() // SlurmdbV0041PostAccountsAssociationRequest | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAccountsAssociation(context.Background()).SlurmdbV0041PostAccountsAssociationRequest(slurmdbV0041PostAccountsAssociationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAccountsAssociation`: SlurmdbV0041PostAccountsAssociation200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slurmdbV0041PostAccountsAssociationRequest** | [**SlurmdbV0041PostAccountsAssociationRequest**](SlurmdbV0041PostAccountsAssociationRequest.md) | Add list of accounts with conditional association | 

### Return type

[**SlurmdbV0041PostAccountsAssociation200Response**](SlurmdbV0041PostAccountsAssociation200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostAssociations

> V0041OpenapiResp SlurmdbV0041PostAssociations(ctx).V0041OpenapiAssocsResp(v0041OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041OpenapiAssocsResp := *openapiclient.NewV0041OpenapiAssocsResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespAssociationsInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespAssociationsInner("User_example")}) // V0041OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostAssociations(context.Background()).V0041OpenapiAssocsResp(v0041OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostAssociations`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiAssocsResp** | [**V0041OpenapiAssocsResp**](V0041OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostClusters

> V0041OpenapiResp SlurmdbV0041PostClusters(ctx).UpdateTime(updateTime).V0041OpenapiClustersResp(v0041OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter reservations since update timestamp (optional)
	v0041OpenapiClustersResp := *openapiclient.NewV0041OpenapiClustersResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespClustersInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespClustersInner()}) // V0041OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostClusters(context.Background()).UpdateTime(updateTime).V0041OpenapiClustersResp(v0041OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostClusters`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter reservations since update timestamp | 
 **v0041OpenapiClustersResp** | [**V0041OpenapiClustersResp**](V0041OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostConfig

> V0041OpenapiResp SlurmdbV0041PostConfig(ctx).V0041OpenapiSlurmdbdConfigResp(v0041OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041OpenapiSlurmdbdConfigResp := *openapiclient.NewV0041OpenapiSlurmdbdConfigResp() // V0041OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostConfig(context.Background()).V0041OpenapiSlurmdbdConfigResp(v0041OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostConfig`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiSlurmdbdConfigResp** | [**V0041OpenapiSlurmdbdConfigResp**](V0041OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostQos

> V0041OpenapiResp SlurmdbV0041PostQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).V0041OpenapiSlurmdbdQosResp(v0041OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)
	v0041OpenapiSlurmdbdQosResp := *openapiclient.NewV0041OpenapiSlurmdbdQosResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespQosInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespQosInner()}) // V0041OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).V0041OpenapiSlurmdbdQosResp(v0041OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostQos`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **withDeleted** | **string** | Include deleted QOS | 
 **v0041OpenapiSlurmdbdQosResp** | [**V0041OpenapiSlurmdbdQosResp**](V0041OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostTres

> V0041OpenapiResp SlurmdbV0041PostTres(ctx).V0041OpenapiTresResp(v0041OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041OpenapiTresResp := *openapiclient.NewV0041OpenapiTresResp([]openapiclient.SlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner{*openapiclient.NewSlurmdbV0041PostUsersAssociationRequestAssociationConditionAssociationGrptresInner("Type_example")}) // V0041OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostTres(context.Background()).V0041OpenapiTresResp(v0041OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostTres`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiTresResp** | [**V0041OpenapiTresResp**](V0041OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostUsers

> V0041OpenapiResp SlurmdbV0041PostUsers(ctx).V0041OpenapiUsersResp(v0041OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0041OpenapiUsersResp := *openapiclient.NewV0041OpenapiUsersResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespUsersInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespUsersInner("Name_example")}) // V0041OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostUsers(context.Background()).V0041OpenapiUsersResp(v0041OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostUsers`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0041OpenapiUsersResp** | [**V0041OpenapiUsersResp**](V0041OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostUsersAssociation

> SlurmdbV0041PostUsersAssociation200Response SlurmdbV0041PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).SlurmdbV0041PostUsersAssociationRequest(slurmdbV0041PostUsersAssociationRequest).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Filter partitions since update timestamp (optional)
	flags := "flags_example" // string | Query flags (optional)
	slurmdbV0041PostUsersAssociationRequest := *openapiclient.NewSlurmdbV0041PostUsersAssociationRequest(*openapiclient.NewSlurmdbV0041PostUsersAssociationRequestAssociationCondition([]string{"Users_example"}), *openapiclient.NewSlurmdbV0041PostUsersAssociationRequestUser()) // SlurmdbV0041PostUsersAssociationRequest | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).SlurmdbV0041PostUsersAssociationRequest(slurmdbV0041PostUsersAssociationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostUsersAssociation`: SlurmdbV0041PostUsersAssociation200Response
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Filter partitions since update timestamp | 
 **flags** | **string** | Query flags | 
 **slurmdbV0041PostUsersAssociationRequest** | [**SlurmdbV0041PostUsersAssociationRequest**](SlurmdbV0041PostUsersAssociationRequest.md) | Create users with conditional association | 

### Return type

[**SlurmdbV0041PostUsersAssociation200Response**](SlurmdbV0041PostUsersAssociation200Response.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0041PostWckeys

> V0041OpenapiResp SlurmdbV0041PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0041OpenapiWckeyResp(v0041OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted wckeys (optional)
	v0041OpenapiWckeyResp := *openapiclient.NewV0041OpenapiWckeyResp([]openapiclient.V0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner{*openapiclient.NewV0041OpenapiSlurmdbdConfigRespUsersInnerWckeysInner("Cluster_example", "Name_example", "User_example")}) // V0041OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0041PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0041OpenapiWckeyResp(v0041OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0041PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0041PostWckeys`: V0041OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0041PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0041PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted wckeys | 
 **v0041OpenapiWckeyResp** | [**V0041OpenapiWckeyResp**](V0041OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0041OpenapiResp**](V0041OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteAccount

> V0042OpenapiAccountsRemovedResp SlurmdbV0042DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteAccount`: V0042OpenapiAccountsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiAccountsRemovedResp**](V0042OpenapiAccountsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteAssociation

> V0042OpenapiAssocsRemovedResp SlurmdbV0042DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteAssociation`: V0042OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0042OpenapiAssocsRemovedResp**](V0042OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteAssociations

> V0042OpenapiAssocsRemovedResp SlurmdbV0042DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteAssociations`: V0042OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0042OpenapiAssocsRemovedResp**](V0042OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteCluster

> V0042OpenapiClustersRemovedResp SlurmdbV0042DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteCluster`: V0042OpenapiClustersRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0042OpenapiClustersRemovedResp**](V0042OpenapiClustersRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteSingleQos

> V0042OpenapiSlurmdbdQosRemovedResp SlurmdbV0042DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteSingleQos`: V0042OpenapiSlurmdbdQosRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiSlurmdbdQosRemovedResp**](V0042OpenapiSlurmdbdQosRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteUser

> V0042OpenapiResp SlurmdbV0042DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteUser`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042DeleteWckey

> V0042OpenapiWckeyRemovedResp SlurmdbV0042DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042DeleteWckey`: V0042OpenapiWckeyRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiWckeyRemovedResp**](V0042OpenapiWckeyRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetAccount

> V0042OpenapiAccountsResp SlurmdbV0042GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetAccount`: V0042OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0042OpenapiAccountsResp**](V0042OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetAccounts

> V0042OpenapiAccountsResp SlurmdbV0042GetAccounts(ctx).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	dELETED := "dELETED_example" // string | include deleted associations (optional)
	withAssociations := "withAssociations_example" // string | query includes associations (optional)
	withCoordinators := "withCoordinators_example" // string | query includes coordinators (optional)
	noUsersAreCoords := "noUsersAreCoords_example" // string | remove users as coordinators (optional)
	usersAreCoords := "usersAreCoords_example" // string | users are coordinators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetAccounts(context.Background()).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetAccounts`: V0042OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **dELETED** | **string** | include deleted associations | 
 **withAssociations** | **string** | query includes associations | 
 **withCoordinators** | **string** | query includes coordinators | 
 **noUsersAreCoords** | **string** | remove users as coordinators | 
 **usersAreCoords** | **string** | users are coordinators | 

### Return type

[**V0042OpenapiAccountsResp**](V0042OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetAssociation

> V0042OpenapiAssocsResp SlurmdbV0042GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetAssociation`: V0042OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0042OpenapiAssocsResp**](V0042OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetAssociations

> V0042OpenapiAssocsResp SlurmdbV0042GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetAssociations`: V0042OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0042OpenapiAssocsResp**](V0042OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetCluster

> V0042OpenapiClustersResp SlurmdbV0042GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetCluster`: V0042OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0042OpenapiClustersResp**](V0042OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetClusters

> V0042OpenapiClustersResp SlurmdbV0042GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetClusters`: V0042OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0042OpenapiClustersResp**](V0042OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetConfig

> V0042OpenapiSlurmdbdConfigResp SlurmdbV0042GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetConfig`: V0042OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetConfigRequest struct via the builder pattern


### Return type

[**V0042OpenapiSlurmdbdConfigResp**](V0042OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetDiag

> V0042OpenapiSlurmdbdStatsResp SlurmdbV0042GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetDiag`: V0042OpenapiSlurmdbdStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetDiagRequest struct via the builder pattern


### Return type

[**V0042OpenapiSlurmdbdStatsResp**](V0042OpenapiSlurmdbdStatsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetInstance

> V0042OpenapiInstancesResp SlurmdbV0042GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetInstance`: V0042OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0042OpenapiInstancesResp**](V0042OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetInstances

> V0042OpenapiInstancesResp SlurmdbV0042GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetInstances`: V0042OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0042OpenapiInstancesResp**](V0042OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetJob

> V0042OpenapiSlurmdbdJobsResp SlurmdbV0042GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetJob`: V0042OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiSlurmdbdJobsResp**](V0042OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetJobs

> V0042OpenapiSlurmdbdJobsResp SlurmdbV0042GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV WCKey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetJobs`: V0042OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV WCKey list | 

### Return type

[**V0042OpenapiSlurmdbdJobsResp**](V0042OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetPing

> V0042OpenapiSlurmdbdPingResp SlurmdbV0042GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetPing`: V0042OpenapiSlurmdbdPingResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetPingRequest struct via the builder pattern


### Return type

[**V0042OpenapiSlurmdbdPingResp**](V0042OpenapiSlurmdbdPingResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetQos

> V0042OpenapiSlurmdbdQosResp SlurmdbV0042GetQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetQos`: V0042OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 

### Return type

[**V0042OpenapiSlurmdbdQosResp**](V0042OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetSingleQos

> V0042OpenapiSlurmdbdQosResp SlurmdbV0042GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetSingleQos`: V0042OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0042OpenapiSlurmdbdQosResp**](V0042OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetTres

> V0042OpenapiTresResp SlurmdbV0042GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetTres`: V0042OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetTresRequest struct via the builder pattern


### Return type

[**V0042OpenapiTresResp**](V0042OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetUser

> V0042OpenapiUsersResp SlurmdbV0042GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetUser`: V0042OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include WCKeys | 

### Return type

[**V0042OpenapiUsersResp**](V0042OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetUsers

> V0042OpenapiUsersResp SlurmdbV0042GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default WCKey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With WCKeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetUsers`: V0042OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default WCKey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With WCKeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0042OpenapiUsersResp**](V0042OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetWckey

> V0042OpenapiWckeyResp SlurmdbV0042GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetWckey`: V0042OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0042OpenapiWckeyResp**](V0042OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042GetWckeys

> V0042OpenapiWckeyResp SlurmdbV0042GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042GetWckeys`: V0042OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 

### Return type

[**V0042OpenapiWckeyResp**](V0042OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostAccounts

> V0042OpenapiResp SlurmdbV0042PostAccounts(ctx).V0042OpenapiAccountsResp(v0042OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042OpenapiAccountsResp := *openapiclient.NewV0042OpenapiAccountsResp([]openapiclient.V0042Account{*openapiclient.NewV0042Account("Description_example", "Name_example", "Organization_example")}) // V0042OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostAccounts(context.Background()).V0042OpenapiAccountsResp(v0042OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostAccounts`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiAccountsResp** | [**V0042OpenapiAccountsResp**](V0042OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostAccountsAssociation

> V0042OpenapiAccountsAddCondRespStr SlurmdbV0042PostAccountsAssociation(ctx).V0042OpenapiAccountsAddCondResp(v0042OpenapiAccountsAddCondResp).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042OpenapiAccountsAddCondResp := *openapiclient.NewV0042OpenapiAccountsAddCondResp() // V0042OpenapiAccountsAddCondResp | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostAccountsAssociation(context.Background()).V0042OpenapiAccountsAddCondResp(v0042OpenapiAccountsAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostAccountsAssociation`: V0042OpenapiAccountsAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiAccountsAddCondResp** | [**V0042OpenapiAccountsAddCondResp**](V0042OpenapiAccountsAddCondResp.md) | Add list of accounts with conditional association | 

### Return type

[**V0042OpenapiAccountsAddCondRespStr**](V0042OpenapiAccountsAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostAssociations

> V0042OpenapiResp SlurmdbV0042PostAssociations(ctx).V0042OpenapiAssocsResp(v0042OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042OpenapiAssocsResp := *openapiclient.NewV0042OpenapiAssocsResp([]openapiclient.V0042Assoc{*openapiclient.NewV0042Assoc("User_example")}) // V0042OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostAssociations(context.Background()).V0042OpenapiAssocsResp(v0042OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostAssociations`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiAssocsResp** | [**V0042OpenapiAssocsResp**](V0042OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostClusters

> V0042OpenapiResp SlurmdbV0042PostClusters(ctx).UpdateTime(updateTime).V0042OpenapiClustersResp(v0042OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)
	v0042OpenapiClustersResp := *openapiclient.NewV0042OpenapiClustersResp([]openapiclient.V0042ClusterRec{*openapiclient.NewV0042ClusterRec()}) // V0042OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostClusters(context.Background()).UpdateTime(updateTime).V0042OpenapiClustersResp(v0042OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostClusters`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 
 **v0042OpenapiClustersResp** | [**V0042OpenapiClustersResp**](V0042OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostConfig

> V0042OpenapiResp SlurmdbV0042PostConfig(ctx).V0042OpenapiSlurmdbdConfigResp(v0042OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042OpenapiSlurmdbdConfigResp := *openapiclient.NewV0042OpenapiSlurmdbdConfigResp() // V0042OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostConfig(context.Background()).V0042OpenapiSlurmdbdConfigResp(v0042OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostConfig`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiSlurmdbdConfigResp** | [**V0042OpenapiSlurmdbdConfigResp**](V0042OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostQos

> V0042OpenapiResp SlurmdbV0042PostQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0042OpenapiSlurmdbdQosResp(v0042OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	v0042OpenapiSlurmdbdQosResp := *openapiclient.NewV0042OpenapiSlurmdbdQosResp([]openapiclient.V0042Qos{*openapiclient.NewV0042Qos()}) // V0042OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0042OpenapiSlurmdbdQosResp(v0042OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostQos`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **v0042OpenapiSlurmdbdQosResp** | [**V0042OpenapiSlurmdbdQosResp**](V0042OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostTres

> V0042OpenapiResp SlurmdbV0042PostTres(ctx).V0042OpenapiTresResp(v0042OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042OpenapiTresResp := *openapiclient.NewV0042OpenapiTresResp([]openapiclient.V0042Tres{*openapiclient.NewV0042Tres("Type_example")}) // V0042OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostTres(context.Background()).V0042OpenapiTresResp(v0042OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostTres`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiTresResp** | [**V0042OpenapiTresResp**](V0042OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostUsers

> V0042OpenapiResp SlurmdbV0042PostUsers(ctx).V0042OpenapiUsersResp(v0042OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0042OpenapiUsersResp := *openapiclient.NewV0042OpenapiUsersResp([]openapiclient.V0042User{*openapiclient.NewV0042User("Name_example")}) // V0042OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostUsers(context.Background()).V0042OpenapiUsersResp(v0042OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostUsers`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0042OpenapiUsersResp** | [**V0042OpenapiUsersResp**](V0042OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostUsersAssociation

> V0042OpenapiUsersAddCondRespStr SlurmdbV0042PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).V0042OpenapiUsersAddCondResp(v0042OpenapiUsersAddCondResp).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query partitions updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)
	v0042OpenapiUsersAddCondResp := *openapiclient.NewV0042OpenapiUsersAddCondResp(*openapiclient.NewV0042UsersAddCond([]string{"Users_example"}), *openapiclient.NewV0042UserShort()) // V0042OpenapiUsersAddCondResp | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).V0042OpenapiUsersAddCondResp(v0042OpenapiUsersAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostUsersAssociation`: V0042OpenapiUsersAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 
 **v0042OpenapiUsersAddCondResp** | [**V0042OpenapiUsersAddCondResp**](V0042OpenapiUsersAddCondResp.md) | Create users with conditional association | 

### Return type

[**V0042OpenapiUsersAddCondRespStr**](V0042OpenapiUsersAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0042PostWckeys

> V0042OpenapiResp SlurmdbV0042PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0042OpenapiWckeyResp(v0042OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)
	v0042OpenapiWckeyResp := *openapiclient.NewV0042OpenapiWckeyResp([]openapiclient.V0042Wckey{*openapiclient.NewV0042Wckey("Cluster_example", "Name_example", "User_example")}) // V0042OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0042PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0042OpenapiWckeyResp(v0042OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0042PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0042PostWckeys`: V0042OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0042PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0042PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 
 **v0042OpenapiWckeyResp** | [**V0042OpenapiWckeyResp**](V0042OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0042OpenapiResp**](V0042OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043DeleteAccount

> V0043OpenapiAccountsRemovedResp SlurmdbV0043DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043DeleteAccount`: V0043OpenapiAccountsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0043OpenapiAccountsRemovedResp**](V0043OpenapiAccountsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043DeleteAssociation

> V0043OpenapiAssocsRemovedResp SlurmdbV0043DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043DeleteAssociation`: V0043OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0043OpenapiAssocsRemovedResp**](V0043OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043DeleteAssociations

> V0043OpenapiAssocsRemovedResp SlurmdbV0043DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043DeleteAssociations`: V0043OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0043OpenapiAssocsRemovedResp**](V0043OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043DeleteCluster

> V0043OpenapiClustersRemovedResp SlurmdbV0043DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043DeleteCluster`: V0043OpenapiClustersRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0043OpenapiClustersRemovedResp**](V0043OpenapiClustersRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043DeleteSingleQos

> V0043OpenapiSlurmdbdQosRemovedResp SlurmdbV0043DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043DeleteSingleQos`: V0043OpenapiSlurmdbdQosRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0043OpenapiSlurmdbdQosRemovedResp**](V0043OpenapiSlurmdbdQosRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043DeleteUser

> V0043OpenapiResp SlurmdbV0043DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043DeleteUser`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043DeleteWckey

> V0043OpenapiWckeyRemovedResp SlurmdbV0043DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043DeleteWckey`: V0043OpenapiWckeyRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0043OpenapiWckeyRemovedResp**](V0043OpenapiWckeyRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetAccount

> V0043OpenapiAccountsResp SlurmdbV0043GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetAccount`: V0043OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0043OpenapiAccountsResp**](V0043OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetAccounts

> V0043OpenapiAccountsResp SlurmdbV0043GetAccounts(ctx).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	dELETED := "dELETED_example" // string | include deleted associations (optional)
	withAssociations := "withAssociations_example" // string | query includes associations (optional)
	withCoordinators := "withCoordinators_example" // string | query includes coordinators (optional)
	noUsersAreCoords := "noUsersAreCoords_example" // string | remove users as coordinators (optional)
	usersAreCoords := "usersAreCoords_example" // string | users are coordinators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetAccounts(context.Background()).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetAccounts`: V0043OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **dELETED** | **string** | include deleted associations | 
 **withAssociations** | **string** | query includes associations | 
 **withCoordinators** | **string** | query includes coordinators | 
 **noUsersAreCoords** | **string** | remove users as coordinators | 
 **usersAreCoords** | **string** | users are coordinators | 

### Return type

[**V0043OpenapiAccountsResp**](V0043OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetAssociation

> V0043OpenapiAssocsResp SlurmdbV0043GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetAssociation`: V0043OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0043OpenapiAssocsResp**](V0043OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetAssociations

> V0043OpenapiAssocsResp SlurmdbV0043GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetAssociations`: V0043OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0043OpenapiAssocsResp**](V0043OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetCluster

> V0043OpenapiClustersResp SlurmdbV0043GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetCluster`: V0043OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0043OpenapiClustersResp**](V0043OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetClusters

> V0043OpenapiClustersResp SlurmdbV0043GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetClusters`: V0043OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0043OpenapiClustersResp**](V0043OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetConfig

> V0043OpenapiSlurmdbdConfigResp SlurmdbV0043GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetConfig`: V0043OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetConfigRequest struct via the builder pattern


### Return type

[**V0043OpenapiSlurmdbdConfigResp**](V0043OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetDiag

> V0043OpenapiSlurmdbdStatsResp SlurmdbV0043GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetDiag`: V0043OpenapiSlurmdbdStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetDiagRequest struct via the builder pattern


### Return type

[**V0043OpenapiSlurmdbdStatsResp**](V0043OpenapiSlurmdbdStatsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetInstance

> V0043OpenapiInstancesResp SlurmdbV0043GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetInstance`: V0043OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0043OpenapiInstancesResp**](V0043OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetInstances

> V0043OpenapiInstancesResp SlurmdbV0043GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetInstances`: V0043OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0043OpenapiInstancesResp**](V0043OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetJob

> V0043OpenapiSlurmdbdJobsResp SlurmdbV0043GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetJob`: V0043OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0043OpenapiSlurmdbdJobsResp**](V0043OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetJobs

> V0043OpenapiSlurmdbdJobsResp SlurmdbV0043GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV WCKey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetJobs`: V0043OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV WCKey list | 

### Return type

[**V0043OpenapiSlurmdbdJobsResp**](V0043OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetPing

> V0043OpenapiSlurmdbdPingResp SlurmdbV0043GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetPing`: V0043OpenapiSlurmdbdPingResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetPingRequest struct via the builder pattern


### Return type

[**V0043OpenapiSlurmdbdPingResp**](V0043OpenapiSlurmdbdPingResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetQos

> V0043OpenapiSlurmdbdQosResp SlurmdbV0043GetQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetQos`: V0043OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 

### Return type

[**V0043OpenapiSlurmdbdQosResp**](V0043OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetSingleQos

> V0043OpenapiSlurmdbdQosResp SlurmdbV0043GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetSingleQos`: V0043OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0043OpenapiSlurmdbdQosResp**](V0043OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetTres

> V0043OpenapiTresResp SlurmdbV0043GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetTres`: V0043OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetTresRequest struct via the builder pattern


### Return type

[**V0043OpenapiTresResp**](V0043OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetUser

> V0043OpenapiUsersResp SlurmdbV0043GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetUser`: V0043OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include WCKeys | 

### Return type

[**V0043OpenapiUsersResp**](V0043OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetUsers

> V0043OpenapiUsersResp SlurmdbV0043GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default WCKey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With WCKeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetUsers`: V0043OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default WCKey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With WCKeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0043OpenapiUsersResp**](V0043OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetWckey

> V0043OpenapiWckeyResp SlurmdbV0043GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetWckey`: V0043OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0043OpenapiWckeyResp**](V0043OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043GetWckeys

> V0043OpenapiWckeyResp SlurmdbV0043GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043GetWckeys`: V0043OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 

### Return type

[**V0043OpenapiWckeyResp**](V0043OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostAccounts

> V0043OpenapiResp SlurmdbV0043PostAccounts(ctx).V0043OpenapiAccountsResp(v0043OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0043OpenapiAccountsResp := *openapiclient.NewV0043OpenapiAccountsResp([]openapiclient.V0043Account{*openapiclient.NewV0043Account("Description_example", "Name_example", "Organization_example")}) // V0043OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostAccounts(context.Background()).V0043OpenapiAccountsResp(v0043OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostAccounts`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043OpenapiAccountsResp** | [**V0043OpenapiAccountsResp**](V0043OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostAccountsAssociation

> V0043OpenapiAccountsAddCondRespStr SlurmdbV0043PostAccountsAssociation(ctx).V0043OpenapiAccountsAddCondResp(v0043OpenapiAccountsAddCondResp).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0043OpenapiAccountsAddCondResp := *openapiclient.NewV0043OpenapiAccountsAddCondResp() // V0043OpenapiAccountsAddCondResp | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostAccountsAssociation(context.Background()).V0043OpenapiAccountsAddCondResp(v0043OpenapiAccountsAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostAccountsAssociation`: V0043OpenapiAccountsAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043OpenapiAccountsAddCondResp** | [**V0043OpenapiAccountsAddCondResp**](V0043OpenapiAccountsAddCondResp.md) | Add list of accounts with conditional association | 

### Return type

[**V0043OpenapiAccountsAddCondRespStr**](V0043OpenapiAccountsAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostAssociations

> V0043OpenapiResp SlurmdbV0043PostAssociations(ctx).V0043OpenapiAssocsResp(v0043OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0043OpenapiAssocsResp := *openapiclient.NewV0043OpenapiAssocsResp([]openapiclient.V0043Assoc{*openapiclient.NewV0043Assoc("User_example")}) // V0043OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostAssociations(context.Background()).V0043OpenapiAssocsResp(v0043OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostAssociations`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043OpenapiAssocsResp** | [**V0043OpenapiAssocsResp**](V0043OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostClusters

> V0043OpenapiResp SlurmdbV0043PostClusters(ctx).UpdateTime(updateTime).V0043OpenapiClustersResp(v0043OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)
	v0043OpenapiClustersResp := *openapiclient.NewV0043OpenapiClustersResp([]openapiclient.V0043ClusterRec{*openapiclient.NewV0043ClusterRec()}) // V0043OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostClusters(context.Background()).UpdateTime(updateTime).V0043OpenapiClustersResp(v0043OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostClusters`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 
 **v0043OpenapiClustersResp** | [**V0043OpenapiClustersResp**](V0043OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostConfig

> V0043OpenapiResp SlurmdbV0043PostConfig(ctx).V0043OpenapiSlurmdbdConfigResp(v0043OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0043OpenapiSlurmdbdConfigResp := *openapiclient.NewV0043OpenapiSlurmdbdConfigResp() // V0043OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostConfig(context.Background()).V0043OpenapiSlurmdbdConfigResp(v0043OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostConfig`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043OpenapiSlurmdbdConfigResp** | [**V0043OpenapiSlurmdbdConfigResp**](V0043OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostQos

> V0043OpenapiResp SlurmdbV0043PostQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0043OpenapiSlurmdbdQosResp(v0043OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	v0043OpenapiSlurmdbdQosResp := *openapiclient.NewV0043OpenapiSlurmdbdQosResp([]openapiclient.V0043Qos{*openapiclient.NewV0043Qos()}) // V0043OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0043OpenapiSlurmdbdQosResp(v0043OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostQos`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **v0043OpenapiSlurmdbdQosResp** | [**V0043OpenapiSlurmdbdQosResp**](V0043OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostTres

> V0043OpenapiResp SlurmdbV0043PostTres(ctx).V0043OpenapiTresResp(v0043OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0043OpenapiTresResp := *openapiclient.NewV0043OpenapiTresResp([]openapiclient.V0043Tres{*openapiclient.NewV0043Tres("Type_example")}) // V0043OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostTres(context.Background()).V0043OpenapiTresResp(v0043OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostTres`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043OpenapiTresResp** | [**V0043OpenapiTresResp**](V0043OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostUsers

> V0043OpenapiResp SlurmdbV0043PostUsers(ctx).V0043OpenapiUsersResp(v0043OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0043OpenapiUsersResp := *openapiclient.NewV0043OpenapiUsersResp([]openapiclient.V0043User{*openapiclient.NewV0043User("Name_example")}) // V0043OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostUsers(context.Background()).V0043OpenapiUsersResp(v0043OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostUsers`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0043OpenapiUsersResp** | [**V0043OpenapiUsersResp**](V0043OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostUsersAssociation

> V0043OpenapiUsersAddCondRespStr SlurmdbV0043PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).V0043OpenapiUsersAddCondResp(v0043OpenapiUsersAddCondResp).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query partitions updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)
	v0043OpenapiUsersAddCondResp := *openapiclient.NewV0043OpenapiUsersAddCondResp(*openapiclient.NewV0043UsersAddCond([]string{"Users_example"}), *openapiclient.NewV0043UserShort()) // V0043OpenapiUsersAddCondResp | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).V0043OpenapiUsersAddCondResp(v0043OpenapiUsersAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostUsersAssociation`: V0043OpenapiUsersAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 
 **v0043OpenapiUsersAddCondResp** | [**V0043OpenapiUsersAddCondResp**](V0043OpenapiUsersAddCondResp.md) | Create users with conditional association | 

### Return type

[**V0043OpenapiUsersAddCondRespStr**](V0043OpenapiUsersAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0043PostWckeys

> V0043OpenapiResp SlurmdbV0043PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0043OpenapiWckeyResp(v0043OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)
	v0043OpenapiWckeyResp := *openapiclient.NewV0043OpenapiWckeyResp([]openapiclient.V0043Wckey{*openapiclient.NewV0043Wckey("Cluster_example", "Name_example", "User_example")}) // V0043OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0043PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0043OpenapiWckeyResp(v0043OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0043PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0043PostWckeys`: V0043OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0043PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0043PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 
 **v0043OpenapiWckeyResp** | [**V0043OpenapiWckeyResp**](V0043OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0043OpenapiResp**](V0043OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteAccount

> V0044OpenapiAccountsRemovedResp SlurmdbV0044DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteAccount`: V0044OpenapiAccountsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiAccountsRemovedResp**](V0044OpenapiAccountsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteAssociation

> V0044OpenapiAssocsRemovedResp SlurmdbV0044DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteAssociation`: V0044OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0044OpenapiAssocsRemovedResp**](V0044OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteAssociations

> V0044OpenapiAssocsRemovedResp SlurmdbV0044DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteAssociations`: V0044OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0044OpenapiAssocsRemovedResp**](V0044OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteCluster

> V0044OpenapiClustersRemovedResp SlurmdbV0044DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteCluster`: V0044OpenapiClustersRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0044OpenapiClustersRemovedResp**](V0044OpenapiClustersRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteSingleQos

> V0044OpenapiSlurmdbdQosRemovedResp SlurmdbV0044DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteSingleQos`: V0044OpenapiSlurmdbdQosRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiSlurmdbdQosRemovedResp**](V0044OpenapiSlurmdbdQosRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteUser

> V0044OpenapiResp SlurmdbV0044DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteUser`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteWckey

> V0044OpenapiWckeyRemovedResp SlurmdbV0044DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteWckey`: V0044OpenapiWckeyRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiWckeyRemovedResp**](V0044OpenapiWckeyRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetAccount

> V0044OpenapiAccountsResp SlurmdbV0044GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetAccount`: V0044OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0044OpenapiAccountsResp**](V0044OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetAccounts

> V0044OpenapiAccountsResp SlurmdbV0044GetAccounts(ctx).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	dELETED := "dELETED_example" // string | include deleted associations (optional)
	withAssociations := "withAssociations_example" // string | query includes associations (optional)
	withCoordinators := "withCoordinators_example" // string | query includes coordinators (optional)
	noUsersAreCoords := "noUsersAreCoords_example" // string | remove users as coordinators (optional)
	usersAreCoords := "usersAreCoords_example" // string | users are coordinators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetAccounts(context.Background()).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetAccounts`: V0044OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **dELETED** | **string** | include deleted associations | 
 **withAssociations** | **string** | query includes associations | 
 **withCoordinators** | **string** | query includes coordinators | 
 **noUsersAreCoords** | **string** | remove users as coordinators | 
 **usersAreCoords** | **string** | users are coordinators | 

### Return type

[**V0044OpenapiAccountsResp**](V0044OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetAssociation

> V0044OpenapiAssocsResp SlurmdbV0044GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetAssociation`: V0044OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0044OpenapiAssocsResp**](V0044OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetAssociations

> V0044OpenapiAssocsResp SlurmdbV0044GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetAssociations`: V0044OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0044OpenapiAssocsResp**](V0044OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetCluster

> V0044OpenapiClustersResp SlurmdbV0044GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetCluster`: V0044OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0044OpenapiClustersResp**](V0044OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetClusters

> V0044OpenapiClustersResp SlurmdbV0044GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetClusters`: V0044OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0044OpenapiClustersResp**](V0044OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetConfig

> V0044OpenapiSlurmdbdConfigResp SlurmdbV0044GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetConfig`: V0044OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetConfigRequest struct via the builder pattern


### Return type

[**V0044OpenapiSlurmdbdConfigResp**](V0044OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetDiag

> V0044OpenapiSlurmdbdStatsResp SlurmdbV0044GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetDiag`: V0044OpenapiSlurmdbdStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetDiagRequest struct via the builder pattern


### Return type

[**V0044OpenapiSlurmdbdStatsResp**](V0044OpenapiSlurmdbdStatsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetInstance

> V0044OpenapiInstancesResp SlurmdbV0044GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetInstance`: V0044OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0044OpenapiInstancesResp**](V0044OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetInstances

> V0044OpenapiInstancesResp SlurmdbV0044GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetInstances`: V0044OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0044OpenapiInstancesResp**](V0044OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetJob

> V0044OpenapiSlurmdbdJobsResp SlurmdbV0044GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetJob`: V0044OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiSlurmdbdJobsResp**](V0044OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetJobs

> V0044OpenapiSlurmdbdJobsResp SlurmdbV0044GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).JobAltered(jobAltered).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	jobAltered := "jobAltered_example" // string | Job record has been altered (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV WCKey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).JobAltered(jobAltered).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetJobs`: V0044OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **jobAltered** | **string** | Job record has been altered | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV WCKey list | 

### Return type

[**V0044OpenapiSlurmdbdJobsResp**](V0044OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetPing

> V0044OpenapiSlurmdbdPingResp SlurmdbV0044GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetPing`: V0044OpenapiSlurmdbdPingResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetPingRequest struct via the builder pattern


### Return type

[**V0044OpenapiSlurmdbdPingResp**](V0044OpenapiSlurmdbdPingResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetQos

> V0044OpenapiSlurmdbdQosResp SlurmdbV0044GetQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetQos`: V0044OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 

### Return type

[**V0044OpenapiSlurmdbdQosResp**](V0044OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetSingleQos

> V0044OpenapiSlurmdbdQosResp SlurmdbV0044GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetSingleQos`: V0044OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0044OpenapiSlurmdbdQosResp**](V0044OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetTres

> V0044OpenapiTresResp SlurmdbV0044GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetTres`: V0044OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetTresRequest struct via the builder pattern


### Return type

[**V0044OpenapiTresResp**](V0044OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetUser

> V0044OpenapiUsersResp SlurmdbV0044GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetUser`: V0044OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include WCKeys | 

### Return type

[**V0044OpenapiUsersResp**](V0044OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetUsers

> V0044OpenapiUsersResp SlurmdbV0044GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default WCKey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With WCKeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetUsers`: V0044OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default WCKey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With WCKeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0044OpenapiUsersResp**](V0044OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetWckey

> V0044OpenapiWckeyResp SlurmdbV0044GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetWckey`: V0044OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiWckeyResp**](V0044OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetWckeys

> V0044OpenapiWckeyResp SlurmdbV0044GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetWckeys`: V0044OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 

### Return type

[**V0044OpenapiWckeyResp**](V0044OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostAccounts

> V0044OpenapiResp SlurmdbV0044PostAccounts(ctx).V0044OpenapiAccountsResp(v0044OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiAccountsResp := *openapiclient.NewV0044OpenapiAccountsResp([]openapiclient.V0044Account{*openapiclient.NewV0044Account("Description_example", "Name_example", "Organization_example")}) // V0044OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostAccounts(context.Background()).V0044OpenapiAccountsResp(v0044OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostAccounts`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiAccountsResp** | [**V0044OpenapiAccountsResp**](V0044OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostAccountsAssociation

> V0044OpenapiAccountsAddCondRespStr SlurmdbV0044PostAccountsAssociation(ctx).V0044OpenapiAccountsAddCondResp(v0044OpenapiAccountsAddCondResp).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiAccountsAddCondResp := *openapiclient.NewV0044OpenapiAccountsAddCondResp(*openapiclient.NewV0044AccountsAddCond([]string{"Accounts_example"})) // V0044OpenapiAccountsAddCondResp | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostAccountsAssociation(context.Background()).V0044OpenapiAccountsAddCondResp(v0044OpenapiAccountsAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostAccountsAssociation`: V0044OpenapiAccountsAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiAccountsAddCondResp** | [**V0044OpenapiAccountsAddCondResp**](V0044OpenapiAccountsAddCondResp.md) | Add list of accounts with conditional association | 

### Return type

[**V0044OpenapiAccountsAddCondRespStr**](V0044OpenapiAccountsAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostAssociations

> V0044OpenapiResp SlurmdbV0044PostAssociations(ctx).V0044OpenapiAssocsResp(v0044OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiAssocsResp := *openapiclient.NewV0044OpenapiAssocsResp([]openapiclient.V0044Assoc{*openapiclient.NewV0044Assoc("User_example")}) // V0044OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostAssociations(context.Background()).V0044OpenapiAssocsResp(v0044OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostAssociations`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiAssocsResp** | [**V0044OpenapiAssocsResp**](V0044OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostClusters

> V0044OpenapiResp SlurmdbV0044PostClusters(ctx).UpdateTime(updateTime).V0044OpenapiClustersResp(v0044OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)
	v0044OpenapiClustersResp := *openapiclient.NewV0044OpenapiClustersResp([]openapiclient.V0044ClusterRec{*openapiclient.NewV0044ClusterRec()}) // V0044OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostClusters(context.Background()).UpdateTime(updateTime).V0044OpenapiClustersResp(v0044OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostClusters`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 
 **v0044OpenapiClustersResp** | [**V0044OpenapiClustersResp**](V0044OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostConfig

> V0044OpenapiResp SlurmdbV0044PostConfig(ctx).V0044OpenapiSlurmdbdConfigResp(v0044OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiSlurmdbdConfigResp := *openapiclient.NewV0044OpenapiSlurmdbdConfigResp() // V0044OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostConfig(context.Background()).V0044OpenapiSlurmdbdConfigResp(v0044OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostConfig`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiSlurmdbdConfigResp** | [**V0044OpenapiSlurmdbdConfigResp**](V0044OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostJob

> V0044OpenapiJobModifyResp SlurmdbV0044PostJob(ctx, jobId).V0044JobModify(v0044JobModify).Execute()

Update job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	v0044JobModify := *openapiclient.NewV0044JobModify() // V0044JobModify | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostJob(context.Background(), jobId).V0044JobModify(v0044JobModify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostJob`: V0044OpenapiJobModifyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0044JobModify** | [**V0044JobModify**](V0044JobModify.md) | Job update description | 

### Return type

[**V0044OpenapiJobModifyResp**](V0044OpenapiJobModifyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostJobs

> V0044OpenapiJobModifyResp SlurmdbV0044PostJobs(ctx).V0044OpenapiJobModifyReq(v0044OpenapiJobModifyReq).Execute()

Update jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiJobModifyReq := *openapiclient.NewV0044OpenapiJobModifyReq() // V0044OpenapiJobModifyReq | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostJobs(context.Background()).V0044OpenapiJobModifyReq(v0044OpenapiJobModifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostJobs`: V0044OpenapiJobModifyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiJobModifyReq** | [**V0044OpenapiJobModifyReq**](V0044OpenapiJobModifyReq.md) | Job update description | 

### Return type

[**V0044OpenapiJobModifyResp**](V0044OpenapiJobModifyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostQos

> V0044OpenapiResp SlurmdbV0044PostQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0044OpenapiSlurmdbdQosResp(v0044OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	v0044OpenapiSlurmdbdQosResp := *openapiclient.NewV0044OpenapiSlurmdbdQosResp([]openapiclient.V0044Qos{*openapiclient.NewV0044Qos()}) // V0044OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0044OpenapiSlurmdbdQosResp(v0044OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostQos`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **v0044OpenapiSlurmdbdQosResp** | [**V0044OpenapiSlurmdbdQosResp**](V0044OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostTres

> V0044OpenapiResp SlurmdbV0044PostTres(ctx).V0044OpenapiTresResp(v0044OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiTresResp := *openapiclient.NewV0044OpenapiTresResp([]openapiclient.V0044Tres{*openapiclient.NewV0044Tres("Type_example")}) // V0044OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostTres(context.Background()).V0044OpenapiTresResp(v0044OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostTres`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiTresResp** | [**V0044OpenapiTresResp**](V0044OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostUsers

> V0044OpenapiResp SlurmdbV0044PostUsers(ctx).V0044OpenapiUsersResp(v0044OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiUsersResp := *openapiclient.NewV0044OpenapiUsersResp([]openapiclient.V0044User{*openapiclient.NewV0044User("Name_example")}) // V0044OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostUsers(context.Background()).V0044OpenapiUsersResp(v0044OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostUsers`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiUsersResp** | [**V0044OpenapiUsersResp**](V0044OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostUsersAssociation

> V0044OpenapiUsersAddCondRespStr SlurmdbV0044PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).V0044OpenapiUsersAddCondResp(v0044OpenapiUsersAddCondResp).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query partitions updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)
	v0044OpenapiUsersAddCondResp := *openapiclient.NewV0044OpenapiUsersAddCondResp(*openapiclient.NewV0044UsersAddCond([]string{"Users_example"}), *openapiclient.NewV0044UserShort()) // V0044OpenapiUsersAddCondResp | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).V0044OpenapiUsersAddCondResp(v0044OpenapiUsersAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostUsersAssociation`: V0044OpenapiUsersAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 
 **v0044OpenapiUsersAddCondResp** | [**V0044OpenapiUsersAddCondResp**](V0044OpenapiUsersAddCondResp.md) | Create users with conditional association | 

### Return type

[**V0044OpenapiUsersAddCondRespStr**](V0044OpenapiUsersAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostWckeys

> V0044OpenapiResp SlurmdbV0044PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0044OpenapiWckeyResp(v0044OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)
	v0044OpenapiWckeyResp := *openapiclient.NewV0044OpenapiWckeyResp([]openapiclient.V0044Wckey{*openapiclient.NewV0044Wckey("Cluster_example", "Name_example", "User_example")}) // V0044OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0044OpenapiWckeyResp(v0044OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostWckeys`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 
 **v0044OpenapiWckeyResp** | [**V0044OpenapiWckeyResp**](V0044OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

