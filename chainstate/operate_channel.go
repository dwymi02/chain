package chainstate

import (
	"github.com/hacash/core/fields"
	"github.com/hacash/core/stores"
)

//
func (cs *ChainState) Channel(channelid fields.ChannelId) *stores.Channel {
	query, e1 := cs.channelDB.CreateNewQueryInstance(channelid)
	if e1 != nil {
		return nil // error
	}
	defer query.Destroy()
	vdatas, e2 := query.Find()
	if e2 != nil {
		return nil // error
	}
	if vdatas == nil {
		if cs.base != nil {
			return cs.base.Channel(channelid) // check base
		} else {
			return nil // not find
		}
	}
	if len(vdatas) == 0 {
		return nil // error
	}
	var stoitem stores.Channel
	_, e3 := stoitem.Parse(vdatas, 0)
	if e3 != nil {
		return nil // error
	}
	// return ok
	return &stoitem
}

//
func (cs *ChainState) ChannelCreate(channel_id fields.ChannelId, channel *stores.Channel) error {
	query, e1 := cs.channelDB.CreateNewQueryInstance(channel_id)
	if e1 != nil {
		return e1 // error
	}
	defer query.Destroy()
	stodatas, e3 := channel.Serialize()
	if e3 != nil {
		return e3 // error
	}
	e4 := query.Save(stodatas)
	if e4 != nil {
		return e4 // error
	}
	// ok
	return nil
}

//
func (cs *ChainState) ChannelUpdate(channel_id fields.ChannelId, channel *stores.Channel) error {
	return cs.ChannelCreate(channel_id, channel)
}

//
func (cs *ChainState) ChannelDelete(channel_id fields.ChannelId) error {
	query, e1 := cs.channelDB.CreateNewQueryInstance(channel_id)
	if e1 != nil {
		return e1 // error
	}
	defer query.Destroy()
	e2 := query.Delete()
	if e2 != nil {
		return e2 // error
	}
	// ok
	return nil
}
