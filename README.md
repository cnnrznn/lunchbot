# lunchbot

Lunchbot is a slack app for monitoring channels and setting user status based on text.

## Usage

```bash
ofl
# or
lunch
```

Set your status as "Out for lunch" for 60 minutes.

```bash
coffee
# or
ofc
```

Sets you status as "Out for coffee" for 30 minutes.

```bash
/imback
```

Clear your set status

## Installation - User

Once your admin has created the bot in your workspace with [example] name @Bot:

1. Invite the bot to channels by typing "@Bot"
2. Follow the authorization link provided by the admin to allow @Bot to make status updates to your account.

## Installation - Admin

To install this on a server, you need a few things:

- A certificate and private key file. I got mine from letsencyrpt
- `PEMFILE`, `KEYFILE` set to the paths of the files
- `LUNCHBOT_ID`, `LUNCHBOT_SECRET` set to your Slack bot's `client_id` and
  `client_secret`
  
You also need to create a slack app in your workspace via https://api.slack.com/ with the following features:

- Slash commands
  - https://{your-url}/imback
- Event subscriptions
  - path for handling events
  - Use `/challenge/` directory to handle the verification process
  - https://{your-url}/event
- Bot user "Bots"
- Permissions
  - https://{your-url}/authorize
  - Scopes
    - `channels:history`
    - `groups:history`
    - `commands`
  - User Token Scopes
    - `users.profile:write`
