package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetDialParams creates a new GetDialParams object
// with the default values initialized.
func NewGetDialParams() *GetDialParams {
	var ()
	return &GetDialParams{}
}

/*GetDialParams contains all the parameters to send to the API endpoint
for the get dial operation typically these are written to a http.Request
*/
type GetDialParams struct {

	/*DialID
	  ID of the dial

	*/
	DialID string
}

// WithDialID adds the dialId to the get dial params
func (o *GetDialParams) WithDialID(DialID string) *GetDialParams {
	o.DialID = DialID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetDialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param dial_id
	if err := r.SetPathParam("dial_id", o.DialID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
