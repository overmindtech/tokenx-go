/*
Overmind API

API for managing your Overmind account

API version: 0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overmind

import (
	"encoding/json"
)

// checks if the AdminCreateSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdminCreateSourceRequest{}

// AdminCreateSourceRequest Details to update for a source
type AdminCreateSourceRequest struct {
	// The descriptive name of the source
	DescriptiveName string `json:"descriptive_name"`
	// What source to configure. Currently either \"stdlib\" or \"aws\"
	Type string `json:"type"`
	// Config for this source. See the source documentation for what source-specific config is available/required. This will be supplied directly to viper via a config file at `/etc/srcman/config/source.yaml`
	Config map[string]interface{} `json:"config,omitempty"`
	// Additional config options that should be passed to the source. The keys of this object should be file names, and the values should be their content. These files will be made available to the source at runtime. Check the source's documentation for what to configure here if required
	AdditionalConfig *map[string]string `json:"additional_config,omitempty"`
}

// NewAdminCreateSourceRequest instantiates a new AdminCreateSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdminCreateSourceRequest(descriptiveName string, type_ string) *AdminCreateSourceRequest {
	this := AdminCreateSourceRequest{}
	this.DescriptiveName = descriptiveName
	this.Type = type_
	return &this
}

// NewAdminCreateSourceRequestWithDefaults instantiates a new AdminCreateSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdminCreateSourceRequestWithDefaults() *AdminCreateSourceRequest {
	this := AdminCreateSourceRequest{}
	return &this
}

// GetDescriptiveName returns the DescriptiveName field value
func (o *AdminCreateSourceRequest) GetDescriptiveName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DescriptiveName
}

// GetDescriptiveNameOk returns a tuple with the DescriptiveName field value
// and a boolean to check if the value has been set.
func (o *AdminCreateSourceRequest) GetDescriptiveNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescriptiveName, true
}

// SetDescriptiveName sets field value
func (o *AdminCreateSourceRequest) SetDescriptiveName(v string) {
	o.DescriptiveName = v
}

// GetType returns the Type field value
func (o *AdminCreateSourceRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AdminCreateSourceRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AdminCreateSourceRequest) SetType(v string) {
	o.Type = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *AdminCreateSourceRequest) GetConfig() map[string]interface{} {
	if o == nil || IsNil(o.Config) {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminCreateSourceRequest) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *AdminCreateSourceRequest) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *AdminCreateSourceRequest) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetAdditionalConfig returns the AdditionalConfig field value if set, zero value otherwise.
func (o *AdminCreateSourceRequest) GetAdditionalConfig() map[string]string {
	if o == nil || IsNil(o.AdditionalConfig) {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalConfig
}

// GetAdditionalConfigOk returns a tuple with the AdditionalConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdminCreateSourceRequest) GetAdditionalConfigOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.AdditionalConfig) {
		return nil, false
	}
	return o.AdditionalConfig, true
}

// HasAdditionalConfig returns a boolean if a field has been set.
func (o *AdminCreateSourceRequest) HasAdditionalConfig() bool {
	if o != nil && !IsNil(o.AdditionalConfig) {
		return true
	}

	return false
}

// SetAdditionalConfig gets a reference to the given map[string]string and assigns it to the AdditionalConfig field.
func (o *AdminCreateSourceRequest) SetAdditionalConfig(v map[string]string) {
	o.AdditionalConfig = &v
}

func (o AdminCreateSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdminCreateSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["descriptive_name"] = o.DescriptiveName
	toSerialize["type"] = o.Type
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.AdditionalConfig) {
		toSerialize["additional_config"] = o.AdditionalConfig
	}
	return toSerialize, nil
}

type NullableAdminCreateSourceRequest struct {
	value *AdminCreateSourceRequest
	isSet bool
}

func (v NullableAdminCreateSourceRequest) Get() *AdminCreateSourceRequest {
	return v.value
}

func (v *NullableAdminCreateSourceRequest) Set(val *AdminCreateSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAdminCreateSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAdminCreateSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdminCreateSourceRequest(val *AdminCreateSourceRequest) *NullableAdminCreateSourceRequest {
	return &NullableAdminCreateSourceRequest{value: val, isSet: true}
}

func (v NullableAdminCreateSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdminCreateSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


