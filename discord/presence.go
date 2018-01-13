package discord

type Presence struct {
	User    *User     `json:"user"`
	Roles   []uint64  `json:"roles"`
	Game    *Activity `json:"activty"`
	GuildID uint64    `json:"guild_id"`
	Status  string    `json:"status"`
}

func NewPresence() *Presence {
	return &Presence{}
}

func (p *Presence) Update(status string) {
	// Update the presence.
	// talk to the discord api
}

func (p *Presence) String() string {
	return p.Status
}