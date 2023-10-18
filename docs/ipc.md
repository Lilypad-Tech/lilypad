## IPC geth replacement

requirements:

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

restart terminal

```bash
curl -L https://foundry.paradigm.xyz | bash
```

restart terminal

```bash
git clone git@github.com:consensus-shipyard/fendermint.git
git clone git@github.com:consensus-shipyard/ipc-solidity-actors.git
cd fendermint
export CARGO_NET_GIT_FETCH_WITH_CLI=true
sudo apt-get install protobuf-compiler clang
foundryup
make build docker-build
```