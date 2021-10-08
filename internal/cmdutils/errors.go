package cmdutils

import (
	"fmt"
	"net/http"
)

func OutputError(httpResponse *http.Response, err error) {
	if httpResponse.StatusCode == 400 {
		fmt.Println(err)
	} else if httpResponse.StatusCode == 401 {
		fmt.Println("Unauthorized access please check the user exists, and credentials are valid.")
	} else if httpResponse.StatusCode == 403 {
		fmt.Println("This operation is not allowed for your user account.")
	} else if httpResponse.StatusCode == 404 {
		fmt.Println("Requested resource could not be found.")
	} else if httpResponse.StatusCode == 500 {
		fmt.Println("Internal server error.")
	}
}
