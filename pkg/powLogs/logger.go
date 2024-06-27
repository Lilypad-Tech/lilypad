package powLogs

import (
	"encoding/json"
	"os"

	"github.com/lilypad-tech/lilypad/pkg/http"
)

type PowLog struct {
	POWTime string // YYYY-MM-DD HH::MM:ss (triggers are every hour on the hour and randomly within the hour)
	// t := time.Now(); t.Format(time.RFC3339)
	Caller   string // cron or rp
	CallerID string // cron or rp public address
	Step     string // text instead of numbers, this way we can query by giving a sort order to the events if we know all the steps, and we do not need to reorder all events if we add one in between
	Payload  string // a json payload
	Status   string // log/info/warn/error a way to split an event payload into multiple entries to then be able to filter logs
}

const namespace = "pow-logs"
const eventsEndpoint = "events"

var host = os.Getenv("API_HOST")

func TrackEvent(data PowLog) {
	if host == "" {
		return
	}

	var url = host + namespace + "/" + eventsEndpoint
	byts, _ := json.Marshal(data)
	payload := string(byts)

	http.GenericJSONPostClient(url, payload)
}
