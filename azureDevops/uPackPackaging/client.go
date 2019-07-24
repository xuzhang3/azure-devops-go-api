﻿// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package uPackPackaging

import (
    "bytes"
    "context"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/microsoft/azure-devops-go-api/azureDevops"
    "net/http"
    "net/url"
)

var ResourceAreaId, _ = uuid.Parse("d397749b-f115-4027-b6dd-77a65dd10d21")

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

// [Preview API]
func (client Client) AddPackage(ctx context.Context, args AddPackageArgs) error {
    if args.Metadata == nil {
        return &azureDevops.ArgumentNilError{ArgumentName: "metadata"}
    }
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *args.PackageVersion

    body, marshalErr := json.Marshal(*args.Metadata)
    if marshalErr != nil {
        return marshalErr
    }
    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    _, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
    if err != nil {
        return err
    }

    return nil
}

// Arguments for the AddPackage function
type AddPackageArgs struct {
    // (required)
    Metadata *UPackPackagePushMetadata
    // (required)
    FeedId *string
    // (required)
    PackageName *string
    // (required)
    PackageVersion *string
}

// [Preview API]
func (client Client) GetPackageMetadata(ctx context.Context, args GetPackageMetadataArgs) (*UPackPackageMetadata, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName
    if args.PackageVersion == nil || *args.PackageVersion == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageVersion"} 
    }
    routeValues["packageVersion"] = *args.PackageVersion

    queryParams := url.Values{}
    if args.Intent != nil {
        queryParams.Add("intent", string(*args.Intent))
    }
    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackPackageMetadata
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPackageMetadata function
type GetPackageMetadataArgs struct {
    // (required)
    FeedId *string
    // (required)
    PackageName *string
    // (required)
    PackageVersion *string
    // (optional)
    Intent *UPackGetPackageMetadataIntent
}

// [Preview API]
func (client Client) GetPackageVersionsMetadata(ctx context.Context, args GetPackageVersionsMetadataArgs) (*UPackLimitedPackageMetadataListResponse, error) {
    routeValues := make(map[string]string)
    if args.FeedId == nil || *args.FeedId == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "feedId"} 
    }
    routeValues["feedId"] = *args.FeedId
    if args.PackageName == nil || *args.PackageName == "" {
        return nil, &azureDevops.ArgumentNilOrEmptyError{ArgumentName: "packageName"} 
    }
    routeValues["packageName"] = *args.PackageName

    locationId, _ := uuid.Parse("4cdb2ced-0758-4651-8032-010f070dd7e5")
    resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
    if err != nil {
        return nil, err
    }

    var responseValue UPackLimitedPackageMetadataListResponse
    err = client.Client.UnmarshalBody(resp, &responseValue)
    return &responseValue, err
}

// Arguments for the GetPackageVersionsMetadata function
type GetPackageVersionsMetadataArgs struct {
    // (required)
    FeedId *string
    // (required)
    PackageName *string
}

