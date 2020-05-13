package api

import (
	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/utils/httputil"
	"github.com/diamondburned/arikawa/utils/json"
)

func (c *Client) Member(guildID, userID discord.Snowflake) (*discord.Member, error) {
	var m *discord.Member
	err := c.RequestJSON(&m, "GET", EndpointGuilds+guildID.String()+"/members/"+userID.String())
	if err != nil {
		return nil, err
	}

	// Set the guild id on the member, because Discord doesn't.
	if m != nil {
		m.GuildID = guildID
	}

	return m, nil
}

// Members returns members until it reaches max. This function automatically
// paginates, meaning the normal 1000 limit is handled internally. Max can be 0,
// in which the function will try and fetch everything.
func (c *Client) Members(guildID discord.Snowflake, max uint) ([]discord.Member, error) {
	var members []discord.Member
	var after discord.Snowflake = 0

	const hardLimit int = 1000

	for fetch := uint(hardLimit); max > 0; fetch = uint(hardLimit) {
		if max > 0 {
			if fetch > max {
				fetch = max
			}
			max -= fetch
		}

		m, err := c.MembersAfter(guildID, after, fetch)
		if err != nil {
			return members, err
		}
		members = append(members, m...)

		// There aren't any to fetch, even if this is less than max.
		if len(members) < hardLimit {
			break
		}

		after = members[hardLimit-1].User.ID
	}

	return members, nil
}

// MembersAfter returns a list of all guild members, from 1-1000 for limits. The
// default limit is 1 and the maximum limit is 1000.
func (c *Client) MembersAfter(guildID, after discord.Snowflake, limit uint) ([]discord.Member, error) {
	if limit == 0 {
		limit = 1
	}

	if limit > 1000 {
		limit = 1000
	}

	var param struct {
		After discord.Snowflake `schema:"after,omitempty"`

		Limit uint `schema:"limit"`
	}

	param.Limit = limit
	param.After = after

	var mems []discord.Member
	return mems, c.RequestJSON(
		&mems, "GET",
		EndpointGuilds+guildID.String()+"/members",
		httputil.WithSchema(c, param),
	)
}

// AnyMemberData, all fields are optional.
type AnyMemberData struct {
	Nick json.OptionString `json:"nick,omitempty"`
	Mute json.OptionBool   `json:"mute,omitempty"`
	Deaf json.OptionBool   `json:"deaf,omitempty"`

	Roles *[]discord.Snowflake `json:"roles,omitempty"`

	// Only for ModifyMember, requires MOVE_MEMBER
	VoiceChannel discord.Snowflake `json:"channel_id,omitempty"`
}

// AddMember requires access(Token).
func (c *Client) AddMember(guildID, userID discord.Snowflake, token string,
	data AnyMemberData) (*discord.Member, error) {

	// VoiceChannel doesn't belong here.
	data.VoiceChannel = discord.NullSnowflake

	var param struct {
		Token string `json:"access_token"`
		AnyMemberData
	}

	param.Token = token
	param.AnyMemberData = data

	var mem *discord.Member

	return mem, c.RequestJSON(
		&mem, "PUT",
		EndpointGuilds+guildID.String()+"/members/"+userID.String(),
		httputil.WithJSONBody(param),
	)
}

func (c *Client) ModifyMember(guildID, userID discord.Snowflake, data AnyMemberData) error {
	return c.FastRequest(
		"PATCH",
		EndpointGuilds+guildID.String()+"/members/"+userID.String(),
		httputil.WithJSONBody(data),
	)
}

// PruneCount returns the number of members that would be removed in a prune
// operation. Requires KICK_MEMBERS. Days must be 1 or more, default 7.
func (c *Client) PruneCount(guildID discord.Snowflake, days uint) (uint, error) {
	if days == 0 {
		days = 7
	}

	var param struct {
		Days uint `schema:"days"`
	}

	param.Days = days

	var resp struct {
		Pruned uint `json:"pruned"`
	}

	return resp.Pruned, c.RequestJSON(
		&resp, "GET",
		EndpointGuilds+guildID.String()+"/prune",
		httputil.WithSchema(c, param),
	)
}

// Prune returns the number of members that is removed. Requires KICK_MEMBERS.
// Days must be 1 or more, default 7.
func (c *Client) Prune(guildID discord.Snowflake, days uint) (uint, error) {
	if days == 0 {
		days = 7
	}

	var param struct {
		Count    uint `schema:"count"`
		RetCount bool `schema:"compute_prune_count"`
	}

	param.Count = days
	param.RetCount = true // maybe expose this later?

	var resp struct {
		Pruned uint `json:"pruned"`
	}

	return resp.Pruned, c.RequestJSON(
		&resp, "POST",
		EndpointGuilds+guildID.String()+"/prune",
		httputil.WithSchema(c, param),
	)
}

// Kick requires KICK_MEMBERS.
func (c *Client) Kick(guildID, userID discord.Snowflake) error {
	return c.FastRequest("DELETE", EndpointGuilds+guildID.String()+"/members/"+userID.String())
}

func (c *Client) Bans(guildID discord.Snowflake) ([]discord.Ban, error) {
	var bans []discord.Ban
	return bans, c.RequestJSON(&bans, "GET", EndpointGuilds+guildID.String()+"/bans")
}

func (c *Client) GetBan(guildID, userID discord.Snowflake) (*discord.Ban, error) {
	var ban *discord.Ban
	return ban, c.RequestJSON(&ban, "GET", EndpointGuilds+guildID.String()+"/bans/"+userID.String())
}

// Ban requires the BAN_MEMBERS permission. Days is the days back for Discord
// to delete the user's message, maximum 7 days.
func (c *Client) Ban(guildID, userID discord.Snowflake, days uint, reason string) error {
	if days > 7 {
		days = 7
	}

	var param struct {
		DeleteDays uint   `schema:"delete_message_days,omitempty"`
		Reason     string `schema:"reason,omitempty"`
	}

	param.DeleteDays = days
	param.Reason = reason

	return c.FastRequest(
		"PUT",
		EndpointGuilds+guildID.String()+"/bans/"+userID.String(),
		httputil.WithSchema(c, param),
	)
}

// Unban requires BAN_MEMBERS.
func (c *Client) Unban(guildID, userID discord.Snowflake) error {
	return c.FastRequest("DELETE", EndpointGuilds+guildID.String()+"/bans/"+userID.String())
}
