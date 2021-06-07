package response

import "fmt"

// Error http error response struct
type Error struct {
	Code        int    `json:"-"`
	Description string `json:"error_description"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.Description)
}
