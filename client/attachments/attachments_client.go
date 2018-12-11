// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new attachments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for attachments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAttachmentsIDLogs get attachments ID logs API
*/
func (a *Client) GetAttachmentsIDLogs(params *GetAttachmentsIDLogsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAttachmentsIDLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAttachmentsIDLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAttachmentsIDLogs",
		Method:             "GET",
		PathPattern:        "/attachments/{id}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAttachmentsIDLogsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAttachmentsIDLogsOK), nil

}

/*
PostAttachments post attachments API
*/
func (a *Client) PostAttachments(params *PostAttachmentsParams, authInfo runtime.ClientAuthInfoWriter) (*PostAttachmentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAttachmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAttachments",
		Method:             "POST",
		PathPattern:        "/attachments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAttachmentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAttachmentsCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
