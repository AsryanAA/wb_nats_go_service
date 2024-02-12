package pkg

const (
	MsgOrderCreated    string = "A new order successfully created"
	MsgNotParseJSON    string = "Can not parse this json"
	MsgDoNotSendToNats string = "Can not send to queue from nats"
	MsgDoNotInitNats   string = "Can not connect to Nats"
	MsgNotDataFound    string = "Request not data found"
)
