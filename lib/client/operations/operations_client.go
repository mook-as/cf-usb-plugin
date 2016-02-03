package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new operations API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
Create a dial for
*/
func (a *Client) CreateDial(params *CreateDialParams, authInfo client.AuthInfoWriter) (*CreateDialCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDialParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "createDial",
		Method:             "POST",
		PathPattern:        "/dials",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDialReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDialCreated), nil
}

/*
Create a new driver

*/
func (a *Client) CreateDriver(params *CreateDriverParams, authInfo client.AuthInfoWriter) (*CreateDriverCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDriverParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "createDriver",
		Method:             "POST",
		PathPattern:        "/drivers",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDriverReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDriverCreated), nil
}

/*
Create a driver instance
*/
func (a *Client) CreateDriverInstance(params *CreateDriverInstanceParams, authInfo client.AuthInfoWriter) (*CreateDriverInstanceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDriverInstanceParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "createDriverInstance",
		Method:             "POST",
		PathPattern:        "/driver_instances",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateDriverInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDriverInstanceCreated), nil
}

/*
Create a plan
*/
func (a *Client) CreateServicePlan(params *CreateServicePlanParams, authInfo client.AuthInfoWriter) (*CreateServicePlanCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateServicePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "createServicePlan",
		Method:             "POST",
		PathPattern:        "/plans",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateServicePlanReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateServicePlanCreated), nil
}

/*
Delets the `dial` with the **dial_id**
*/
func (a *Client) DeleteDial(params *DeleteDialParams, authInfo client.AuthInfoWriter) (*DeleteDialNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDialParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteDial",
		Method:             "DELETE",
		PathPattern:        "/dials/{dial_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDialReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDialNoContent), nil
}

/*
Update driver
*/
func (a *Client) DeleteDriver(params *DeleteDriverParams, authInfo client.AuthInfoWriter) (*DeleteDriverNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDriverParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteDriver",
		Method:             "DELETE",
		PathPattern:        "/drivers/{driver_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDriverReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDriverNoContent), nil
}

/*
Delete a driver instance
*/
func (a *Client) DeleteDriverInstance(params *DeleteDriverInstanceParams, authInfo client.AuthInfoWriter) (*DeleteDriverInstanceNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDriverInstanceParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteDriverInstance",
		Method:             "DELETE",
		PathPattern:        "/driver_instances/{driver_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteDriverInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDriverInstanceNoContent), nil
}

/*
Delets the `plan` with the **planID** for the **serviceID**
*/
func (a *Client) DeleteServicePlan(params *DeleteServicePlanParams, authInfo client.AuthInfoWriter) (*DeleteServicePlanNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteServicePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteServicePlan",
		Method:             "DELETE",
		PathPattern:        "/plans/{plan_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteServicePlanReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteServicePlanNoContent), nil
}

/*
Gets `dials`
*/
func (a *Client) GetAllDials(params *GetAllDialsParams, authInfo client.AuthInfoWriter) (*GetAllDialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllDialsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getAllDials",
		Method:             "GET",
		PathPattern:        "/dials",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllDialsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllDialsOK), nil
}

/*
Gets the `dial` with the **dial_id**
*/
func (a *Client) GetDial(params *GetDialParams, authInfo client.AuthInfoWriter) (*GetDialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDialParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getDial",
		Method:             "GET",
		PathPattern:        "/dials/{dial_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDialReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDialOK), nil
}

/*
Get dial schema
*/
func (a *Client) GetDialSchema(params *GetDialSchemaParams, authInfo client.AuthInfoWriter) (*GetDialSchemaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDialSchemaParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getDialSchema",
		Method:             "GET",
		PathPattern:        "/drivers/{driver_id}/dial_schema",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDialSchemaReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDialSchemaOK), nil
}

/*
Get driver
*/
func (a *Client) GetDriver(params *GetDriverParams, authInfo client.AuthInfoWriter) (*GetDriverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDriverParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getDriver",
		Method:             "GET",
		PathPattern:        "/drivers/{driver_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDriverReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDriverOK), nil
}

/*
Gets driver configurations

*/
func (a *Client) GetDriverInstance(params *GetDriverInstanceParams, authInfo client.AuthInfoWriter) (*GetDriverInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDriverInstanceParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getDriverInstance",
		Method:             "GET",
		PathPattern:        "/driver_instances/{driver_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDriverInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDriverInstanceOK), nil
}

/*
Gets driver instances for a driver

*/
func (a *Client) GetDriverInstances(params *GetDriverInstancesParams, authInfo client.AuthInfoWriter) (*GetDriverInstancesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDriverInstancesParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getDriverInstances",
		Method:             "GET",
		PathPattern:        "/driver_instances",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDriverInstancesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDriverInstancesOK), nil
}

