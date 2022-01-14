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

// InlineResponse20015DataMagnets struct for InlineResponse20015DataMagnets
type InlineResponse20015DataMagnets struct {
	Magnet *string `json:"magnet,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Instant *bool `json:"instant,omitempty"`
	Error *InlineResponse20010DataError `json:"error,omitempty"`
}

// NewInlineResponse20015DataMagnets instantiates a new InlineResponse20015DataMagnets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20015DataMagnets() *InlineResponse20015DataMagnets {
	this := InlineResponse20015DataMagnets{}
	return &this
}

// NewInlineResponse20015DataMagnetsWithDefaults instantiates a new InlineResponse20015DataMagnets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20015DataMagnetsWithDefaults() *InlineResponse20015DataMagnets {
	this := InlineResponse20015DataMagnets{}
	return &this
}

// GetMagnet returns the Magnet field value if set, zero value otherwise.
func (o *InlineResponse20015DataMagnets) GetMagnet() string {
	if o == nil || o.Magnet == nil {
		var ret string
		return ret
	}
	return *o.Magnet
}

// GetMagnetOk returns a tuple with the Magnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015DataMagnets) GetMagnetOk() (*string, bool) {
	if o == nil || o.Magnet == nil {
		return nil, false
	}
	return o.Magnet, true
}

// HasMagnet returns a boolean if a field has been set.
func (o *InlineResponse20015DataMagnets) HasMagnet() bool {
	if o != nil && o.Magnet != nil {
		return true
	}

	return false
}

// SetMagnet gets a reference to the given string and assigns it to the Magnet field.
func (o *InlineResponse20015DataMagnets) SetMagnet(v string) {
	o.Magnet = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *InlineResponse20015DataMagnets) GetHash() string {
	if o == nil || o.Hash == nil {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015DataMagnets) GetHashOk() (*string, bool) {
	if o == nil || o.Hash == nil {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *InlineResponse20015DataMagnets) HasHash() bool {
	if o != nil && o.Hash != nil {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *InlineResponse20015DataMagnets) SetHash(v string) {
	o.Hash = &v
}

// GetInstant returns the Instant field value if set, zero value otherwise.
func (o *InlineResponse20015DataMagnets) GetInstant() bool {
	if o == nil || o.Instant == nil {
		var ret bool
		return ret
	}
	return *o.Instant
}

// GetInstantOk returns a tuple with the Instant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015DataMagnets) GetInstantOk() (*bool, bool) {
	if o == nil || o.Instant == nil {
		return nil, false
	}
	return o.Instant, true
}

// HasInstant returns a boolean if a field has been set.
func (o *InlineResponse20015DataMagnets) HasInstant() bool {
	if o != nil && o.Instant != nil {
		return true
	}

	return false
}

// SetInstant gets a reference to the given bool and assigns it to the Instant field.
func (o *InlineResponse20015DataMagnets) SetInstant(v bool) {
	o.Instant = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse20015DataMagnets) GetError() InlineResponse20010DataError {
	if o == nil || o.Error == nil {
		var ret InlineResponse20010DataError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015DataMagnets) GetErrorOk() (*InlineResponse20010DataError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse20015DataMagnets) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given InlineResponse20010DataError and assigns it to the Error field.
func (o *InlineResponse20015DataMagnets) SetError(v InlineResponse20010DataError) {
	o.Error = &v
}

func (o InlineResponse20015DataMagnets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Magnet != nil {
		toSerialize["magnet"] = o.Magnet
	}
	if o.Hash != nil {
		toSerialize["hash"] = o.Hash
	}
	if o.Instant != nil {
		toSerialize["instant"] = o.Instant
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20015DataMagnets struct {
	value *InlineResponse20015DataMagnets
	isSet bool
}

func (v NullableInlineResponse20015DataMagnets) Get() *InlineResponse20015DataMagnets {
	return v.value
}

func (v *NullableInlineResponse20015DataMagnets) Set(val *InlineResponse20015DataMagnets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20015DataMagnets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20015DataMagnets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20015DataMagnets(val *InlineResponse20015DataMagnets) *NullableInlineResponse20015DataMagnets {
	return &NullableInlineResponse20015DataMagnets{value: val, isSet: true}
}

func (v NullableInlineResponse20015DataMagnets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20015DataMagnets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

