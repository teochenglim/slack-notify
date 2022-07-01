package slack_notify

import (
  "fmt"
  "bytes"
  "os"
  "time"

  "encoding/json"
	"net/http"

  "github.com/hashicorp/go-retryablehttp"
  "github.com/teochenglim/slack-notify/pkg/slack-notify/config"

)

type Slack struct {
  Config  *config.Config
}

func New() *Slack {
  c := config.LoadConfigurations()
  s := &Slack{
    Config: c,
  }
	return s
}

type Message struct {
	Text            string       `json:"text"`
  UserName        string       `json:"username,omitempty"`
  IconURL         string       `json:"icon_url,omitempty"`
  IconEmoji       string       `json:"icon_emoji,omitempty"`
  Channel         string       `json:"channel,omitempty"`
  UnfurlLinks     bool         `json:"unfurl_links"`
  Attachments     []Attachment `json:"attachments,omitempty"`
  MarkdownSupport bool         `json:"mrkdwn,omitempty"`
}

type Attachment struct {
	Fallback string  `json:"fallback"`
	Pretext  string  `json:"pretext,omitempty"`
	Color    string  `json:"color,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value,omitempty"`
	Short bool   `json:short"`
}

func (s *Slack) SendMessage() (error) {

  if s.Config.SLACK_WEBHOOK == "" {
    fmt.Fprintf(os.Stderr, "Slack Webhook is required. ")
    os.Exit(1)
  }

  if s.Config.SLACK_MESSAGE == "" {
    fmt.Fprintf(os.Stderr, "Slack Message is required. ")
    os.Exit(1)
  }

  message := Message{
    UserName:        s.Config.SLACK_USERNAME,
    IconURL:         s.Config.SLACK_ICON,
    Channel:         s.Config.SLACK_CHANNEL,
    MarkdownSupport: s.Config.SLACK_MARKDOWN,
    Attachments: []Attachment{
      {
        Fallback: s.Config.SLACK_MESSAGE,
        Color:    s.Config.SLACK_COLOR,
        Fields: []Field{
          {
            Title: s.Config.SLACK_TITLE,
            Value: s.Config.SLACK_MESSAGE,
          },
        },
      },
    },
  }

  err := send(s.Config.SLACK_WEBHOOK, message, s.Config.SLACK_VERBOSE)

  if err != nil {
    fmt.Fprintf(os.Stderr, "Error sending message: %s\n", err)
    os.Exit(2)
  }
  return err
}

func send(webhookURL string, message Message, verbose bool) (error) {

  var (
  		response         *http.Response
  		request          *retryablehttp.Request
  )

  enc, err := json.Marshal(message)
	if err != nil {
		return err
	}

	b := bytes.NewBuffer(enc)

  client := retryablehttp.NewClient()
  client.RetryMax = 3
  client.RetryWaitMin = 1 * time.Second
  client.RetryWaitMax = 5 * time.Second
  client.HTTPClient.Timeout = 1 * time.Minute
  client.Logger = nil

  if request, err = retryablehttp.NewRequest("POST", webhookURL, b); err != nil {
		return err
	}
  request.Header.Set("Content-Type", "application/json")

  if response, err = client.Do(request); err != nil {
		return err
	}

  if verbose == true {
	   fmt.Println("Message sent! status: ", response.Status)
  }

	return nil
}
