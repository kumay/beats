// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aws

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aws", asset.ModuleFieldsPri, AssetAws); err != nil {
		panic(err)
	}
}

// AssetAws returns asset data.
// This is the base64 encoded gzipped contents of module/aws.
func AssetAws() string {
	return "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"
}
