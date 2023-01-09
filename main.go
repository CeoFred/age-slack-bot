package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/shomali11/slacker"
	"strconv"
)

func main() {
		os.Setenv("SLACK_BOT_TOKEN", "") // oatuh token
		os.Setenv("SLACK_APP_TOKEN", "") // socket mode app token

		bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

		go printCommandEvents(bot.CommandEvents())

		bot.Command("my yob is <year>", &slacker.CommandDefinition{
			Description: "yob calculator",
			Examples: []string{"my yob is 2000"},
			Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter){
					year := request.Param("year")
					yob, err := strconv.Atoi(year)
					if err!= nil {
						log.Fatal(err)
            return
          }
					age := 2022-yob
					r := fmt.Sprintf("age is %d", age)
					response.Reply(r)
			},
		})
 
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := bot.Listen(ctx)

    if err!= nil {
			log.Fatal(err)
		}

}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent){
	for event := range analyticsChannel {
    fmt.Printf("Command: %s\n", event.Command)
		fmt.Printf("Timestamp: %s\n", event.Timestamp)
		fmt.Printf("Parameters: %s\n", event.Parameters)
		fmt.Printf("Event: %s\n", event.Event)
		fmt.Println("--------------------------------")
	}
}