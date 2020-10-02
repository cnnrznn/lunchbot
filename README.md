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

## Installation

To install this on a server, you need a few things:

- A certificate and private key file. I got mine from letsencyrpt
- `PEMFILE`, `KEYFILE` set to the paths of the files
- `LUNCHBOT_ID`, `LUNCHBOT_SECRET` set to your Slack bot's `client_id` and
  `client_secret`
