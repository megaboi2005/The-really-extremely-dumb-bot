package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"strings"
	"github.com/bwmarrin/discordgo"
    "math/rand"
    "time"
	"strconv"
)


var (
	Token string
	Users = []string {}
)
type ReactData struct {
	input string
	reaction string
}
func gennumber(min,max int) int{
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
func init() {

	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func contains(slice []string, str string) bool{
	for _,value := range slice {
		if value == str {
			return true
		}
	}
	return false
}
func main() {

	
	dg, err := discordgo.New("Bot " + "token here")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	
	dg.AddHandler(messageCreate)
	dg.AddHandler(Ready)
	
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	
	dg.Close()
}

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	s.UpdateGameStatus(0, "Just be nice and say please")
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	var splitmsg = strings.Split(strings.ToLower(m.Content)," ")
	if splitmsg[0] == "hey" && splitmsg[1] == "bot" {
		for _, value := range Users {
			if m.Author.ID == value { 
				s.ChannelMessageSend(m.ChannelID,"hey what!!!???")
				return 
			}
		}
		s.ChannelMessageSend(m.ChannelID,"hey, whats up?")
		Users = append(Users, m.Author.ID)
		fmt.Println(Users)
	}

	if splitmsg[0] == "please" {
		if contains(Users,m.Author.ID) {} else {
			s.ChannelMessageSend(m.ChannelID,"don't be rude and get my attention with \"hey bot\" first")
			return
		}
		var Userstemp = []string {}
	
		for _, value := range Users {
			if value != m.Author.ID {
				Userstemp = append(Userstemp, value)
			}
		}
		Users = Userstemp
		if len(splitmsg) == 1 {
			s.ChannelMessageSend(m.ChannelID, "please what")
			return
		}
		switch splitmsg[1] {
			case "rate":
				s.ChannelMessageSend(m.ChannelID, strconv.Itoa(gennumber(0,100)) + " out of 100")
			case "help":
				s.ChannelMessageSend(m.ChannelID, "find it out yourself")
			case "echo","say":
				s.ChannelMessageSend(m.ChannelID, strings.TrimPrefix(m.Content, "please " + splitmsg[1]))
			case "8ball":
				var ballout = [...]string {"As I see it, yes.", "Ask again later.", "Better not tell you now.", "Cannot predict now.", "Concentrate and ask again.", "Don’t count on it.", "It is certain.", "It is decidedly so.", "Most likely.", "My reply is no.", "My sources say no.", "Outlook not so good.", "Outlook good.", "Reply hazy, try again.","Signs point to yes.", "Very doubtful.", "Without a doubt.", "Yes.", "Yes – definitely.", "You may rely on it."}
				s.ChannelMessageSend(m.ChannelID, ballout[gennumber(0,len(ballout))])
			case "dice":
				s.ChannelMessageSend(m.ChannelID, strconv.Itoa(gennumber(1,6)))
			case "killme":
				var verbs = [...]string {" jumped on a "," ran into a "," tripped on a "," slid on a "," sat on a "," stepped on a "," ate a "}
				var objects = [...]string {"ball","dildo","guitar","trombone","apple","pen","vibrator","mouse","laptop","pencil","marker","speaker","baby","door"}
				var injuries = [...]string {" and broke their neck"," and lost their virginity"," and destroyed their kidney", " and died from cancer"," and died from Dysentery", " and had a heart attack", " and had sex with it", " and exploded"}
				s.ChannelMessageSend(m.ChannelID, "<@"+m.Author.ID+">" + verbs[gennumber(0,len(verbs))] + objects[gennumber(0,len(objects))] + injuries[gennumber(0,len(injuries))])
			default:
				s.ChannelMessageSend(m.ChannelID, "please what")
		}
		//if splitmsg[1] == "rate" {
		//	s.ChannelMessageSend(m.ChannelID, gennumber(0,100) + " out of 100")
		//}
		//s.ChannelMessageSend(m.ChannelID, "hello world")
		return
	}
	var inputs = [...]ReactData{
		{input: "hi", reaction: "how are you, " + m.Author.Username},
		{input: "wtf", reaction: "wtf is right"},
		{input: "rip", reaction: "true, rip"},
		{input: "balls", reaction: "!??!"},
		{input: "megaboi", reaction: "death"},
		{input: "tr-ed", reaction: "second best text editor (first goes to pona)"},
		{input: "cum", reaction: "calm down bro"},
		{input: "rei", reaction: "haii >w< hiii"},
		{input: "good", reaction: "good??? me too."},
		{input: "yeti", reaction: "𝔒𝔥 𝔪𝔦𝔫𝔢 𝔤𝔬𝔡, ℑ 𝔞𝔠𝔠𝔦𝔡𝔢𝔫𝔱𝔞𝔩𝔩𝔶 𝔥𝔞𝔱𝔥 𝔰𝔢𝔫𝔱 𝔱𝔥𝔬𝔲 𝔞 𝔭𝔦𝔠𝔱𝔲𝔯𝔢 𝔬𝔣 𝔪𝔦𝔫𝔢 𝔠𝔬𝔠𝔨 𝔞𝔫𝔡 𝔟𝔞𝔩𝔩𝔰...𝔭𝔯𝔦𝔱𝔥𝔢𝔢 𝔡𝔢𝔩𝔢𝔱𝔢 𝔦𝔱!! '𝔏𝔢𝔰𝔱...𝔱𝔥𝔬𝔲 𝔡𝔢𝔰𝔦𝔯𝔢 𝔱𝔬 𝔩𝔬𝔬𝔨? 𝔥𝔞𝔥𝔞 ℑ 𝔧𝔢𝔰𝔱, 𝔡𝔢𝔩𝔢𝔱𝔢 𝔦𝔱...𝔰𝔥𝔬𝔲𝔩𝔡 𝔱𝔥𝔢𝔢 𝔠𝔯𝔞𝔳𝔢... 𝔥𝔞𝔥𝔞 𝔫𝔞𝔶, 𝔟𝔞𝔫𝔦𝔰𝔥 𝔦𝔱...'𝔩𝔢𝔰𝔱?"},
		{input: "no", reaction: "yes"},
		{input: "yes", reaction: "no"},
		{input: "wtf", reaction: "yeah..WTF!!"},
		{input: "stfu", reaction: "no u"},
		{input: "bitch", reaction: "who is the female dog here?"},
		{input: "gay", reaction: "k"},
		{input: "k", reaction: "kk"},
		{input: "fortnite", reaction: "fortnite is a pretty bad game ngl"},
		{input: "roblox", reaction: "The game with creepy pedophiliac moderators? SOUNDS FUN"},
		{input: "aids", reaction: "I agree."},
		{input: "bad", reaction: "true"},
		{input: "help", reaction: "errrr, sorry pal, I can't help you rn."},
		{input: "hello", reaction: "hello"},
	}



	for _, value1 := range splitmsg {
		for _, value2 := range inputs {
			if strings.ToLower(value1) == strings.ToLower(value2.input) {
				s.ChannelMessageSend(m.ChannelID, value2.reaction)
				
			}
		}
	}


}