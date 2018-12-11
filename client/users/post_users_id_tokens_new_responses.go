// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostUsersIDTokensNewReader is a Reader for the PostUsersIDTokensNew structure.
type PostUsersIDTokensNewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersIDTokensNewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostUsersIDTokensNewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUsersIDTokensNewCreated creates a PostUsersIDTokensNewCreated with default headers values
func NewPostUsersIDTokensNewCreated() *PostUsersIDTokensNewCreated {
	return &PostUsersIDTokensNewCreated{}
}

/*PostUsersIDTokensNewCreated handles this case with default header values.

PostUsersIDTokensNewCreated post users Id tokens new created
*/
type PostUsersIDTokensNewCreated struct {
}

func (o *PostUsersIDTokensNewCreated) Error() string {
	return fmt.Sprintf("[POST /users/{id}/tokens/new][%d] postUsersIdTokensNewCreated ", 201)
}

func (o *PostUsersIDTokensNewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
