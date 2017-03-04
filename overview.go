package rabbitclient

type Overview struct {
	Contexts []struct {
		Description string `json:"description"`
		Node        string `json:"node"`
		Path        string `json:"path"`
		Port        string `json:"port"`
	} `json:"contexts"`
	ErlangVersion string `json:"erlang_version"`
	ExchangeTypes []struct {
		Description string `json:"description"`
		Enabled     bool   `json:"enabled"`
		Name        string `json:"name"`
	} `json:"exchange_types"`
	Listeners []struct {
		IpAdress string `json:"ip_address"`
		Node     string `json:"node"`
		Port     int    `json:"port"`
		Protocol string `json:"protocol"`
	} `json:"listeners"`
	ManagementVersion string   `json:"management_version"`
	MessageStats struct {
		Deliver int64 `json:"deliver"`
		DeliverDetails struct {
			Rate float64 `json:"rate"`
		} `json:"deliver_details"`
		DeliverNoAck int64 `json:"deliver_no_ack"`
		DeliverNoAckDetails struct{
			Rate float64 `json:"rate"`
		} `json:"deliver_no_ack_details"`
		Get int `json:"get"`
		GetDetails struct{
			Rate float64 `json:"rate"`
		} `json:"get_details"`
		GetNoAck int64 `json:"get_no_ack"`
		GetNoAckDetails struct{
			Rate float64 `json:"rate"`
		} `json:"get_no_ack_details"`
		Publish int64 `json:"publish"`
		PublishDetails struct{
			Rate float64 `json:"rate"`
		} `json:"publish_details"`
		PublishIn int64 `json:"publish_in"`
		PublishInDetails struct{
			Rate float64 `json:"rate"`
		} `json:"publish_in_details"`
		PublishOut int64 `json:"publish_out"`
		PublishOutDetails struct{
			Rate float64 `json:"rate"`
		} `json:"publish_out_details"`
		Ack int64 `json:"ack"`
		AckDetails struct{
			Rate float64 `json:"rate"`
		} `json:"ack_details"`
		DeliverGet int64 `json:"deliver_get"`
		DeliverGetDetails struct{
			Rate float64 `json:"rate"`
		} `json:"deliver_get_details"`
		Confirm int64 `json:"confirm"`
		ConfirmDetails struct{
			Rate float64 `json:"rate"`
		} `json:"confirm_details"`
		ReturnUnroutable int64 `json:"return_unroutable"`
		ReturnUnroutableDetails struct{
			Rate float64 `json:"rate"`
		} `json:"return_unroutable_details"`
		Redeliver int64 `json:"redeliver"`
		RedeliverDetails struct{
			Rate float64 `json:"rate"`
		} `json:"redeliver_details"`
	}`json:"message_stats"`
	Node string   `json:"node"`
	ObjectTotals struct {
		Channels    int `json:"channels"`
		Connections int `json:"connections"`
		Consumers   int `json:"consumers"`
		Exchanges   int `json:"exchanges"`
		Queues      int `json:"queues"`
	} `json:"object_totals"`
	QueueTotals struct {
		Messages int `json:"messages"`
		MessagesDetails struct {
			Interval  int     `json:"interval"`
			LastEvent int     `json:"last_event"`
			Rate      float64 `json:"rate"`
		} `json:"messages_details"`
		MessagesReady int `json:"messages_ready"`
		MessagesReadyDetails struct {
			Interval  int     `json:"interval"`
			LastEvent int     `json:"last_event"`
			Rate      float64 `json:"rate"`
		} `json:"messages_ready_details"`
		MessagesUnacknowledged int `json:"messages_unacknowledged"`
		MessagesUnacknowledgedDetails struct {
			Interval  int     `json:"interval"`
			LastEvent int     `json:"last_event"`
			Rate      float64 `json:"rate"`
		} `json:"messages_unacknowledged_details"`
	} `json:"queue_totals"`
	RabbitmqVersion  string `json:"rabbitmq_version"`
	StatisticsDbNode string `json:"statistics_db_node"`
	StatisticsLevel  string `json:"statistics_level"`
}