/*
Get driver config schema
*/
func (a *Client) GetDriverSchema(params *GetDriverSchemaParams, authInfo client.AuthInfoWriter) (*GetDriverSchemaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDriverSchemaParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getDriverSchema",
		Method:             "GET",
		PathPattern:        "/drivers/{driver_id}/config_schema",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDriverSchemaReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDriverSchemaOK), nil
}

/*
Gets information about the available `drivers`

*/
func (a *Client) GetDrivers(params *GetDriversParams, authInfo client.AuthInfoWriter) (*GetDriversOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDriversParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getDrivers",
		Method:             "GET",
		PathPattern:        "/drivers",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDriversReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDriversOK), nil
}

/*
Gets information about the USB.

*/
func (a *Client) GetInfo(params *GetInfoParams, authInfo client.AuthInfoWriter) (*GetInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInfoParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getInfo",
		Method:             "GET",
		PathPattern:        "/info",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetInfoOK), nil
}

/*
Gets the `service` with the id **serviceID**
*/
func (a *Client) GetService(params *GetServiceParams, authInfo client.AuthInfoWriter) (*GetServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getService",
		Method:             "GET",
		PathPattern:        "/services/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceOK), nil
}

/*
Gets the existing `service`

*/
func (a *Client) GetServiceByInstanceID(params *GetServiceByInstanceIDParams, authInfo client.AuthInfoWriter) (*GetServiceByInstanceIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceByInstanceIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getServiceByInstanceId",
		Method:             "GET",
		PathPattern:        "/services",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServiceByInstanceIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServiceByInstanceIDOK), nil
}

/*
Gets the `plan` with the **planID**
*/
func (a *Client) GetServicePlan(params *GetServicePlanParams, authInfo client.AuthInfoWriter) (*GetServicePlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServicePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getServicePlan",
		Method:             "GET",
		PathPattern:        "/plans/{plan_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServicePlanReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServicePlanOK), nil
}

/*
Gets `plans`
*/
func (a *Client) GetServicePlans(params *GetServicePlansParams, authInfo client.AuthInfoWriter) (*GetServicePlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServicePlansParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getServicePlans",
		Method:             "GET",
		PathPattern:        "/plans",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServicePlansReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetServicePlansOK), nil
}

/*
Pings the driver

*/
func (a *Client) PingDriverInstance(params *PingDriverInstanceParams, authInfo client.AuthInfoWriter) (*PingDriverInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPingDriverInstanceParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "pingDriverInstance",
		Method:             "GET",
		PathPattern:        "/driver_instances/{driver_instance_id}/ping",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PingDriverInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PingDriverInstanceOK), nil
}

/*
Updates the broker catalog

*/
func (a *Client) UpdateCatalog(params *UpdateCatalogParams, authInfo client.AuthInfoWriter) (*UpdateCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCatalogParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "updateCatalog",
		Method:             "POST",
		PathPattern:        "/update_catalog",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCatalogReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateCatalogOK), nil
}

/*
Updates the dial with the id **dial_id**
*/
func (a *Client) UpdateDial(params *UpdateDialParams, authInfo client.AuthInfoWriter) (*UpdateDialOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDialParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "updateDial",
		Method:             "PUT",
		PathPattern:        "/dials/{dial_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDialReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDialOK), nil
}

/*
Update driver
*/
func (a *Client) UpdateDriver(params *UpdateDriverParams, authInfo client.AuthInfoWriter) (*UpdateDriverOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDriverParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "updateDriver",
		Method:             "PUT",
		PathPattern:        "/drivers/{driver_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDriverReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDriverOK), nil
}

/*
Update a driver instance

*/
func (a *Client) UpdateDriverInstance(params *UpdateDriverInstanceParams, authInfo client.AuthInfoWriter) (*UpdateDriverInstanceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDriverInstanceParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "updateDriverInstance",
		Method:             "PUT",
		PathPattern:        "/driver_instances/{driver_instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateDriverInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateDriverInstanceOK), nil
}

/*
Updates the `service` with the id **serviceID**
*/
func (a *Client) UpdateService(params *UpdateServiceParams, authInfo client.AuthInfoWriter) (*UpdateServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServiceParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "updateService",
		Method:             "PUT",
		PathPattern:        "/services/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateServiceOK), nil
}

/*
Updates the plan with the id **planID** for the service id **serviceID**
*/
func (a *Client) UpdateServicePlan(params *UpdateServicePlanParams, authInfo client.AuthInfoWriter) (*UpdateServicePlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateServicePlanParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "updateServicePlan",
		Method:             "PUT",
		PathPattern:        "/plans/{plan_id}",
		ProducesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateServicePlanReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateServicePlanOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
