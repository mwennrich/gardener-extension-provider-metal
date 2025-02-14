// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new machine API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for machine API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddProvisioningEvent adds a machine provisioning event
*/
func (a *Client) AddProvisioningEvent(params *AddProvisioningEventParams, authInfo runtime.ClientAuthInfoWriter) (*AddProvisioningEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddProvisioningEventParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addProvisioningEvent",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/event",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddProvisioningEventReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddProvisioningEventOK), nil

}

/*
AllocateMachine allocates a machine
*/
func (a *Client) AllocateMachine(params *AllocateMachineParams, authInfo runtime.ClientAuthInfoWriter) (*AllocateMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "allocateMachine",
		Method:             "POST",
		PathPattern:        "/v1/machine/allocate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AllocateMachineOK), nil

}

/*
ChassisIdentifyLEDOff sends a power off to the chassis identify l e d
*/
func (a *Client) ChassisIdentifyLEDOff(params *ChassisIdentifyLEDOffParams, authInfo runtime.ClientAuthInfoWriter) (*ChassisIdentifyLEDOffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChassisIdentifyLEDOffParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "chassisIdentifyLEDOff",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/chassis-identify-led-off/{description}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChassisIdentifyLEDOffReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChassisIdentifyLEDOffOK), nil

}

/*
ChassisIdentifyLEDOn sends a power on to the chassis identify l e d
*/
func (a *Client) ChassisIdentifyLEDOn(params *ChassisIdentifyLEDOnParams, authInfo runtime.ClientAuthInfoWriter) (*ChassisIdentifyLEDOnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChassisIdentifyLEDOnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "chassisIdentifyLEDOn",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/chassis-identify-led-on",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ChassisIdentifyLEDOnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChassisIdentifyLEDOnOK), nil

}

/*
CheckMachineLiveliness externals trigger for evaluating machine liveliness
*/
func (a *Client) CheckMachineLiveliness(params *CheckMachineLivelinessParams, authInfo runtime.ClientAuthInfoWriter) (*CheckMachineLivelinessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckMachineLivelinessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkMachineLiveliness",
		Method:             "POST",
		PathPattern:        "/v1/machine/liveliness",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckMachineLivelinessReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckMachineLivelinessOK), nil

}

/*
FinalizeAllocation finalizes the allocation of the machine by reconfiguring the switch sent on successful image installation
*/
func (a *Client) FinalizeAllocation(params *FinalizeAllocationParams, authInfo runtime.ClientAuthInfoWriter) (*FinalizeAllocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFinalizeAllocationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "finalizeAllocation",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/finalize-allocation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FinalizeAllocationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FinalizeAllocationOK), nil

}

/*
FindMachine gets machine by id
*/
func (a *Client) FindMachine(params *FindMachineParams, authInfo runtime.ClientAuthInfoWriter) (*FindMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findMachine",
		Method:             "GET",
		PathPattern:        "/v1/machine/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindMachineOK), nil

}

/*
FindMachines searches machines
*/
func (a *Client) FindMachines(params *FindMachinesParams, authInfo runtime.ClientAuthInfoWriter) (*FindMachinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindMachinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findMachines",
		Method:             "POST",
		PathPattern:        "/v1/machine/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindMachinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FindMachinesOK), nil

}

/*
FreeMachine frees a machine
*/
func (a *Client) FreeMachine(params *FreeMachineParams, authInfo runtime.ClientAuthInfoWriter) (*FreeMachineOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFreeMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "freeMachine",
		Method:             "DELETE",
		PathPattern:        "/v1/machine/{id}/free",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FreeMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FreeMachineOK), nil

}

/*
GetProvisioningEventContainer gets the current machine provisioning event container
*/
func (a *Client) GetProvisioningEventContainer(params *GetProvisioningEventContainerParams, authInfo runtime.ClientAuthInfoWriter) (*GetProvisioningEventContainerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProvisioningEventContainerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProvisioningEventContainer",
		Method:             "GET",
		PathPattern:        "/v1/machine/{id}/event",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetProvisioningEventContainerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProvisioningEventContainerOK), nil

}

/*
IPMIData returns the IP m i connection data for a machine
*/
func (a *Client) IPMIData(params *IPMIDataParams, authInfo runtime.ClientAuthInfoWriter) (*IPMIDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIPMIDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ipmiData",
		Method:             "GET",
		PathPattern:        "/v1/machine/{id}/ipmi",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IPMIDataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*IPMIDataOK), nil

}

/*
ListMachines gets all known machines
*/
func (a *Client) ListMachines(params *ListMachinesParams, authInfo runtime.ClientAuthInfoWriter) (*ListMachinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMachinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listMachines",
		Method:             "GET",
		PathPattern:        "/v1/machine",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListMachinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListMachinesOK), nil

}

/*
MachineBios boots machine into b i o s on next reboot
*/
func (a *Client) MachineBios(params *MachineBiosParams, authInfo runtime.ClientAuthInfoWriter) (*MachineBiosOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineBiosParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineBios",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/bios",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineBiosReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MachineBiosOK), nil

}

/*
MachineOff sends a power off to the machine
*/
func (a *Client) MachineOff(params *MachineOffParams, authInfo runtime.ClientAuthInfoWriter) (*MachineOffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineOffParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineOff",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/off",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineOffReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MachineOffOK), nil

}

/*
MachineOn sends a power on to the machine
*/
func (a *Client) MachineOn(params *MachineOnParams, authInfo runtime.ClientAuthInfoWriter) (*MachineOnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineOnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineOn",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/on",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineOnReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MachineOnOK), nil

}

/*
MachineReset sends a reset to the machine
*/
func (a *Client) MachineReset(params *MachineResetParams, authInfo runtime.ClientAuthInfoWriter) (*MachineResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMachineResetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "machineReset",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/power/reset",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MachineResetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MachineResetOK), nil

}

/*
RegisterMachine registers a machine
*/
func (a *Client) RegisterMachine(params *RegisterMachineParams, authInfo runtime.ClientAuthInfoWriter) (*RegisterMachineOK, *RegisterMachineCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterMachineParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "registerMachine",
		Method:             "POST",
		PathPattern:        "/v1/machine/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterMachineReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RegisterMachineOK:
		return value, nil, nil
	case *RegisterMachineCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
SetChassisIdentifyLEDState sets the state of a chassis identify l e d
*/
func (a *Client) SetChassisIdentifyLEDState(params *SetChassisIdentifyLEDStateParams, authInfo runtime.ClientAuthInfoWriter) (*SetChassisIdentifyLEDStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetChassisIdentifyLEDStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setChassisIdentifyLEDState",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/chassis-identify-led-state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetChassisIdentifyLEDStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetChassisIdentifyLEDStateOK), nil

}

/*
SetMachineState sets the state of a machine
*/
func (a *Client) SetMachineState(params *SetMachineStateParams, authInfo runtime.ClientAuthInfoWriter) (*SetMachineStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetMachineStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setMachineState",
		Method:             "POST",
		PathPattern:        "/v1/machine/{id}/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetMachineStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetMachineStateOK), nil

}

/*
WaitForAllocation waits for an allocation of this machine
*/
func (a *Client) WaitForAllocation(params *WaitForAllocationParams, authInfo runtime.ClientAuthInfoWriter) (*WaitForAllocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWaitForAllocationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "waitForAllocation",
		Method:             "GET",
		PathPattern:        "/v1/machine/{id}/wait",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WaitForAllocationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*WaitForAllocationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
