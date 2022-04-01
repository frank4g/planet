## Build
cat.|instr.|output
-|-|-
```scaffold.chain(planet)```|*```starport scaffold chain github.com/deep2essence/planet --no-module```*|*```app/,cmd,/docs/,testutil/,vue/,config.yml```<br><br>```go.mod,go.sum,readme.md,.github/,.gitignore```*
```scaffold.module(blog)```|*```starport scaffold module blog --ibc```*|*generated:```proto/blog/+packet.proto```<br>```x/blog/+module_ibc.go```<br>```testutil/keeper/blog.go```<br><br>modified:```app/app.go,docs/static/openapi.yml```*
```scaffold.list(blog.post)```|```starport scaffold list post title content --module blog```|*generated:```proto/blog/post.proto```<br>```x/blog/keeper/grpc+,msg_server+,+.go```<br>```x/blog/client/cli/tx+,query+.go```<br>```x/blog/types```<br><br>modified:```proto/blog/tx,query,genesis.proto```<br>```x/blog/genesis,handler,module.go```<br>```query,tx,genesis,keys,codec```*
```scaffold.list(blog.sentPost)```|```starport scaffold list sentPost postID title chain --module blog```|*generated:```proto/blog/sent_post.proto```<br>-*
```scaffold.list(blog.timedoutPost)```|```starport scaffold list timedoutPost title chain --module blog```|*generated:```proto/blog/timeout_post.proto```<br>-*
```scaffold.packet(blog.packet)```|```starport scaffold packet ibcPost title content --ack postID --module blog```|*```proto/blog/packet.proto:+PacketData,+PacketAck```<br>```proto/blog/tx.proto:SendIbcPost,MsgSendIbcPost```<br>```handler,tx,moduel,codec,ibc,event,message,packet```*

## Q & A
Q|A
-|-
```CRUD actions```?|
```starport scaffold module```vs.```starport scaffold list```|
```starport scaffold list```vs.```starport scaffold message```|
```starport scaffold module --ibc```|
```starport scaffold module --dep bank```|
```starport scaffold list --module blog```|*The ```--module``` flag defines which module the new transaction type is added to. This optional flag lets you manage multiple modules within your Starport app. When the flag is not present, the type is scaffolded in the module that matches the name of the repo.*
```starport scaffold list --no-message```|*When a new type is scaffolded, the default behavior is to scaffold messages that can be sent by users for **CRUD operations**. The --no-message flag disables this feature. Disable the messages option for the app since you want the posts to be created upon reception of IBC packets and not directly created from a user's messages.*
```starport scaffold packet ```
```starport scaffold packet --ack```|*The ```--ack``` flag defines which identifier is returned to the sending blockchain*