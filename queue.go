package rabbitclient

import (
	"context"
	"encoding/json"
	"net/http"
)

type Queue struct {
	Memory                        int64                  `json:"memory"`
	MessageStats                  *MessageStats          `json:"message_stats"`
	Messages                      int64                  `json:"messages"`
	MessagesDetails               *StatsDetails          `json:"messages_details"`
	MessagesReady                 int64                  `json:"messages_ready"`
	MessagesReadyDetails          *StatsDetails          `json:"messages_ready_details"`
	MessagesUnacknowledged        int64                  `json:"messages_unacknowledged"`
	MessagesUnacknowledgedDetails *StatsDetails          `json:"messages_unacknowledged_details"`
	IdleSince                     string                 `json:"idle_since"`
	ConsumerUtilization           string                 `json:"consumer_utilization"`
	Policy                        string                 `json:"policy"`
	ExclusiveConsumerTag          string                 `json:"exclusive_consumer_tag"`
	Consumers                     int                    `json:"consumers"`
	RecoverableSlaves             []string                 `json:"recoverable_slaves"`
	State                         string                 `json:"state"`
	MessagesRAM                   int64                  `json:"messages_ram"`
	MessagesReadyRAM              int64                  `json:"messages_ready_ram"`
	MessagesUnacknowledgedRAM     int64                  `json:"messages_unacknowledged_ram"`
	MessageBytes                  int64                  `json:"message_bytes"`
	MessageBytesReady             int64                  `json:"message_bytes_ready"`
	MessageBytesUnacknowledged    int64                  `json:"message_bytes_unacknowledged"`
	MessageBytesRAM               int64                  `json:"message_bytes_ram"`
	MessageBytesPersistent        int64                  `json:"message_bytes_persistent"`
	DiskReads                     int64                  `json:"disk_reads"`
	DiskWrites                    int64                  `json:"disk_writes"`
	BackingQueueStatus            *QueueStatus           `json:"backing_queue_status"`
	Incoming                      []interface{}          `json:"incoming"`
	Deliveries                    []interface{}          `json:"deliveries"`
	ConsumerDetails               []interface{}          `json:"consumer_details"`
	Name                          string                 `json:"name"`
	VHost                         string                 `json:"vhost"`
	Type                          string                 `json:"type"`
	Durable                       bool                   `json:"durable"`
	AutoDelete                    bool                   `json:"auto_delete"`
	Internal                      bool                   `json:"internal"`
	Arguments                     map[string]interface{} `json:"arguments"`
	Node                          string                 `json:"node"`
}

type QueueStatus struct {
	Q1                int64         `json:"q1"`
	Q2                int64         `json:"q2"`
	Q3                int64         `json:"q3"`
	Q4                int64         `json:"q4"`
	Delta             []interface{} `json:"delta"` // e.g. ["delta", "undefined", 0, "undefined"]
	Len               int64         `json:"len"`
	TargetRAMCount    string        `json:"target_ram_count"` // e.g. "infinity"
	NextSeqID         int64         `json:"next_seq_id"`
	AvgIngressRate    float64       `json:"avg_ingress_rate"`
	AvgEgressRate     float64       `json:"avg_egress_rate"`
	AvgAckIngressRate float64       `json:"avg_ack_ingress_rate"`
	AvgAckEgressRate  float64       `json:"avg_ack_egress_rate"`
}

func (c *Conn) GetQueues(ctx context.Context, outC chan<- []Queue, errC chan<- error) {
	err := c.get(ctx, "queues", func(c context.Context, resp *http.Response) error {
		var queue []Queue
		err := json.NewDecoder(resp.Body).Decode(&queue)
		if err != nil {
			return err
		}
		outC <- queue
		return nil
	})
	if err != nil {
		errC <- err
	}
}

func (c *Conn) GetVhostQueue(ctx context.Context, vhost string, outC chan<- []Queue, errC chan<- error) {
	if vhost == "/" {
		vhost = "%2f"
	}
	err := c.get(ctx, "queues/"+vhost, func(c context.Context, resp *http.Response) error {
		var queue []Queue
		err := json.NewDecoder(resp.Body).Decode(&queue)
		if err != nil {
			return err
		}
		outC <- queue
		return nil
	})
	if err != nil {
		errC <- err
	}
}

func (c *Conn) GetQueue(ctx context.Context, vhost, name string, outC chan<- Queue, errC chan<- error) {
	if vhost == "/" {
		vhost = "%2f"
	}
	err := c.get(ctx, "queues/"+vhost+"/"+name, func(c context.Context, resp *http.Response) error {
		var queue Queue
		err := json.NewDecoder(resp.Body).Decode(&queue)
		if err != nil {
			return err
		}
		outC <- queue
		return nil
	})
	if err != nil {
		errC <- err
	}
}
