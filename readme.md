#### 1. What does this tool do?

  > This lightweight tools is designed to send slack message written in Golang without any dependancy.
  >
  > You need to setup SLACK_WEBHOOK which you need to get from slack website.
  > Slack Webhook is a authenticated URL. Detail explaination is here https://api.slack.com/messaging/webhooks
  >
  > The minimum inputs are SLACK_WEBHOOK and SLACK_MESSAGE which is the text you want to send.

#### 2. What are the configuration option and priority?

  2.1 using config file, file name is "slack-notify.yaml" (Priority lowest)

    2.1.1 config file search from running directory (./config/slack-notify.yaml)
    2.1.2 config file search from running directory (./slack-notify.yaml)
    2.1.2 config file search from running directory ($HOME/slack-notify.yaml)

  ```yaml
  SLACK_WEBHOOK: "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  SLACK_ICON: "https://slack.com/img/icons/app-57.png"
  SLACK_CHANNEL: "#[chatroom_name]"
  SLACK_TITLE: "Title text"
  SLACK_MESSAGE: "Hello slack"
  SLACK_COLOR: "#00FFFF"
  SLACK_USERNAME: "CLTEO"
  SLACK_VERBOSE: false

  ```

  2.2 using env (Priority higher than config file)

  ```yaml
  SLACK_WEBHOOK="https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
  SLACK_ICON="https://slack.com/img/icons/app-57.png"
  SLACK_CHANNEL="#[chatroom_name]"
  SLACK_TITLE="Title text"
  SLACK_MESSAGE="Hello slack"
  SLACK_COLOR="#00FFFF"
  SLACK_USERNAME="CLTEO"
  SLACK_VERBOSE=false

  ```

  2.3 using cli options (Priority highest)

  ```
  $ slack-notify -h
  slack is CLI client tools to send slack message by given simple inputs.
  The main input is SLACK_WEBHOOK which you need to get from slack website.
  Web hook explain here https://api.slack.com/messaging/webhooks
  The minimum inputs are SLACK_WEBHOOK and SLACK_MESSAGE which is the text you want to send.

  Usage:
    slack-notify [flags]

  Flags:
    -c, --channel string    Slack channel name (default "#chatops")
    -o, --color string      Slack message color #7CD197
    -h, --help              help for slack
    -i, --icon string       Slack icon URL https://slack.com/img/icons/app-57.png
    -m, --message string    Slack message text (default "My slack message")
    -t, --title string      Slack message title
    -u, --username string   Slack message display name
    -v, --verbose           verbose output
    -w, --webook string     Slack webhook url https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX

  ```

#### 3. Usage

  ##### Running `slack-notify` in a shell prompt goes like this:

  ```console
  $ export SLACK_WEBHOOK=https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/xxxxxxxx
  $ SLACK_MESSAGE="hello" slack-notify

  ```

  ##### Running the Docker container goes like this:

  ```console
  $ export SLACK_WEBHOOK=https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/xxxxxxxx
  $ docker run -e SLACK_WEBHOOK=$SLACK_WEBHOOK -e SLACK_MESSAGE="hello" -e SLACK_CHANNEL=#chatops teochenglim/slack-notify
  ```

  ##### Gitlab **.gitlab-ci.yml**

  >You may to setup variables via [CI/CD variables]([https://gitlab.com/help/ci/variables/README#variables)

  ```yaml
  stages:
  - notify

  variables:
  SLACK_WEBHOOK: https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/xxxxxxxx
  SLACK_CHANNEL: deploy

  notify:
    stage: notify
    image: teochenglim/slack-notify:latest
    script:
      - 'SLACK_MESSAGE="Message" slack-notify'

  ```

  ##### Environment Variables

  ```shell
  # The Slack-assigned webhook
  SLACK_WEBHOOK=https://hooks.slack.com/services/Txxxxxx/Bxxxxxx/xxxxxxxx
  # A URL to an icon
  SLACK_ICON=http://example.com/icon.png
  # The channel to send the message to (if omitted, use Slack-configured default)
  SLACK_CHANNEL=example
  # The title of the message
  SLACK_TITLE="Hello World"
  # The body of the message
  SLACK_MESSAGE="Today is a fine day"
  # RGB color to for message formatting. (Slack determines what is colored by this)
  SLACK_COLOR="#efefef"
  # The name of the sender of the message. Does not need to be a "real" username
  SLACK_USERNAME="Bot"
  ```

#### 4. docker build

  ##### Build It

  Compile:

  ```
  make build
  ```

  Publish to DockerHub

  ```
  make docker-build docker-push
  ```

#### 5. Special thanks

  > Jeremy Liu https://github.com/Ksloveyuan

  > https://github.com/krom/slack-notify/

  > https://github.com/easonlin404/go-slack/
