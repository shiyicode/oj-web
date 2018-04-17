#!/bin/bash

workspace=$(cd $(dirname $0) && pwd -P)
cd $workspace

#const
app="oj-web"
cfg=./cfg


function build() {
	local go="/usr/local/go"
	if [ -d "$go" ]; then
		export GOROOT="$go"
		export PATH="$PATH:$GOROOT/bin"
		export GOPATH="$GOROOT:$(pwd):$(pwd)/deps"
	else
		echo "Go文件不存在!"
		exit 1
	fi
	echo "`go version`"
	go build
}

function make_output {
	local output="./output"
	rm -rf $output &>/dev/null
	mkdir -p $output &>/dev/null
	(
		cp -vrf $app $output && # 拷贝 二进制文件 至output目录
		cp -vrf $cfg $output && # 拷贝 cfg配置文件目录 至output目录
		echo -e "make output ok"
	) || { echo -e "make output error"; rm -rf "./output"; exit 2; }
}

build

make_output

echo "build down"
exit 0