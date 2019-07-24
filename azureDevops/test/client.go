﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package test

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "io"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

var ResourceAreaId, _ = uuid.Parse("c2aa639c-3ccc-4740-b3b6-ce2a1e1d984e")

type Client struct {
    Client azureDevops.Client
}

func NewClient(ctx context.Context, connection azureDevops.Connection) (*Client, error) {
    client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
    if err != nil {
        return nil, err
    }
    return &Client {
        Client: *client,
    }, nil
}

// Gets the action results for an iteration in a test result.
func (client Client) GetActionResults(ctx context.Context, args GetActionResultsArgs) (*[]TestActionResultModel, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.IterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*args.IterationId)
    if args.ActionPath != nil && *args.ActionPath != "" {
        routeValues["actionPath"] = *args.ActionPath
    }

    locationId, _ := uuid.Parse("eaf40c31-ff84-4062-aafd-d5664be11a37")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestActionResultModel
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetActionResults function
type GetActionResultsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test result that contains the iterations.
    TestCaseResultId *int
    // (required) ID of the iteration that contains the actions.
    IterationId *int
    // (optional) Path of a specific action, used to get just that action.
    ActionPath *string
}

// [Preview API] Attach a file to a test result.
func (client Client) CreateTestResultAttachment(ctx context.Context, args CreateTestResultAttachmentArgs) (*TestAttachmentReference, error) {
    if args.AttachmentRequestModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestResultAttachment function
type CreateTestResultAttachmentArgs struct {
    // (required) Attachment details TestAttachmentRequestModel
    AttachmentRequestModel *TestAttachmentRequestModel
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test result against which attachment has to be uploaded.
    TestCaseResultId *int
}

// [Preview API] Attach a file to a test result
func (client Client) CreateTestSubResultAttachment(ctx context.Context, args CreateTestSubResultAttachmentArgs) (*TestAttachmentReference, error) {
    if args.AttachmentRequestModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    queryParams := url.Values{}
    if args.TestSubResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*args.TestSubResultId))
    body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestSubResultAttachment function
type CreateTestSubResultAttachmentArgs struct {
    // (required) Attachment Request Model.
    AttachmentRequestModel *TestAttachmentRequestModel
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test results that contains sub result.
    TestCaseResultId *int
    // (required) ID of the test sub results against which attachment has to be uploaded.
    TestSubResultId *int
}

// [Preview API] Download a test result attachment by its ID.
func (client Client) GetTestResultAttachmentContent(ctx context.Context, args GetTestResultAttachmentContentArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestResultAttachmentContent function
type GetTestResultAttachmentContentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the testCaseResultId.
    RunId *int
    // (required) ID of the test result whose attachment has to be downloaded.
    TestCaseResultId *int
    // (required) ID of the test result attachment to be downloaded.
    AttachmentId *int
}

// [Preview API] Get list of test result attachments reference.
func (client Client) GetTestResultAttachments(ctx context.Context, args GetTestResultAttachmentsArgs) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultAttachments function
type GetTestResultAttachmentsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test result.
    TestCaseResultId *int
}

// [Preview API] Download a test result attachment by its ID.
func (client Client) GetTestResultAttachmentZip(ctx context.Context, args GetTestResultAttachmentZipArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestResultAttachmentZip function
type GetTestResultAttachmentZipArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the testCaseResultId.
    RunId *int
    // (required) ID of the test result whose attachment has to be downloaded.
    TestCaseResultId *int
    // (required) ID of the test result attachment to be downloaded.
    AttachmentId *int
}

// [Preview API] Download a test sub result attachment
func (client Client) GetTestSubResultAttachmentContent(ctx context.Context, args GetTestSubResultAttachmentContentArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    queryParams := url.Values{}
    if args.TestSubResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*args.TestSubResultId))
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestSubResultAttachmentContent function
type GetTestSubResultAttachmentContentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test results that contains sub result.
    TestCaseResultId *int
    // (required) ID of the test result attachment to be downloaded
    AttachmentId *int
    // (required) ID of the test sub result whose attachment has to be downloaded
    TestSubResultId *int
}

// [Preview API] Get list of test sub result attachments
func (client Client) GetTestSubResultAttachments(ctx context.Context, args GetTestSubResultAttachmentsArgs) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    queryParams := url.Values{}
    if args.TestSubResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*args.TestSubResultId))
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestSubResultAttachments function
type GetTestSubResultAttachmentsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test results that contains sub result.
    TestCaseResultId *int
    // (required) ID of the test sub result whose attachment has to be downloaded
    TestSubResultId *int
}

