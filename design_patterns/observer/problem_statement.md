### You are building a YouTube-like platform.

Users can subscribe to channels.

When a channel uploads a new video:
```
"Go Concurrency Explained"
```
all subscribers should automatically receive a notification.

### Requirements
#### Subscriber Operations

A user should be able to:
```Go
Subscribe(channel)
Unsubscribe(channel)
```

#### Channel Operations

A channel should be able to:
```Go
UploadVideo(title string)
```

#### Notification Behavior

Suppose:

Channel: TechWorld

Subscribers:
- Aditya
- Rahul
- Priya

When:
```Go
channel.UploadVideo("Go Concurrency Explained")
```
Expected:
```
Notify Aditya
Notify Rahul
Notify Priya
```
### Future Requirements

Tomorrow product team may introduce:
```
EmailSubscriber
Send Email
MobileSubscriber
Send Push Notification
SlackSubscriber
Send Slack Message
```
The channel should not need modification whenever a new subscriber type is introduced.