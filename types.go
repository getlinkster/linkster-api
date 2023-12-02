package main

type ConfigInfo struct {
	Port            int    `mapstructure:"PORT"`
	EventSchemaId   string `mapstructure:"EVENT_SCHEMA_ID"`
	ProfileSchemaId string `mapstructure:"PROFILE_SCHEMA_ID"`
	IssuerUname     string `mapstructure:"ISSUER_UNAME"`
	IssuerPass      string `mapstructure:"ISSUER_PASS"`
	IssuerApi       string `mapstructure:"ISSUER_API"`
}

type NewEventRequest struct {
	ValidUntil string `json:"valid-until,omitempty"`
	ClaimLimit int    `json:"claim-limit,omitempty"`
	EventInfo  Event  `json:"event"`
}

type Event struct {
	Name           string `json:"event-name"`
	Location       string `json:"event-location"`
	Date           string `json:"event-date"`
	AdditionalInfo string `json:"additional-info"`
}

type Profile struct {
	Name       string `json:"name"`
	Wallet     string `json:"wallet"`
	Profession string `json:"profession"`
	Company    string `json:"company"`
	Telegram   string `json:"telegram"`
	/*
		Linkedin     string `json:"linkedin"`
		Location     string `json:"location"`
	*/
}

type EventCredentialLinkRequest struct {
	SchemaId            string `json:"schemaID"`
	ClaimLinkExpiration string `json:"claimLinkExpiration,omitempty"`
	LimitedClaims       int    `json:"limitedClaims,omitempty"`
	SignatureProof      bool   `json:"signatureProof"`
	MtProof             bool   `json:"mtProof"`
	CredentialSubject   Event  `json:"credentialSubject"`
}

type ProfileCredentialLinkRequest struct {
	SchemaId          string  `json:"schemaID"`
	SignatureProof    bool    `json:"signatureProof"`
	MtProof           bool    `json:"mtProof"`
	CredentialSubject Profile `json:"credentialSubject"`
}

type CreateLinkResponse struct {
	Id string `json:"id"`
}

type QrCodeRequestResponse struct {
	QrCode    string `json:"qrCode"`
	SessionID string `json:"sessionID"`
}
