# Super simple telegram wrapper

# Starting the bot

## 1. Create custom commands
``` golang

telegram.COMMAND_MAP["!himom"] = func() string {
	return fmt.Sprint("Hello my child")
}

```

## 2. Start the bot
``` golang

ChannelID := 234234244
ApiKey := "YOUR-API-KEY"
err := telegram.Start(ChannelID, ApiKey, 0, 60)
if err != nil {
    log.Println(err)
}

```

## 3. Send a message 
``` golang

err := telegram.SendMessage("hi mom")
if err != nil {
    log.Println(err)
}

```