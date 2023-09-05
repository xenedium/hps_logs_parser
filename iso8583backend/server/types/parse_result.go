package types

import (
	"time"

	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
)

/*
id: string;
name: string;
date: Date;
status: 'downloading' | 'parsing' | 'done' | 'error';
type: 'ssh' | 'upload'
messages: Message[];
*/

type ParseResult struct {
	Id       string                    `json:"id"`
	Name     string                    `json:"name"`
	Date     time.Time                 `json:"date"`
	Status   string                    `json:"status"`
	Type     string                    `json:"type"`
	Messages []*protocolBuffer.Message `json:"messages"`
}
