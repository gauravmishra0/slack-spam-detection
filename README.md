# Slack spam detection

This is a Go-based project for detecting spam messages.


## Usage
To use this project, run:
```shell
export SLACK_WEBHOOK_URL="slack_webhook_url"
go run main.go
```
## Example
Input Payload
```js
{
        "RecordType": "Bounce",
        "Type": "SpamNotification",
        "TypeCode": 512,
        "Name": "Spam notification",
        "Tag": "",
        "MessageStream": "outbound",
        "Description": "The message was delivered, but was either blocked by the user, or classified as spam, bulk mail, or had rejected content.",
        "Email": "zaphod@example.com",
        "From": "notifications@honeybadger.io",
        "BouncedAt": "2023-02-27T21:41:30Z",
}
```
If Type is SpamNotification it will send alert on slack channel using slack webhook url. 
