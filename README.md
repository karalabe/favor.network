# Favor Network

*"Those who want respect, give respet."* ~Tony Soprano

The favor network is the ultimate *decentralized* and *trustless* ledger to keep track of personal favors you owe to others and those that others owe to you. Not only does it allow you to honour all the promises you've made, but it gradually builds up a stove of reputation data, giving you an idea how *respectable* someone is when it comes to returning a favor.

However, the favor network is not a simple one-to-one ledger of promises, but rather a *living* and *organic* network where favors and promises may flow freely between all participants. Instead of promising someone to return the favor, you might reward them with a pledge you received yourself, leaving it to them to redeem that favor.

The raison d'être of the favor network is to break the one sided relationships in human interactions and permit promises to find their way back to their origins, promoting a culture of generosity and preventing the formation of [givers and takers](https://www.ted.com/talks/adam_grant_are_you_a_giver_or_a_taker).

## Status Network + Favor Network = :heart:

When was the last time you opened your browser, loaded a website and logged in just to ask your friend something? Probably a long time ago! Most of our day-to-day interactions revolve around chat rooms through our mobile phones. To make tracking favors feasible, it must be invisible yet omnipresent in our discussions. *It must be as simple as plain old asking for it.*

Enter the [Status Network](https://status.im/). On the surface Status seems little more than a mobile messanger application, however it is a fully decentralized end-to-end encrypted communication platform with access to both the Ethereum blockchain as well as the Swarm file system.

By integrating the [Favor Network](https://favor.network/) seamlessly into Status, users can ask, accept and honour favors as seamlessly as sending messages to one another. All of this while preserving privacy (Whisper), guaranteeing correctness (Ethereum) and minimizing consensus costs (Swarm).

[![Favor Network](https://img.youtube.com/vi/lO4AsBTiXB8/0.jpg)](https://www.youtube.com/watch?v=lO4AsBTiXB8)

## Missing features

## Status bugs and suggestions

#### Allow chatbots to have event loops or chain callbacks

The chatbots can currently be directed to issue transactions on the users' behalf. The best solution I could come up with to notify the user about this is to have the bot send a message with the transaction details (app specific, not technical), and send a second message with a URL linking the transaction to etherscan. This live URL allows the user to track the progress of the transaction.

A much nicer solution would be to allow the bot to monitor the blockchain for transaction inclusion events and do something when it happens. One possibility is to extend Status with a new chatbot event, which gets triggered if a transaction is included that was sent by the bot. This would allow the bot to react when something it initiated has completed.

Another possibility is to permit chatbots to have their own life cycles, which would allow them to periodically query the chain for needed data and ping the user whenever it deems useful. The benefit of this second approach is that it covers both events initiated locally as well as remote contract updates.

#### Allow chatbots to post into human-to-human chats

Currently Status allows a chatbot to register a `global` command that is made available on human-to-human chat windows, permitting both participants to interact with the bot itself while chatting with each other. The bot however is not allowed to send a message back into this conversation.

To understand why such a feature would be really powerful, consider the following interaction:

 * Alice pings the favor bot to open a new request to Bob
 * A preview is shown, but the transaction is still pending
   * Bob cannot react until the transction finishes
 * *Transaction executes sucessfully*
   * *Nobody knows about it, so it's trial and error*
 * **If the favor bot could notify both parties**
   * **Bob could properly react to the request**

An alternative could be that if a command from a human-to-human chat results in a transaction, Status could automatically inject a message sent via `Console` that could show a progression icon, changing to green/red when the transaction finishes/fails. This would help users notice that they can continue some interaction suspended by Ethereum's transaction processing.

Such a feature would also be beneficial for `/send` and `/receive` where users could be notified not just that the other party executed the command, but also of it's current execution status.

#### Message formatting breaks URL detection

If a chatbot sends a message to the user via `status.sendMessage` which contains a URL, that is automatically detected and converted into a `@browser URL` tappable command. However this functionality breaks if the message being sent contains any format modifiers (`*` or `~`).

Beside fixing the issue so that format modifiers don't break URL detection, I'd venture into suggesting support for a few more modifiers:

 * Explicitly mark a URL as such, perhaps supporting an alternate display string. You could make this feature only available for chat bots but not for users so users can't spoof links but chatbots can display short and sweet versions.
 * Explicitly mark an Ethereum address as such, shortening the display to `0x123456...000000` for example, and creating a tapable `@browse https://etherscan.io/address/0x...` link out of it. The hash could be extended with a contact's name if it's an address known to us.
 * Explicitly mark an Ethereum transaction and block hash as such, shortening and linking them to a block explorer. Perhaps a visual que could also be used to differentiate between the two (note, differentiation cannot be done automatically, it needs a hint from the creator of the message).

#### Allow formatting modifiers in command previews

Currently a chatbot can send a message via `status.sendMessage`, where formatting is done via small modifiers (e.g. `*` for bold and `~` for italic). On the other hand a chatbot's command is displayed via `component.preview`, where formatting is done via full blown reactive components.

This is unfortunate because if I need to display the same information in both a sent message as well as a command preview, I need to make two formatters. Furthermore it breaks the style because I can't imagine an easy way to make them look the same.

My suggestion would be to provide a pre-built component from the status SDK that does the same formatting a used in `status.sendMessage`. Of course making something explicitly catered for each command is arguably better, having a nice default out of the box might help adoption by allowing faster prototyping.

#### Allow chatbots to send multiple command pills

Currently Status features a limited support for sending `requests`, which is basically a plea for the other side to call a pre-parametrized `command` (such are requesting the other side to send some Ether). This construct is useful for human-to-human interactions, but it's a bit limited for bot-to-human messages.

The limitation for chatbots is that a bot may suggest multiple executable commands in a single chat message to make the information more compact (opposed to sending many messages, each containing one suggestion). In this case it's not possible to make all of the commands actively clickable.

My suggestion would be to allow chatbots to send multiple commands like in the example below, where all of the commands get turned into pills (clickable inline buttons). This way we can preserve the "feel" of chatting with someone while at the same time providing seamless contextual aid. The pills should allow for hidden parameters expanded on click of course.

![ChatBot Pills](http://i.imgur.com/Gl5s4fM.png)

#### Allow deleting synchronized chain data

When setting up multiple emualtors at once, usually some of them started syncing from zero. Not sure why this happened (possibly a bad peer?), but this results in a very slow startup and a lot of disk space. I know that ultimately Geth needs to be smarter here (would appreciate a bug report), but perhaps having some menu to "delete chain data" would help both fix this error as well as possibly clear out old data when a new version of Status is released with a new CHT.