// [Preview API] Download a test sub result attachment
func (client Client) GetTestSubResultAttachmentZip(ctx context.Context, args GetTestSubResultAttachmentZipArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.AttachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    queryParams := url.Values{}
    if args.TestSubResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSubResultId"}
    }
    queryParams.Add("testSubResultId", strconv.Itoa(*args.TestSubResultId))
    locationId, _ := uuid.Parse("2bffebe9-2f0f-4639-9af8-56129e9fed2d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestSubResultAttachmentZip function
type GetTestSubResultAttachmentZipArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test results that contains sub result.
    TestCaseResultId *int
    // (required) ID of the test result attachment to be downloaded
    AttachmentId *int
    // (required) ID of the test sub result whose attachment has to be downloaded
    TestSubResultId *int
}

// [Preview API] Attach a file to a test run.
func (client Client) CreateTestRunAttachment(ctx context.Context, args CreateTestRunAttachmentArgs) (*TestAttachmentReference, error) {
    if args.AttachmentRequestModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentRequestModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    body, marshalErr := json.Marshal(*args.AttachmentRequestModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4f004af4-a507-489c-9b13-cb62060beb11")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestAttachmentReference
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestRunAttachment function
type CreateTestRunAttachmentArgs struct {
    // (required) Attachment details TestAttachmentRequestModel
    AttachmentRequestModel *TestAttachmentRequestModel
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run against which attachment has to be uploaded.
    RunId *int
}

// [Preview API] Download a test run attachment by its ID.
func (client Client) GetTestRunAttachmentContent(ctx context.Context, args GetTestRunAttachmentContentArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.AttachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("4f004af4-a507-489c-9b13-cb62060beb11")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestRunAttachmentContent function
type GetTestRunAttachmentContentArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run whose attachment has to be downloaded.
    RunId *int
    // (required) ID of the test run attachment to be downloaded.
    AttachmentId *int
}

// [Preview API] Get list of test run attachments reference.
func (client Client) GetTestRunAttachments(ctx context.Context, args GetTestRunAttachmentsArgs) (*[]TestAttachment, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    locationId, _ := uuid.Parse("4f004af4-a507-489c-9b13-cb62060beb11")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestAttachment
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunAttachments function
type GetTestRunAttachmentsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run.
    RunId *int
}

// [Preview API] Download a test run attachment by its ID.
func (client Client) GetTestRunAttachmentZip(ctx context.Context, args GetTestRunAttachmentZipArgs) (io.ReadCloser, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.AttachmentId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "attachmentId"} 
    }
    routeValues["attachmentId"] = strconv.Itoa(*args.AttachmentId)

    locationId, _ := uuid.Parse("4f004af4-a507-489c-9b13-cb62060beb11")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
    if err != nil {
        return nil, err
    }

    return resp.Body, err
}

// Arguments for the GetTestRunAttachmentZip function
type GetTestRunAttachmentZipArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run whose attachment has to be downloaded.
    RunId *int
    // (required) ID of the test run attachment to be downloaded.
    AttachmentId *int
}

// [Preview API] Get code coverage data for a build.
func (client Client) GetBuildCodeCoverage(ctx context.Context, args GetBuildCodeCoverageArgs) (*[]BuildCoverage, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "buildId"}
    }
    queryParams.Add("buildId", strconv.Itoa(*args.BuildId))
    if args.Flags == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "flags"}
    }
    queryParams.Add("flags", strconv.Itoa(*args.Flags))
    locationId, _ := uuid.Parse("77560e8a-4e8c-4d59-894e-a5f264c24444")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []BuildCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetBuildCodeCoverage function
type GetBuildCodeCoverageArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the build for which code coverage data needs to be fetched.
    BuildId *int
    // (required) Value of flags determine the level of code coverage details to be fetched. Flags are additive. Expected Values are 1 for Modules, 2 for Functions, 4 for BlockData.
    Flags *int
}

// [Preview API] Get code coverage data for a test run
func (client Client) GetTestRunCodeCoverage(ctx context.Context, args GetTestRunCodeCoverageArgs) (*[]TestRunCoverage, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.Flags == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "flags"}
    }
    queryParams.Add("flags", strconv.Itoa(*args.Flags))
    locationId, _ := uuid.Parse("9629116f-3b89-4ed8-b358-d4694efda160")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRunCoverage
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunCodeCoverage function
type GetTestRunCodeCoverageArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run for which code coverage data needs to be fetched.
    RunId *int
    // (required) Value of flags determine the level of code coverage details to be fetched. Flags are additive. Expected Values are 1 for Modules, 2 for Functions, 4 for BlockData.
    Flags *int
}

