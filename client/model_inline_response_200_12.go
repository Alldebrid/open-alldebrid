/*
Alldebrid API

Welcome to the OpenAPI Alldebrid API v4 !<br /> You can use this API to access various Alldebrid services from custom applications or scripts.<br /> <br /> API is organized around REST,<br /> returns JSON-encoded responses and use standard HTTP response codes.<br /> <br /> All calls are to be made on the HTTPS endpoints.<br /> Some are public, others require to be authentificated with an apikey (see Authentication).<br /> <br /> You must also identify your apps or script with a meaningfull agent parameter.<br /> <br /> This API version is namespaced as v4, as such all endpoint start with /v4/,<br /> such like http://api.alldebrid.com/v4/ping?agent=apiShowcase.<br /> <br /> This API v4 should be the final version regarding general response format and errors (hopefully).<br />

API version: 4.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20012 struct for InlineResponse20012
type InlineResponse20012 struct {
	Status *string `json:"status,omitempty"`
	Data *InlineResponse20012Data `json:"data,omitempty"`
	Error *InlineResponse20010Error `json:"error,omitempty"`
}

// NewInlineResponse20012 instantiates a new InlineResponse20012 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20012() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// NewInlineResponse20012WithDefaults instantiates a new InlineResponse20012 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20012WithDefaults() *InlineResponse20012 {
	this := InlineResponse20012{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20012) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20012) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20012) SetStatus(v string) {
	o.Status = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20012) GetData() InlineResponse20012Data {
	if o == nil || o.Data == nil {
		var ret InlineResponse20012Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetDataOk() (*InlineResponse20012Data, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20012) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given InlineResponse20012Data and assigns it to the Data field.
func (o *InlineResponse20012) SetData(v InlineResponse20012Data) {
	o.Data = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse20012) GetError() InlineResponse20010Error {
	if o == nil || o.Error == nil {
		var ret InlineResponse20010Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20012) GetErrorOk() (*InlineResponse20010Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse20012) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given InlineResponse20010Error and assigns it to the Error field.
func (o *InlineResponse20012) SetError(v InlineResponse20010Error) {
	o.Error = &v
}

func (o InlineResponse20012) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20012 struct {
	value *InlineResponse20012
	isSet bool
}

func (v NullableInlineResponse20012) Get() *InlineResponse20012 {
	return v.value
}

func (v *NullableInlineResponse20012) Set(val *InlineResponse20012) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20012) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20012) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20012(val *InlineResponse20012) *NullableInlineResponse20012 {
	return &NullableInlineResponse20012{value: val, isSet: true}
}

func (v NullableInlineResponse20012) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20012) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

