package goaftership

type EnvelopeMeta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

type Checkpoint struct {
	Location string `json:"location"`
	Message  string `json:"message"`
	Tag      string `json:"tag"`
	Time     string `json:"checkpoint_time"`
}

type Tracking struct {
	Id             string       `json:"id"`
	TrackingNumber string       `json:"tracking_number"`
	Slug           string       `json:"slug"`
	Checkpoints    []Checkpoint `json:"checkpoints"`
	Tag            string       `json:"tag"`
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
	Tracking TrackingRequest `json:"tracking"`
}

type TrackingRequest struct {
	TrackingNumber string `json:"tracking_number"`
}
