package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Siva_Nutakki/DialBRX-FastAGI/config"
	"github.com/Siva_Nutakki/DialBRX-FastAGI/internal/models"
)

func UpdateCallConnectDetails(callDetails map[string]interface{}) error {
	updatePostConnectDetails := models.UpdatePostConnectDetails{
		UID:        callDetails["uid"].(string),
		CallStatus: callDetails["status_code"].(string),
		HangupBy:   callDetails["hangup_by"].(string),
	}
	url := config.AppConfig.PushCallDetails
	updatePostConnectDetailsJSON, err := json.Marshal(updatePostConnectDetails)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(updatePostConnectDetailsJSON)))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to push call details, status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
}
