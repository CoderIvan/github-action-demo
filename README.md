# github-action-demo

* 如果分支test有变动，构建镜像版本`test`，用于部署至测试网

* 如果分支main有变动，构建镜像版本`main`，用于部署至现网

* 如果发现有新的tag(va.b.c)，构建镜像版本`a.b.c`，用于根据版本变更
