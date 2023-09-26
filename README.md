# Lilypad v2 🍃

This cloud is just someone else's computer.

![image](https://github.com/bacalhau-project/lilypad/assets/264658/d91dad9a-ca46-43d4-a94b-d33454efc7ae)

# Getting started

Welcome to the prerelease series of Lilypad v2.

## Aurora Testnet

The testnet has a base curency of ETH and you will also get LP to pay for jobs (and nodes to stake).

Metamask:

```
Network name: Lilypad v2 Aurora testnet
New RPC URL: http://testnetv2.arewehotshityet.com:8545
Chain ID: 1337
Currency symbol: ETH
Block explorer URL: (leave blank)
```

### Fund your wallet with ETH and LP

To obtain funds, go to [http://testnetv2.arewehotshityet.com:8080](http://testnetv2.arewehotshityet.com:8080)

The faucet will give you both ETH (to pay for gas) and LP (to stake and pay for jobs).

## Install CLI

Currently only supports x86_64 Linux

```
curl -sSL -o lilypad https://github.com/bacalhau-project/lilypad/releases/download/v2.0.0-b7e9e04/lilypad
chmod +x lilypad
sudo mv lilypad /usr/local/bin
```

## Run a job

```
export WEB3_PRIVATE_KEY=<your private key>
```
(or arrange for the key to be in your environment in a more secure way that doesn't get written to your shell history)


### Cows

```
lilypad run cowsay:v0.0.1 -i Message="hey beautiful"
```


### SDXL

```
lilypad run sdxl:v0.9-lilypad1 -i PromptEnv="PROMPT=record player in reykjavik"
```

![image-42](https://github.com/bacalhau-project/lilypad/assets/264658/d48bb897-79a0-4f3a-b938-e85a8cfa3f0e)

## Run a node, earn LP

```
lilypad serve
```

systemd units & more details [here](https://github.com/bacalhau-project/lilypad/tree/main/ops)

## Available modules

Check the github releases page for each module or just use the git hash as the tag.

* [sdxl](https://github.com/bacalhau-project/lilypad-module-sdxl)
* [stable-diffusion](https://github.com/bacalhau-project/lilypad-module-stable-diffusion)
* [duckdb](https://github.com/bacalhau-project/lilypad-module-duckdb)
* [fastchat](https://github.com/bacalhau-project/lilypad-module-fastchat)
* [lora-inference](https://github.com/bacalhau-project/lilypad-module-lora-inference)
* [lora-training](https://github.com/bacalhau-project/lilypad-module-lora-training)
* [filecoin-data-prep](https://github.com/bacalhau-project/lilypad-module-filecoin-data-prep)
* [wasm](https://github.com/bacalhau-project/lilypad-module-wasm)
* [cowsay](https://github.com/bacalhau-project/lilypad-module-cowsay)


## Write a module

A module is just a git repo.

Module versions are just git tags.

In your repo, create a file called `lilypad_module.json.tmpl`

See [cowsay](https://github.com/bacalhau-project/lilypad-module-cowsay) for example

This is a json template with Go text/template style `{{.Message}}` sections which will be replaced by Lilypad with valid JSON strings which are passed as input to modules. You can also do fancy things with go templates like setting defaults, see cowsay for example. While developing a module, you can use the git hash to test it.

Pass inputs as:

```
lilypad run github.com/username/repo:tag -i Message=moo
```

Inputs are a map of strings to strings.

**YOU MUST MAKE YOUR MODULE DETERMINISTIC**

Tips:
* Make the output reproducible, for example for the diffusers library, see [here](https://huggingface.co/docs/diffusers/using-diffusers/reproducibility)
* Strip timestamps and time measurements out of the output, including to stdout/stderr
* Don't read any sources of entropy (e.g. /dev/random)
* When referencing docker images, you MUST specify their sha256 hashes, as shown in this example

If your module is not deterministic, compute providers will not adopt it and add it to their allowlists.