// Get iteration for a result
func (client Client) GetTestIteration(ctx context.Context, args GetTestIterationArgs) (*TestIterationDetailsModel, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.IterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*args.IterationId)

    queryParams := url.Values{}
    if args.IncludeActionResults != nil {
        queryParams.Add("includeActionResults", strconv.FormatBool(*args.IncludeActionResults))
    }
    locationId, _ := uuid.Parse("73eb9074-3446-4c44-8296-2f811950ff8d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestIterationDetailsModel
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestIteration function
type GetTestIterationArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test result that contains the iterations.
    TestCaseResultId *int
    // (required) Id of the test results Iteration.
    IterationId *int
    // (optional) Include result details for each action performed in the test iteration. ActionResults refer to outcome (pass/fail) of test steps that are executed as part of a running a manual test. Including the ActionResults flag gets the outcome of test steps in the actionResults section and test parameters in the parameters section for each test iteration.
    IncludeActionResults *bool
}

// Get iterations for a result
func (client Client) GetTestIterations(ctx context.Context, args GetTestIterationsArgs) (*[]TestIterationDetailsModel, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    queryParams := url.Values{}
    if args.IncludeActionResults != nil {
        queryParams.Add("includeActionResults", strconv.FormatBool(*args.IncludeActionResults))
    }
    locationId, _ := uuid.Parse("73eb9074-3446-4c44-8296-2f811950ff8d")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestIterationDetailsModel
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestIterations function
type GetTestIterationsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test result that contains the iterations.
    TestCaseResultId *int
    // (optional) Include result details for each action performed in the test iteration. ActionResults refer to outcome (pass/fail) of test steps that are executed as part of a running a manual test. Including the ActionResults flag gets the outcome of test steps in the actionResults section and test parameters in the parameters section for each test iteration.
    IncludeActionResults *bool
}

// Get a list of parameterized results
func (client Client) GetResultParameters(ctx context.Context, args GetResultParametersArgs) (*[]TestResultParameterModel, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)
    if args.IterationId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "iterationId"} 
    }
    routeValues["iterationId"] = strconv.Itoa(*args.IterationId)

    queryParams := url.Values{}
    if args.ParamName != nil {
        queryParams.Add("paramName", *args.ParamName)
    }
    locationId, _ := uuid.Parse("7c69810d-3354-4af3-844a-180bd25db08a")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestResultParameterModel
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetResultParameters function
type GetResultParametersArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test run that contains the result.
    RunId *int
    // (required) ID of the test result that contains the iterations.
    TestCaseResultId *int
    // (required) ID of the iteration that contains the parameterized results.
    IterationId *int
    // (optional) Name of the parameter.
    ParamName *string
}

// Get a test point.
func (client Client) GetPoint(ctx context.Context, args GetPointArgs) (*TestPoint, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.PointIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pointIds"} 
    }
    routeValues["pointIds"] = strconv.Itoa(*args.PointIds)

    queryParams := url.Values{}
    if args.WitFields != nil {
        queryParams.Add("witFields", *args.WitFields)
    }
    locationId, _ := uuid.Parse("3bcfd5c8-be62-488e-b1da-b8289ce9299c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestPoint
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPoint function
type GetPointArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan.
    PlanId *int
    // (required) ID of the suite that contains the point.
    SuiteId *int
    // (required) ID of the test point to get.
    PointIds *int
    // (optional) Comma-separated list of work item field names.
    WitFields *string
}

// Get a list of test points.
func (client Client) GetPoints(ctx context.Context, args GetPointsArgs) (*[]TestPoint, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)

    queryParams := url.Values{}
    if args.WitFields != nil {
        queryParams.Add("witFields", *args.WitFields)
    }
    if args.ConfigurationId != nil {
        queryParams.Add("configurationId", *args.ConfigurationId)
    }
    if args.TestCaseId != nil {
        queryParams.Add("testCaseId", *args.TestCaseId)
    }
    if args.TestPointIds != nil {
        queryParams.Add("testPointIds", *args.TestPointIds)
    }
    if args.IncludePointDetails != nil {
        queryParams.Add("includePointDetails", strconv.FormatBool(*args.IncludePointDetails))
    }
    if args.Skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*args.Skip))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    locationId, _ := uuid.Parse("3bcfd5c8-be62-488e-b1da-b8289ce9299c")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPoints function
type GetPointsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan.
    PlanId *int
    // (required) ID of the suite that contains the points.
    SuiteId *int
    // (optional) Comma-separated list of work item field names.
    WitFields *string
    // (optional) Get test points for specific configuration.
    ConfigurationId *string
    // (optional) Get test points for a specific test case, valid when configurationId is not set.
    TestCaseId *string
    // (optional) Get test points for comma-separated list of test point IDs, valid only when configurationId and testCaseId are not set.
    TestPointIds *string
    // (optional) Include all properties for the test point.
    IncludePointDetails *bool
    // (optional) Number of test points to skip..
    Skip *int
    // (optional) Number of test points to return.
    Top *int
}

// Update test points.
func (client Client) UpdateTestPoints(ctx context.Context, args UpdateTestPointsArgs) (*[]TestPoint, error) {
    if args.PointUpdateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "pointUpdateModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.PointIds == nil || *args.PointIds == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "pointIds"} 
    }
    routeValues["pointIds"] = *args.PointIds

    body, marshalErr := json.Marshal(*args.PointUpdateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("3bcfd5c8-be62-488e-b1da-b8289ce9299c")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestPoint
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateTestPoints function
type UpdateTestPointsArgs struct {
    // (required) Data to update.
    PointUpdateModel *PointUpdateModel
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan.
    PlanId *int
    // (required) ID of the suite that contains the points.
    SuiteId *int
    // (required) ID of the test point to get. Use a comma-separated list of IDs to update multiple test points.
    PointIds *string
}

// [Preview API] Get test points using query.
func (client Client) GetPointsByQuery(ctx context.Context, args GetPointsByQueryArgs) (*TestPointsQuery, error) {
    if args.Query == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "query"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.Skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*args.Skip))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    body, marshalErr := json.Marshal(*args.Query)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("b4264fd0-a5d1-43e2-82a5-b9c46b7da9ce")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestPointsQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPointsByQuery function
type GetPointsByQueryArgs struct {
    // (required) TestPointsQuery to get test points.
    Query *TestPointsQuery
    // (required) Project ID or project name
    Project *string
    // (optional) Number of test points to skip..
    Skip *int
    // (optional) Number of test points to return.
    Top *int
}

