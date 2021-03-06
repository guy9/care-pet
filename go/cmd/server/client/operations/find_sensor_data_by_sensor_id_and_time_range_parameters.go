// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindSensorDataBySensorIDAndTimeRangeParams creates a new FindSensorDataBySensorIDAndTimeRangeParams object
// with the default values initialized.
func NewFindSensorDataBySensorIDAndTimeRangeParams() *FindSensorDataBySensorIDAndTimeRangeParams {
	var ()
	return &FindSensorDataBySensorIDAndTimeRangeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindSensorDataBySensorIDAndTimeRangeParamsWithTimeout creates a new FindSensorDataBySensorIDAndTimeRangeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindSensorDataBySensorIDAndTimeRangeParamsWithTimeout(timeout time.Duration) *FindSensorDataBySensorIDAndTimeRangeParams {
	var ()
	return &FindSensorDataBySensorIDAndTimeRangeParams{

		timeout: timeout,
	}
}

// NewFindSensorDataBySensorIDAndTimeRangeParamsWithContext creates a new FindSensorDataBySensorIDAndTimeRangeParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindSensorDataBySensorIDAndTimeRangeParamsWithContext(ctx context.Context) *FindSensorDataBySensorIDAndTimeRangeParams {
	var ()
	return &FindSensorDataBySensorIDAndTimeRangeParams{

		Context: ctx,
	}
}

// NewFindSensorDataBySensorIDAndTimeRangeParamsWithHTTPClient creates a new FindSensorDataBySensorIDAndTimeRangeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindSensorDataBySensorIDAndTimeRangeParamsWithHTTPClient(client *http.Client) *FindSensorDataBySensorIDAndTimeRangeParams {
	var ()
	return &FindSensorDataBySensorIDAndTimeRangeParams{
		HTTPClient: client,
	}
}

/*FindSensorDataBySensorIDAndTimeRangeParams contains all the parameters to send to the API endpoint
for the find sensor data by sensor id and time range operation typically these are written to a http.Request
*/
type FindSensorDataBySensorIDAndTimeRangeParams struct {

	/*From
	  from timestamp

	*/
	From strfmt.DateTime
	/*ID
	  ID of a sensor to fetch

	*/
	ID strfmt.UUID
	/*P
	  sensor parameter id

	*/
	P strfmt.UUID
	/*To
	  to timestamp

	*/
	To strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) WithTimeout(timeout time.Duration) *FindSensorDataBySensorIDAndTimeRangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) WithContext(ctx context.Context) *FindSensorDataBySensorIDAndTimeRangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) WithHTTPClient(client *http.Client) *FindSensorDataBySensorIDAndTimeRangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFrom adds the from to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) WithFrom(from strfmt.DateTime) *FindSensorDataBySensorIDAndTimeRangeParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) SetFrom(from strfmt.DateTime) {
	o.From = from
}

// WithID adds the id to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) WithID(id strfmt.UUID) *FindSensorDataBySensorIDAndTimeRangeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithP adds the p to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) WithP(p strfmt.UUID) *FindSensorDataBySensorIDAndTimeRangeParams {
	o.SetP(p)
	return o
}

// SetP adds the p to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) SetP(p strfmt.UUID) {
	o.P = p
}

// WithTo adds the to to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) WithTo(to strfmt.DateTime) *FindSensorDataBySensorIDAndTimeRangeParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the find sensor data by sensor id and time range params
func (o *FindSensorDataBySensorIDAndTimeRangeParams) SetTo(to strfmt.DateTime) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *FindSensorDataBySensorIDAndTimeRangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param from
	qrFrom := o.From
	qFrom := qrFrom.String()
	if qFrom != "" {
		if err := r.SetQueryParam("from", qFrom); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param p
	if err := r.SetPathParam("p", o.P.String()); err != nil {
		return err
	}

	// query param to
	qrTo := o.To
	qTo := qrTo.String()
	if qTo != "" {
		if err := r.SetQueryParam("to", qTo); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
