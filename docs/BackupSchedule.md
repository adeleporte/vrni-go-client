# BackupSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | True, to enable scheduled backup | [optional] [default to true]
**SchedulePeriod** | [**BackupSchedulePeriod**](BackupSchedulePeriod.md) |  | [optional] 
**Minute** | **int32** | The minute at which backup needs to run (permitted values 0 - 59) | [optional] 
**Hour** | **int32** | The hour at which backup needs to run (permitted values 0 - 23) | [optional] 
**DayOfWeek** | **int32** | The day of the week when backup to be scheduled (permitted values 1{Sunday} - 7{Saturday}) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


