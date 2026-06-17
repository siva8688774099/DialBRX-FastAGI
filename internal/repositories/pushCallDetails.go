package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	fmt.Println("URL for pushing call details:", url)
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
	CallDurationStr, ok := callDetails["agi_arg_2"].(string)
	DialedDurationStr, ok := callDetails["agi_arg_3"].(string)

	if !ok {
		return fmt.Errorf("invalid call duration")
	}
	callDuration, err := strconv.Atoi(CallDurationStr)
	if err != nil {
		return fmt.Errorf("invalid call duration: %v", err)
	}
	if !ok {
		return fmt.Errorf("invalid dialed duration")
	}
	dialedDuration, err := strconv.Atoi(DialedDurationStr)
	if err != nil {
		return fmt.Errorf("invalid dialed duration: %v", err)
	}
	pushCallbackDetails := models.PushCallbackDetails{
		UID:            callDetails["agi_arg_1"].(string),
		CallDuration:   callDuration,
		DialedDuration: dialedDuration,
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
