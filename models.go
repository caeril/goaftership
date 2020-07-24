package goaftership

type StatusTag string

const ST_Pending = StatusTag("Pending")
const ST_InfoReceived = StatusTag("InfoReceived")
const ST_InTransit = StatusTag("InTransit")
const ST_OutForDelivery = StatusTag("OutForDelivery")
const ST_AttemptFail = StatusTag("AttemptFail")
const ST_Delivered = StatusTag("Delivered")
const ST_AvailableForPickup = StatusTag("AvailableForPickup")
const ST_Exception = StatusTag("Exception")
const ST_Expired = StatusTag("Expired")

type EnvelopeMeta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Checkpoint struct {
	Location string    `json:"location"`
	Message  string    `json:"message"`
	Tag      StatusTag `json:"tag"`
	Time     string    `json:"checkpoint_time"`
}

type Tracking struct {
	Id             string       `json:"id"`
	TrackingNumber string       `json:"tracking_number"`
	Slug           string       `json:"slug"`
	Checkpoints    []Checkpoint `json:"checkpoints"`
	Tag            StatusTag    `json:"tag"`
}

type EnvelopeData struct {
	Tracking Tracking `json:"tracking"`
}

type ResponseEnvelope struct {
	Meta EnvelopeMeta `json:"meta"`
	Data EnvelopeData `json:"data"`
}

// --

type RequestEnvelope struct {
	Tracking     TrackingRequest     `json:"tracking,omitempty"`
	Notification NotificationRequest `json:"notification,omitempty"`
}

type TrackingRequest struct {
	TrackingNumber string `json:"tracking_number"`
}

// -- Notification Models

type NotificationRequest struct {
	Emails []string `json:"emails"`
	SMSes  []string `json:"smses"`
}