// [Preview API] Get test result retention settings
func (client Client) GetResultRetentionSettings(ctx context.Context, args GetResultRetentionSettingsArgs) (*ResultRetentionSettings, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    locationId, _ := uuid.Parse("a3206d9e-fa8d-42d3-88cb-f75c51e69cde")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ResultRetentionSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetResultRetentionSettings function
type GetResultRetentionSettingsArgs struct {
    // (required) Project ID or project name
    Project *string
}

// [Preview API] Update test result retention settings
func (client Client) UpdateResultRetentionSettings(ctx context.Context, args UpdateResultRetentionSettingsArgs) (*ResultRetentionSettings, error) {
    if args.RetentionSettings == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "retentionSettings"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.RetentionSettings)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a3206d9e-fa8d-42d3-88cb-f75c51e69cde")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue ResultRetentionSettings
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateResultRetentionSettings function
type UpdateResultRetentionSettingsArgs struct {
    // (required) Test result retention settings details to be updated
    RetentionSettings *ResultRetentionSettings
    // (required) Project ID or project name
    Project *string
}

// Add test results to a test run.
func (client Client) AddTestResultsToTestRun(ctx context.Context, args AddTestResultsToTestRunArgs) (*[]TestCaseResult, error) {
    if args.Results == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "results"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    body, marshalErr := json.Marshal(*args.Results)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4637d869-3a76-4468-8057-0bb02aa385cf")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the AddTestResultsToTestRun function
type AddTestResultsToTestRunArgs struct {
    // (required) List of test results to add.
    Results *[]TestCaseResult
    // (required) Project ID or project name
    Project *string
    // (required) Test run ID into which test results to add.
    RunId *int
}

// Get a test result for a test run.
func (client Client) GetTestResultById(ctx context.Context, args GetTestResultByIdArgs) (*TestCaseResult, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)
    if args.TestCaseResultId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseResultId"} 
    }
    routeValues["testCaseResultId"] = strconv.Itoa(*args.TestCaseResultId)

    queryParams := url.Values{}
    if args.DetailsToInclude != nil {
        queryParams.Add("detailsToInclude", string(*args.DetailsToInclude))
    }
    locationId, _ := uuid.Parse("4637d869-3a76-4468-8057-0bb02aa385cf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestCaseResult
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResultById function
type GetTestResultByIdArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Test run ID of a test result to fetch.
    RunId *int
    // (required) Test result ID.
    TestCaseResultId *int
    // (optional) Details to include with test results. Default is None. Other values are Iterations, WorkItems and SubResults.
    DetailsToInclude *ResultDetails
}

// Get test results for a test run.
func (client Client) GetTestResults(ctx context.Context, args GetTestResultsArgs) (*[]TestCaseResult, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.DetailsToInclude != nil {
        queryParams.Add("detailsToInclude", string(*args.DetailsToInclude))
    }
    if args.Skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*args.Skip))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    if args.Outcomes != nil {
        var stringList []string
        for _, item := range *args.Outcomes {
            stringList = append(stringList, string(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    locationId, _ := uuid.Parse("4637d869-3a76-4468-8057-0bb02aa385cf")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestResults function
type GetTestResultsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Test run ID of test results to fetch.
    RunId *int
    // (optional) Details to include with test results. Default is None. Other values are Iterations and WorkItems.
    DetailsToInclude *ResultDetails
    // (optional) Number of test results to skip from beginning.
    Skip *int
    // (optional) Number of test results to return. Maximum is 1000 when detailsToInclude is None and 200 otherwise.
    Top *int
    // (optional) Comma separated list of test outcomes to filter test results.
    Outcomes *[]TestOutcome
}

// Update test results in a test run.
func (client Client) UpdateTestResults(ctx context.Context, args UpdateTestResultsArgs) (*[]TestCaseResult, error) {
    if args.Results == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "results"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    body, marshalErr := json.Marshal(*args.Results)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("4637d869-3a76-4468-8057-0bb02aa385cf")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestCaseResult
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateTestResults function
type UpdateTestResultsArgs struct {
    // (required) List of test results to update.
    Results *[]TestCaseResult
    // (required) Project ID or project name
    Project *string
    // (required) Test run ID whose test results to update.
    RunId *int
}

// Get test run statistics , used when we want to get summary of a run by outcome.
func (client Client) GetTestRunStatistics(ctx context.Context, args GetTestRunStatisticsArgs) (*TestRunStatistic, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    locationId, _ := uuid.Parse("0a42c424-d764-4a16-a2d5-5c85f87d0ae8")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRunStatistic
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunStatistics function
type GetTestRunStatisticsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the run to get.
    RunId *int
}

// Create new test run.
func (client Client) CreateTestRun(ctx context.Context, args CreateTestRunArgs) (*TestRun, error) {
    if args.TestRun == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testRun"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.TestRun)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestRun function
type CreateTestRunArgs struct {
    // (required) Run details RunCreateModel
    TestRun *RunCreateModel
    // (required) Project ID or project name
    Project *string
}

// Delete a test run by its ID.
func (client Client) DeleteTestRun(ctx context.Context, args DeleteTestRunArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteTestRun function
type DeleteTestRunArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the run to delete.
    RunId *int
}

// Get a test run by its ID.
func (client Client) GetTestRunById(ctx context.Context, args GetTestRunByIdArgs) (*TestRun, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    queryParams := url.Values{}
    if args.IncludeDetails != nil {
        queryParams.Add("includeDetails", strconv.FormatBool(*args.IncludeDetails))
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRunById function
type GetTestRunByIdArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the run to get.
    RunId *int
    // (optional) Defualt value is true. It includes details like run statistics,release,build,Test enviornment,Post process state and more
    IncludeDetails *bool
}

// Get a list of test runs.
func (client Client) GetTestRuns(ctx context.Context, args GetTestRunsArgs) (*[]TestRun, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.BuildUri != nil {
        queryParams.Add("buildUri", *args.BuildUri)
    }
    if args.Owner != nil {
        queryParams.Add("owner", *args.Owner)
    }
    if args.TmiRunId != nil {
        queryParams.Add("tmiRunId", *args.TmiRunId)
    }
    if args.PlanId != nil {
        queryParams.Add("planId", strconv.Itoa(*args.PlanId))
    }
    if args.IncludeRunDetails != nil {
        queryParams.Add("includeRunDetails", strconv.FormatBool(*args.IncludeRunDetails))
    }
    if args.Automated != nil {
        queryParams.Add("automated", strconv.FormatBool(*args.Automated))
    }
    if args.Skip != nil {
        queryParams.Add("$skip", strconv.Itoa(*args.Skip))
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestRuns function
type GetTestRunsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional) URI of the build that the runs used.
    BuildUri *string
    // (optional) Team foundation ID of the owner of the runs.
    Owner *string
    // (optional)
    TmiRunId *string
    // (optional) ID of the test plan that the runs are a part of.
    PlanId *int
    // (optional) If true, include all the properties of the runs.
    IncludeRunDetails *bool
    // (optional) If true, only returns automated runs.
    Automated *bool
    // (optional) Number of test runs to skip.
    Skip *int
    // (optional) Number of test runs to return.
    Top *int
}

// Query Test Runs based on filters. Mandatory fields are minLastUpdatedDate and maxLastUpdatedDate.
func (client Client) QueryTestRuns(ctx context.Context, args QueryTestRunsArgs) (*[]TestRun, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    queryParams := url.Values{}
    if args.MinLastUpdatedDate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "minLastUpdatedDate"}
    }
    queryParams.Add("minLastUpdatedDate", (*args.MinLastUpdatedDate).String())
    if args.MaxLastUpdatedDate == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "maxLastUpdatedDate"}
    }
    queryParams.Add("maxLastUpdatedDate", (*args.MaxLastUpdatedDate).String())
    if args.State != nil {
        queryParams.Add("state", string(*args.State))
    }
    if args.PlanIds != nil {
        var stringList []string
        for _, item := range *args.PlanIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.IsAutomated != nil {
        queryParams.Add("isAutomated", strconv.FormatBool(*args.IsAutomated))
    }
    if args.PublishContext != nil {
        queryParams.Add("publishContext", string(*args.PublishContext))
    }
    if args.BuildIds != nil {
        var stringList []string
        for _, item := range *args.BuildIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.BuildDefIds != nil {
        var stringList []string
        for _, item := range *args.BuildDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.BranchName != nil {
        queryParams.Add("branchName", *args.BranchName)
    }
    if args.ReleaseIds != nil {
        var stringList []string
        for _, item := range *args.ReleaseIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.ReleaseDefIds != nil {
        var stringList []string
        for _, item := range *args.ReleaseDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.ReleaseEnvIds != nil {
        var stringList []string
        for _, item := range *args.ReleaseEnvIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.ReleaseEnvDefIds != nil {
        var stringList []string
        for _, item := range *args.ReleaseEnvDefIds {
            stringList = append(stringList, strconv.Itoa(item))
        }
        listAsString := strings.Join((stringList)[:], ",")
        queryParams.Add("definitions", listAsString)
    }
    if args.RunTitle != nil {
        queryParams.Add("runTitle", *args.RunTitle)
    }
    if args.Top != nil {
        queryParams.Add("$top", strconv.Itoa(*args.Top))
    }
    if args.ContinuationToken != nil {
        queryParams.Add("continuationToken", *args.ContinuationToken)
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestRun
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestRuns function
type QueryTestRunsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Minimum Last Modified Date of run to be queried (Mandatory).
    MinLastUpdatedDate *time.Time
    // (required) Maximum Last Modified Date of run to be queried (Mandatory, difference between min and max date can be atmost 7 days).
    MaxLastUpdatedDate *time.Time
    // (optional) Current state of the Runs to be queried.
    State *TestRunState
    // (optional) Plan Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
    PlanIds *[]int
    // (optional) Automation type of the Runs to be queried.
    IsAutomated *bool
    // (optional) PublishContext of the Runs to be queried.
    PublishContext *TestRunPublishContext
    // (optional) Build Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
    BuildIds *[]int
    // (optional) Build Definition Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
    BuildDefIds *[]int
    // (optional) Source Branch name of the Runs to be queried.
    BranchName *string
    // (optional) Release Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
    ReleaseIds *[]int
    // (optional) Release Definition Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
    ReleaseDefIds *[]int
    // (optional) Release Environment Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
    ReleaseEnvIds *[]int
    // (optional) Release Environment Definition Ids of the Runs to be queried, comma seperated list of valid ids (limit no. of ids 10).
    ReleaseEnvDefIds *[]int
    // (optional) Run Title of the Runs to be queried.
    RunTitle *string
    // (optional) Number of runs to be queried. Limit is 100
    Top *int
    // (optional) continuationToken received from previous batch or null for first batch. It is not supposed to be created (or altered, if received from last batch) by user.
    ContinuationToken *string
}

// Update test run by its ID.
func (client Client) UpdateTestRun(ctx context.Context, args UpdateTestRunArgs) (*TestRun, error) {
    if args.RunUpdateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runUpdateModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.RunId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "runId"} 
    }
    routeValues["runId"] = strconv.Itoa(*args.RunId)

    body, marshalErr := json.Marshal(*args.RunUpdateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("cadb3810-d47d-4a3c-a234-fe5f3be50138")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestRun
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateTestRun function
type UpdateTestRunArgs struct {
    // (required) Run details RunUpdateModel
    RunUpdateModel *RunUpdateModel
    // (required) Project ID or project name
    Project *string
    // (required) ID of the run to update.
    RunId *int
}

// [Preview API] Create a test session
func (client Client) CreateTestSession(ctx context.Context, args CreateTestSessionArgs) (*TestSession, error) {
    if args.TestSession == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSession"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.Team != nil && *args.Team != "" {
        routeValues["team"] = *args.Team
    }

    body, marshalErr := json.Marshal(*args.TestSession)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1500b4b4-6c69-4ca6-9b18-35e9e97fe2ac")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSession
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the CreateTestSession function
type CreateTestSessionArgs struct {
    // (required) Test session details for creation
    TestSession *TestSession
    // (required) Project ID or project name
    Project *string
    // (optional) Team ID or team name
    Team *string
}

// [Preview API] Get a list of test sessions
func (client Client) GetTestSessions(ctx context.Context, args GetTestSessionsArgs) (*[]TestSession, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.Team != nil && *args.Team != "" {
        routeValues["team"] = *args.Team
    }

    queryParams := url.Values{}
    if args.Period != nil {
        queryParams.Add("period", strconv.Itoa(*args.Period))
    }
    if args.AllSessions != nil {
        queryParams.Add("allSessions", strconv.FormatBool(*args.AllSessions))
    }
    if args.IncludeAllProperties != nil {
        queryParams.Add("includeAllProperties", strconv.FormatBool(*args.IncludeAllProperties))
    }
    if args.Source != nil {
        queryParams.Add("source", string(*args.Source))
    }
    if args.IncludeOnlyCompletedSessions != nil {
        queryParams.Add("includeOnlyCompletedSessions", strconv.FormatBool(*args.IncludeOnlyCompletedSessions))
    }
    locationId, _ := uuid.Parse("1500b4b4-6c69-4ca6-9b18-35e9e97fe2ac")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []TestSession
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestSessions function
type GetTestSessionsArgs struct {
    // (required) Project ID or project name
    Project *string
    // (optional) Team ID or team name
    Team *string
    // (optional) Period in days from now, for which test sessions are fetched.
    Period *int
    // (optional) If false, returns test sessions for current user. Otherwise, it returns test sessions for all users
    AllSessions *bool
    // (optional) If true, it returns all properties of the test sessions. Otherwise, it returns the skinny version.
    IncludeAllProperties *bool
    // (optional) Source of the test session.
    Source *TestSessionSource
    // (optional) If true, it returns test sessions in completed state. Otherwise, it returns test sessions for all states
    IncludeOnlyCompletedSessions *bool
}

// [Preview API] Update a test session
func (client Client) UpdateTestSession(ctx context.Context, args UpdateTestSessionArgs) (*TestSession, error) {
    if args.TestSession == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testSession"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.Team != nil && *args.Team != "" {
        routeValues["team"] = *args.Team
    }

    body, marshalErr := json.Marshal(*args.TestSession)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("1500b4b4-6c69-4ca6-9b18-35e9e97fe2ac")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestSession
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateTestSession function
type UpdateTestSessionArgs struct {
    // (required) Test session details for update
    TestSession *TestSession
    // (required) Project ID or project name
    Project *string
    // (optional) Team ID or team name
    Team *string
}

// Add test cases to suite.
func (client Client) AddTestCasesToSuite(ctx context.Context, args AddTestCasesToSuiteArgs) (*[]SuiteTestCase, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.TestCaseIds == nil || *args.TestCaseIds == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = *args.TestCaseIds
    routeValues["action"] = "TestCases"

    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SuiteTestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the AddTestCasesToSuite function
type AddTestCasesToSuiteArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan that contains the suite.
    PlanId *int
    // (required) ID of the test suite to which the test cases must be added.
    SuiteId *int
    // (required) IDs of the test cases to add to the suite. Ids are specified in comma separated format.
    TestCaseIds *string
}

// Get a specific test case in a test suite with test case id.
func (client Client) GetTestCaseById(ctx context.Context, args GetTestCaseByIdArgs) (*SuiteTestCase, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.TestCaseIds == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = strconv.Itoa(*args.TestCaseIds)
    routeValues["action"] = "TestCases"

    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue SuiteTestCase
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestCaseById function
type GetTestCaseByIdArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan that contains the suites.
    PlanId *int
    // (required) ID of the suite that contains the test case.
    SuiteId *int
    // (required) ID of the test case to get.
    TestCaseIds *int
}

// Get all test cases in a suite.
func (client Client) GetTestCases(ctx context.Context, args GetTestCasesArgs) (*[]SuiteTestCase, error) {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    routeValues["action"] = "TestCases"

    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SuiteTestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetTestCases function
type GetTestCasesArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan that contains the suites.
    PlanId *int
    // (required) ID of the suite to get.
    SuiteId *int
}

// The test points associated with the test cases are removed from the test suite. The test case work item is not deleted from the system. See test cases resource to delete a test case permanently.
func (client Client) RemoveTestCasesFromSuiteUrl(ctx context.Context, args RemoveTestCasesFromSuiteUrlArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.TestCaseIds == nil || *args.TestCaseIds == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = *args.TestCaseIds
    routeValues["action"] = "TestCases"

    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the RemoveTestCasesFromSuiteUrl function
type RemoveTestCasesFromSuiteUrlArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan that contains the suite.
    PlanId *int
    // (required) ID of the suite to get.
    SuiteId *int
    // (required) IDs of the test cases to remove from the suite.
    TestCaseIds *string
}

// Updates the properties of the test case association in a suite.
func (client Client) UpdateSuiteTestCases(ctx context.Context, args UpdateSuiteTestCasesArgs) (*[]SuiteTestCase, error) {
    if args.SuiteTestCaseUpdateModel == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteTestCaseUpdateModel"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.PlanId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "planId"} 
    }
    routeValues["planId"] = strconv.Itoa(*args.PlanId)
    if args.SuiteId == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "suiteId"} 
    }
    routeValues["suiteId"] = strconv.Itoa(*args.SuiteId)
    if args.TestCaseIds == nil || *args.TestCaseIds == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "testCaseIds"} 
    }
    routeValues["testCaseIds"] = *args.TestCaseIds
    routeValues["action"] = "TestCases"

    body, marshalErr := json.Marshal(*args.SuiteTestCaseUpdateModel)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("a4a1ec1c-b03f-41ca-8857-704594ecf58e")
    resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue []SuiteTestCase
    err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the UpdateSuiteTestCases function
type UpdateSuiteTestCasesArgs struct {
    // (required) Model for updation of the properties of test case suite association.
    SuiteTestCaseUpdateModel *SuiteTestCaseUpdateModel
    // (required) Project ID or project name
    Project *string
    // (required) ID of the test plan that contains the suite.
    PlanId *int
    // (required) ID of the test suite to which the test cases must be added.
    SuiteId *int
    // (required) IDs of the test cases to add to the suite. Ids are specified in comma separated format.
    TestCaseIds *string
}

// [Preview API] Delete a test case.
func (client Client) DeleteTestCase(ctx context.Context, args DeleteTestCaseArgs) error {
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project
    if args.TestCaseId == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "testCaseId"} 
    }
    routeValues["testCaseId"] = strconv.Itoa(*args.TestCaseId)

    locationId, _ := uuid.Parse("4d472e0f-e32c-4ef8-adf4-a4078772889c")
    _, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the DeleteTestCase function
type DeleteTestCaseArgs struct {
    // (required) Project ID or project name
    Project *string
    // (required) Id of test case to delete.
    TestCaseId *int
}

// [Preview API] Get history of a test method using TestHistoryQuery
func (client Client) QueryTestHistory(ctx context.Context, args QueryTestHistoryArgs) (*TestHistoryQuery, error) {
    if args.Filter == nil {
        return nil, &azureDevops.ArgumentNilError{ArgumentName: "filter"}
    }
    routeValues := make(map[string]string)
    if args.Project == nil || *args.Project == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "project"} 
    }
    routeValues["project"] = *args.Project

    body, marshalErr := json.Marshal(*args.Filter)
    if marshalErr != nil {
        return nil, marshalErr
    }
    locationId, _ := uuid.Parse("929fd86c-3e38-4d8c-b4b6-90df256e5971")
    resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.2", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue TestHistoryQuery
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the QueryTestHistory function
type QueryTestHistoryArgs struct {
    // (required) TestHistoryQuery to get history
    Filter *TestHistoryQuery
    // (required) Project ID or project name
    Project *string
}

