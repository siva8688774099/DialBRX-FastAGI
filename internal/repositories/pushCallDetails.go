package repositories

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"strconv"

	"github.com/Siva_Nutakki/DialBRX-FastAGI/config"
	"github.com/Siva_Nutakki/DialBRX-FastAGI/internal/models"
)

func UpdateCallConnectDetails(callDetails map[string]string) error {
	updatePostConnectDetails := models.UpdatePostConnectDetails{
		UID:        callDetails["agi_arg_1"],
		CallStatus: callDetails["agi_arg_2"],
		HangupBy:   callDetails["agi_arg_3"],
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

func PushCallBackDetails(callDetails map[string]string) error {
	callDuration, err := strconv.Atoi(callDetails["agi_arg_2"])
	if err != nil {
		fmt.Println("Error converting agi_arg_2:", err)
		return fmt.Errorf("invalid integer value for agi_arg_2: %v", err)
	}

	dialedDuration, err := strconv.Atoi(callDetails["agi_arg_3"])
	if err != nil {
		fmt.Println("Error converting agi_arg_3:", err)
		return fmt.Errorf("invalid integer value for agi_arg_3: %v", err)
	}
	pushCallbackDetails := models.PushCallbackDetails{
		UID:            callDetails["agi_arg_1"],
		CallDuration:   callDuration,
		DialedDuration: dialedDuration,
		CallStatus:     callDetails["agi_arg_4"],
		HangupBy:       callDetails["agi_arg_5"],
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
