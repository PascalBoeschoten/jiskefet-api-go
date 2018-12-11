// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLogsIDReader is a Reader for the GetLogsID structure.
type GetLogsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLogsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLogsIDOK creates a GetLogsIDOK with default headers values
func NewGetLogsIDOK() *GetLogsIDOK {
	return &GetLogsIDOK{}
}

/*GetLogsIDOK handles this case with default header values.

GetLogsIDOK get logs Id o k
*/
type GetLogsIDOK struct {
}

func (o *GetLogsIDOK) Error() string {
	return fmt.Sprintf("[GET /logs/{id}][%d] getLogsIdOK ", 200)
}

func (o *GetLogsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
