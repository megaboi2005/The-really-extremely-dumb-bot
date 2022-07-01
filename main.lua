local discordia = require('discordia')
local client = discordia.Client()
local inputs = {
"pona",
"tred",
"wow",
"how are you",
"gm",
"gn",
"shut up",
"women",
"$help",
"megagay",
"megayboi",
"megaboi",
"chromebook",
"python",
"con",
"who are you",
"i am ",
"hello",
"pp",
"poopoo",
"sever",
"node.js",
"javascript",
"hi",
"based",
"goodbye",
"bye"
}
local outcomes = {
"me.",
"based",
"agreed.",
"I am feeling pretty good, how about you?",
"Good morning xpingx.",
"Good night xpingx.",
"thats not very nice :(",
"hahaaaaahahahaha",
"wtf, I am a chatbot not a moderation bot",
"STOP",
"SHUT UP",
"deadname",
"ok chromie",
"snake language",
"push pop",
"I am the really extremely dumb bot",
"Nice to meet you xnamex",
"hello xpingx",
"poopoo",
"pp",
"I PeeD My Self",
"bad.js",
"true == \"true\" // -> false;",
"Hi xpingx",
"Based?? Based on what?!?!?!",
"Goodbye ily",
"Goodbye queen"
}
function split(s, delimiter)
    result = {};
    for match in (s..delimiter):gmatch("(.-)"..delimiter) do
        table.insert(result, match);
    end
    return result;
end

function wordfind(sentence,word)
	list = split(sentence," ")
	for i = 0, #list do
		if list[i] == word then return true; end
	end
	return false
end
usermsgcount = {}


client:on('messageCreate', function(message)
	
	if usermsgcount[message.author.tag] == nil then usermsgcount[message.author.tag] = 0 else usermsgcount[message.author.tag] = usermsgcount[message.author.tag] + 1 end
	
	math.randomseed(os.time())
	math.random(); math.random(); math.random()
    if message.author.bot then return end
    local finalmessage
    if message.content == "*secret" then
        local compilation = ""
        for a = 1 , #inputs do
            compilation = compilation .. " " .. inputs[a] .. ","
        end
        message.channel:send(compilation)
    end
    if message.content:lower():find("please", 1,false) then
        splitmsg = split(message.content:lower()," ")
        if splitmsg[2] == "repeat" then
            if type(tonumber(splitmsg[3])) == "nil" then return end
           	if tonumber(splitmsg[3]) > 10 then message.channel:send("number is too big"); return;  end
  			message.channel:send(string.rep(message.content:lower():sub(string.len(splitmsg[1]) + string.len(splitmsg[2]) + string.len(splitmsg[3]) + 4):gsub("@","(at)"):gsub("\n","") .. "\n",tonumber(splitmsg[3])))
                
           
        elseif splitmsg[2] == "echo" then
        	if message.content:sub(13) == "" then message.channel:send("Hello World."); return; end
			message.channel:send(message.content:sub(13):gsub("xuserx", message.author.name))
        
		elseif splitmsg[2] == "reverse" then
			message.channel:send(string.reverse(message.content:lower():sub(16)))

		elseif splitmsg[2] == "shake" then

			if splitmsg[3] == "8ball" then
				balloutcomes = {"As I see it, yes.", "Ask again later.", "Better not tell you now.", "Cannot predict now.", "Concentrate and ask again.", "Don’t count on it.", "It is certain.", "It is decidedly so.", "Most likely.", "My reply is no.", "My sources say no.", "Outlook not so good.", "Outlook good.", "Reply hazy, try again.","Signs point to yes.", "Very doubtful.", "Without a doubt.", "Yes.", "Yes – definitely.", "You may rely on it."}
				message.channel:send(balloutcomes[math.random(1,#balloutcomes)])

			elseif splitmsg[3] == "dice" then
				print(tonumber(splitmsg[4]))
				if tonumber(splitmsg[4]) < 0 then message.channel:send("no negative numbers"); return; elseif tonumber(splitmsg[4]) == nil then message.channel:send("not a valid number"); return; end
				message.channel:send(tostring(math.random(1,tonumber(splitmsg[4]))))
			end

		elseif splitmsg[2] == "help" then file = io.open("help.txt", "r"):read("*a"); message.channel:send(file);

		elseif splitmsg[2] == "give" then
				if splitmsg[3] == "msgcount" then 
					message.channel:send(usermsgcount[message.author.tag]) 
				elseif splitmsg[3] == "profile" then
					if splitmsg[4] == nil then 
						message.channel:send(message.author:getAvatarURL())
					else
						user = client:getUser(splitmsg[4]:gsub("<",""):gsub(">",""):gsub("@",""))
						if user == nil then message.channel:send("that user isn't in the server") end
						message.channel:send(user:getAvatarURL()) 
					end
				
				end
		else
			message.channel:send("please what??!?")
        end
        return
    end
    for a = 1 , #inputs do 
        if wordfind(message.content:lower(),inputs[a]) then
			cutmessage = message.content:lower():gsub(inputs[a], "")
            finalmessage = outcomes[a]:gsub("xnamex", cutmessage):gsub("xpingx", message.author.name):gsub("@","(at)")
            message.channel:send(finalmessage)
        end
    end


end)



client:run("Bot (token)");
