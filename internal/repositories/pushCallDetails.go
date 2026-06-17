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
		UID:        callDetails["agi_arg_1"].(string),
		CallStatus: callDetails["agi_arg_2"].(string),
		HangupBy:   callDetails["agi_arg_3"].(string),
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

func PushCallBackDetails(callDetails map[string]interface{}) error {
	pushCallbackDetails := models.PushCallbackDetails{
		UID:            callDetails["agi_arg_1"].(string),
		CallDuration:   int(callDetails["agi_arg_2"].(float64)),
		DialedDuration: int(callDetails["agi_arg_3"].(float64)),
		CallStatus:     callDetails["agi_arg_4"].(string),
		HangupBy:       callDetails["agi_arg_5"].(string),
	}
	url := config.AppConfig.PushCallDetails
	pushCallbackDetailsJSON, err := json.Marshal(pushCallbackDetails)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(string(pushCallbackDetailsJSON)))
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
