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

> 如果您在中国大陆，需要翻墙。而且不只是要浏览器翻墙，命令行也要使用代理翻墙。同时，localhost 不应该有代理：在终端命令行输入 `export no_proxy=localhost,127.0.0.0/8`.

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

# Build a Job Module

A Lilypad module is a Git repository that can be used to perform various tasks using predefined templates and inputs.&#x20;

This guide will walk you through the process of creating a Lilypad module, including defining a JSON template, handling inputs, ensuring determinism, and other best practices.

## Module Structure

1. Start by creating a Git repository for your Lilypad module. The module's versions will be represented as Git tags.
2. Inside your module's repository, create a file named `lilypad_module.json.tmpl`. This file will serve as a JSON template with Go text/template style sections, like `{{.Message}}`, which will be replaced by Lilypad with JSON-encoded inputs.
3. You can also use Go templates to set defaults and perform other template-related operations. Refer to [the `cowsay` example](https://docs.lilypad.tech/lilypad/lilypad-milky-way-examples/hello-cow-world) for inspiration.

## Handling Inputs

Inputs are passed to your Lilypad module as a map of strings to strings. You can pass inputs using the following command:

```bash
lilypad run github.com/username/repo:tag -i Message=moo
```

Ensure that your module can accept and process inputs as specified.

## Ensuring Determinism

To ensure that your Lilypad module is deterministic, follow these guidelines:

* Make the output of your module reproducible.
* Strip timestamps and time measurements from the output, including to `stdout` and `stderr`.
* Avoid reading from sources of entropy, such as /dev/random.
* When referencing Docker images, specify their sha256 hashes.

If your module is not deterministic, it may not be adopted by compute providers, and it won't be added to their allowlists.

## Testing Your Module

During development, you can use the Git hash to test your module. Ensure that your module functions correctly and produces deterministic results.

## Examples

Here are some example Lilypad modules for reference:

* [**Cowsay**](https://github.com/Lilypad-Tech/lilypad-module-cowsay)**:** Lilypad "Hello World" example
* [**SDXL v0.9/v1.0**](https://github.com/Lilypad-Tech/lilypad-module-sdxl-pipeline): Text to image generation.
* [**SDV v1.0/1.1**](https://github.com/Lilypad-Tech/lilypad-module-sdv-pipeline): Text to video generation.

Deprecated examples:

* [**lora-training**](https://github.com/Lilypad-Tech/lilypad-module-lora-training)**:** An example module for LoRa training tasks.
* [**lora-inference**](https://github.com/Lilypad-Tech/lilypad-module-lora-inference)**:** An example module for LoRa inference tasks.
* [**duckdb**](https://github.com/Lilypad-Tech/lilypad-module-duckdb)**:** An example module related to DuckDB.

These examples can help you understand how to structure your Lilypad modules and follow best practices.

## Conclusion

In this guide, we've covered the essential steps to create a Lilypad module, including defining a JSON template, handling inputs, ensuring determinism, and testing your module. Follow these best practices to create reliable and reusable modules for Lilypad.

For more information and examples, refer to the official Lilypad documentation and the cowsay example module.
