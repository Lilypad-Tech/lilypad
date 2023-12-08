First export your private key from meta-mask:

```bash
export WEB3_PRIVATE_KEY=XXX
```

clone the repo (if you have not):

```bash
git clone https://github.com/bacalhau-project/lilypad.git
cd lilypad
# note - remove this once the branch is merged
git checkout example-onchain-client
```

now we install node modules:

```bash
cd hardhat
yarn install
```

then run the example script:

```bash
npx hardhat run --network prod scripts/run-cowsay-onchain.ts
```

