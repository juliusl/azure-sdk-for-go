package policyinsightsapi

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
	"github.com/Azure/azure-sdk-for-go/services/policyinsights/mgmt/2019-10-01/policyinsights"
	"github.com/Azure/go-autorest/autorest/date"
)

// PolicyTrackedResourcesClientAPI contains the set of methods on the PolicyTrackedResourcesClient type.
type PolicyTrackedResourcesClientAPI interface {
	ListQueryResultsForManagementGroup(ctx context.Context, managementGroupName string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsPage, err error)
	ListQueryResultsForResource(ctx context.Context, resourceID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsPage, err error)
	ListQueryResultsForResourceGroup(ctx context.Context, resourceGroupName string, subscriptionID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsPage, err error)
	ListQueryResultsForSubscription(ctx context.Context, subscriptionID string, top *int32, filter string) (result policyinsights.PolicyTrackedResourcesQueryResultsPage, err error)
}

var _ PolicyTrackedResourcesClientAPI = (*policyinsights.PolicyTrackedResourcesClient)(nil)

// RemediationsClientAPI contains the set of methods on the RemediationsClient type.
type RemediationsClientAPI interface {
	CancelAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string) (result policyinsights.Remediation, err error)
	CancelAtResource(ctx context.Context, resourceID string, remediationName string) (result policyinsights.Remediation, err error)
	CancelAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string) (result policyinsights.Remediation, err error)
	CancelAtSubscription(ctx context.Context, subscriptionID string, remediationName string) (result policyinsights.Remediation, err error)
	CreateOrUpdateAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string, parameters policyinsights.Remediation) (result policyinsights.Remediation, err error)
	CreateOrUpdateAtResource(ctx context.Context, resourceID string, remediationName string, parameters policyinsights.Remediation) (result policyinsights.Remediation, err error)
	CreateOrUpdateAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string, parameters policyinsights.Remediation) (result policyinsights.Remediation, err error)
	CreateOrUpdateAtSubscription(ctx context.Context, subscriptionID string, remediationName string, parameters policyinsights.Remediation) (result policyinsights.Remediation, err error)
	DeleteAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string) (result policyinsights.Remediation, err error)
	DeleteAtResource(ctx context.Context, resourceID string, remediationName string) (result policyinsights.Remediation, err error)
	DeleteAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string) (result policyinsights.Remediation, err error)
	DeleteAtSubscription(ctx context.Context, subscriptionID string, remediationName string) (result policyinsights.Remediation, err error)
	GetAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string) (result policyinsights.Remediation, err error)
	GetAtResource(ctx context.Context, resourceID string, remediationName string) (result policyinsights.Remediation, err error)
	GetAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string) (result policyinsights.Remediation, err error)
	GetAtSubscription(ctx context.Context, subscriptionID string, remediationName string) (result policyinsights.Remediation, err error)
	ListDeploymentsAtManagementGroup(ctx context.Context, managementGroupID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultPage, err error)
	ListDeploymentsAtResource(ctx context.Context, resourceID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultPage, err error)
	ListDeploymentsAtResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultPage, err error)
	ListDeploymentsAtSubscription(ctx context.Context, subscriptionID string, remediationName string, top *int32) (result policyinsights.RemediationDeploymentsListResultPage, err error)
	ListForManagementGroup(ctx context.Context, managementGroupID string, top *int32, filter string) (result policyinsights.RemediationListResultPage, err error)
	ListForResource(ctx context.Context, resourceID string, top *int32, filter string) (result policyinsights.RemediationListResultPage, err error)
	ListForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, filter string) (result policyinsights.RemediationListResultPage, err error)
	ListForSubscription(ctx context.Context, subscriptionID string, top *int32, filter string) (result policyinsights.RemediationListResultPage, err error)
}

var _ RemediationsClientAPI = (*policyinsights.RemediationsClient)(nil)

// PolicyEventsClientAPI contains the set of methods on the PolicyEventsClient type.
type PolicyEventsClientAPI interface {
	GetMetadata(ctx context.Context, scope string) (result policyinsights.String, err error)
	ListQueryResultsForManagementGroup(ctx context.Context, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForPolicyDefinition(ctx context.Context, subscriptionID string, policyDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForPolicySetDefinition(ctx context.Context, subscriptionID string, policySetDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForResource(ctx context.Context, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx context.Context, subscriptionID string, resourceGroupName string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForSubscription(ctx context.Context, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
	ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx context.Context, subscriptionID string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyEventsQueryResults, err error)
}

var _ PolicyEventsClientAPI = (*policyinsights.PolicyEventsClient)(nil)

// PolicyStatesClientAPI contains the set of methods on the PolicyStatesClient type.
type PolicyStatesClientAPI interface {
	ListQueryResultsForManagementGroup(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, managementGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForPolicyDefinition(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policyDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForPolicySetDefinition(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policySetDefinitionName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForResource(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, resourceID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string, expand string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForResourceGroup(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, resourceGroupName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForResourceGroupLevelPolicyAssignment(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, resourceGroupName string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForSubscription(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	ListQueryResultsForSubscriptionLevelPolicyAssignment(ctx context.Context, policyStatesResource policyinsights.PolicyStatesResource, subscriptionID string, policyAssignmentName string, top *int32, orderBy string, selectParameter string, from *date.Time, toParameter *date.Time, filter string, apply string) (result policyinsights.PolicyStatesQueryResults, err error)
	SummarizeForManagementGroup(ctx context.Context, managementGroupName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForPolicyDefinition(ctx context.Context, subscriptionID string, policyDefinitionName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForPolicySetDefinition(ctx context.Context, subscriptionID string, policySetDefinitionName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForResource(ctx context.Context, resourceID string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForResourceGroup(ctx context.Context, subscriptionID string, resourceGroupName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForResourceGroupLevelPolicyAssignment(ctx context.Context, subscriptionID string, resourceGroupName string, policyAssignmentName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForSubscription(ctx context.Context, subscriptionID string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
	SummarizeForSubscriptionLevelPolicyAssignment(ctx context.Context, subscriptionID string, policyAssignmentName string, top *int32, from *date.Time, toParameter *date.Time, filter string) (result policyinsights.SummarizeResults, err error)
}

var _ PolicyStatesClientAPI = (*policyinsights.PolicyStatesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result policyinsights.OperationsListResults, err error)
}

var _ OperationsClientAPI = (*policyinsights.OperationsClient)(nil)

// PolicyMetadataClientAPI contains the set of methods on the PolicyMetadataClient type.
type PolicyMetadataClientAPI interface {
	GetResource(ctx context.Context, resourceName string) (result policyinsights.PolicyMetadata, err error)
	List(ctx context.Context, top *int32) (result policyinsights.PolicyMetadataCollectionPage, err error)
}

var _ PolicyMetadataClientAPI = (*policyinsights.PolicyMetadataClient)(nil)
