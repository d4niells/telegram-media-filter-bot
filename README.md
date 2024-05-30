# Telegram Media Filter Bot

A Telegram bot written in Go that filters specific types of messages (photos, videos, documents, text) in group chats. This bot allows group administrators to dynamically set the type of messages to be filtered using commands.

## Features

- **Dynamic Filtering:** Administrators can set the bot to filter photos, videos, documents, or text messages using commands.
- **Group Compatibility:** Designed to work seamlessly in Telegram group chats.
- **Customizable Responses:** Sends a warning message whenever a filtered message is deleted.

## Commands

- `/setfilter photo` - Set the bot to filter photo messages.
- `/setfilter video` - Set the bot to filter video messages.
- `/setfilter document` - Set the bot to filter document messages.
- `/setfilter text` - Set the bot to filter text messages.
- `/setfilter none` - Disable message filtering.

## Setup

1. **Get a Bot Token from BotFather:**
    - Open Telegram and search for BotFather.
    - Start a chat with BotFather and use the `/newbot` command to create a new bot.
    - Follow the prompts to name your bot and get the bot token.

2. **Clone the Repository:**
   ```sh
   git clone https://github.com/yourusername/telegram-media-filter-bot.git
   cd telegram-media-filter-bot
   ```
   
3. **Set Up Environment Variables:**
    - Create a .env file in the project directory.
    - Add your Telegram bot token to the .env file:
      ```sh
      TELEGRAM_BOT_TOKEN=your_bot_token
      ```

4. **Run the Bot:**
    ```sh
    go run main.go
    ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with your improvements.

## License

This project is licensed under the MIT License.
