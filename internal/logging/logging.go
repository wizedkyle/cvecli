package logging

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

func Console() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	log := zerolog.New(output).With().Timestamp().Logger()
	return &log
}

func DebugHttpResponse(response *http.Response) {
	// Log response status and response body
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		Console().Fatal().Err(err)
	}
	Console().Debug().Msg("Response Status: " + response.Status)
	Console().Debug().Msg("Response Body: " + string(bytes))
}

func SetLoggingLevel(debug *bool) {
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
