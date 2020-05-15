package discord

type Guild struct {
	ID     Snowflake `json:"id,string"`
	Name   string    `json:"name"`
	Icon   Hash      `json:"icon"`
	Splash Hash      `json:"splash,omitempty"` // server invite bg

	Owner   bool      `json:"owner,omitempty"` // self is owner
	OwnerID Snowflake `json:"owner_id,string"`

	Permissions Permissions `json:"permissions,omitempty"`

	VoiceRegion string `json:"region"`

	AFKChannelID Snowflake `json:"afk_channel_id,string,omitempty"`
	AFKTimeout   Seconds   `json:"afk_timeout"`

	Embeddable     bool      `json:"embed_enabled,omitempty"`
	EmbedChannelID Snowflake `json:"embed_channel_id,string,omitempty"`

	Verification   Verification   `json:"verification_level"`
	Notification   Notification   `json:"default_message_notifications"`
	ExplicitFilter ExplicitFilter `json:"explicit_content_filter"`

	Roles    []Role         `json:"roles"`
	Emojis   []Emoji        `json:"emojis"`
	Features []GuildFeature `json:"guild_features"`

	MFA MFALevel `json:"mfa"`

	AppID Snowflake `json:"application_id,string,omitempty"`

	Widget bool `json:"widget_enabled,omitempty"`

	WidgetChannelID Snowflake `json:"widget_channel_id,string,omitempty"`
	SystemChannelID Snowflake `json:"system_channel_id,string,omitempty"`

	// It's DefaultMaxPresences when MaxPresences is 0.
	MaxPresences uint64 `json:"max_presences,omitempty"`
	MaxMembers   uint64 `json:"max_members,omitempty"`

	VanityURLCode string `json:"vanity_url_code,omitempty"`
	Description   string `json:"description,omitempty"`
	Banner        Hash   `json:"banner,omitempty"`

	NitroBoost    NitroBoost `json:"premium_tier"`
	NitroBoosters uint64     `json:"premium_subscription_count,omitempty"`

	// Defaults to en-US, only set if guild has DISCOVERABLE
	PreferredLocale string `json:"preferred_locale"`

	// Only presented if WithCounts is true.
	ApproximateMembers   uint64 `json:"approximate_member_count,omitempty"`
	ApproximatePresences uint64 `json:"approximate_presence_count,omitempty"`
}

// IconURL returns the URL to the guild icon and auto detects a suitable type.
// An empty string is returned if there's no icon.
func (g Guild) IconURL() string {
	return g.IconURLWithType(AutoImage)
}

// IconURLWithType returns the URL to the guild icon using the passed ImageType. An
// empty string is returned if there's no icon.
//
// Supported ImageTypes: PNG, JPEG, WebP, GIF
func (g Guild) IconURLWithType(t ImageType) string {
	if g.Icon == "" {
		return ""
	}

	return "https://cdn.discordapp.com/icons/" + g.ID.String() + "/" + t.format(g.Icon)
}

// BannerURL returns the URL to the banner, which is the image on top of the
// channels list. This will always return a link to a PNG file.
func (g Guild) BannerURL() string {
	return g.BannerURLWithType(PNGImage)
}

// BannerURLWithType returns the URL to the banner, which is the image on top of the
// channels list using the passed image type.
//
// Supported ImageTypes: PNG, JPEG, WebP
func (g Guild) BannerURLWithType(t ImageType) string {
	if g.Banner == "" {
		return ""
	}

	return "https://cdn.discordapp.com/banners/" +
		g.ID.String() + "/" + t.format(g.Banner)
}

// SplashURL returns the URL to the guild splash, which is the invite page's
// background. This will always return a link to a PNG file.
func (g Guild) SplashURL() string {
	if g.Splash == "" {
		return ""
	}

	return "https://cdn.discordapp.com/splashes/" +
		g.ID.String() + "/" + g.Splash + ".png"
}

// SplashURLWithType returns the URL to the guild splash, which is the invite page's
// background, using the passed ImageType.
//
// Supported ImageTypes: PNG, JPEG, WebP
func (g Guild) SplashURLWithType(t ImageType) string {
	if g.Splash == "" {
		return ""
	}

	return "https://cdn.discordapp.com/splashes/" +
		g.ID.String() + "/" + t.format(g.Splash)
}

// https://discord.com/developers/docs/resources/guild#guild-preview-object
type GuildPreview struct {
	// ID is the guild id.
	ID Snowflake `json:"id"`
	// Name is the guild name (2-100 characters).
	Name string `json:"name"`

	// Icon is the icon hash.
	Icon Hash `json:"icon"`
	// Splash is the splash hash.
	Splash Hash `json:"splash"`
	// DiscoverySplash is the discovery splash hash.
	DiscoverySplash Hash `json:"discovery_splash"`

	// Emojis are the custom guild emojis.
	Emojis []Emoji `json:"emojis"`
	// Features are the enabled guild features.
	Features []GuildFeature `json:"guild_features"`

	// ApproximateMembers is the approximate number of members in this guild.
	ApproximateMembers uint64 `json:"approximate_member_count"`
	// ApproximatePresences is the approximate number of online members in this
	// guild.
	ApproximatePresences uint64 `json:"approximate_presence_count"`

	// Description is the description for the guild.
	Description string `json:"description,omitempty"`
}

