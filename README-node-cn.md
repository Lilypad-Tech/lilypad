# Getting started running a Node on Lilypad 🍃 Network
<p align="left">
  <a href="https://lilypad.team/discord">
    <img src="https://img.shields.io/badge/chat-Discord-7289DA?logo=discord" alt="GaiaNet Discord">
  </a>
  <a href="https://twitter.com/Lilypad_Tech">
    <img src="https://img.shields.io/badge/Twitter-1DA1F2?logo=twitter&amp;logoColor=white" alt="GaiaNet Twitter">
  </a>
   <a href="https://lilypad.tech">
    <img src="https://img.shields.io/website?up_message=Website&url=https://www.gaianet.ai/" alt="Gaianet website">
  </a>
</p>

> 如果您在中国大陆，需要翻墙。而且不只是要浏览器翻墙，命令行也要使用代理翻墙。同时，localhost 不应该有代理：在终端命令行输入 `export no_proxy=localhost,127.0.0.0/8`.

[English(Native)](README-node-en.md) | [Japanese(日本語)](README-node-ja.md) | [Chinese(中文)](README-node-cn) | [Turkish (Türkçe)](README-node-tr.md) | We need your help to translate this README into your native language.



Lilypad enables users to run containerised AI workloads easily in a decentralized GPU network, where anyone can get paid to connect their compute nodes to the network and run container jobs. Users have access to easily run jobs such as Stable Diffusion XL and cutting edge open source LLMs both on chain, from CLI and via [Lilypad AI Studio](https://app.lilypad.tech) on the web.

Visit the [Lilypad Docs](https://docs.lilypad.tech/) site for a more comprehensive overview to getting up and running including a [Quick Start Guide](https://lilypad.team/quickstart)

Like this project? ⭐ Star us!

## Getting started running container jobs on Lilypad



Jobs (containers) can be run on Lilypad by utilising the [Installable CLI](https://lilypad.team/cli), also available for installation through the [Go toolchain](https://lilypad.team/cligo). After setting up the necessary pre-requisites, the CLI enables users to run jobs as described below:

```
lilypad run cowsay:v0.0.4 -i Message="moo"
```

[![Watch the video](https://img.youtube.com/vi/Ep9ML9h8DTE/0.jpg)](https://www.youtube.com/watch?v=Ep9ML9h8DTE)

The current list of modules can be found in the following repositories: 


* [lilysay](https://github.com/Lilypad-Tech/lilypad-module-lilysay)
* [cowsay](https://github.com/lilypad-tech/lilypad-module-cowsay)
* [Stable Diffusion XL](https://github.com/Lilypad-Tech/lilypad-module-sdxl-pipeline/)
* [Stable Diffusion Video](https://github.com/Lilypad-Tech/lilypad-module-sdv-pipeline)
* [wasm](https://github.com/lilypad-tech/lilypad-module-wasm)

Containerised job modules can be built and added to the available module list; for more details visit the [building a job](https://lilypad.team/building) documentation. If you would like to contribute, open a pull request on this repository to add your link to the list above.

## Getting started running a Node on Lilypad Network

[Japanese(日本語)](README-ja.md) | [Chinese(中文)](README-node-cn) | [Turkish (Türkçe)](README-tr.md) | We need your help to translate this README into your native language.

As a distributed network Lilypad also brings with it the ability to run as a node and contribute to the GPU and compute capabilities. See the documentation on [running a node](https://lilypad.team/node) which contains more details instructions and overview for getting set up. 

## The Lilypad Community 

[Read our Blog](https://lilypad.team/blog)

[Join the Discord](https://lilypad.team/discord)

[Follow us on Twitter/X](https://lilypad.team/x)

[Check out our videos on YouTube](https://lilypad.team/youtube)

