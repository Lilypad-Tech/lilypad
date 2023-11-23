# Lilypad v2 🍃

This cloud is just someone else's computer.

![image](https://github.com/bacalhau-project/lilypad/assets/264658/d91dad9a-ca46-43d4-a94b-d33454efc7ae)

Lilypad enables users to run AI workloads easily in a decentralized GPU network where anyone can get paid to connect their compute nodes to the network and run jobs. Users have access to easy Stable Diffusion XL and cutting edge open source LLMs both on chain, from CLI and via [Lilypad AI Studio](https://lilypad.tech) on the web.

# Getting started

Welcome to the prerelease series of Lilypad v2.

## Aurora Testnet

The testnet has a base curency of ETH and you will also get LP to pay for jobs (and nodes to stake).

Metamask:

```
Network name: Lilypad v2 Aurora testnet
New RPC URL: http://testnet.lilypad.tech:8545
Chain ID: 1337
Currency symbol: ETH
Block explorer URL: (leave blank)
```

### Fund your wallet with ETH and LP

To obtain funds, go to [http://faucet.lilypad.tech:8080](http://faucet.lilypad.tech:8080)

The faucet will give you both ETH (to pay for gas) and LP (to stake and pay for jobs).

## Install CLI

Download the latest release of Lilypad for your platform. Both the amd64/x86_64 and arm64 variants of macOS and Linux are supported. (If you are on Apple Silicon, you'll want arm64).

The commands below will automatically detect your OS and processor architecture and download the correct Lilypad build for your machine.

```
# Detect your machine's architecture and set it as $OSARCH
OSARCH=$(uname -m | awk '{if ($0 ~ /arm64|aarch64/) print "arm64"; else if ($0 ~ /x86_64|amd64/) print "amd64"; else print "unsupported_arch"}') && export OSARCH
# Detect your operating system and set it as $OSNAME
OSNAME=$(uname -s | awk '{if ($1 == "Darwin") print "darwin"; else if ($1 == "Linux") print "linux"; else print "unsupported_os"}') && export OSNAME;
# Download the latest production build
curl -sSL -o lilypad https://github.com/bacalhau-project/lilypad/releases/download/v2.0.0-d63a7ff/lilypad-$OSNAME-$OSARCH
# Make Lilypad executable and install it
chmod +x lilypad
sudo mv lilypad /usr/local/bin/lilypad
```

You can also, at your option, choose to compile Lilypad using Go and install it that way on any machine that supports the Go toolchain.

## Run a job

```
export WEB3_PRIVATE_KEY=<your private key>
```
(or arrange for the key to be in your environment in a more secure way that doesn't get written to your shell history)


### Cows

```
lilypad run cowsay:v0.0.1 -i Message="moo"
```


### SDXL

```
lilypad run sdxl:v0.9-lilypad1 -i PromptEnv="PROMPT=beautiful view of iceland with a record player"
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

This is a json template with Go text/template style `{{.Message}}` sections which will be replaced by Lilypad with json encoded inputs to modules. You can also do fancy things with go templates like setting defaults, see cowsay for example. While developing a module, you can use the git hash to test it.

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


### Writing Advanced Modules

1. `subt`:
The `subt` function allows for substitutions in your template, a feature that addresses the issue outlined in [#14](https://github.com/bacalhau-project/lilypad/issues/14).

This function is a workaround for the lack of direct substitution support in the module. It implements the [printf](https://pkg.go.dev/text/template#Template.Funcs) function under the hood, which allows you to format strings with placeholders.

<details>
  <summary> 
    Usage   
  </summary>
    The `subt` function can be used in the same way as the `printf` function in Go. You pass in a format string, followed by values that correspond to the placeholders in the format string.
    ```
    const templateText = `
    {{ subt "Hello %s" .name }}
    `
    ```
</details>

[Example Code](https://go.dev/play/p/oBgc2Cetug3)
