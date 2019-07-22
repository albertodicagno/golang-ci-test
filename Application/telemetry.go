package main
import (
"github.com/Microsoft/ApplicationInsights-Go/appinsights"
"os"
)

// StartTelemetry starts sending telemetry to the AppInsights service
func StartTelemetry(){
	client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))
	request := appinsights.NewRequestTelemetry("GET", "https://myapp.azurewebsites.net/", 1 , "Success")
	client.Track(request)	
}
