package rabbitclient

import (
	"context"
	"encoding/json"
	"net/http"
)

type Node struct {
	Name                                  string `json:"name"`
	Type                                  string `json:"type"`
	Running                               bool `json:"running"`
	EnabledPlugins                        []string `json:"enabled_plugins"`
	NetTicktime                           int `json:"net_ticktime"`
	ConfigFiles                           []string `json:"config_files"`
	DbDir                                 string `json:"db_dir"`
	SaslLogFile                           string `json:"sasl_log_file"`
	LogFile                               string `json:"log_file"`
	Contexts                              []Context `json:"contexts"`
	Applications                          []Application `json:"applications"`
	AuthMechanisms                        []AuthMechanism `json:"auth_mechanisms"`
	ExchangeTypes                         []ExchangeType `json:"exchange_types"`
	Processors                            int64 `json:"processors"`
	RunQueue                              int64 `json:"run_queue"`
	Uptime                                int64 `json:"uptime"`
	RatesMode                             string `json:"rates_mode"`
	ProcTotal                             int64 `json:"proc_total"`
	OsOid                                 int64 `json:"os_oid"`
	FdTotal                               int64 `json:"fd_total"`
	SocketsTotal                          int64 `json:"sockets_total"`
	MemLimit                              int64 `json:"mem_limit"`
	MemAlarm                              bool `json:"mem_alarm"`
	DiskFreeLimit                         int `json:"disk_free_limit"`
	DiskFreeAlarm                         bool `json:"disk_free_alarm"`
	Partitions                            []interface{} `json:"partitions"`
	ClusterLinks                          []interface{} `json:"cluster_links"`
	MemUsed                               int64 `json:"mem_used"`
	MemUsedDetails                        StatsDetails `json:"mem_used_details"`
	FdUsed                                int64 `json:"fd_used"`
	FdUsedDetails                         StatsDetails `json:"fd_used_details"`
	SocketsUsed                           int64 `json:"sockets_used"`
	SocketsUsedDetails                    StatsDetails `json:"sockets_used_details"`
	ProcUsed                              int64 `json:"proc_used"`
	ProcUsedDetails                       StatsDetails `json:"proc_used_details"`
	DiskFree                              int64 `json:"disk_free"`
	DiskFreeDetails                       StatsDetails `json:"disk_free_details"`
	IoReadCount                           int64 `json:"io_read_count"`
	IoReadCountDetails                    StatsDetails `json:"io_read_count_details"`
	IoReadBytes                           int64 `json:"io_read_bytes"`
	IoReadBytesDetails                    StatsDetails `json:"io_read_bytes_details"`
	IoReadAvgTime                         float64 `json:"io_read_avg_time"`
	IoReadAvgTimeDetails                  StatsDetails `json:"io_read_avg_time_details"`
	IoWriteCount                          int64 `json:"io_write_count"`
	IoWriteCountDetails                   StatsDetails `json:"io_write_count_details"`
	IoWriteBytes                          int64 `json:"io_write_bytes"`
	IoWriteBytesDetails                   StatsDetails `json:"io_write_bytes_details"`
	IoWriteAvgTime                        float64 `json:"io_write_avg_time"`
	IoWriteAvgTimeDetails                 StatsDetails `json:"io_write_avg_time_details"`
	IoSyncCount                           int64 `json:"io_sync_count"`
	IoSyncCountDetail                     StatsDetails `json:"io_sync_count_detail"`
	IoSyncAvgTime                         float64 `json:"io_sync_avg_time"`
	IoSyncAvgTimeDetails                  StatsDetails `json:"io_sync_avg_time_details"`
	IoSeekCount                           int64 `json:"io_seek_count"`
	IoSeekCountDetail                     StatsDetails `json:"io_seek_count_detail"`
	IoSeekAvgTime                         float64 `json:"io_seek_avg_time"`
	IoSeekAvgTimeDetail                   StatsDetails `json:"io_seek_avg_time_detail"`
	MnesiaRamTxCount                      int64 `json:"mnesia_ram_tx_count"`
	MnesiaRamTxCountDetails               StatsDetails `json:"mnesia_ram_tx_count_details"`
	MnesiaDiskTxCount                     int64 `json:"mnesia_disk_tx_count"`
	MnesiaDiskTxCountDetails              StatsDetails `json:"mnesia_disk_tx_count_details"`
	MsgStoreReadCount                     int64 `json:"msg_store_read_count"`
	MsgStoreReadCountDetails              StatsDetails `json:"msg_store_read_count_details"`
	MsgStoreWriteCount                    int64 `json:"msg_store_write_count"`
	MsgStoreWriteCountDetails             StatsDetails `json:"msg_store_write_count_details"`
	QueueIndexJournalWriteCount           int64 `json:"queue_index_journal_write_count"`
	QueueIndexJournalWriteCountDetails    StatsDetails `json:"queue_index_journal_write_count_details"`
	QueueIndexWriteCount                  int64 `json:"queue_index_write_count"`
	QueueIndexWriteCountDetails           StatsDetails `json:"queue_index_write_count_details"`
	QueueIndexReadCount                   int64 `json:"queue_index_read_count"`
	QueueIndexReadCountDetails            StatsDetails `json:"queue_index_read_count_details"`
	GcNum                                 int64 `json:"gc_num"`
	GcNumDetails                          StatsDetails `json:"gc_num_details"`
	GcBytesReclaimed                      int64 `json:"gc_bytes_reclaimed"`
	GcBytesReclaimedDetails               StatsDetails `json:"gc_bytes_reclaimed_details"`
	ContextSwitches                       int64 `json:"context_switches"`
	ContextSwitchesDetails                StatsDetails `json:"context_switches_details"`
	IoFileHandleOpenAttemptCount          int64 `json:"io_file_handle_open_attempt_count"`
	IoFileHandleOpenAttemptCountDetails   StatsDetails `json:"io_file_handle_open_attempt_count_details"`
	IoFileHandleOpenAttemptAvgTime        float64 `json:"io_file_handle_open_attempt_avg_time"`
	IoFileHandleOpenAttemptAvgTimeDetails StatsDetails `json:"io_file_handle_open_attempt_avg_time_details"`
}

type Application struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type AuthMechanism struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Enable      bool `json:"enable"`
}

func (c *Conn) GetNodes(ctx context.Context, host string, outC chan<- []Node, errC chan<- error) {
	err := c.get(ctx, host, "nodes", func(c context.Context, resp *http.Response) error {
		var nodes []Node
		err := json.NewDecoder(resp.Body).Decode(&nodes)
		if err != nil {
			return err
		}
		outC <- nodes
		return nil
	})
	if err != nil {
		errC <- err
	}
}

func (c *Conn) GetNode(ctx context.Context, host, node string, outC chan<- Node, errC chan<- error) {
	err := c.get(ctx, host, "nodes/"+node, func(c context.Context, resp *http.Response) error {
		var node Node
		err := json.NewDecoder(resp.Body).Decode(&node)
		if err != nil {
			return err
		}
		outC <- node
		return nil
	})
	if err != nil {
		errC <- err
	}
}
