﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package cloudLoadTest

import (
    "bytes"
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
    "strconv"
)

var ResourceAreaId, _ = uuid.Parse("7ae6d0a6-cda5-44cf-a261-28c392bed25c")

type Client struct {
    Client azureDevops.Client
}

func NewClient(connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// group (required): Agent group to be created
func (client Client) CreateAgentGroup(group *AgentGroup) (*AgentGroup, error) {
    if group == nil {
        return nil, errors.New("group is a required parameter")
    }
    body, marshalErr := json.Marshal(*group)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("ab8d91c1-12d9-4ec5-874d-1ddb23e17720")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue AgentGroup
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// agentGroupId (optional): The agent group indentifier
// machineSetupInput (optional)
// machineAccessData (optional)
// outgoingRequestUrls (optional)
// agentGroupName (optional): Name of the agent group
func (client Client) GetAgentGroups(agentGroupId *string, machineSetupInput *bool, machineAccessData *bool, outgoingRequestUrls *bool, agentGroupName *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if agentGroupId != nil && *agentGroupId != "" {
        routeValues["agentGroupId"] = *agentGroupId
    }

    queryParams := url.Values{}
    if machineSetupInput != nil {
        queryParams.Add("machineSetupInput", strconv.FormatBool(*machineSetupInput))
    }
    if machineAccessData != nil {
        queryParams.Add("machineAccessData", strconv.FormatBool(*machineAccessData))
    }
    if outgoingRequestUrls != nil {
        queryParams.Add("outgoingRequestUrls", strconv.FormatBool(*outgoingRequestUrls))
    }
    if agentGroupName != nil {
        queryParams.Add("agentGroupName", *agentGroupName)
    }
    locationId, _ := uuid.Parse("ab8d91c1-12d9-4ec5-874d-1ddb23e17720")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// agentGroupId (required): The agent group identifier
// agentName (required): Name of the static agent
func (client Client) DeleteStaticAgent(agentGroupId *string, agentName *string) (*string, error) {
    routeValues := make(map[string]string)
    if agentGroupId == nil || *agentGroupId == "" {
        return nil, errors.New("agentGroupId is a required parameter")
    }
    routeValues["agentGroupId"] = *agentGroupId

    queryParams := url.Values{}
    if agentName == nil {
        return nil, errors.New("agentName is a required parameter")
    }
    queryParams.Add("agentName", *agentName)
    locationId, _ := uuid.Parse("87e4b63d-7142-4b50-801e-72ba9ff8ee9b")
    resp, err := client.Client.Send(http.MethodDelete, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue string
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// agentGroupId (required): The agent group identifier
// agentName (optional): Name of the static agent
func (client Client) GetStaticAgents(agentGroupId *string, agentName *string) (*interface{}, error) {
    routeValues := make(map[string]string)
    if agentGroupId == nil || *agentGroupId == "" {
        return nil, errors.New("agentGroupId is a required parameter")
    }
    routeValues["agentGroupId"] = *agentGroupId

    queryParams := url.Values{}
    if agentName != nil {
        queryParams.Add("agentName", *agentName)
    }
    locationId, _ := uuid.Parse("87e4b63d-7142-4b50-801e-72ba9ff8ee9b")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// applicationId (required): Filter by APM application identifier.
func (client Client) GetApplication(applicationId *string) (*Application, error) {
    routeValues := make(map[string]string)
    if applicationId == nil || *applicationId == "" {
        return nil, errors.New("applicationId is a required parameter")
    }
    routeValues["applicationId"] = *applicationId

    locationId, _ := uuid.Parse("2c986dce-8e8d-4142-b541-d016d5aff764")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue Application
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// type_ (optional): Filters the results based on the plugin type.
func (client Client) GetApplications(type_ *string) (*[]Application, error) {
    queryParams := url.Values{}
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    locationId, _ := uuid.Parse("2c986dce-8e8d-4142-b541-d016d5aff764")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []Application
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// testRunId (required): The test run identifier
// groupNames (required): Comma separated names of counter groups, such as 'Application', 'Performance' and 'Throughput'
// includeSummary (optional)
func (client Client) GetCounters(testRunId *string, groupNames *string, includeSummary *bool) (*[]TestRunCounterInstance, error) {
    routeValues := make(map[string]string)
    if testRunId == nil || *testRunId == "" {
        return nil, errors.New("testRunId is a required parameter")
    }
    routeValues["testRunId"] = *testRunId

    queryParams := url.Values{}
    if groupNames == nil {
        return nil, errors.New("groupNames is a required parameter")
    }
    queryParams.Add("groupNames", *groupNames)
    if includeSummary != nil {
        queryParams.Add("includeSummary", strconv.FormatBool(*includeSummary))
    }
    locationId, _ := uuid.Parse("29265ea4-b5a5-4b2e-b054-47f5f6f00183")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRunCounterInstance
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// applicationId (optional): Filter by APM application identifier.
// plugintype (optional): Currently ApplicationInsights is the only available plugin type.
func (client Client) GetApplicationCounters(applicationId *string, plugintype *string) (*[]ApplicationCounters, error) {
    queryParams := url.Values{}
    if applicationId != nil {
        queryParams.Add("applicationId", *applicationId)
    }
    if plugintype != nil {
        queryParams.Add("plugintype", *plugintype)
    }
    locationId, _ := uuid.Parse("c1275ce9-6d26-4bc6-926b-b846502e812d")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ApplicationCounters
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// counterSampleQueryDetails (required)
// testRunId (required): The test run identifier
func (client Client) GetCounterSamples(counterSampleQueryDetails *azureDevops.VssJsonCollectionWrapper, testRunId *string) (*CounterSamplesResult, error) {
    if counterSampleQueryDetails == nil {
        return nil, errors.New("counterSampleQueryDetails is a required parameter")
    }
    routeValues := make(map[string]string)
    if testRunId == nil || *testRunId == "" {
        return nil, errors.New("testRunId is a required parameter")
    }
    routeValues["testRunId"] = *testRunId

    body, marshalErr := json.Marshal(*counterSampleQueryDetails)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("bad18480-7193-4518-992a-37289c5bb92d")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue CounterSamplesResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// testRunId (required): The test run identifier
// type_ (optional): Filter for the particular type of errors.
// subType (optional): Filter for a particular subtype of errors. You should not provide error subtype without error type.
// detailed (optional): To include the details of test errors such as messagetext, request, stacktrace, testcasename, scenarioname, and lasterrordate.
func (client Client) GetLoadTestRunErrors(testRunId *string, type_ *string, subType *string, detailed *bool) (*LoadTestErrors, error) {
    routeValues := make(map[string]string)
    if testRunId == nil || *testRunId == "" {
        return nil, errors.New("testRunId is a required parameter")
    }
    routeValues["testRunId"] = *testRunId

    queryParams := url.Values{}
    if type_ != nil {
        queryParams.Add("type_", *type_)
    }
    if subType != nil {
        queryParams.Add("subType", *subType)
    }
    if detailed != nil {
        queryParams.Add("detailed", strconv.FormatBool(*detailed))
    }
    locationId, _ := uuid.Parse("b52025a7-3fb4-4283-8825-7079e75bd402")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue LoadTestErrors
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// testRunId (required): Id of the test run
func (client Client) GetTestRunMessages(testRunId *string) (*[]TestRunMessage, error) {
    routeValues := make(map[string]string)
    if testRunId == nil || *testRunId == "" {
        return nil, errors.New("testRunId is a required parameter")
    }
    routeValues["testRunId"] = *testRunId

    locationId, _ := uuid.Parse("2e7ba122-f522-4205-845b-2d270e59850a")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRunMessage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// type_ (required): Currently ApplicationInsights is the only available plugin type.
func (client Client) GetPlugin(type_ *string) (*ApplicationType, error) {
    routeValues := make(map[string]string)
    if type_ == nil || *type_ == "" {
        return nil, errors.New("type_ is a required parameter")
    }
    routeValues["type_"] = *type_

    locationId, _ := uuid.Parse("7dcb0bb2-42d5-4729-9958-c0401d5e7693")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ApplicationType
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

func (client Client) GetPlugins() (*[]ApplicationType, error) {
    locationId, _ := uuid.Parse("7dcb0bb2-42d5-4729-9958-c0401d5e7693")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []ApplicationType
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// testRunId (required): The test run identifier
func (client Client) GetLoadTestResult(testRunId *string) (*TestResults, error) {
    routeValues := make(map[string]string)
    if testRunId == nil || *testRunId == "" {
        return nil, errors.New("testRunId is a required parameter")
    }
    routeValues["testRunId"] = *testRunId

    locationId, _ := uuid.Parse("5ed69bd8-4557-4cec-9b75-1ad67d0c257b")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestResults
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// testDefinition (required): Test definition to be created
func (client Client) CreateTestDefinition(testDefinition *TestDefinition) (*TestDefinition, error) {
    if testDefinition == nil {
        return nil, errors.New("testDefinition is a required parameter")
    }
    body, marshalErr := json.Marshal(*testDefinition)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a8f9b135-f604-41ea-9d74-d9a5fd32fcd8")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// testDefinitionId (required): The test definition identifier
func (client Client) GetTestDefinition(testDefinitionId *string) (*TestDefinition, error) {
    routeValues := make(map[string]string)
    if testDefinitionId == nil || *testDefinitionId == "" {
        return nil, errors.New("testDefinitionId is a required parameter")
    }
    routeValues["testDefinitionId"] = *testDefinitionId

    locationId, _ := uuid.Parse("a8f9b135-f604-41ea-9d74-d9a5fd32fcd8")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// fromDate (optional): Date after which test definitions were created
// toDate (optional): Date before which test definitions were crated
// top (optional)
func (client Client) GetTestDefinitions(fromDate *string, toDate *string, top *int) (*[]TestDefinitionBasic, error) {
    queryParams := url.Values{}
    if fromDate != nil {
        queryParams.Add("fromDate", *fromDate)
    }
    if toDate != nil {
        queryParams.Add("toDate", *toDate)
    }
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    locationId, _ := uuid.Parse("a8f9b135-f604-41ea-9d74-d9a5fd32fcd8")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestDefinitionBasic
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// testDefinition (required)
func (client Client) UpdateTestDefinition(testDefinition *TestDefinition) (*TestDefinition, error) {
    if testDefinition == nil {
        return nil, errors.New("testDefinition is a required parameter")
    }
    body, marshalErr := json.Marshal(*testDefinition)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a8f9b135-f604-41ea-9d74-d9a5fd32fcd8")
    resp, err := client.Client.Send(http.MethodPut, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestDefinition
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// webTestDrop (required): Test drop to be created
func (client Client) CreateTestDrop(webTestDrop *TestDrop) (*TestDrop, error) {
    if webTestDrop == nil {
        return nil, errors.New("webTestDrop is a required parameter")
    }
    body, marshalErr := json.Marshal(*webTestDrop)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("d89d0e08-505c-4357-96f6-9729311ce8ad")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestDrop
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// testDropId (required): The test drop identifier
func (client Client) GetTestDrop(testDropId *string) (*TestDrop, error) {
    routeValues := make(map[string]string)
    if testDropId == nil || *testDropId == "" {
        return nil, errors.New("testDropId is a required parameter")
    }
    routeValues["testDropId"] = *testDropId

    locationId, _ := uuid.Parse("d89d0e08-505c-4357-96f6-9729311ce8ad")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestDrop
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// webTestRun (required)
func (client Client) CreateTestRun(webTestRun *TestRun) (*TestRun, error) {
    if webTestRun == nil {
        return nil, errors.New("webTestRun is a required parameter")
    }
    body, marshalErr := json.Marshal(*webTestRun)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b41a84ff-ff03-4ac1-b76e-e7ea25c92aba")
    resp, err := client.Client.Send(http.MethodPost, locationId, "5.1", nil, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// testRunId (required): Unique ID of the test run
func (client Client) GetTestRun(testRunId *string) (*TestRun, error) {
    routeValues := make(map[string]string)
    if testRunId == nil || *testRunId == "" {
        return nil, errors.New("testRunId is a required parameter")
    }
    routeValues["testRunId"] = *testRunId

    locationId, _ := uuid.Parse("b41a84ff-ff03-4ac1-b76e-e7ea25c92aba")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Returns test runs based on the filter specified. Returns all runs of the tenant if there is no filter.
// name (optional): Name for the test run. Names are not unique. Test runs with same name are assigned sequential rolling numbers.
// requestedBy (optional): Filter by the user who requested the test run. Here requestedBy should be the display name of the user.
// status (optional): Filter by the test run status.
// runType (optional): Valid values include: null, one of TestRunType, or "*"
// fromDate (optional): Filter by the test runs that have been modified after the fromDate timestamp.
// toDate (optional): Filter by the test runs that have been modified before the toDate timestamp.
// detailed (optional): Include the detailed test run attributes.
// top (optional): The maximum number of test runs to return.
// runsourceidentifier (optional)
// retentionState (optional)
func (client Client) GetTestRuns(name *string, requestedBy *string, status *string, runType *string, fromDate *string, toDate *string, detailed *bool, top *int, runsourceidentifier *string, retentionState *string) (*interface{}, error) {
    queryParams := url.Values{}
    if name != nil {
        queryParams.Add("name", *name)
    }
    if requestedBy != nil {
        queryParams.Add("requestedBy", *requestedBy)
    }
    if status != nil {
        queryParams.Add("status", *status)
    }
    if runType != nil {
        queryParams.Add("runType", *runType)
    }
    if fromDate != nil {
        queryParams.Add("fromDate", *fromDate)
    }
    if toDate != nil {
        queryParams.Add("toDate", *toDate)
    }
    if detailed != nil {
        queryParams.Add("detailed", strconv.FormatBool(*detailed))
    }
    if top != nil {
        queryParams.Add("top", strconv.Itoa(*top))
    }
    if runsourceidentifier != nil {
        queryParams.Add("runsourceidentifier", *runsourceidentifier)
    }
    if retentionState != nil {
        queryParams.Add("retentionState", *retentionState)
    }
    locationId, _ := uuid.Parse("b41a84ff-ff03-4ac1-b76e-e7ea25c92aba")
    resp, err := client.Client.Send(http.MethodGet, locationId, "5.1", nil, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue interface{}
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// webTestRun (required)
// testRunId (required)
func (client Client) UpdateTestRun(webTestRun *TestRun, testRunId *string) error {
    if webTestRun == nil {
        return errors.New("webTestRun is a required parameter")
    }
    routeValues := make(map[string]string)
    if testRunId == nil || *testRunId == "" {
        return errors.New("testRunId is a required parameter")
    }
    routeValues["testRunId"] = *testRunId

    body, marshalErr := json.Marshal(*webTestRun)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("b41a84ff-ff03-4ac1-b76e-e7ea25c92aba")
    _, err := client.Client.Send(http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

