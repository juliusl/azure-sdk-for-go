package personalizerapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/personalizer/v1.0/personalizer"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	Rank(ctx context.Context, rankRequest personalizer.RankRequest) (result personalizer.RankResponse, err error)
}

var _ BaseClientAPI = (*personalizer.BaseClient)(nil)

// ServiceConfigurationClientAPI contains the set of methods on the ServiceConfigurationClient type.
type ServiceConfigurationClientAPI interface {
	Get(ctx context.Context) (result personalizer.ServiceConfiguration, err error)
	Update(ctx context.Context, config personalizer.ServiceConfiguration) (result personalizer.ServiceConfiguration, err error)
}

var _ ServiceConfigurationClientAPI = (*personalizer.ServiceConfigurationClient)(nil)

// PolicyClientAPI contains the set of methods on the PolicyClient type.
type PolicyClientAPI interface {
	Get(ctx context.Context) (result personalizer.PolicyContract, err error)
	Reset(ctx context.Context) (result personalizer.PolicyContract, err error)
	Update(ctx context.Context, policy personalizer.PolicyContract) (result personalizer.PolicyContract, err error)
}

var _ PolicyClientAPI = (*personalizer.PolicyClient)(nil)

// EvaluationsClientAPI contains the set of methods on the EvaluationsClient type.
type EvaluationsClientAPI interface {
	Create(ctx context.Context, evaluation personalizer.EvaluationContract) (result personalizer.Evaluation, err error)
	Delete(ctx context.Context, evaluationID string) (result autorest.Response, err error)
	Get(ctx context.Context, evaluationID string) (result personalizer.Evaluation, err error)
	List(ctx context.Context) (result personalizer.ListEvaluation, err error)
}

var _ EvaluationsClientAPI = (*personalizer.EvaluationsClient)(nil)

// EventsClientAPI contains the set of methods on the EventsClient type.
type EventsClientAPI interface {
	Activate(ctx context.Context, eventID string) (result autorest.Response, err error)
	Reward(ctx context.Context, eventID string, reward personalizer.RewardRequest) (result autorest.Response, err error)
}

var _ EventsClientAPI = (*personalizer.EventsClient)(nil)

// LogClientAPI contains the set of methods on the LogClient type.
type LogClientAPI interface {
	Delete(ctx context.Context) (result autorest.Response, err error)
	GetProperties(ctx context.Context) (result personalizer.LogsProperties, err error)
}

var _ LogClientAPI = (*personalizer.LogClient)(nil)

// ModelClientAPI contains the set of methods on the ModelClient type.
type ModelClientAPI interface {
	Get(ctx context.Context) (result personalizer.ReadCloser, err error)
	GetProperties(ctx context.Context) (result personalizer.ModelProperties, err error)
	Reset(ctx context.Context) (result autorest.Response, err error)
}

var _ ModelClientAPI = (*personalizer.ModelClient)(nil)
