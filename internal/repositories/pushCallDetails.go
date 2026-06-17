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

func getStringValue(callDetails map[string]interface{}, key string) (string, error) {
	value, ok := callDetails[key]
	if !ok {
		return "", fmt.Errorf("missing %s", key)
	}

	strValue, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("expected %s to be string, got %T", key, value)
	}

	return strings.TrimSpace(strValue), nil
}

func getIntValue(callDetails map[string]interface{}, key string) (int, error) {
	value, ok := callDetails[key]
	if !ok {
		return 0, fmt.Errorf("missing %s", key)
	}

	switch v := value.(type) {
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case float64:
		return int(v), nil
	case string:
		s := strings.TrimSpace(v)
		if s == "" {
			return 0, fmt.Errorf("%s is empty", key)
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return 0, fmt.Errorf("invalid %s value %q: %w", key, s, err)
		}
		return n, nil
	default:
		return 0, fmt.Errorf("expected %s to be numeric or string, got %T", key, value)
	}
}

func UpdateCallConnectDetails(callDetails map[string]interface{}) error {
	uid, err := getStringValue(callDetails, "agi_arg_1")
	if err != nil {
		return err
	}

	callStatus, err := getStringValue(callDetails, "agi_arg_2")
	if err != nil {
		return err
	}

	hangupBy, err := getStringValue(callDetails, "agi_arg_3")
	if err != nil {
		return err
	}

	updatePostConnectDetails := models.UpdatePostConnectDetails{
		UID:        uid,
		CallStatus: callStatus,
		HangupBy:   hangupBy,
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
	uid, err := getStringValue(callDetails, "agi_arg_1")
	if err != nil {
		return err
	}

	callDuration, err := getIntValue(callDetails, "agi_arg_2")
	if err != nil {
		return err
	}

	dialedDuration, err := getIntValue(callDetails, "agi_arg_3")
	if err != nil {
		return err
	}

	callStatus, err := getStringValue(callDetails, "agi_arg_4")
	if err != nil {
		return err
	}

	hangupBy, err := getStringValue(callDetails, "agi_arg_5")
	if err != nil {
		return err
	}

	pushCallbackDetails := models.PushCallbackDetails{
		UID:            uid,
		CallDuration:   callDuration,
		DialedDuration: dialedDuration,
		CallStatus:     callStatus,
		HangupBy:       hangupBy,
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
