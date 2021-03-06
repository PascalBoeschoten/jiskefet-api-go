// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PatchLogsIDRunsReader is a Reader for the PatchLogsIDRuns structure.
type PatchLogsIDRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLogsIDRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchLogsIDRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLogsIDRunsOK creates a PatchLogsIDRunsOK with default headers values
func NewPatchLogsIDRunsOK() *PatchLogsIDRunsOK {
	return &PatchLogsIDRunsOK{}
}

/*PatchLogsIDRunsOK handles this case with default header values.

OK
*/
type PatchLogsIDRunsOK struct {
	Payload interface{}
}

func (o *PatchLogsIDRunsOK) Error() string {
	return fmt.Sprintf("[PATCH /logs/{id}/runs][%d] patchLogsIdRunsOK  %+v", 200, o.Payload)
}

func (o *PatchLogsIDRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
