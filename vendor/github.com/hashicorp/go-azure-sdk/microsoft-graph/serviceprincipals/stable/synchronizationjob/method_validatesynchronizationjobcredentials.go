package synchronizationjob

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateSynchronizationJobCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ValidateSynchronizationJobCredentialsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultValidateSynchronizationJobCredentialsOperationOptions() ValidateSynchronizationJobCredentialsOperationOptions {
	return ValidateSynchronizationJobCredentialsOperationOptions{}
}

func (o ValidateSynchronizationJobCredentialsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ValidateSynchronizationJobCredentialsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ValidateSynchronizationJobCredentialsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ValidateSynchronizationJobCredentials - Invoke action validateCredentials. Validate that the credentials are valid in
// the tenant.
func (c SynchronizationJobClient) ValidateSynchronizationJobCredentials(ctx context.Context, id stable.ServicePrincipalIdSynchronizationJobId, input ValidateSynchronizationJobCredentialsRequest, options ValidateSynchronizationJobCredentialsOperationOptions) (result ValidateSynchronizationJobCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/validateCredentials", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}