// IconURL returns the URL to the guild icon and auto detects a suitable type.
// An empty string is returned if there's no icon.
func (g GuildPreview) IconURL() string {
	return g.IconURLWithType(AutoImage)
}

// IconURLWithType returns the URL to the guild icon using the passed ImageType. An
// empty string is returned if there's no icon.
//
// Supported ImageTypes: PNG, JPEG, WebP, GIF
func (g GuildPreview) IconURLWithType(t ImageType) string {
	if g.Icon == "" {
		return ""
	}

	return "https://cdn.discordapp.com/icons/" + g.ID.String() + "/" + t.format(g.Icon)
}

// SplashURL returns the URL to the guild splash, which is the invite page's
// background. This will always return a link to a PNG file.
func (g GuildPreview) SplashURL() string {
	if g.Splash == "" {
		return ""
	}

	return "https://cdn.discordapp.com/splashes/" +
		g.ID.String() + "/" + g.Splash + ".png"
}

// SplashURLWithType returns the URL to the guild splash, which is the invite page's
// background, using the passed ImageType.
//
// Supported ImageTypes: PNG, JPEG, WebP
func (g GuildPreview) SplashURLWithType(t ImageType) string {
	if g.Splash == "" {
		return ""
	}

	return "https://cdn.discordapp.com/splashes/" +
		g.ID.String() + "/" + t.format(g.Splash)
}

// DiscoverySplashURL returns the URL to the guild splash, which is the invite page's
// background. This will always return a link to a PNG file.
func (g GuildPreview) DiscoverySplashURL() string {
	if g.Splash == "" {
		return ""
	}

	return "https://cdn.discordapp.com/splashes/" +
		g.ID.String() + "/" + g.Splash + ".png"
}

// DiscoverySplashURLWithType returns the URL to the guild splash, which is the invite page's
// background, using the passed ImageType.
//
// Supported ImageTypes: PNG, JPEG, WebP
func (g GuildPreview) DiscoverySplashURLWithType(t ImageType) string {
	if g.Splash == "" {
		return ""
	}

	return "https://cdn.discordapp.com/splashes/" +
		g.ID.String() + "/" + t.format(g.Splash)
}

type Role struct {
	ID   Snowflake `json:"id,string"`
	Name string    `json:"name"`

	Color    Color `json:"color"`
	Hoist    bool  `json:"hoist"` // if the role is separated
	Position int   `json:"position"`

	Permissions Permissions `json:"permissions"`

	Managed     bool `json:"managed"`
	Mentionable bool `json:"mentionable"`
}

func (r Role) Mention() string {
	return "<&" + r.ID.String() + ">"
}

type Presence struct {
	User    User        `json:"user"`
	RoleIDs []Snowflake `json:"roles"`

	// These fields are only filled in gateway events, according to the
	// documentation.

	Nick    string    `json:"nick"`
	GuildID Snowflake `json:"guild_id"`

	PremiumSince Timestamp `json:"premium_since,omitempty"`

	Game       *Activity  `json:"game"`
	Activities []Activity `json:"activities"`

	Status       Status `json:"status"`
	ClientStatus struct {
		Desktop Status `json:"desktop,omitempty"`
		Mobile  Status `json:"mobile,omitempty"`
		Web     Status `json:"web,omitempty"`
	} `json:"client_status"`
}

type Member struct {
	User    User        `json:"user"`
	Nick    string      `json:"nick,omitempty"`
	RoleIDs []Snowflake `json:"roles"`

	Joined       Timestamp `json:"joined_at"`
	BoostedSince Timestamp `json:"premium_since,omitempty"`

	Deafened bool      `json:"deaf"`
	Muted    bool      `json:"mute"`
	GuildID  Snowflake `json:"guild_id"`
}

func (m Member) Mention() string {
	return "<@!" + m.User.ID.String() + ">"
}

func (m *Member) IsSelf(user User) bool {
	return m.User.ID.String() == user.ID.String()
}

type Ban struct {
	Reason string `json:"reason,omitempty"`
	User   User   `json:"user"`
}

type Integration struct {
	ID   Snowflake `json:"id"`
	Name string    `json:"name"`
	Type Service   `json:"type"`

	Enabled bool `json:"enabled"`
	Syncing bool `json:"syncing"`

	// used for subscribers
	RoleID Snowflake `json:"role_id"`

	ExpireBehavior    ExpireBehavior `json:"expire_behavior"`
	ExpireGracePeriod int            `json:"expire_grace_period"`

	User    User `json:"user"`
	Account struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"account"`

	SyncedAt Timestamp `json:"synced_at"`
}

type GuildEmbed struct {
	Enabled   bool      `json:"enabled"`
	ChannelID Snowflake `json:"channel_id,omitempty"`
}

// DefaultMemberColor is the color used for members without colored roles.
var DefaultMemberColor Color = 0x0

func MemberColor(guild Guild, member Member) Color {
	var c = DefaultMemberColor
	var pos int

	for _, r := range guild.Roles {
		for _, mr := range member.RoleIDs {
			if mr != r.ID {
				continue
			}

			if r.Color > 0 && r.Position > pos {
				c = r.Color
				pos = r.Position
			}
		}
	}

	return c
}
