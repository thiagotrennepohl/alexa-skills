package models

type JSON map[string]interface{}

//AlexaProgressiveRequest is the basic structure to send a feedback to the user while processing a skill
type AlexaProgressiveRequest struct {
	Header struct {
		RequestID string `json:"requestId"`
	} `json:"header"`
	Directive struct {
		Type   string `json:"type"`
		Speech string `json:"speech"`
	} `json:"directive"`
}

//AlexaIncomingPayload is the payload sent by amazon when a skill is invoked
type AlexaIncomingPayload struct {
	Version string `json:"version"`
	Request struct {
		DialogState string `json:"dialogstate"`
		Type        string `json:"type"`
		RequestID   string `json:"requestId"`
		Time        string `json:"timestamp"`
		Intent      struct {
			Name               string                 `json:"name"`
			ConfirmationStatus string                 `json:"confirmationstatus"`
			Slots              map[string]interface{} `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
	Context struct {
		System AlexaSkillContext `json:"system"`
	}
}

type Containers struct {
	ContainerNames []string
}

//AlexaSkillContext keeps authorization tokens and the applicationID, this structure is used when sending a progressive request
type AlexaSkillContext struct {
	Application struct {
		ApplicationID string `json:"applicationId"`
	}
	APIAccessToken string `json:"apiAccessToken"`
	APIEndPoint    string `json:"apiEndPoint"`
}

//AlexaResponse is the skill response that will be sent to the user
type AlexaResponse struct {
	Version          string `json:"version"`
	ShouldEndSession bool   `json:"shouldEndSession"`
	Response         struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
			SSML string `json:"ssml"`
		} `json:"outputSpeech"`
	} `json:"response"`
}

func (a *AlexaResponse) SetPlainTextErrorResponse(errorMessage string) {
	a.Response.OutputSpeech.Text = errorMessage
	a.Response.OutputSpeech.Type = "PlainText"
	a.ShouldEndSession = true
}

func (a *AlexaResponse) SetPlainTextResponse(message string) {
	a.Response.OutputSpeech.Text = message
	a.Response.OutputSpeech.Type = "PlainText"
	a.ShouldEndSession = true
}

func (a *AlexaResponse) SetSSMLResponse(message string) {
	a.Response.OutputSpeech.SSML = message
	a.Response.OutputSpeech.Type = "SSML"
	a.ShouldEndSession = true
}
