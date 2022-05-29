local discordia = require('discordia')
local client = discordia.Client()
local inputs = {
"tred",
"yedit",
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
"sever"
}
local outcomes = {
"me.",
"907740337705984010",
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
"I PeeD My Self"
}
function split(s, delimiter)
    result = {};
    for match in (s..delimiter):gmatch("(.-)"..delimiter) do
        table.insert(result, match);
    end
    return result;
end
client:on('messageCreate', function(message)
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
  				message.channel:send(string.rep(message.content:lower():sub(string.len(splitmsg[1]) + string.len(splitmsg[2]) + string.len(splitmsg[3]) + 4):gsub("@","(at)") .. "\n",tonumber(splitmsg[3])))
                
           
        elseif splitmsg[2] == "echo" then
			message.channel:send(message.content:lower())
        
		elseif splitmsg[2] == "reverse" then
			message.channel:send(string.reverse(message.content:lower():sub(16)))
        end
        return
    end
    for a = 1 , #inputs do
        if string.find(message.content:lower(),inputs[a]) then
			cutmessage = message.content:lower():gsub(inputs[a], "")
            finalmessage = outcomes[a]:gsub("xnamex", cutmessage):gsub("xpingx", message.author.name)
            message.channel:send(finalmessage)
        end
    end
end
client:run("Bot (insert token)");
