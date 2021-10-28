package cmdutils

import (
	"github.com/pterm/pterm"
	"net/http"
)

func OutputError(httpResponse *http.Response, err error) {
	if httpResponse.StatusCode == 400 {
		pterm.Error.Println(err)
	} else if httpResponse.StatusCode == 401 {
		pterm.Error.Println("Unauthorized access please check the user exists, and credentials are valid.")
	} else if httpResponse.StatusCode == 403 {
		pterm.Error.Println("This operation is not allowed for your user account.")
	} else if httpResponse.StatusCode == 404 {
		pterm.Error.Println("Requested resource could not be found.")
	} else if httpResponse.StatusCode == 500 {
		pterm.Error.Println("Internal server error.")
	}
}
