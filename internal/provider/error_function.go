// Copyright (c) The OpenTofu Authors
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = ErrorFunction{}
)

func NewErrorFunction() function.Function {
	return ErrorFunction{}
}

type ErrorFunction struct{}

func (e ErrorFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "error"
}

func (e ErrorFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Error function",
		MarkdownDescription: "Causes an error",
		Parameters:          []function.Parameter{},
		Return:              function.DynamicReturn{},
	}
}

func (e ErrorFunction) Run(_ context.Context, _ function.RunRequest, resp *function.RunResponse) {
	resp.Error = function.NewFuncError("error function threw an error")
}
