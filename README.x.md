- https://docs.starport.com/guide/ibc.html

## Build
cat.|instr.|output
-|-|-
```scaffold.chain(planet)```|*```starport scaffold chain github.com/deep2essence/planet --no-module```*|*```app/,cmd,/docs/,testutil/,vue/,config.yml```<br><br>```go.mod,go.sum,readme.md,.github/,.gitignore```*
```scaffold.module(blog)```|*```starport scaffold module blog --ibc```*|*generated:```proto/blog/+packet.proto```<br>```x/blog/+module_ibc.go```<br>```testutil/keeper/blog.go```<br><br>modified:```app/app.go,docs/static/openapi.yml```*
```scaffold.list(blog.post)```|```starport scaffold list post title content --module blog```|*generated:```proto/blog/post.proto```<br>```x/blog/keeper/grpc+,msg_server+,+.go```<br>```x/blog/client/cli/tx+,query+.go```<br>```x/blog/types```<br><br>modified:```proto/blog/tx,query,genesis.proto```<br>```x/blog/genesis,handler,module.go```<br>```query,tx,genesis,keys,codec```*
```scaffold.list(blog.sentPost)```|```starport scaffold list sentPost postID title chain --module blog```|*generated:```proto/blog/sent_post.proto```<br>-*
```scaffold.list(blog.timedoutPost)```|```starport scaffold list timedoutPost title chain --module blog```|*generated:```proto/blog/timeout_post.proto```<br>-*
```scaffold.packet(blog.ibcPost)```|```starport scaffold packet ibcPost title content --ack postID --module blog```|```proto/blog/tx.proto:msg.SendIbcPost,MsgSendIbcPost```<br>```proto/blog/packet.proto```:```BlogPacketData```,<br>```IbcPostPacketData```,<br>```IbcPostPacketAck```<br><br>```starport scaffold packet``` generates one message:```MsgSendIbcPost```  and two packet data: ```IbcPostPacketData``` & ```IbcPostPacketAck```. Msgs orignally have "creator" field, Packet Msgs have inherited fields - "port", "channelID", "timeoutTimestamp" excluding customized fields and "creator". But packet(Data+Ack) don't have "creator" field inherited. Problem originates from this. When you send message to ibc packet, you can't deliver creator. So, you should insert the creator field into packetData in proto.file and also need to implement corresponding packet.pb.go files generated from packet.proto.<br><br>, where question is - Is it possible to regenerate packet.pb.go from modified packet.proto using buf or cosmos-sdk protobuf toolkit? Is it worth? Is it okay to manually modify pb.go file?<br><br>SendingPacket is almost perfect in ```msg_server_ibc_post.go```<br>,but ReceivingPacket need to be complemented in ```ibc_post.go```<br>```OnRecvIbcPostPacket```,```OnAcknowledgementIbcPostPacket```,```OnTimeoutIbcPostPacket``` modified.
## Q & A
Q|A
-|-
```CRUD actions```?|
## Test
k|v
-|-
```clean existing relayer```|```rm -rf ~/.starport/relayer```
```run earth blockchain app```|
```run earth blockchain app```|
```Configure and start the relayer```|```starport relayer configure -a \```<br>```--source-rpc "http://0.0.0.0:26657" \```<br>```--source-faucet "http://0.0.0.0:4500" \```<br>```--source-port "blog" \```<br>```--source-version "blog-1" \```<br>```--source-gasprice "0.0000025stake" \```<br>```--source-prefix "cosmos" \```<br>```--source-gaslimit 300000 \```<br>```--target-rpc "http://0.0.0.0:26659" \```<br>```--target-faucet "http://0.0.0.0:4501" \```<br>```--target-port "blog" \```<br>```--target-version "blog-1" \```<br>```--target-gasprice "0.0000025stake" \```<br>```--target-prefix "cosmos" \```<br>```--target-gaslimit 300000```
```earth->mars```|```planetd tx blog send-ibc-post blog channel-0 "Hello" "Hello Mars, I'm Alice from Earth" --from alice --chain-id earth --home ~/.earth```
```list```|```planetd q blog list-post --node tcp://localhost:26659```
```list```|```planetd q blog list-sent-post```
```timestamped```|```planetd tx blog send-ibc-post blog channel-0 "Sorry" "Sorry Mars, you will never see this post" --from alice --chain-id earth --home ~/.earth --packet-timeout-timestamp 1```
```list```|```planetd q blog list-timedout-post```
```mars->earth```|```planetd tx blog send-ibc-post blog channel-0 "Hello" "Hello Earth, I'm Alice from Mars" --from alice --chain-id mars --home ~/.mars --node tcp://localhost:26659```
```list```|```planetd q blog list-post```
