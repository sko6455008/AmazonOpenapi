// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package gen

// Amazon defines model for Amazon.
type Amazon struct {
	Asin   string `json:"asin"`
	Maker  string `json:"maker,omitempty"`
	Name   string `json:"name,omitempty"`
	Price  int64  `json:"price"`
	Reason string `json:"reason"`
	Url    string `json:"url"`
}

// AmazonPatch defines model for AmazonPatch.
type AmazonPatch struct {
	Maker  *string `json:"maker,omitempty"`
	Name   *string `json:"name,omitempty"`
	Price  *int64  `json:"price,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Url    *string `json:"url,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// AddAmazonJSONBody defines parameters for AddAmazon.
type AddAmazonJSONBody Amazon

// PatchAmazonJSONBody defines parameters for PatchAmazon.
type PatchAmazonJSONBody AmazonPatch

// PutAmazonJSONBody defines parameters for PutAmazon.
type PutAmazonJSONBody Amazon

// AddAmazonJSONRequestBody defines body for AddAmazon for application/json ContentType.
type AddAmazonJSONRequestBody AddAmazonJSONBody

// PatchAmazonJSONRequestBody defines body for PatchAmazon for application/json ContentType.
type PatchAmazonJSONRequestBody PatchAmazonJSONBody

// PutAmazonJSONRequestBody defines body for PutAmazon for application/json ContentType.
type PutAmazonJSONRequestBody PutAmazonJSONBody
