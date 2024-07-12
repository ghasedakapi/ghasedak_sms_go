
# Ghasedak API Client

This is a simple Go client for interacting with the Ghasedak API. The client supports various functionalities including checking SMS status, retrieving account information, sending different types of SMS messages, and more.

## Installation


```sh
go get github.com/ghasedakapi/ghasedak-go
```

[//]: # (## Usage)

### Setting Up


```sh
apikey := "b7ee4eacexxxxxxxxxxxxxxxxxxxxx"
line := "1234xxxxx"
receptor := "09xxxxxxxxx"
client := NewGhasedak(apikey)
```

## Features

### 1. Check SMS Status

```go
checkSmsStatusQuery := map[string]interface{}{
    "type": 1,
    "ids":  []string{"2366799", "2366805"},
}
checkSmsStatusOutput, err := client.CheckSmsStatus(checkSmsStatusQuery)
if err != nil {
    log.Fatalf("CheckSmsStatus Error: %v\n", err)
}
fmt.Printf("CheckSmsStatus Output: %v\n", checkSmsStatusOutput)
```

### 2. Get Account Information

```go
getAccountInformationOutput, err := client.GetAccountInformation()
if err != nil {
    log.Fatalf("GetAccountInformation Error: %v\n", err)
}
fmt.Printf("GetAccountInformation Output: %v\n", getAccountInformationOutput)
```

### 3. Get Received SMSes

```go
getReceivedSmsesQuery := map[string]interface{}{
    "line_number": line,
    "is_read":     false,
}
getReceivedSmsesOutput, err := client.GetReceivedSmses(getReceivedSmsesQuery)
if err != nil {
    log.Fatalf("GetReceivedSmses Error: %v\n", err)
}
fmt.Printf("GetReceivedSmses Output: %v\n", getReceivedSmsesOutput)
```

### 4. Get Received SMSes with Paging

```go
getReceivedSmsesPagingQuery := map[string]interface{}{
    "line_number": line,
    "is_read":     false,
    "page_size":   10,
    "page_number": 1,
}
getReceivedSmsesPagingOutput, err := client.GetReceivedSmsesPaging(getReceivedSmsesPagingQuery)
if err != nil {
    log.Fatalf("GetReceivedSmsesPaging Error: %v\n", err)
}
fmt.Printf("GetReceivedSmsesPaging Output: %v\n", getReceivedSmsesPagingOutput)
```

### 5. Get OTP Template Parameters

```go
templateName := "oldOTP"
getOtpTemplateParametersOutput, err := client.GetOtpTemplateParameters(templateName)
if err != nil {
    log.Fatalf("GetOtpTemplateParameters Error: %v\n", err)
}
fmt.Printf("GetOtpTemplateParameters Output: %v\n", getOtpTemplateParametersOutput)
```

### 6. Send Simple SMS

```go
sendSimpleSmsCommand := map[string]interface{}{
    "message":    "Hello World!",
    "receptor":   receptor,
    "linenumber": line,
}
sendSimpleSmsOutput, err := client.SendSimpleSms(sendSimpleSmsCommand)
if err != nil {
    log.Fatalf("SendSimpleSms Error: %v\n", err)
}
fmt.Printf("SendSimpleSms Output: %v\n", sendSimpleSmsOutput)
```

### 7. Send Bulk SMS

```go
sendBulkSmsCommand := map[string]interface{}{
    "message":    "Hello Everyone!",
    "receptors":  []string{receptor, receptor},
    "linenumber": line,
}
sendBulkSmsOutput, err := client.SendBulkSms(sendBulkSmsCommand)
if err != nil {
    log.Fatalf("SendBulkSms Error: %v\n", err)
}
fmt.Printf("SendBulkSms Output: %v\n", sendBulkSmsOutput)
```

### 8. Send Pair-to-Pair SMS

```go
sendPairToPairSmsCommand := map[string]interface{}{
    "items": []map[string]interface{}{
        {
            "lineNumber": line,
            "receptor":   receptor,
            "message":    "string",
        },
        {
            "lineNumber": line,
            "receptor":   receptor,
            "message":    "string3",
        },
    },
    "udh": false,
}
sendPairToPairSmsOutput, err := client.SendPairToPairSms(sendPairToPairSmsCommand)
if err != nil {
    log.Fatalf("SendPairToPairSms Error: %v\n", err)
}
fmt.Printf("SendPairToPairSms Output: %v\n", sendPairToPairSmsOutput)
```

### 9. Send OTP SMS (Old Method)

```go
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
sendOtpSmsOldOutput, err := client.SendOtpSmsOld(sendOtpSmsOldCommand)
if err != nil {
    log.Fatalf("SendOtpSmsOld Error: %v\n", err)
}
fmt.Printf("SendOtpSmsOld Output: %v\n", sendOtpSmsOldOutput)
```

### 10. Send OTP SMS (new Method)

```go
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
sendOtpSmsOutput, err := client.SendOtpSms(sendOtpSmsCommand)
if err != nil {
    log.Fatalf("SendOtpSms Error: %v\n", err)
}
fmt.Printf("SendOtpSms Output: %v\n", sendOtpSmsOutput)
```

## Contributing

Feel free to open issues or submit pull requests for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.

---

Replace `"yourusername"` in the installation section with your actual GitHub username. Also, ensure you implement the `NewGhasedak` function and other related functions (`CheckSmsStatus`, `GetAccountInformation`, `GetReceivedSmses`, `GetReceivedSmsesPaging`, `GetOtpTemplateParameters`, `SendSimpleSms`, `SendBulkSms`, `SendPairToPairSms`, `SendOtpSmsOld`, `SendOtpSms`) in your code to make the client work as expected.
