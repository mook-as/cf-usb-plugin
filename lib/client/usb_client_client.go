package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
)

// Default usb client HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new usb client HTTP client.
func NewHTTPClient(formats strfmt.Registry) *UsbClient {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("localhost", "/", []string{"http"})
	return New(transport, formats)
}

// New creates a new usb client client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *UsbClient {
	cli := new(UsbClient)
	cli.Transport = transport

	cli.Operations = operations.New(transport, formats)

	return cli
}

// UsbClient is a client for usb client
type UsbClient struct {
	Operations *operations.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *UsbClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Operations.SetTransport(transport)

}
