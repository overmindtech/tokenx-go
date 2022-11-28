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

// checks if the Organization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Organization{}

// Organization An organization
type Organization struct {
	// Unique ID of the org
	OrgId *string `json:"org_id,omitempty"`
	// Name of the organization you would like to create. This is the name that an end-user would type in the pre-login prompt to identify which organization they wanted to log in through. Unique logical identifier. May contain lowercase alphabetical characters, numbers, underscores (_), and dashes (-). Can start with a number. Must be between 3 and 50 characters.
	Name *string `json:"name,omitempty"`
	// User-friendly name of the organizations that can be displayed in the login flow and email templates.
	DisplayName *string `json:"display_name,omitempty"`
	// The public Nkey which signs all NATS \"user\" tokens
	PublicNkey *string `json:"public_nkey,omitempty"`
}

// NewOrganization instantiates a new Organization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganization() *Organization {
	this := Organization{}
	return &this
}

// NewOrganizationWithDefaults instantiates a new Organization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationWithDefaults() *Organization {
	this := Organization{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Organization) GetOrgId() string {
	if o == nil || isNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetOrgIdOk() (*string, bool) {
	if o == nil || isNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Organization) HasOrgId() bool {
	if o != nil && !isNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Organization) SetOrgId(v string) {
	o.OrgId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Organization) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Organization) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Organization) SetName(v string) {
	o.Name = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Organization) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Organization) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Organization) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPublicNkey returns the PublicNkey field value if set, zero value otherwise.
func (o *Organization) GetPublicNkey() string {
	if o == nil || isNil(o.PublicNkey) {
		var ret string
		return ret
	}
	return *o.PublicNkey
}

// GetPublicNkeyOk returns a tuple with the PublicNkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Organization) GetPublicNkeyOk() (*string, bool) {
	if o == nil || isNil(o.PublicNkey) {
		return nil, false
	}
	return o.PublicNkey, true
}

// HasPublicNkey returns a boolean if a field has been set.
func (o *Organization) HasPublicNkey() bool {
	if o != nil && !isNil(o.PublicNkey) {
		return true
	}

	return false
}

// SetPublicNkey gets a reference to the given string and assigns it to the PublicNkey field.
func (o *Organization) SetPublicNkey(v string) {
	o.PublicNkey = &v
}

func (o Organization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Organization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.PublicNkey) {
		toSerialize["public_nkey"] = o.PublicNkey
	}
	return toSerialize, nil
}

type NullableOrganization struct {
	value *Organization
	isSet bool
}

func (v NullableOrganization) Get() *Organization {
	return v.value
}

func (v *NullableOrganization) Set(val *Organization) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganization(val *Organization) *NullableOrganization {
	return &NullableOrganization{value: val, isSet: true}
}

func (v NullableOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


