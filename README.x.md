cat.|instr.|output
-|-|-
```scaffold.chain(planet)```|*```starport scaffold chain github.com/deep2essence/planet --no-module```*|*```app/,cmd,/docs/,testutil/,vue/,config.yml```<br><br>```go.mod,go.sum,readme.md,.github/,.gitignore```*
```scaffold.module(blog)```|*```starport scaffold module blog --ibc```*|*generated:```proto/blog/+packet.proto```<br>```x/blog/+module_ibc.go```<br>```testutil/keeper/blog.go```<br><br>modified:```app/app.go,docs/static/openapi.yml```*