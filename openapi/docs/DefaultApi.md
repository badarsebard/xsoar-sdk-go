# \DefaultApi

All URIs are relative to *https://hostname:443*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAdHocTask**](DefaultApi.md#AddAdHocTask) | **Post** /inv-playbook/task/add/{investigationId} | Add ad-hoc task
[**CloseIncidentsBatch**](DefaultApi.md#CloseIncidentsBatch) | **Post** /incident/batchClose | Batch close incidents
[**CompleteTask**](DefaultApi.md#CompleteTask) | **Post** /inv-playbook/task/complete | [Deprecated] Complete a task
[**CompleteTaskV2**](DefaultApi.md#CompleteTaskV2) | **Post** /v2/inv-playbook/task/complete | Complete a task
[**CopyScript**](DefaultApi.md#CopyScript) | **Post** /automation/copy | Copy automation
[**CreateAccount**](DefaultApi.md#CreateAccount) | **Post** /account | Create an account
[**CreateDockerImage**](DefaultApi.md#CreateDockerImage) | **Post** /settings/docker-images | Create Image
[**CreateFeedIndicatorsJson**](DefaultApi.md#CreateFeedIndicatorsJson) | **Post** /indicators/feed/json | Create feed indicators from JSON
[**CreateHAGroup**](DefaultApi.md#CreateHAGroup) | **Post** /ha-group/create | 
[**CreateHAInstaller**](DefaultApi.md#CreateHAInstaller) | **Post** /host/build/{id} | 
[**CreateHostInstaller**](DefaultApi.md#CreateHostInstaller) | **Post** /host/build | 
[**CreateIncident**](DefaultApi.md#CreateIncident) | **Post** /incident | Create single incident
[**CreateIncidentJson**](DefaultApi.md#CreateIncidentJson) | **Post** /incident/json | Create incident from JSON
[**CreateIncidentsBatch**](DefaultApi.md#CreateIncidentsBatch) | **Post** /incident/batch | Batch create incidents
[**CreateOrUpdateIncidentType**](DefaultApi.md#CreateOrUpdateIncidentType) | **Post** /incidenttype | Create new Incident Type
[**CreateOrUpdateWhitelisted**](DefaultApi.md#CreateOrUpdateWhitelisted) | **Post** /indicators/whitelist/update | Create whitelisted
[**CreateUpdateClassifier**](DefaultApi.md#CreateUpdateClassifier) | **Post** /classifier | Create or update a classifier
[**CreateUpdateClassifierAccount**](DefaultApi.md#CreateUpdateClassifierAccount) | **Post** /{acc}/classifier | Create or update a classifier
[**CreateUpdateIntegrationInstance**](DefaultApi.md#CreateUpdateIntegrationInstance) | **Put** /settings/integration | Create/update an integration instance
[**CreateUpdateIntegrationInstanceAccount**](DefaultApi.md#CreateUpdateIntegrationInstanceAccount) | **Put** /{acc}/settings/integration | Create/update an integration instance
[**DeleteAccount**](DefaultApi.md#DeleteAccount) | **Delete** /account/purge/{accountname} | 
[**DeleteAdHocTask**](DefaultApi.md#DeleteAdHocTask) | **Post** /inv-playbook/task/delete/{investigationId}/{invPBTaskId} | Delete ad-hoc task
[**DeleteAutomationScript**](DefaultApi.md#DeleteAutomationScript) | **Post** /automation/delete | Delete existing automation
[**DeleteClassifier**](DefaultApi.md#DeleteClassifier) | **Delete** /classifier/{id} | Delete a classifier
[**DeleteClassifierAccount**](DefaultApi.md#DeleteClassifierAccount) | **Delete** /{acc}/classifier/{id} | Delete a classifier for account
[**DeleteEvidenceOp**](DefaultApi.md#DeleteEvidenceOp) | **Post** /evidence/delete | delete evidence
[**DeleteHAGroup**](DefaultApi.md#DeleteHAGroup) | **Delete** /ha-group/{id} | 
[**DeleteHost**](DefaultApi.md#DeleteHost) | **Delete** /host/{id} | 
[**DeleteIncidentsBatch**](DefaultApi.md#DeleteIncidentsBatch) | **Post** /incident/batchDelete | Batch delete incidents
[**DeleteIndicatorsBatch**](DefaultApi.md#DeleteIndicatorsBatch) | **Post** /indicators/batchDelete | Batch whitelist or delete indicators
[**DeleteIntegrationInstance**](DefaultApi.md#DeleteIntegrationInstance) | **Delete** /settings/integration/{id} | Delete integration instance
[**DeleteIntegrationInstanceAccount**](DefaultApi.md#DeleteIntegrationInstanceAccount) | **Delete** /{acc}/settings/integration/{id} | Delete integration instance
[**DeleteWidget**](DefaultApi.md#DeleteWidget) | **Delete** /widgets/{id} | Remove existing widget
[**DownloadFile**](DefaultApi.md#DownloadFile) | **Get** /entry/download/{entryid} | Download file
[**DownloadLatestReport**](DefaultApi.md#DownloadLatestReport) | **Get** /report/{id}/latest | Get latest report by ID
[**EditAdHocTask**](DefaultApi.md#EditAdHocTask) | **Post** /inv-playbook/task/edit/{investigationId} | Edit ad-hoc task
[**EntryExportArtifact**](DefaultApi.md#EntryExportArtifact) | **Post** /entry/exportArtifact | Export Artifact
[**ExecuteReport**](DefaultApi.md#ExecuteReport) | **Post** /report/{id}/{requestId}/execute | Execute report
[**ExportIncidentsToCsvBatch**](DefaultApi.md#ExportIncidentsToCsvBatch) | **Post** /incident/batch/exportToCsv | Batch export incidents to csv
[**ExportIndicatorsToCsvBatch**](DefaultApi.md#ExportIndicatorsToCsvBatch) | **Post** /indicators/batch/exportToCsv | Batch export indicators to csv
[**ExportIndicatorsToStixBatch**](DefaultApi.md#ExportIndicatorsToStixBatch) | **Post** /indicators/batch/export/stix | Batch export indicators to STIX
[**GetAllReports**](DefaultApi.md#GetAllReports) | **Get** /reports | Get all reports
[**GetAllWidgets**](DefaultApi.md#GetAllWidgets) | **Get** /widgets | 
[**GetAudits**](DefaultApi.md#GetAudits) | **Post** /settings/audits | Get Audits
[**GetAutomationScripts**](DefaultApi.md#GetAutomationScripts) | **Post** /automation/search | Search Automation (aka scripts)
[**GetContainers**](DefaultApi.md#GetContainers) | **Get** /health/containers | Get health containers
[**GetDockerImages**](DefaultApi.md#GetDockerImages) | **Get** /settings/docker-images | Get Docker Images
[**GetHAGroup**](DefaultApi.md#GetHAGroup) | **Get** /ha-group/{id} | 
[**GetHAInstaller**](DefaultApi.md#GetHAInstaller) | **Get** /host/download/{id} | 
[**GetHostInstaller**](DefaultApi.md#GetHostInstaller) | **Get** /host/download | 
[**GetIncidentAsCsv**](DefaultApi.md#GetIncidentAsCsv) | **Get** /incident/csv/{id} | Get incident as CSV
[**GetIncidentsFieldsByIncidentType**](DefaultApi.md#GetIncidentsFieldsByIncidentType) | **Get** /incidentfields/associatedTypes/{type} | Get all incident fields associated with incident type
[**GetIndicatorsAsCsv**](DefaultApi.md#GetIndicatorsAsCsv) | **Get** /indicators/csv/{id} | Get indicators as CSV
[**GetIndicatorsAsSTIX**](DefaultApi.md#GetIndicatorsAsSTIX) | **Get** /indicators/stix/v2/{id} | Get indicators as STIX V2
[**GetReportByID**](DefaultApi.md#GetReportByID) | **Get** /reports/{id} | Get report by ID
[**GetStatsForDashboard**](DefaultApi.md#GetStatsForDashboard) | **Post** /v2/statistics/dashboards/query | Get Dashboard Statistics
[**GetStatsForDashboardOldFormat**](DefaultApi.md#GetStatsForDashboardOldFormat) | **Post** /statistics/dashboards/query | [Deprecated] Get Dashboard Statistics
[**GetStatsForWidget**](DefaultApi.md#GetStatsForWidget) | **Post** /v2/statistics/widgets/query | Get Widget Statistics
[**GetStatsForWidgetOldFormat**](DefaultApi.md#GetStatsForWidgetOldFormat) | **Post** /statistics/widgets/query | [Deprecated] Get Widget Statistics
[**GetWidget**](DefaultApi.md#GetWidget) | **Get** /widgets/{id} | Get widget by ID
[**HealthHandler**](DefaultApi.md#HealthHandler) | **Get** /health | Check if Cortex XSOAR server is available
[**ImportClassifier**](DefaultApi.md#ImportClassifier) | **Post** /classifier/import | Import a classifier
[**ImportDashboard**](DefaultApi.md#ImportDashboard) | **Post** /dashboards/import | Import a dashboard
[**ImportIncidentFields**](DefaultApi.md#ImportIncidentFields) | **Post** /incidentfields/import | Import an incident field
[**ImportIncidentTypesHandler**](DefaultApi.md#ImportIncidentTypesHandler) | **Post** /incidenttypes/import | Import an incident type
[**ImportScript**](DefaultApi.md#ImportScript) | **Post** /automation/import | Import an automation
[**ImportWidget**](DefaultApi.md#ImportWidget) | **Post** /widgets/import | Import a widget
[**IncidentFileUpload**](DefaultApi.md#IncidentFileUpload) | **Post** /incident/upload/{id} | 
[**IndicatorWhitelist**](DefaultApi.md#IndicatorWhitelist) | **Post** /indicator/whitelist | Whitelists or deletes Indicator
[**IndicatorsCreate**](DefaultApi.md#IndicatorsCreate) | **Post** /indicator/create | Create Indicator
[**IndicatorsCreateBatch**](DefaultApi.md#IndicatorsCreateBatch) | **Post** /indicators/upload | Create indicators
[**IndicatorsEdit**](DefaultApi.md#IndicatorsEdit) | **Post** /indicator/edit | Edit Indicator
[**IndicatorsSearch**](DefaultApi.md#IndicatorsSearch) | **Post** /indicators/search | Search indicators
[**IndicatorsTimelineDelete**](DefaultApi.md#IndicatorsTimelineDelete) | **Post** /indicators/timeline/delete | Delete indicators timeline
[**IntegrationUpload**](DefaultApi.md#IntegrationUpload) | **Post** /settings/integration-conf/upload | Upload an integration
[**InvestigationAddEntriesSync**](DefaultApi.md#InvestigationAddEntriesSync) | **Post** /entry/execute/sync | Create new entry in existing investigation
[**InvestigationAddEntryHandler**](DefaultApi.md#InvestigationAddEntryHandler) | **Post** /entry | Create new entry in existing investigation
[**InvestigationAddFormattedEntryHandler**](DefaultApi.md#InvestigationAddFormattedEntryHandler) | **Post** /entry/formatted | Create new formatted entry in existing investigation
[**ListAccounts**](DefaultApi.md#ListAccounts) | **Get** /accounts | List accounts
[**ListAccountsDetails**](DefaultApi.md#ListAccountsDetails) | **Get** /accounts/data | Detailed accounts
[**ListClassifiers**](DefaultApi.md#ListClassifiers) | **Post** /classifier/search | search classifiers
[**ListClassifiersAccount**](DefaultApi.md#ListClassifiersAccount) | **Post** /{acc}/classifier/search | search classifiers
[**ListHAGroups**](DefaultApi.md#ListHAGroups) | **Get** /ha-groups | 
[**ListHosts**](DefaultApi.md#ListHosts) | **Get** /hosts | 
[**ListIntegrations**](DefaultApi.md#ListIntegrations) | **Post** /settings/integration/search | List integrations
[**ListIntegrationsAccount**](DefaultApi.md#ListIntegrationsAccount) | **Post** /{acc}/settings/integration/search | List integrations
[**ListMainHosts**](DefaultApi.md#ListMainHosts) | **Get** /health/appservers | List the main hosts
[**LogoutEveryoneHandler**](DefaultApi.md#LogoutEveryoneHandler) | **Post** /logout/everyone | Sign out all open users sessions
[**LogoutMyselfHandler**](DefaultApi.md#LogoutMyselfHandler) | **Post** /logout/myself | Sign out all my open sessions
[**LogoutMyselfOtherSessionsHandler**](DefaultApi.md#LogoutMyselfOtherSessionsHandler) | **Post** /logout/myself/other | Sign out all my other open sessions
[**LogoutUserSessionsHandler**](DefaultApi.md#LogoutUserSessionsHandler) | **Post** /logout/user/{username} | Sign out all sessions of the provided username
[**OverridePlaybookYaml**](DefaultApi.md#OverridePlaybookYaml) | **Post** /playbook/save/yaml | Import and override playbook
[**ResetROIWidget**](DefaultApi.md#ResetROIWidget) | **Delete** /statistics/application/roi | Reset ROI widget
[**RevokeUserAPIKey**](DefaultApi.md#RevokeUserAPIKey) | **Post** /apikeys/revoke/user/{username} | 
[**SaveEvidence**](DefaultApi.md#SaveEvidence) | **Post** /evidence | Save evidence
[**SaveOrUpdateScript**](DefaultApi.md#SaveOrUpdateScript) | **Post** /automation | Create or update automation
[**SaveWidget**](DefaultApi.md#SaveWidget) | **Post** /widgets | Add or update a widget
[**SearchEvidence**](DefaultApi.md#SearchEvidence) | **Post** /evidence/search | Search evidence
[**SearchIncidents**](DefaultApi.md#SearchIncidents) | **Post** /incidents/search | Search incidents by filter
[**SearchInvestigations**](DefaultApi.md#SearchInvestigations) | **Post** /investigations/search | Search investigations by filter
[**SetTagsField**](DefaultApi.md#SetTagsField) | **Post** /incidentfield/tags/reset/{id} | Set tags field
[**SimpleCompleteTask**](DefaultApi.md#SimpleCompleteTask) | **Post** /inv-playbook/task/complete/simple | Complete task simple (no file)
[**StartAccounts**](DefaultApi.md#StartAccounts) | **Post** /accounts/start | Start accounts
[**StopAccounts**](DefaultApi.md#StopAccounts) | **Post** /accounts/stop | Stop accounts
[**SubmitTaskForm**](DefaultApi.md#SubmitTaskForm) | **Post** /v2/inv-playbook/task/form/submit | Complete a task
[**TaskAddComment**](DefaultApi.md#TaskAddComment) | **Post** /inv-playbook/task/note/add | Task add comment
[**TaskAssign**](DefaultApi.md#TaskAssign) | **Post** /inv-playbook/task/assign | Assign task
[**TaskSetDue**](DefaultApi.md#TaskSetDue) | **Post** /inv-playbook/task/due | Set task due date
[**TaskUnComplete**](DefaultApi.md#TaskUnComplete) | **Post** /inv-playbook/task/uncomplete | Un complete a task
[**UpdateAccount**](DefaultApi.md#UpdateAccount) | **Post** /account/update/{accountname} | 
[**UpdateAccountHost**](DefaultApi.md#UpdateAccountHost) | **Post** /host/move/{accountname}/{hostgroupid} | 
[**UpdateEntryNote**](DefaultApi.md#UpdateEntryNote) | **Post** /entry/note | Mark entry as note
[**UpdateEntryTagsOp**](DefaultApi.md#UpdateEntryTagsOp) | **Post** /entry/tags | Set entry tags
[**WorkersStatusHandler**](DefaultApi.md#WorkersStatusHandler) | **Get** /workers/status | Get workers status



## AddAdHocTask

> InvestigationPlaybook AddAdHocTask(ctx, investigationId).InvPlaybookTaskData(invPlaybookTaskData).Execute()

Add ad-hoc task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    investigationId := "investigationId_example" // string | investigation ID
    invPlaybookTaskData := *openapiclient.NewInvPlaybookTaskData() // InvPlaybookTaskData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.AddAdHocTask(context.Background(), investigationId).InvPlaybookTaskData(invPlaybookTaskData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AddAdHocTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAdHocTask`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AddAdHocTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **string** | investigation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAdHocTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invPlaybookTaskData** | [**InvPlaybookTaskData**](InvPlaybookTaskData.md) |  | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseIncidentsBatch

> IncidentSearchResponseWrapper CloseIncidentsBatch(ctx).UpdateDataBatch(updateDataBatch).Execute()

Batch close incidents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateDataBatch := *openapiclient.NewUpdateDataBatch() // UpdateDataBatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CloseIncidentsBatch(context.Background()).UpdateDataBatch(updateDataBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CloseIncidentsBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloseIncidentsBatch`: IncidentSearchResponseWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CloseIncidentsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseIncidentsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataBatch** | [**UpdateDataBatch**](UpdateDataBatch.md) |  | 

### Return type

[**IncidentSearchResponseWrapper**](IncidentSearchResponseWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteTask

> InvestigationPlaybook CompleteTask(ctx).InvestigationId(investigationId).FileComment(fileComment).TaskId(taskId).TaskInput(taskInput).Version(version).File(file).FileName(fileName).Execute()

[Deprecated] Complete a task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    investigationId := "investigationId_example" // string | investigation ID
    fileComment := "fileComment_example" // string | file comment
    taskId := "taskId_example" // string | Task Id
    taskInput := "taskInput_example" // string | task input
    version := "version_example" // string | Version
    file := os.NewFile(1234, "some_file") // *os.File | file
    fileName := "fileName_example" // string | file name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CompleteTask(context.Background()).InvestigationId(investigationId).FileComment(fileComment).TaskId(taskId).TaskInput(taskInput).Version(version).File(file).FileName(fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CompleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteTask`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CompleteTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigationId** | **string** | investigation ID | 
 **fileComment** | **string** | file comment | 
 **taskId** | **string** | Task Id | 
 **taskInput** | **string** | task input | 
 **version** | **string** | Version | 
 **file** | ***os.File** | file | 
 **fileName** | **string** | file name | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteTaskV2

> InvestigationPlaybook CompleteTaskV2(ctx).InvestigationId(investigationId).TaskId(taskId).TaskInput(taskInput).Version(version).File(file).TaskComment(taskComment).FileNames(fileNames).FileComments(fileComments).Execute()

Complete a task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    investigationId := "investigationId_example" // string | investigation ID
    taskId := "taskId_example" // string | Task Id
    taskInput := "taskInput_example" // string | Task input
    version := "version_example" // string | Version
    file := os.NewFile(1234, "some_file") // *os.File | Files to attach to the task
    taskComment := "taskComment_example" // string | Task comment or command to run (optional)
    fileNames := "fileNames_example" // string | file names separated by %###% (only if files provided) (optional)
    fileComments := "fileComments_example" // string | file comment separated by %###% (only if files provided) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CompleteTaskV2(context.Background()).InvestigationId(investigationId).TaskId(taskId).TaskInput(taskInput).Version(version).File(file).TaskComment(taskComment).FileNames(fileNames).FileComments(fileComments).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CompleteTaskV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteTaskV2`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CompleteTaskV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTaskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigationId** | **string** | investigation ID | 
 **taskId** | **string** | Task Id | 
 **taskInput** | **string** | Task input | 
 **version** | **string** | Version | 
 **file** | ***os.File** | Files to attach to the task | 
 **taskComment** | **string** | Task comment or command to run | 
 **fileNames** | **string** | file names separated by %###% (only if files provided) | 
 **fileComments** | **string** | file comment separated by %###% (only if files provided) | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyScript

> AutomationScriptResult CopyScript(ctx).AutomationScriptFilterWrapper(automationScriptFilterWrapper).Execute()

Copy automation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    automationScriptFilterWrapper := *openapiclient.NewAutomationScriptFilterWrapper() // AutomationScriptFilterWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CopyScript(context.Background()).AutomationScriptFilterWrapper(automationScriptFilterWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CopyScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CopyScript`: AutomationScriptResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CopyScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCopyScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationScriptFilterWrapper** | [**AutomationScriptFilterWrapper**](AutomationScriptFilterWrapper.md) |  | 

### Return type

[**AutomationScriptResult**](AutomationScriptResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> AccountsWrapper CreateAccount(ctx).CreateAccountRequest(createAccountRequest).Execute()

Create an account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createAccountRequest := *openapiclient.NewCreateAccountRequest() // CreateAccountRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateAccount(context.Background()).CreateAccountRequest(createAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: AccountsWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountRequest** | [**CreateAccountRequest**](CreateAccountRequest.md) |  | 

### Return type

[**AccountsWrapper**](AccountsWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDockerImage

> NewDockerImageResult CreateDockerImage(ctx).NewDockerImageRequest(newDockerImageRequest).Execute()

Create Image



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    newDockerImageRequest := *openapiclient.NewNewDockerImageRequest() // NewDockerImageRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateDockerImage(context.Background()).NewDockerImageRequest(newDockerImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDockerImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDockerImage`: NewDockerImageResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDockerImage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDockerImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newDockerImageRequest** | [**NewDockerImageRequest**](NewDockerImageRequest.md) |  | 

### Return type

[**NewDockerImageResult**](NewDockerImageResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFeedIndicatorsJson

> CreateFeedIndicatorsJson(ctx).FeedIndicatorsRequest(feedIndicatorsRequest).Execute()

Create feed indicators from JSON



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    feedIndicatorsRequest := *openapiclient.NewFeedIndicatorsRequest() // FeedIndicatorsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateFeedIndicatorsJson(context.Background()).FeedIndicatorsRequest(feedIndicatorsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateFeedIndicatorsJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFeedIndicatorsJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **feedIndicatorsRequest** | [**FeedIndicatorsRequest**](FeedIndicatorsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHAGroup

> CreateUpdateHAGroup CreateHAGroup(ctx).CreateHAGroupRequest(createHAGroupRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createHAGroupRequest := *openapiclient.NewCreateHAGroupRequest() // CreateHAGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateHAGroup(context.Background()).CreateHAGroupRequest(createHAGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateHAGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHAGroup`: CreateUpdateHAGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateHAGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHAGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createHAGroupRequest** | [**CreateHAGroupRequest**](CreateHAGroupRequest.md) |  | 

### Return type

[**CreateUpdateHAGroup**](CreateUpdateHAGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHAInstaller

> string CreateHAInstaller(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | HA group ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateHAInstaller(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateHAInstaller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHAInstaller`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateHAInstaller`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | HA group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHAInstallerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHostInstaller

> string CreateHostInstaller(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateHostInstaller(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateHostInstaller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateHostInstaller`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateHostInstaller`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostInstallerRequest struct via the builder pattern


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncident

> IncidentWrapper CreateIncident(ctx).CreateIncidentRequest(createIncidentRequest).Execute()

Create single incident



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createIncidentRequest := *openapiclient.NewCreateIncidentRequest() // CreateIncidentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIncident(context.Background()).CreateIncidentRequest(createIncidentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncident``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncident`: IncidentWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncident`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIncidentRequest** | [**CreateIncidentRequest**](CreateIncidentRequest.md) |  | 

### Return type

[**IncidentWrapper**](IncidentWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncidentJson

> IncidentWrapper CreateIncidentJson(ctx).Execute()

Create incident from JSON



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIncidentJson(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncidentJson``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncidentJson`: IncidentWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncidentJson`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentJsonRequest struct via the builder pattern


### Return type

[**IncidentWrapper**](IncidentWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIncidentsBatch

> IncidentSearchResponseWrapper CreateIncidentsBatch(ctx).UpdateDataBatch(updateDataBatch).Execute()

Batch create incidents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateDataBatch := *openapiclient.NewUpdateDataBatch() // UpdateDataBatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateIncidentsBatch(context.Background()).UpdateDataBatch(updateDataBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIncidentsBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIncidentsBatch`: IncidentSearchResponseWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIncidentsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIncidentsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataBatch** | [**UpdateDataBatch**](UpdateDataBatch.md) |  | 

### Return type

[**IncidentSearchResponseWrapper**](IncidentSearchResponseWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateIncidentType

> IncidentType CreateOrUpdateIncidentType(ctx).IncidentType(incidentType).Execute()

Create new Incident Type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incidentType := *openapiclient.NewIncidentType() // IncidentType |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateOrUpdateIncidentType(context.Background()).IncidentType(incidentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrUpdateIncidentType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateIncidentType`: IncidentType
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrUpdateIncidentType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateIncidentTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incidentType** | [**IncidentType**](IncidentType.md) |  | 

### Return type

[**IncidentType**](IncidentType.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdateWhitelisted

> WhitelistedIndicator CreateOrUpdateWhitelisted(ctx).WhitelistedIndicator(whitelistedIndicator).Execute()

Create whitelisted



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    whitelistedIndicator := *openapiclient.NewWhitelistedIndicator() // WhitelistedIndicator | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateOrUpdateWhitelisted(context.Background()).WhitelistedIndicator(whitelistedIndicator).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrUpdateWhitelisted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateWhitelisted`: WhitelistedIndicator
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrUpdateWhitelisted`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateWhitelistedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **whitelistedIndicator** | [**WhitelistedIndicator**](WhitelistedIndicator.md) |  | 

### Return type

[**WhitelistedIndicator**](WhitelistedIndicator.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUpdateClassifier

> InstanceClassifier CreateUpdateClassifier(ctx).CreateUpdateClassifierRequest(createUpdateClassifierRequest).Execute()

Create or update a classifier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createUpdateClassifierRequest := CreateUpdateClassifierRequest(987) // CreateUpdateClassifierRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUpdateClassifier(context.Background()).CreateUpdateClassifierRequest(createUpdateClassifierRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUpdateClassifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUpdateClassifier`: InstanceClassifier
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUpdateClassifier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateClassifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateClassifierRequest** | **CreateUpdateClassifierRequest** |  | 

### Return type

[**InstanceClassifier**](InstanceClassifier.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUpdateClassifierAccount

> InstanceClassifier CreateUpdateClassifierAccount(ctx, acc).CreateUpdateClassifierAccountRequest(createUpdateClassifierAccountRequest).Execute()

Create or update a classifier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    acc := "acc_example" // string | 
    createUpdateClassifierAccountRequest := CreateUpdateClassifierRequest(987) // CreateUpdateClassifierRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUpdateClassifierAccount(context.Background(), acc).CreateUpdateClassifierAccountRequest(createUpdateClassifierAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUpdateClassifierAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUpdateClassifierAccount`: InstanceClassifier
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUpdateClassifierAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acc** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateClassifierAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUpdateClassifierAccountRequest** | **CreateUpdateClassifierRequest** |  | 

### Return type

[**InstanceClassifier**](InstanceClassifier.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUpdateIntegrationInstance

> map[string]interface{} CreateUpdateIntegrationInstance(ctx).CreateIntegrationRequest(createIntegrationRequest).Execute()

Create/update an integration instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createIntegrationRequest := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUpdateIntegrationInstance(context.Background()).CreateIntegrationRequest(createIntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUpdateIntegrationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUpdateIntegrationInstance`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUpdateIntegrationInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateIntegrationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIntegrationRequest** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUpdateIntegrationInstanceAccount

> map[string]interface{} CreateUpdateIntegrationInstanceAccount(ctx, acc).CreateIntegrationRequest(createIntegrationRequest).Execute()

Create/update an integration instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    acc := "acc_example" // string | 
    createIntegrationRequest := map[string]interface{}(Object) // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateUpdateIntegrationInstanceAccount(context.Background(), acc).CreateIntegrationRequest(createIntegrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUpdateIntegrationInstanceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUpdateIntegrationInstanceAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateUpdateIntegrationInstanceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acc** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUpdateIntegrationInstanceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIntegrationRequest** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> AccountsWrapper DeleteAccount(ctx, accountname).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAccount(context.Background(), accountname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccount`: AccountsWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountsWrapper**](AccountsWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAdHocTask

> InvestigationPlaybook DeleteAdHocTask(ctx, investigationId, invPBTaskId).Execute()

Delete ad-hoc task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    investigationId := "investigationId_example" // string | investigation ID
    invPBTaskId := "invPBTaskId_example" // string | ad-hoc task ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAdHocTask(context.Background(), investigationId, invPBTaskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAdHocTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAdHocTask`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAdHocTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **string** | investigation ID | 
**invPBTaskId** | **string** | ad-hoc task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdHocTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutomationScript

> DeleteAutomationScript(ctx).AutomationScriptFilterWrapper(automationScriptFilterWrapper).Execute()

Delete existing automation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    automationScriptFilterWrapper := *openapiclient.NewAutomationScriptFilterWrapper() // AutomationScriptFilterWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteAutomationScript(context.Background()).AutomationScriptFilterWrapper(automationScriptFilterWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAutomationScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationScriptFilterWrapper** | [**AutomationScriptFilterWrapper**](AutomationScriptFilterWrapper.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClassifier

> DeleteClassifier(ctx, id).Execute()

Delete a classifier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteClassifier(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteClassifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClassifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClassifierAccount

> DeleteClassifierAccount(ctx, id, acc).Execute()

Delete a classifier for account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    acc := "acc_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteClassifierAccount(context.Background(), id, acc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteClassifierAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**acc** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClassifierAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEvidenceOp

> DeleteEvidenceOp(ctx).DeleteEvidenceId(deleteEvidenceId).Execute()

delete evidence



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    deleteEvidenceId := *openapiclient.NewDeleteEvidence() // DeleteEvidence |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteEvidenceOp(context.Background()).DeleteEvidenceId(deleteEvidenceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEvidenceOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEvidenceOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteEvidenceId** | [**DeleteEvidence**](DeleteEvidence.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHAGroup

> string DeleteHAGroup(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | HA group ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteHAGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteHAGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHAGroup`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteHAGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | HA group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHAGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHost

> string DeleteHost(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Host ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteHost(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteHost`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIncidentsBatch

> IncidentSearchResponseWrapper DeleteIncidentsBatch(ctx).UpdateDataBatch(updateDataBatch).Execute()

Batch delete incidents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateDataBatch := *openapiclient.NewUpdateDataBatch() // UpdateDataBatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIncidentsBatch(context.Background()).UpdateDataBatch(updateDataBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIncidentsBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIncidentsBatch`: IncidentSearchResponseWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIncidentsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIncidentsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataBatch** | [**UpdateDataBatch**](UpdateDataBatch.md) |  | 

### Return type

[**IncidentSearchResponseWrapper**](IncidentSearchResponseWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIndicatorsBatch

> UpdateResponse DeleteIndicatorsBatch(ctx).GenericIndicatorUpdateBatch(genericIndicatorUpdateBatch).Execute()

Batch whitelist or delete indicators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    genericIndicatorUpdateBatch := *openapiclient.NewGenericIndicatorUpdateBatch() // GenericIndicatorUpdateBatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIndicatorsBatch(context.Background()).GenericIndicatorUpdateBatch(genericIndicatorUpdateBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIndicatorsBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIndicatorsBatch`: UpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteIndicatorsBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIndicatorsBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericIndicatorUpdateBatch** | [**GenericIndicatorUpdateBatch**](GenericIndicatorUpdateBatch.md) |  | 

### Return type

[**UpdateResponse**](UpdateResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegrationInstance

> DeleteIntegrationInstance(ctx, id).Execute()

Delete integration instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIntegrationInstance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIntegrationInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIntegrationInstanceAccount

> DeleteIntegrationInstanceAccount(ctx, id, acc).Execute()

Delete integration instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | 
    acc := "acc_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteIntegrationInstanceAccount(context.Background(), id, acc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIntegrationInstanceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**acc** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIntegrationInstanceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWidget

> DeleteWidget(ctx, id).Execute()

Remove existing widget



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Widget id to remove (returned from widget save or widgets get)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWidget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Widget id to remove (returned from widget save or widgets get) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadFile

> *os.File DownloadFile(ctx, entryid).Execute()

Download file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    entryid := "entryid_example" // string | Entry ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadFile(context.Background(), entryid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadFile`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryid** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLatestReport

> *os.File DownloadLatestReport(ctx, id).Execute()

Get latest report by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | the ID of the report to get

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DownloadLatestReport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadLatestReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadLatestReport`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadLatestReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the ID of the report to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLatestReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAdHocTask

> InvestigationPlaybook EditAdHocTask(ctx, investigationId).InvPlaybookTaskData(invPlaybookTaskData).Execute()

Edit ad-hoc task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    investigationId := "investigationId_example" // string | investigation ID
    invPlaybookTaskData := *openapiclient.NewInvPlaybookTaskData() // InvPlaybookTaskData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.EditAdHocTask(context.Background(), investigationId).InvPlaybookTaskData(invPlaybookTaskData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditAdHocTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAdHocTask`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EditAdHocTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**investigationId** | **string** | investigation ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAdHocTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invPlaybookTaskData** | [**InvPlaybookTaskData**](InvPlaybookTaskData.md) |  | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EntryExportArtifact

> EntryExportArtifact(ctx).DownloadEntry(downloadEntry).Execute()

Export Artifact



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    downloadEntry := *openapiclient.NewDownloadEntry() // DownloadEntry |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.EntryExportArtifact(context.Background()).DownloadEntry(downloadEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EntryExportArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEntryExportArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **downloadEntry** | [**DownloadEntry**](DownloadEntry.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteReport

> ExecuteReport(ctx, id, requestId).Execute()

Execute report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | the ID of the report to get
    requestId := "requestId_example" // string | the ID to register the request under

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ExecuteReport(context.Background(), id, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExecuteReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the ID of the report to get | 
**requestId** | **string** | the ID to register the request under | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportIncidentsToCsvBatch

> string ExportIncidentsToCsvBatch(ctx).UpdateDataBatch(updateDataBatch).Execute()

Batch export incidents to csv



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateDataBatch := *openapiclient.NewUpdateDataBatch() // UpdateDataBatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ExportIncidentsToCsvBatch(context.Background()).UpdateDataBatch(updateDataBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportIncidentsToCsvBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportIncidentsToCsvBatch`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExportIncidentsToCsvBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportIncidentsToCsvBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateDataBatch** | [**UpdateDataBatch**](UpdateDataBatch.md) |  | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportIndicatorsToCsvBatch

> string ExportIndicatorsToCsvBatch(ctx).GenericIndicatorUpdateBatch(genericIndicatorUpdateBatch).Execute()

Batch export indicators to csv



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    genericIndicatorUpdateBatch := *openapiclient.NewGenericIndicatorUpdateBatch() // GenericIndicatorUpdateBatch | Required parameters from `genericIndicatorUpdateBatch`: `columns`, `filter`. You should also include either `all` or `ids`  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ExportIndicatorsToCsvBatch(context.Background()).GenericIndicatorUpdateBatch(genericIndicatorUpdateBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportIndicatorsToCsvBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportIndicatorsToCsvBatch`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExportIndicatorsToCsvBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportIndicatorsToCsvBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericIndicatorUpdateBatch** | [**GenericIndicatorUpdateBatch**](GenericIndicatorUpdateBatch.md) | Required parameters from &#x60;genericIndicatorUpdateBatch&#x60;: &#x60;columns&#x60;, &#x60;filter&#x60;. You should also include either &#x60;all&#x60; or &#x60;ids&#x60;  | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportIndicatorsToStixBatch

> string ExportIndicatorsToStixBatch(ctx).GenericIndicatorUpdateBatch(genericIndicatorUpdateBatch).Execute()

Batch export indicators to STIX



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    genericIndicatorUpdateBatch := *openapiclient.NewGenericIndicatorUpdateBatch() // GenericIndicatorUpdateBatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ExportIndicatorsToStixBatch(context.Background()).GenericIndicatorUpdateBatch(genericIndicatorUpdateBatch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ExportIndicatorsToStixBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportIndicatorsToStixBatch`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ExportIndicatorsToStixBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportIndicatorsToStixBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genericIndicatorUpdateBatch** | [**GenericIndicatorUpdateBatch**](GenericIndicatorUpdateBatch.md) |  | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllReports

> []Report GetAllReports(ctx).Execute()

Get all reports



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAllReports(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllReports`: []Report
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllReports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReportsRequest struct via the builder pattern


### Return type

[**[]Report**](Report.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllWidgets

> []Widget GetAllWidgets(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAllWidgets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAllWidgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllWidgets`: []Widget
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAllWidgets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllWidgetsRequest struct via the builder pattern


### Return type

[**[]Widget**](Widget.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAudits

> AuditResult GetAudits(ctx).Filter(filter).Execute()

Get Audits



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := *openapiclient.NewGenericStringDateFilter() // GenericStringDateFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAudits(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAudits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAudits`: AuditResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAudits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**GenericStringDateFilter**](GenericStringDateFilter.md) |  | 

### Return type

[**AuditResult**](AuditResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationScripts

> AutomationScriptResult GetAutomationScripts(ctx).AutomationScriptFilter(automationScriptFilter).Execute()

Search Automation (aka scripts)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    automationScriptFilter := *openapiclient.NewAutomationScriptFilter() // AutomationScriptFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetAutomationScripts(context.Background()).AutomationScriptFilter(automationScriptFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAutomationScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutomationScripts`: AutomationScriptResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAutomationScripts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationScriptFilter** | [**AutomationScriptFilter**](AutomationScriptFilter.md) |  | 

### Return type

[**AutomationScriptResult**](AutomationScriptResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainers

> ContainersInfo GetContainers(ctx).Execute()

Get health containers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetContainers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainers`: ContainersInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetContainers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersRequest struct via the builder pattern


### Return type

[**ContainersInfo**](ContainersInfo.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDockerImages

> DockerImagesResult GetDockerImages(ctx).Execute()

Get Docker Images



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDockerImages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDockerImages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDockerImages`: DockerImagesResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDockerImages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDockerImagesRequest struct via the builder pattern


### Return type

[**DockerImagesResult**](DockerImagesResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHAGroup

> CreateUpdateHAGroup GetHAGroup(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | HA group ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetHAGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetHAGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHAGroup`: CreateUpdateHAGroup
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetHAGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | HA group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHAGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateUpdateHAGroup**](CreateUpdateHAGroup.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHAInstaller

> *os.File GetHAInstaller(ctx, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | HA group ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetHAInstaller(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetHAInstaller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHAInstaller`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetHAInstaller`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | HA group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHAInstallerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostInstaller

> *os.File GetHostInstaller(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetHostInstaller(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetHostInstaller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostInstaller`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetHostInstaller`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostInstallerRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentAsCsv

> *os.File GetIncidentAsCsv(ctx, id).Execute()

Get incident as CSV



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | CSV file to fetch (returned from batch export to csv call)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIncidentAsCsv(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIncidentAsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentAsCsv`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIncidentAsCsv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CSV file to fetch (returned from batch export to csv call) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentAsCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIncidentsFieldsByIncidentType

> []IncidentField GetIncidentsFieldsByIncidentType(ctx, type_).Execute()

Get all incident fields associated with incident type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string | the name (case sensitive) of the incident type

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIncidentsFieldsByIncidentType(context.Background(), type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIncidentsFieldsByIncidentType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIncidentsFieldsByIncidentType`: []IncidentField
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIncidentsFieldsByIncidentType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | the name (case sensitive) of the incident type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIncidentsFieldsByIncidentTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IncidentField**](IncidentField.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndicatorsAsCsv

> *os.File GetIndicatorsAsCsv(ctx, id).Execute()

Get indicators as CSV



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | CSV file to fetch (returned from batch export to csv call)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIndicatorsAsCsv(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIndicatorsAsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndicatorsAsCsv`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIndicatorsAsCsv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | CSV file to fetch (returned from batch export to csv call) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndicatorsAsCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndicatorsAsSTIX

> *os.File GetIndicatorsAsSTIX(ctx, id).Execute()

Get indicators as STIX V2



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | STIX V2 file to fetch (returned from batch export to STIX call)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetIndicatorsAsSTIX(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIndicatorsAsSTIX``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndicatorsAsSTIX`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIndicatorsAsSTIX`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | STIX V2 file to fetch (returned from batch export to STIX call) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndicatorsAsSTIXRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportByID

> Report GetReportByID(ctx, id).Execute()

Get report by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | the ID of the report to get

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetReportByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetReportByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportByID`: Report
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetReportByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the ID of the report to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Report**](Report.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatsForDashboard

> []StatsQueryResponse GetStatsForDashboard(ctx).Execute()

Get Dashboard Statistics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatsForDashboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatsForDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsForDashboard`: []StatsQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatsForDashboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsForDashboardRequest struct via the builder pattern


### Return type

[**[]StatsQueryResponse**](StatsQueryResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatsForDashboardOldFormat

> []StatsQueryResponse GetStatsForDashboardOldFormat(ctx).Execute()

[Deprecated] Get Dashboard Statistics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatsForDashboardOldFormat(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatsForDashboardOldFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsForDashboardOldFormat`: []StatsQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatsForDashboardOldFormat`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsForDashboardOldFormatRequest struct via the builder pattern


### Return type

[**[]StatsQueryResponse**](StatsQueryResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatsForWidget

> map[string]interface{} GetStatsForWidget(ctx).Execute()

Get Widget Statistics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatsForWidget(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatsForWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsForWidget`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatsForWidget`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsForWidgetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatsForWidgetOldFormat

> map[string]interface{} GetStatsForWidgetOldFormat(ctx).Execute()

[Deprecated] Get Widget Statistics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetStatsForWidgetOldFormat(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatsForWidgetOldFormat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatsForWidgetOldFormat`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatsForWidgetOldFormat`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsForWidgetOldFormatRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWidget

> Widget GetWidget(ctx, id).Execute()

Get widget by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The ID of widget to get.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetWidget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWidget`: Widget
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWidget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of widget to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Widget**](Widget.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthHandler

> string HealthHandler(ctx).Execute()

Check if Cortex XSOAR server is available



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.HealthHandler(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.HealthHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthHandler`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.HealthHandler`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthHandlerRequest struct via the builder pattern


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportClassifier

> []InstanceClassifier ImportClassifier(ctx).File(file).Execute()

Import a classifier



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | file

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ImportClassifier(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportClassifier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportClassifier`: []InstanceClassifier
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportClassifier`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportClassifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | file | 

### Return type

[**[]InstanceClassifier**](InstanceClassifier.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportDashboard

> Dashboard ImportDashboard(ctx).File(file).Execute()

Import a dashboard



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | The JSON file of the dashboard to import.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ImportDashboard(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The JSON file of the dashboard to import. | 

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportIncidentFields

> IncidentFieldsWithErrors ImportIncidentFields(ctx).File(file).Execute()

Import an incident field



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | file

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ImportIncidentFields(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportIncidentFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportIncidentFields`: IncidentFieldsWithErrors
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportIncidentFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportIncidentFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | file | 

### Return type

[**IncidentFieldsWithErrors**](IncidentFieldsWithErrors.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportIncidentTypesHandler

> IncidentTypesWithErrors ImportIncidentTypesHandler(ctx).File(file).Execute()

Import an incident type



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | file

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ImportIncidentTypesHandler(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportIncidentTypesHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportIncidentTypesHandler`: IncidentTypesWithErrors
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportIncidentTypesHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportIncidentTypesHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | file | 

### Return type

[**IncidentTypesWithErrors**](IncidentTypesWithErrors.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportScript

> AutomationScript ImportScript(ctx).File(file).Execute()

Import an automation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | file

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ImportScript(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportScript`: AutomationScript
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | file | 

### Return type

[**AutomationScript**](AutomationScript.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportWidget

> Widget ImportWidget(ctx).File(file).Execute()

Import a widget



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | The JSON file of the widget to import.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ImportWidget(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ImportWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImportWidget`: Widget
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ImportWidget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportWidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The JSON file of the widget to import. | 

### Return type

[**Widget**](Widget.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncidentFileUpload

> IncidentWrapper IncidentFileUpload(ctx, id).File(file).FileName(fileName).FileComment(fileComment).Field(field).ShowMediaFile(showMediaFile).Last(last).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Incident id to update
    file := os.NewFile(1234, "some_file") // *os.File | file
    fileName := "fileName_example" // string | file name (optional)
    fileComment := "fileComment_example" // string | file comment (optional)
    field := "field_example" // string | field name to hold the attachment details. If not specified, `attachment` will be used. (optional)
    showMediaFile := true // bool | show media file (optional)
    last := true // bool | If set to true will create an investigation. Used for uploading after creating incident. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.IncidentFileUpload(context.Background(), id).File(file).FileName(fileName).FileComment(fileComment).Field(field).ShowMediaFile(showMediaFile).Last(last).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IncidentFileUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncidentFileUpload`: IncidentWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IncidentFileUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Incident id to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncidentFileUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | file | 
 **fileName** | **string** | file name | 
 **fileComment** | **string** | file comment | 
 **field** | **string** | field name to hold the attachment details. If not specified, &#x60;attachment&#x60; will be used. | 
 **showMediaFile** | **bool** | show media file | 
 **last** | **bool** | If set to true will create an investigation. Used for uploading after creating incident. | 

### Return type

[**IncidentWrapper**](IncidentWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndicatorWhitelist

> UpdateResponse IndicatorWhitelist(ctx).UpdateIndicatorReputationData(updateIndicatorReputationData).Execute()

Whitelists or deletes Indicator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateIndicatorReputationData := *openapiclient.NewUpdateIndicatorReputationData() // UpdateIndicatorReputationData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.IndicatorWhitelist(context.Background()).UpdateIndicatorReputationData(updateIndicatorReputationData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IndicatorWhitelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndicatorWhitelist`: UpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IndicatorWhitelist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndicatorWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateIndicatorReputationData** | [**UpdateIndicatorReputationData**](UpdateIndicatorReputationData.md) |  | 

### Return type

[**UpdateResponse**](UpdateResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndicatorsCreate

> IocObject IndicatorsCreate(ctx).IocObject(iocObject).Execute()

Create Indicator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    iocObject := *openapiclient.NewIndicatorContext() // IndicatorContext |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.IndicatorsCreate(context.Background()).IocObject(iocObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IndicatorsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndicatorsCreate`: IocObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IndicatorsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndicatorsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iocObject** | [**IndicatorContext**](IndicatorContext.md) |  | 

### Return type

[**IocObject**](IocObject.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndicatorsCreateBatch

> IocObjects IndicatorsCreateBatch(ctx).File(file).FileName(fileName).Execute()

Create indicators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | file
    fileName := "fileName_example" // string | file name (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.IndicatorsCreateBatch(context.Background()).File(file).FileName(fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IndicatorsCreateBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndicatorsCreateBatch`: IocObjects
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IndicatorsCreateBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndicatorsCreateBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | file | 
 **fileName** | **string** | file name | 

### Return type

[**IocObjects**](IocObjects.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndicatorsEdit

> IocObject IndicatorsEdit(ctx).IocObject(iocObject).Execute()

Edit Indicator



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    iocObject := *openapiclient.NewIocObject() // IocObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.IndicatorsEdit(context.Background()).IocObject(iocObject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IndicatorsEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndicatorsEdit`: IocObject
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IndicatorsEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndicatorsEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iocObject** | [**IocObject**](IocObject.md) |  | 

### Return type

[**IocObject**](IocObject.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndicatorsSearch

> IndicatorResult IndicatorsSearch(ctx).IndicatorFilter(indicatorFilter).Execute()

Search indicators



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    indicatorFilter := *openapiclient.NewIndicatorFilter() // IndicatorFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.IndicatorsSearch(context.Background()).IndicatorFilter(indicatorFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IndicatorsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndicatorsSearch`: IndicatorResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IndicatorsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndicatorsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indicatorFilter** | [**IndicatorFilter**](IndicatorFilter.md) |  | 

### Return type

[**IndicatorResult**](IndicatorResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IndicatorsTimelineDelete

> IndicatorEditBulkResponse IndicatorsTimelineDelete(ctx).IndicatorFilter(indicatorFilter).Execute()

Delete indicators timeline



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    indicatorFilter := *openapiclient.NewIndicatorFilter() // IndicatorFilter |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.IndicatorsTimelineDelete(context.Background()).IndicatorFilter(indicatorFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IndicatorsTimelineDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IndicatorsTimelineDelete`: IndicatorEditBulkResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IndicatorsTimelineDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIndicatorsTimelineDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indicatorFilter** | [**IndicatorFilter**](IndicatorFilter.md) |  | 

### Return type

[**IndicatorEditBulkResponse**](IndicatorEditBulkResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationUpload

> ModuleConfiguration IntegrationUpload(ctx).File(file).Execute()

Upload an integration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | file

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.IntegrationUpload(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.IntegrationUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IntegrationUpload`: ModuleConfiguration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.IntegrationUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | file | 

### Return type

[**ModuleConfiguration**](ModuleConfiguration.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestigationAddEntriesSync

> []Entry InvestigationAddEntriesSync(ctx).UpdateEntry(updateEntry).Execute()

Create new entry in existing investigation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateEntry := *openapiclient.NewUpdateEntry() // UpdateEntry |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.InvestigationAddEntriesSync(context.Background()).UpdateEntry(updateEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InvestigationAddEntriesSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestigationAddEntriesSync`: []Entry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InvestigationAddEntriesSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationAddEntriesSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEntry** | [**UpdateEntry**](UpdateEntry.md) |  | 

### Return type

[**[]Entry**](Entry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestigationAddEntryHandler

> Entry InvestigationAddEntryHandler(ctx).UpdateEntry(updateEntry).Execute()

Create new entry in existing investigation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateEntry := *openapiclient.NewUpdateEntry() // UpdateEntry |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.InvestigationAddEntryHandler(context.Background()).UpdateEntry(updateEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InvestigationAddEntryHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestigationAddEntryHandler`: Entry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InvestigationAddEntryHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationAddEntryHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEntry** | [**UpdateEntry**](UpdateEntry.md) |  | 

### Return type

[**Entry**](Entry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestigationAddFormattedEntryHandler

> Entry InvestigationAddFormattedEntryHandler(ctx).UploadedEntry(uploadedEntry).Execute()

Create new formatted entry in existing investigation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    uploadedEntry := *openapiclient.NewUploadedEntry() // UploadedEntry |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.InvestigationAddFormattedEntryHandler(context.Background()).UploadedEntry(uploadedEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.InvestigationAddFormattedEntryHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestigationAddFormattedEntryHandler`: Entry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.InvestigationAddFormattedEntryHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvestigationAddFormattedEntryHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadedEntry** | [**UploadedEntry**](UploadedEntry.md) |  | 

### Return type

[**Entry**](Entry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> AccountsWrapper ListAccounts(ctx).Execute()

List accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: AccountsWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


### Return type

[**AccountsWrapper**](AccountsWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccountsDetails

> map[string]interface{} ListAccountsDetails(ctx).Execute()

Detailed accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListAccountsDetails(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAccountsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccountsDetails`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAccountsDetails`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsDetailsRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClassifiers

> InstanceClassifiers ListClassifiers(ctx).SearchClassifiers(searchClassifiers).Execute()

search classifiers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    searchClassifiers := *openapiclient.NewInlineObject() // InlineObject |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListClassifiers(context.Background()).SearchClassifiers(searchClassifiers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListClassifiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClassifiers`: InstanceClassifiers
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListClassifiers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClassifiersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchClassifiers** | [**InlineObject**](InlineObject.md) |  | 

### Return type

[**InstanceClassifiers**](InstanceClassifiers.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClassifiersAccount

> InstanceClassifiers ListClassifiersAccount(ctx, acc).SearchClassifiersAccount(searchClassifiersAccount).Execute()

search classifiers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    acc := "acc_example" // string | 
    searchClassifiersAccount := *openapiclient.NewInlineObject1() // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListClassifiersAccount(context.Background(), acc).SearchClassifiersAccount(searchClassifiersAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListClassifiersAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClassifiersAccount`: InstanceClassifiers
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListClassifiersAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acc** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClassifiersAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchClassifiersAccount** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**InstanceClassifiers**](InstanceClassifiers.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHAGroups

> HAGroups ListHAGroups(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListHAGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListHAGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHAGroups`: HAGroups
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListHAGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHAGroupsRequest struct via the builder pattern


### Return type

[**HAGroups**](HAGroups.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosts

> Hosts ListHosts(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListHosts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListHosts`: Hosts
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListHosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHostsRequest struct via the builder pattern


### Return type

[**Hosts**](Hosts.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrations

> map[string]interface{} ListIntegrations(ctx).Size(size).Execute()

List integrations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    size := *openapiclient.NewInlineObject2() // InlineObject2 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIntegrations(context.Background()).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIntegrations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrations`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationsAccount

> map[string]interface{} ListIntegrationsAccount(ctx, acc).Size(size).Execute()

List integrations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    acc := "acc_example" // string | 
    size := *openapiclient.NewInlineObject3() // InlineObject3 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListIntegrationsAccount(context.Background(), acc).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIntegrationsAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIntegrationsAccount`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIntegrationsAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**acc** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | [**InlineObject3**](InlineObject3.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMainHosts

> MainHosts ListMainHosts(ctx).Execute()

List the main hosts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListMainHosts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMainHosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMainHosts`: MainHosts
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMainHosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMainHostsRequest struct via the builder pattern


### Return type

[**MainHosts**](MainHosts.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutEveryoneHandler

> string LogoutEveryoneHandler(ctx).Execute()

Sign out all open users sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.LogoutEveryoneHandler(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LogoutEveryoneHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogoutEveryoneHandler`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LogoutEveryoneHandler`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutEveryoneHandlerRequest struct via the builder pattern


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutMyselfHandler

> string LogoutMyselfHandler(ctx).Execute()

Sign out all my open sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.LogoutMyselfHandler(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LogoutMyselfHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogoutMyselfHandler`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LogoutMyselfHandler`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutMyselfHandlerRequest struct via the builder pattern


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutMyselfOtherSessionsHandler

> string LogoutMyselfOtherSessionsHandler(ctx).Execute()

Sign out all my other open sessions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.LogoutMyselfOtherSessionsHandler(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LogoutMyselfOtherSessionsHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogoutMyselfOtherSessionsHandler`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LogoutMyselfOtherSessionsHandler`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutMyselfOtherSessionsHandlerRequest struct via the builder pattern


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutUserSessionsHandler

> string LogoutUserSessionsHandler(ctx, username).Execute()

Sign out all sessions of the provided username



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | Username to logout

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.LogoutUserSessionsHandler(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LogoutUserSessionsHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogoutUserSessionsHandler`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LogoutUserSessionsHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Username to logout | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutUserSessionsHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OverridePlaybookYaml

> PlaybookWithWarnings OverridePlaybookYaml(ctx).File(file).Execute()

Import and override playbook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | file

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.OverridePlaybookYaml(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.OverridePlaybookYaml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OverridePlaybookYaml`: PlaybookWithWarnings
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.OverridePlaybookYaml`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOverridePlaybookYamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | file | 

### Return type

[**PlaybookWithWarnings**](PlaybookWithWarnings.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetROIWidget

> ResetROIWidget(ctx).Execute()

Reset ROI widget



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResetROIWidget(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResetROIWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetROIWidgetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeUserAPIKey

> RevokeUserAPIKey(ctx, username).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    username := "username_example" // string | The username which the API keys assigned to

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RevokeUserAPIKey(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RevokeUserAPIKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username which the API keys assigned to | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserAPIKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveEvidence

> Evidence SaveEvidence(ctx).Evidence(evidence).Execute()

Save evidence



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    evidence := *openapiclient.NewEvidence() // Evidence |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SaveEvidence(context.Background()).Evidence(evidence).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SaveEvidence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveEvidence`: Evidence
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SaveEvidence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evidence** | [**Evidence**](Evidence.md) |  | 

### Return type

[**Evidence**](Evidence.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveOrUpdateScript

> AutomationScriptResult SaveOrUpdateScript(ctx).AutomationScriptFilterWrapper(automationScriptFilterWrapper).Execute()

Create or update automation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    automationScriptFilterWrapper := *openapiclient.NewAutomationScriptFilterWrapper() // AutomationScriptFilterWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SaveOrUpdateScript(context.Background()).AutomationScriptFilterWrapper(automationScriptFilterWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SaveOrUpdateScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveOrUpdateScript`: AutomationScriptResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SaveOrUpdateScript`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveOrUpdateScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **automationScriptFilterWrapper** | [**AutomationScriptFilterWrapper**](AutomationScriptFilterWrapper.md) |  | 

### Return type

[**AutomationScriptResult**](AutomationScriptResult.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveWidget

> Widget SaveWidget(ctx).Widget(widget).Execute()

Add or update a widget



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    widget := *openapiclient.NewWidget("Name_example", "WidgetType_example") // Widget |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SaveWidget(context.Background()).Widget(widget).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SaveWidget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveWidget`: Widget
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SaveWidget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveWidgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **widget** | [**Widget**](Widget.md) |  | 

### Return type

[**Widget**](Widget.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEvidence

> EvidencesSearchResponse SearchEvidence(ctx).EvidencesFilterWrapper(evidencesFilterWrapper).Execute()

Search evidence



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    evidencesFilterWrapper := *openapiclient.NewEvidencesFilterWrapper() // EvidencesFilterWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SearchEvidence(context.Background()).EvidencesFilterWrapper(evidencesFilterWrapper).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchEvidence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEvidence`: EvidencesSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchEvidence`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evidencesFilterWrapper** | [**EvidencesFilterWrapper**](EvidencesFilterWrapper.md) |  | 

### Return type

[**EvidencesSearchResponse**](EvidencesSearchResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIncidents

> InlineResponse200 SearchIncidents(ctx).Filter(filter).Execute()

Search incidents by filter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := *openapiclient.NewSearchIncidentsData() // SearchIncidentsData | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SearchIncidents(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchIncidents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIncidents`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchIncidents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIncidentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**SearchIncidentsData**](SearchIncidentsData.md) |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchInvestigations

> InvestigationSearchResponse SearchInvestigations(ctx).Filter(filter).Execute()

Search investigations by filter



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := *openapiclient.NewSearchInvestigationsData() // SearchInvestigationsData |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SearchInvestigations(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchInvestigations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchInvestigations`: InvestigationSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchInvestigations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchInvestigationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**SearchInvestigationsData**](SearchInvestigationsData.md) |  | 

### Return type

[**InvestigationSearchResponse**](InvestigationSearchResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetTagsField

> SetTagsField(ctx, id).Data(data).Execute()

Set tags field



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The machine name of the field prefixed with the type. For example indicator_tags or incident_dbotmirrortags
    data := *openapiclient.NewTagsFieldValues() // TagsFieldValues | The new select values of the field

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SetTagsField(context.Background(), id).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetTagsField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The machine name of the field prefixed with the type. For example indicator_tags or incident_dbotmirrortags | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetTagsFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **data** | [**TagsFieldValues**](TagsFieldValues.md) | The new select values of the field | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SimpleCompleteTask

> InvestigationPlaybook SimpleCompleteTask(ctx).InvTaskInfo(invTaskInfo).Execute()

Complete task simple (no file)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    invTaskInfo := *openapiclient.NewInvTaskInfo() // InvTaskInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SimpleCompleteTask(context.Background()).InvTaskInfo(invTaskInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SimpleCompleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SimpleCompleteTask`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SimpleCompleteTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSimpleCompleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invTaskInfo** | [**InvTaskInfo**](InvTaskInfo.md) |  | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartAccounts

> StartAccounts(ctx).StartAccountsRequest(startAccountsRequest).Execute()

Start accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    startAccountsRequest := *openapiclient.NewAccountsWrapper() // AccountsWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.StartAccounts(context.Background()).StartAccountsRequest(startAccountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StartAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAccountsRequest** | [**AccountsWrapper**](AccountsWrapper.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopAccounts

> StopAccounts(ctx).StopAccountsRequest(stopAccountsRequest).Execute()

Stop accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    stopAccountsRequest := *openapiclient.NewAccountsWrapper() // AccountsWrapper |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.StopAccounts(context.Background()).StopAccountsRequest(stopAccountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StopAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStopAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stopAccountsRequest** | [**AccountsWrapper**](AccountsWrapper.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitTaskForm

> InvestigationPlaybook SubmitTaskForm(ctx).InvestigationId(investigationId).TaskId(taskId).Answers(answers).File(file).Execute()

Complete a task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    investigationId := "investigationId_example" // string | investigation ID
    taskId := "taskId_example" // string | Task Id
    answers := "answers_example" // string | the answers to the task form. Answers are keyed by numerical question id
    file := os.NewFile(1234, "some_file") // *os.File | Files to attach to the task (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SubmitTaskForm(context.Background()).InvestigationId(investigationId).TaskId(taskId).Answers(answers).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubmitTaskForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitTaskForm`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubmitTaskForm`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitTaskFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investigationId** | **string** | investigation ID | 
 **taskId** | **string** | Task Id | 
 **answers** | **string** | the answers to the task form. Answers are keyed by numerical question id | 
 **file** | ***os.File** | Files to attach to the task | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskAddComment

> InvestigationPlaybook TaskAddComment(ctx).InvTaskInfo(invTaskInfo).Execute()

Task add comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    invTaskInfo := *openapiclient.NewInvTaskInfo() // InvTaskInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TaskAddComment(context.Background()).InvTaskInfo(invTaskInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TaskAddComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskAddComment`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TaskAddComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskAddCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invTaskInfo** | [**InvTaskInfo**](InvTaskInfo.md) |  | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskAssign

> InvestigationPlaybook TaskAssign(ctx).InvPlaybookAssignee(invPlaybookAssignee).Execute()

Assign task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    invPlaybookAssignee := *openapiclient.NewInvPlaybookAssignee() // InvPlaybookAssignee |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TaskAssign(context.Background()).InvPlaybookAssignee(invPlaybookAssignee).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TaskAssign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskAssign`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TaskAssign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskAssignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invPlaybookAssignee** | [**InvPlaybookAssignee**](InvPlaybookAssignee.md) |  | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskSetDue

> InvestigationPlaybook TaskSetDue(ctx).InvPlaybookDue(invPlaybookDue).Execute()

Set task due date



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    invPlaybookDue := *openapiclient.NewInvPlaybookDue() // InvPlaybookDue |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TaskSetDue(context.Background()).InvPlaybookDue(invPlaybookDue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TaskSetDue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskSetDue`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TaskSetDue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskSetDueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invPlaybookDue** | [**InvPlaybookDue**](InvPlaybookDue.md) |  | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskUnComplete

> InvestigationPlaybook TaskUnComplete(ctx).InvTaskInfo(invTaskInfo).Execute()

Un complete a task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    invTaskInfo := *openapiclient.NewInvTaskInfo() // InvTaskInfo |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.TaskUnComplete(context.Background()).InvTaskInfo(invTaskInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.TaskUnComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaskUnComplete`: InvestigationPlaybook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.TaskUnComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskUnCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invTaskInfo** | [**InvTaskInfo**](InvTaskInfo.md) |  | 

### Return type

[**InvestigationPlaybook**](InvestigationPlaybook.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> RolesAndPropagationLabelsWrapper UpdateAccount(ctx, accountname).UpdateRolesAndPropagationLabelsRequest(updateRolesAndPropagationLabelsRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | 
    updateRolesAndPropagationLabelsRequest := *openapiclient.NewUpdateRolesAndPropagationLabelsRequest() // UpdateRolesAndPropagationLabelsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAccount(context.Background(), accountname).UpdateRolesAndPropagationLabelsRequest(updateRolesAndPropagationLabelsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: RolesAndPropagationLabelsWrapper
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRolesAndPropagationLabelsRequest** | [**UpdateRolesAndPropagationLabelsRequest**](UpdateRolesAndPropagationLabelsRequest.md) |  | 

### Return type

[**RolesAndPropagationLabelsWrapper**](RolesAndPropagationLabelsWrapper.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccountHost

> UpdateAccountHostResponse UpdateAccountHost(ctx, accountname, hostgroupid).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountname := "accountname_example" // string | Account Name
    hostgroupid := "hostgroupid_example" // string | Host Group Id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateAccountHost(context.Background(), accountname, hostgroupid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccountHost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccountHost`: UpdateAccountHostResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAccountHost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountname** | **string** | Account Name | 
**hostgroupid** | **string** | Host Group Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountHostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UpdateAccountHostResponse**](UpdateAccountHostResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntryNote

> Entry UpdateEntryNote(ctx).UpdateEntry(updateEntry).Execute()

Mark entry as note



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateEntry := *openapiclient.NewUpdateEntry() // UpdateEntry |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateEntryNote(context.Background()).UpdateEntry(updateEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEntryNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntryNote`: Entry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEntryNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntryNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEntry** | [**UpdateEntry**](UpdateEntry.md) |  | 

### Return type

[**Entry**](Entry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntryTagsOp

> Entry UpdateEntryTagsOp(ctx).UpdateEntryTags(updateEntryTags).Execute()

Set entry tags



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    updateEntryTags := *openapiclient.NewUpdateEntryTags() // UpdateEntryTags |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateEntryTagsOp(context.Background()).UpdateEntryTags(updateEntryTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEntryTagsOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntryTagsOp`: Entry
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEntryTagsOp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntryTagsOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEntryTags** | [**UpdateEntryTags**](UpdateEntryTags.md) |  | 

### Return type

[**Entry**](Entry.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkersStatusHandler

> Info WorkersStatusHandler(ctx).Execute()

Get workers status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.WorkersStatusHandler(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.WorkersStatusHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WorkersStatusHandler`: Info
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.WorkersStatusHandler`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorkersStatusHandlerRequest struct via the builder pattern


### Return type

[**Info**](Info.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

