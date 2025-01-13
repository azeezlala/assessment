package model

type NotificationMessage struct {
	Type    string
	Content string
}

type Notification struct {
	ID      string
	UserID  string
	Message NotificationMessage
}
