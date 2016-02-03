package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

// NewCreateDialParams creates a new CreateDialParams object
// with the default values initialized.
func NewCreateDialParams() *CreateDialParams {
	var ()
	return &CreateDialParams{}
}

/*CreateDialParams contains all the parameters to send to the API endpoint
for the create dial operation typically these are written to a http.Request
*/
type CreateDialParams struct {

	/*Dial
	  New dial

	*/
	Dial *models.Dial
}

// WithDial adds the dial to the create dial params
func (o *CreateDialParams) WithDial(dial *models.Dial) *CreateDialParams {
	o.Dial = dial
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateDialParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Dial == nil {
		o.Dial = new(models.Dial)
	}

	if err := r.SetBodyParam(o.Dial); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
