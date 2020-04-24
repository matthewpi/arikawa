package voice

// OPCode 2
// https://discordapp.com/developers/docs/topics/voice-connections#establishing-a-voice-websocket-connection-example-voice-ready-payload
type ReadyEvent struct {
	SSRC        uint32   `json:"ssrc"`
	IP          string   `json:"ip"`
	Port        int      `json:"port"`
	Modes       []string `json:"modes"`
	Experiments []string `json:"experiments"`

	// From Discord's API Docs:
	//
	// `heartbeat_interval` here is an erroneous field and should be ignored.
	// The correct `heartbeat_interval` value comes from the Hello payload.

	// HeartbeatInterval discord.Milliseconds `json:"heartbeat_interval"`
}

// OPCode 4
// https://discordapp.com/developers/docs/topics/voice-connections#establishing-a-voice-udp-connection-example-session-description-payload
type SessionDescriptionEvent struct {
	Mode      string   `json:"mode"`
	SecretKey [32]byte `json:"secret_key"`
}

// OPCode 5
type SpeakingEvent SpeakingData

// OPCode 6
// https://discordapp.com/developers/docs/topics/voice-connections#heartbeating-example-heartbeat-ack-payload
type HeartbeatACKEvent uint64

// OPCode 8
// https://discordapp.com/developers/docs/topics/voice-connections#heartbeating-example-hello-payload-since-v3
type HelloEvent struct {
	HeartbeatInterval float64 `json:"heartbeat_interval"`
}

// OPCode 9
// https://discordapp.com/developers/docs/topics/voice-connections#resuming-voice-connection-example-resumed-payload
type ResumedEvent struct{}