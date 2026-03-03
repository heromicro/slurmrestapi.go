# SlurmV0041GetLicenses200ResponseLicensesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseName** | Pointer to **string** | Name of the license | [optional] 
**Total** | Pointer to **int32** | Total number of licenses present | [optional] 
**Used** | Pointer to **int32** | Number of licenses in use | [optional] 
**Free** | Pointer to **int32** | Number of licenses currently available | [optional] 
**Remote** | Pointer to **bool** | Indicates whether licenses are served by the database | [optional] 
**Reserved** | Pointer to **int32** | Number of licenses reserved | [optional] 
**LastConsumed** | Pointer to **int32** | Last known number of licenses that were consumed in the license manager (Remote Only) | [optional] 
**LastDeficit** | Pointer to **int32** | Number of \&quot;missing licenses\&quot; from the cluster&#39;s perspective | [optional] 
**LastUpdate** | Pointer to **int64** | When the license information was last updated (UNIX Timestamp) | [optional] 

## Methods

### NewSlurmV0041GetLicenses200ResponseLicensesInner

`func NewSlurmV0041GetLicenses200ResponseLicensesInner() *SlurmV0041GetLicenses200ResponseLicensesInner`

NewSlurmV0041GetLicenses200ResponseLicensesInner instantiates a new SlurmV0041GetLicenses200ResponseLicensesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlurmV0041GetLicenses200ResponseLicensesInnerWithDefaults

`func NewSlurmV0041GetLicenses200ResponseLicensesInnerWithDefaults() *SlurmV0041GetLicenses200ResponseLicensesInner`

NewSlurmV0041GetLicenses200ResponseLicensesInnerWithDefaults instantiates a new SlurmV0041GetLicenses200ResponseLicensesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseName

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### GetTotal

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsed

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetUsed() int32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetUsedOk() (*int32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetUsed(v int32)`

SetUsed sets Used field to given value.

### HasUsed

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasUsed() bool`

HasUsed returns a boolean if a field has been set.

### GetFree

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetFree() int32`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetFreeOk() (*int32, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetFree(v int32)`

SetFree sets Free field to given value.

### HasFree

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetRemote

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasRemote() bool`

HasRemote returns a boolean if a field has been set.

### GetReserved

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetReserved() int32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetReservedOk() (*int32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetReserved(v int32)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetLastConsumed

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetLastConsumed() int32`

GetLastConsumed returns the LastConsumed field if non-nil, zero value otherwise.

### GetLastConsumedOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetLastConsumedOk() (*int32, bool)`

GetLastConsumedOk returns a tuple with the LastConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConsumed

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetLastConsumed(v int32)`

SetLastConsumed sets LastConsumed field to given value.

### HasLastConsumed

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasLastConsumed() bool`

HasLastConsumed returns a boolean if a field has been set.

### GetLastDeficit

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetLastDeficit() int32`

GetLastDeficit returns the LastDeficit field if non-nil, zero value otherwise.

### GetLastDeficitOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetLastDeficitOk() (*int32, bool)`

GetLastDeficitOk returns a tuple with the LastDeficit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeficit

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetLastDeficit(v int32)`

SetLastDeficit sets LastDeficit field to given value.

### HasLastDeficit

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasLastDeficit() bool`

HasLastDeficit returns a boolean if a field has been set.

### GetLastUpdate

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetLastUpdate() int64`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) GetLastUpdateOk() (*int64, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) SetLastUpdate(v int64)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *SlurmV0041GetLicenses200ResponseLicensesInner) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


