package hashtreedb

import (
	"bytes"
	"encoding/binary"
)

//var FindValueOffsetItemSize = uint32(5) // 索引项宽度

type FindValueOffsetItem struct {

	// index
	Type               uint8  // 0:默认空 1:枝 2:叶 3:被删除的叶
	ValueSegmentOffset uint32 // value Segment 位置

	IndexMenuSelfSegmentOffset uint32 // 菜单段位置
	IndexItemSelfAlignment     uint32 // 菜单项位置

	// data
	ValueMark byte
	ValueKey  []byte
	ValueBody []byte //

	// opt cnf
	searchHash  []byte
	searchCount int // 搜索次数，从0开始
}

func NewFindValueOffsetItem(ty uint8, ValueSegmentOffset uint32) *FindValueOffsetItem {
	return &FindValueOffsetItem{
		Type:               ty,
		ValueSegmentOffset: ValueSegmentOffset,
	}
}

func (this *FindValueOffsetItem) IncompleteCopy() *FindValueOffsetItem {
	return &FindValueOffsetItem{
		Type:                       this.Type,
		ValueSegmentOffset:         this.ValueSegmentOffset,
		IndexMenuSelfSegmentOffset: this.IndexMenuSelfSegmentOffset,
		IndexItemSelfAlignment:     this.IndexItemSelfAlignment,
		ValueKey:                   nil,
		ValueBody:                  nil,
		searchHash:                 nil,
	}
}

func (this *FindValueOffsetItem) Parse(buf []byte, seek uint32) error {
	this.Type = uint8(buf[seek])
	this.ValueSegmentOffset = binary.BigEndian.Uint32(buf[seek+1 : seek+5])
	return nil
}

func (this *FindValueOffsetItem) Serialize() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte{this.Type})
	var byt1 = make([]byte, 4)
	binary.BigEndian.PutUint32(byt1, this.ValueSegmentOffset)
	buffer.Write(byt1)
	return buffer.Bytes()
}
