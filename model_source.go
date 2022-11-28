/*
Overmind API

Exchanges OAuth tokens for NATS tokens

API version: 0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package overmind

import (
	"encoding/json"
)

// checks if the Source type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Source{}

// Source A source that is capable of discovering items
type Source struct {
	// The descriptive name of the source
	DescriptiveName *string `json:"descriptive_name,omitempty"`
	// Unique ID of the source
	SourceId *string `json:"source_id,omitempty"`
	// OAuth ID token that the source should use
	IdToken *string `json:"id_token,omitempty"`
	// OAuth token that allows the source to respond to requests
	AccessToken *string `json:"access_token,omitempty"`
	// OAuth refresh token that can be exchanged for a new access_token
	RefreshToken *string `json:"refresh_token,omitempty"`
	// How many replicas of the source to run??? Do we need this?
	Replicas *float32 `json:"replicas,omitempty"`
	// Docker image of the source
	Image *string `json:"image,omitempty"`
	// Config for this source. See the source documentation for what config is available/required
	Config *map[string]string `json:"config,omitempty"`
}

// NewSource instantiates a new Source object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSource() *Source {
	this := Source{}
	return &this
}

// NewSourceWithDefaults instantiates a new Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceWithDefaults() *Source {
	this := Source{}
	return &this
}

// GetDescriptiveName returns the DescriptiveName field value if set, zero value otherwise.
func (o *Source) GetDescriptiveName() string {
	if o == nil || isNil(o.DescriptiveName) {
		var ret string
		return ret
	}
	return *o.DescriptiveName
}

// GetDescriptiveNameOk returns a tuple with the DescriptiveName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetDescriptiveNameOk() (*string, bool) {
	if o == nil || isNil(o.DescriptiveName) {
		return nil, false
	}
	return o.DescriptiveName, true
}

// HasDescriptiveName returns a boolean if a field has been set.
func (o *Source) HasDescriptiveName() bool {
	if o != nil && !isNil(o.DescriptiveName) {
		return true
	}

	return false
}

// SetDescriptiveName gets a reference to the given string and assigns it to the DescriptiveName field.
func (o *Source) SetDescriptiveName(v string) {
	o.DescriptiveName = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *Source) GetSourceId() string {
	if o == nil || isNil(o.SourceId) {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetSourceIdOk() (*string, bool) {
	if o == nil || isNil(o.SourceId) {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *Source) HasSourceId() bool {
	if o != nil && !isNil(o.SourceId) {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *Source) SetSourceId(v string) {
	o.SourceId = &v
}

// GetIdToken returns the IdToken field value if set, zero value otherwise.
func (o *Source) GetIdToken() string {
	if o == nil || isNil(o.IdToken) {
		var ret string
		return ret
	}
	return *o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetIdTokenOk() (*string, bool) {
	if o == nil || isNil(o.IdToken) {
		return nil, false
	}
	return o.IdToken, true
}

// HasIdToken returns a boolean if a field has been set.
func (o *Source) HasIdToken() bool {
	if o != nil && !isNil(o.IdToken) {
		return true
	}

	return false
}

// SetIdToken gets a reference to the given string and assigns it to the IdToken field.
func (o *Source) SetIdToken(v string) {
	o.IdToken = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *Source) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *Source) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *Source) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *Source) GetRefreshToken() string {
	if o == nil || isNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetRefreshTokenOk() (*string, bool) {
	if o == nil || isNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *Source) HasRefreshToken() bool {
	if o != nil && !isNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *Source) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *Source) GetReplicas() float32 {
	if o == nil || isNil(o.Replicas) {
		var ret float32
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetReplicasOk() (*float32, bool) {
	if o == nil || isNil(o.Replicas) {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *Source) HasReplicas() bool {
	if o != nil && !isNil(o.Replicas) {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given float32 and assigns it to the Replicas field.
func (o *Source) SetReplicas(v float32) {
	o.Replicas = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Source) GetImage() string {
	if o == nil || isNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetImageOk() (*string, bool) {
	if o == nil || isNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Source) HasImage() bool {
	if o != nil && !isNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Source) SetImage(v string) {
	o.Image = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Source) GetConfig() map[string]string {
	if o == nil || isNil(o.Config) {
		var ret map[string]string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Source) GetConfigOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Source) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]string and assigns it to the Config field.
func (o *Source) SetConfig(v map[string]string) {
	o.Config = &v
}

func (o Source) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Source) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DescriptiveName) {
		toSerialize["descriptive_name"] = o.DescriptiveName
	}
	if !isNil(o.SourceId) {
		toSerialize["source_id"] = o.SourceId
	}
	if !isNil(o.IdToken) {
		toSerialize["id_token"] = o.IdToken
	}
	if !isNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !isNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if !isNil(o.Replicas) {
		toSerialize["replicas"] = o.Replicas
	}
	if !isNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableSource struct {
	value *Source
	isSet bool
}

func (v NullableSource) Get() *Source {
	return v.value
}

func (v *NullableSource) Set(val *Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSource) IsSet() bool {
	return v.isSet
}

func (v *NullableSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSource(val *Source) *NullableSource {
	return &NullableSource{value: val, isSet: true}
}

func (v NullableSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


