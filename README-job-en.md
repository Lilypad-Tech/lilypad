# Getting started running container jobs on Lilypad 🍃
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

[English(Native)](README-node-en.md) | [Japanese(日本語)](README-node-ja.md) | [Chinese(中文)](README-node-cn) | [Turkish (Türkçe)](README-node-tr.md) | We need your help to translate this README into your native language.

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
