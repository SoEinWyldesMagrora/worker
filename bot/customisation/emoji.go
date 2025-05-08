package customisation

import (
	"fmt"

	"github.com/rxdn/gdl/objects"
	"github.com/rxdn/gdl/objects/guild/emoji"
)

type CustomEmoji struct {
	Name     string
	Id       uint64
	Animated bool
}

func NewCustomEmoji(name string, id uint64, animated bool) CustomEmoji {
	return CustomEmoji{
		Name: name,
		Id:   id,
	}
}

func (e CustomEmoji) String() string {
	if e.Animated {
		return fmt.Sprintf("<a:%s:%d>", e.Name, e.Id)
	} else {
		return fmt.Sprintf("<:%s:%d>", e.Name, e.Id)
	}
}

func (e CustomEmoji) BuildEmoji() *emoji.Emoji {
	return &emoji.Emoji{
		Id:       objects.NewNullableSnowflake(e.Id),
		Name:     e.Name,
		Animated: e.Animated,
	}
}

var (
	EmojiId         = NewCustomEmoji("id", 1370090425724960881, false)
	EmojiOpen       = NewCustomEmoji("open", 1370090940081111201, false)
	EmojiOpenTime   = NewCustomEmoji("opentime", 1370090910024466464, false)
	EmojiClose      = NewCustomEmoji("close", 1370089135821553824, false)
	EmojiCloseTime  = NewCustomEmoji("closetime", 1370090591324475422, false)
	EmojiReason     = NewCustomEmoji("reason", 1370090561134137364, false)
	EmojiSubject    = NewCustomEmoji("subject", 1370090518754627725, false)
	EmojiTranscript = NewCustomEmoji("transcript", 1370090483266883614, false)
	EmojiClaim      = NewCustomEmoji("claim", 1370090743884156960, false)
	EmojiPanel      = NewCustomEmoji("panel", 1370090706542137458, false)
	EmojiRating     = NewCustomEmoji("rating", 1370090672937373706, false)
	EmojiStaff      = NewCustomEmoji("staff", 1370090631229341907, false)
	EmojiThread     = NewCustomEmoji("thread", 1370090776012263497, false)
	EmojiBulletLine = NewCustomEmoji("bulletline", 1370090839979589672, false)
	EmojiPatreon    = NewCustomEmoji("patreon", 1370090807469674536, false)
	EmojiDiscord    = NewCustomEmoji("discord", 1370090869398573076, false)
	//EmojiTime       = NewCustomEmoji("time", 974006684622159952, false)
)

// PrefixWithEmoji Useful for whitelabel bots
func PrefixWithEmoji(s string, emoji CustomEmoji, includeEmoji bool) string {
	if includeEmoji {
		return fmt.Sprintf("%s %s", emoji, s)
	} else {
		return s
	}
}
