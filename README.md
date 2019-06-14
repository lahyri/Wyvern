# Wyvern

### This is a Discord Bot specially crafted as a DnD compendium/helper

## Installing

To install Wyvern in your Discord Server you'll need the following:

* [Golang](https://golang.org/)
* [Discord](https://discordapp.com/) (duuh)

After that, you should clone this repository and run the following command:
```
$go get github.com/bwmarrin/discordgo
```
And your own version of Wyvern is ready to roll.

### Observation

Since I'm not planning to host this permanently, You'll have to [create your own bot account](https://discordpy.readthedocs.io/en/rewrite/discord.html), save the generated token to a .env file in the main directory and run the project everytime you want to use it.

## Running

Open your code directory in the terminal and run:
```
$go run main.go
```
If you've done this correctly, you bot should appear online in the server.

## Commands

To see all the commands available for this bot, just type the command ` !help ` in your discord server, and it shall return all the commands and their description of usage.

## Contribuiting

Feel free to use it and modify it as you want. If you'd like to share any new functionality, just branch the rep and send me a pull request with a description of the alteration or functionality you've added.

## To Do

* Everything besides the dice roll

## Developed by
* [Lahyri Aurbach](https://github.com/lahyri)

## Special Thanks
* [bwmarrin](https://github.com/bwmarrin) for his wonderful [DiscordGo](https://github.com/bwmarrin/discordgo) bindings
* My RPG group for inspiring me to create this