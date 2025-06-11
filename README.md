# English Vocabulary Telegram Bot

A simple Telegram bot that provides definitions for English words.

## Features

* Accepts English words from users via Telegram.
* Fetches definitions using an external dictionary API.
* Returns the definition(s) directly in chat.

## How It Works

* The bot listens for any text message.
* It extracts the first word from the message.
* Sends a request to the dictionary API.
* Parses the response and sends the definition(s) back to the user.

## Example Usage

**User:** `hello`

**Bot:**

```
"Hello!" or an equivalent greeting.
To greet with "hello".
A greeting (salutation) said when meeting someone or acknowledging someone’s arrival or presence.
A greeting used when answering the telephone.
A call for response if it is not clear if anyone is present or listening, or if a telephone conversation may have been disconnected.
Used sarcastically to imply that the person addressed or referred to has done something the speaker or writer considers to be foolish.
An expression of puzzlement or discovery.
```

---

## Acknowledgements

* [telebot.v4](https://github.com/tucnak/telebot) – Go Telegram Bot library
* [dictionaryapi.dev](https://dictionaryapi.dev) – Free dictionary API

