package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Info info

swagger:model info
*/
type Info struct {

	/* broker api version

	Required: true
	*/
	BrokerAPIVersion *string `json:"broker_api_version"`

	/* usb version

	Required: true
	*/
	UsbVersion *string `json:"usb_version"`
}

// Validate validates this info
func (m *Info) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrokerAPIVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsbVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Info) validateBrokerAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("broker_api_version", "body", m.BrokerAPIVersion); err != nil {
		return err
	}

	return nil
}

func (m *Info) validateUsbVersion(formats strfmt.Registry) error {

	if err := validate.Required("usb_version", "body", m.UsbVersion); err != nil {
		return err
	}

	return nil
}
