// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/*ListClustersOK handles this case with default header values.

ClusterList
*/
type ListClustersOK struct {
	Payload models.ClusterList
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersOK  %+v", 200, o.Payload)
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersUnauthorized creates a ListClustersUnauthorized with default headers values
func NewListClustersUnauthorized() *ListClustersUnauthorized {
	return &ListClustersUnauthorized{}
}

/*ListClustersUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListClustersUnauthorized struct {
}

func (o *ListClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersUnauthorized ", 401)
}

func (o *ListClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClustersForbidden creates a ListClustersForbidden with default headers values
func NewListClustersForbidden() *ListClustersForbidden {
	return &ListClustersForbidden{}
}

/*ListClustersForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListClustersForbidden struct {
}

func (o *ListClustersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClustersForbidden ", 403)
}

func (o *ListClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClustersDefault creates a ListClustersDefault with default headers values
func NewListClustersDefault(code int) *ListClustersDefault {
	return &ListClustersDefault{
		_statusCode: code,
	}
}

/*ListClustersDefault handles this case with default header values.

ErrorResponse is the default representation of an error
*/
type ListClustersDefault struct {
	_statusCode int

	Payload *models.ErrorDetails
}

// Code gets the status code for the list clusters default response
func (o *ListClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListClustersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters][%d] listClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}