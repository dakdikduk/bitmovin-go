package models

import "github.com/dakdikduk/bitmovin-go/bitmovintypes"

type Detail struct {
	Date  *string                   `json:"date"`
	ID    *string                   `json:"id"`
	Type  bitmovintypes.MessageType `json:"type"`
	Text  *string                   `json:"text"`
	Field *string                   `json:"field"`
	Links []Link                    `json:"links"`
}
