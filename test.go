package main

import (
	"fmt"
	"log"
)

func main() {

	apikey := "b7ee4eacexxxxxxxxxxxxxxxxxxxxx"
	linenumber := "1234xxxxxxx"
	receptor := "09xxxxxxxxx"
	client := NewGhasedak(apikey)

	// CheckSmsStatus
	checkSmsStatusQuery := map[string]interface{}{
		"type": 1,
		"ids":  []string{"2366799", "2366805"},
	}
	fmt.Printf("CheckSmsStatus Input: %v\n", checkSmsStatusQuery)
	checkSmsStatusOutput, err := client.CheckSmsStatus(checkSmsStatusQuery)
	if err != nil {
		log.Fatalf("CheckSmsStatus Error: %v\n", err)
	}
	fmt.Printf("CheckSmsStatus Output: %v\n", checkSmsStatusOutput)

	// GetAccountInformation
	fmt.Println("GetAccountInformation Input: none")
	getAccountInformationOutput, err := client.GetAccountInformation()
	if err != nil {
		log.Fatalf("GetAccountInformation Error: %v\n", err)
	}
	fmt.Printf("GetAccountInformation Output: %v\n", getAccountInformationOutput)

	// GetReceivedSmses
	getReceivedSmsesQuery := map[string]interface{}{
		"line_number": linenumber,
		"is_read":     false,
	}
	fmt.Printf("GetReceivedSmses Input: %v\n", getReceivedSmsesQuery)
	getReceivedSmsesOutput, err := client.GetReceivedSmses(getReceivedSmsesQuery)
	if err != nil {
		log.Fatalf("GetReceivedSmses Error: %v\n", err)
	}
	fmt.Printf("GetReceivedSmses Output: %v\n", getReceivedSmsesOutput)

	//GetReceivedSmsesPaging
	getReceivedSmsesPagingQuery := map[string]interface{}{
		"line_number": linenumber,
		"is_read":     false,
		"page_size":   10,
		"page_number": 1,
	}
	fmt.Printf("GetReceivedSmsesPaging Input: %v\n", getReceivedSmsesPagingQuery)
	getReceivedSmsesPagingOutput, err := client.GetReceivedSmsesPaging(getReceivedSmsesPagingQuery)
	if err != nil {
		log.Fatalf("GetReceivedSmsesPaging Error: %v\n", err)
	}
	fmt.Printf("GetReceivedSmsesPaging Output: %v\n", getReceivedSmsesPagingOutput)

	// GetOtpTemplateParameters
	templateName := "oldOTP"
	fmt.Printf("GetOtpTemplateParameters Input: %v\n", templateName)
	getOtpTemplateParametersOutput, err := client.GetOtpTemplateParameters(templateName)
	if err != nil {
		log.Fatalf("GetOtpTemplateParameters Error: %v\n", err)
	}
	fmt.Printf("GetOtpTemplateParameters Output: %v\n", getOtpTemplateParametersOutput)

	// SendSimpleSms
	sendSimpleSmsCommand := map[string]interface{}{
		//"sendDate":          time.Now().Format(time.RFC3339),
		"message":    "Hello World!",
		"receptor":   receptor,
		"linenumber": linenumber,
		//"clientReferenceId": "stringfsdfsdf",
		//"udh":               false,
	}
	fmt.Printf("SendSimpleSms Input: %v\n", sendSimpleSmsCommand)
	sendSimpleSmsOutput, err := client.SendSimpleSms(sendSimpleSmsCommand)
	if err != nil {
		log.Fatalf("SendSimpleSms Error: %v\n", err)
	}
	fmt.Printf("SendSimpleSms Output: %v\n", sendSimpleSmsOutput)

	// SendBulkSms
	sendBulkSmsCommand := map[string]interface{}{
		"message":    "Hello Everyone!",
		"receptors":  []string{receptor, receptor},
		"linenumber": linenumber,
	}
	fmt.Printf("SendBulkSms Input: %v\n", sendBulkSmsCommand)
	sendBulkSmsOutput, err := client.SendBulkSms(sendBulkSmsCommand)
	if err != nil {
		log.Fatalf("SendBulkSms Error: %v\n", err)
	}
	fmt.Printf("SendBulkSms Output: %v\n", sendBulkSmsOutput)
	// SendPairToPairSms
	sendPairToPairSmsCommand := map[string]interface{}{
		"items": []map[string]interface{}{
			{
				"lineNumber": linenumber,
				"receptor":   receptor,
				"message":    "string",
				//"clientReferenceId": "sdsdfdsftrifdfdng",
				//"sendDate":         time.Now().Format(time.RFC3339),
			},
			{
				"lineNumber": linenumber,
				"receptor":   receptor,
				"message":    "string3",
				//"clientReferenceId": "sdsfdsfsddfddfdsftring",
				//"sendDate":         time.Now().Format(time.RFC3339),
			},
		},
		"udh": false,
	}
	fmt.Printf("SendPairToPairSms Input: %v\n", sendPairToPairSmsCommand)
	sendPairToPairSmsOutput, err := client.SendPairToPairSms(sendPairToPairSmsCommand)
	if err != nil {
		log.Fatalf("SendPairToPairSms Error: %v\n", err)
	}
	fmt.Printf("SendPairToPairSms Output: %v\n", sendPairToPairSmsOutput)

	// SendOtpSmsOld
	sendOtpSmsOldCommand := map[string]interface{}{
		"receptors": []map[string]string{
			{
				"mobile":            receptor,
				"clientReferenceId": "stghfghgfffrxsding",
			},
		},
		"templateName": "oldOTP",
		"param1":       "fgfhfg",
		"param2":       "hfghgf",
		"param3":       "string",
		"param4":       "string",
		"param5":       "string",
		"param6":       "string",
		"param7":       "string",
		"param8":       "string",
		"param9":       "string",
		"param10":      "string",
		"isVoice":      false,
		"udh":          false,
	}
	fmt.Printf("SendOtpSmsOld Input: %v\n", sendOtpSmsOldCommand)
	sendOtpSmsOldOutput, err := client.SendOtpSmsOld(sendOtpSmsOldCommand)
	if err != nil {
		log.Fatalf("SendOtpSmsOld Error: %v\n", err)
	}
	fmt.Printf("SendOtpSmsOld Output: %v\n", sendOtpSmsOldOutput)

	sendOtpSmsCommand := map[string]interface{}{
		"sendDate": "2024-07-11T18:16:05.952Z",
		"receptors": []map[string]string{
			{
				"mobile":            receptor,
				"clientReferenceId": "string",
			},
		},
		"templateName": "newOTP",
		"inputs": []map[string]string{
			{
				"param": "Name",
				"value": "string",
			},
			{
				"param": "Code",
				"value": "string",
			},
		},
		"udh": false,
	}
	fmt.Printf("SendOtpSms Input: %v\n", sendOtpSmsCommand)
	sendOtpSmsOutput, err := client.SendOtpSms(sendOtpSmsCommand)
	if err != nil {
		log.Fatalf("SendOtpSms Error: %v\n", err)
	}
	fmt.Printf("SendOtpSms Output: %v\n", sendOtpSmsOutput)

}
