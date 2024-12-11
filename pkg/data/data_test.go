//go:build unit

package data

import (
	"encoding/json"
	"testing"

	"github.com/mr-tron/base58"
	"pgregory.net/rapid"
)

func TestResultJSONRoundtrip(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate a random Result
		original := Result{
			ID:     generateCID(t),
			DealID: generateCID(t),
			DataID: generateCID(t),
		}

		// Test marshaling to JSON
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		// Test unmarshaling back
		var decoded Result
		err = json.Unmarshal(data, &decoded)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}

		// Verify the roundtrip preserved values
		if original != decoded {
			t.Errorf("Roundtrip failed: got %v, want %v", decoded, original)
		}

		// Verify both fields exist in JSON
		var rawData map[string]interface{}
		err = json.Unmarshal(data, &rawData)
		if err != nil {
			t.Fatalf("Raw unmarshal failed: %v", err)
		}

		// Check both data_id and results_id exist and match
		dataID, hasDataID := rawData["data_id"]
		resultsID, hasResultsID := rawData["results_id"]

		if !hasDataID {
			t.Error("data_id field missing")
		}
		if !hasResultsID {
			t.Error("results_id field missing")
		}
		if dataID != resultsID {
			t.Errorf("data_id and results_id mismatch: %v != %v", dataID, resultsID)
		}
	})
}

func TestResultJSONBackwardsCompatibility(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		expectedDataID := generateCID(t)

		// Test old client format (results_id)
		oldClientJSON := map[string]interface{}{
			"id":         generateCID(t),
			"results_id": expectedDataID,
		}
		oldClientData, _ := json.Marshal(oldClientJSON)

		var resultFromOld Result
		err := json.Unmarshal(oldClientData, &resultFromOld)
		if err != nil {
			t.Fatalf("Unmarshal of old format failed: %v", err)
		}
		if resultFromOld.DataID != expectedDataID {
			t.Errorf("Old format: got DataID %v, want %v", resultFromOld.DataID, expectedDataID)
		}

		// Test new client format (data_id)
		newClientJSON := map[string]interface{}{
			"id":      generateCID(t),
			"data_id": expectedDataID,
		}
		newClientData, _ := json.Marshal(newClientJSON)

		var resultFromNew Result
		err = json.Unmarshal(newClientData, &resultFromNew)
		if err != nil {
			t.Fatalf("Unmarshal of new format failed: %v", err)
		}
		if resultFromNew.DataID != expectedDataID {
			t.Errorf("New format: got DataID %v, want %v", resultFromNew.DataID, expectedDataID)
		}
	})
}

// Generators

func generateCID(t *rapid.T) string {
	bytes := rapid.SliceOfN(rapid.Byte(), 32, 32).Draw(t, "bytes")
	return "Qm" + base58.Encode(bytes)
}
