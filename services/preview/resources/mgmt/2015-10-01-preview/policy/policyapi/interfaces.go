package policyapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2015-10-01-preview/policy"
	"github.com/Azure/go-autorest/autorest"
)

// AssignmentsClientAPI contains the set of methods on the AssignmentsClient type.
type AssignmentsClientAPI interface {
	Create(ctx context.Context, scope string, policyAssignmentName string, parameters policy.Assignment) (result policy.Assignment, err error)
	CreateByID(ctx context.Context, policyAssignmentID string, parameters policy.Assignment) (result policy.Assignment, err error)
	Delete(ctx context.Context, scope string, policyAssignmentName string) (result policy.Assignment, err error)
	DeleteByID(ctx context.Context, policyAssignmentID string) (result policy.Assignment, err error)
	Get(ctx context.Context, scope string, policyAssignmentName string) (result policy.Assignment, err error)
	GetByID(ctx context.Context, policyAssignmentID string) (result policy.Assignment, err error)
	List(ctx context.Context, filter string) (result policy.AssignmentListResultPage, err error)
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result policy.AssignmentListResultPage, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result policy.AssignmentListResultPage, err error)
}

var _ AssignmentsClientAPI = (*policy.AssignmentsClient)(nil)

// DefinitionsClientAPI contains the set of methods on the DefinitionsClient type.
type DefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters policy.Definition) (result policy.Definition, err error)
	Delete(ctx context.Context, policyDefinitionName string) (result autorest.Response, err error)
	Get(ctx context.Context, policyDefinitionName string) (result policy.Definition, err error)
	List(ctx context.Context, filter string) (result policy.DefinitionListResultPage, err error)
}

var _ DefinitionsClientAPI = (*policy.DefinitionsClient)(nil)
