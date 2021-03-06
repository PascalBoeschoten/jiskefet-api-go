// Code generated by go-swagger; DO NOT EDIT.

package runs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostRunsReader is a Reader for the PostRuns structure.
type PostRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostRunsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRunsCreated creates a PostRunsCreated with default headers values
func NewPostRunsCreated() *PostRunsCreated {
	return &PostRunsCreated{}
}

/*PostRunsCreated handles this case with default header values.

PostRunsCreated post runs created
*/
type PostRunsCreated struct {
}

func (o *PostRunsCreated) Error() string {
	return fmt.Sprintf("[POST /runs][%d] postRunsCreated ", 201)
}

func (o *PostRunsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
