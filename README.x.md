## Build
cat.|instr.|output
-|-|-
```scaffold.chain(planet)```|*```starport scaffold chain github.com/deep2essence/planet --no-module```*|*```app/,cmd,/docs/,testutil/,vue/,config.yml```<br><br>```go.mod,go.sum,readme.md,.github/,.gitignore```*
```scaffold.module(blog)```|*```starport scaffold module blog --ibc```*|*generated:```proto/blog/+packet.proto```<br>```x/blog/+module_ibc.go```<br>```testutil/keeper/blog.go```<br><br>modified:```app/app.go,docs/static/openapi.yml```*
```scaffold.list(blog.post)```|```starport scaffold list post title content --module blog```|*generated:```proto/blog/post.proto```<br>```x/blog/keeper/grpc+,msg_server+,+.go```<br>```x/blog/client/cli/tx+,query+.go```<br>```x/blog/types```<br><br>modified:```proto/blog/tx,query,genesis.proto```<br>```x/blog/genesis,handler,module.go```<br>```query,tx,genesis,keys,```*

## Q & A
Q|A
-|-
```CRUD actions```?|