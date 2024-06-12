## Prerequisite
Get Test ETH from Faucet in to your Metamask wallet.
List of faucets available at this URL:
https://arbitrum.faucet.dev/ArbSepolia

The code currently funds each account with 1 Eth, so make sure you have at least 9 Eth before staring the deployment process.

If you do not have Arbitrum Sepolia network added to Maskmask, use the following settings to do so:
```
Network name: Sepolia Testnet
RPC : https://sepolia-rollup.arbitrum.io/rpc
Chain ID: 421614
Currency symbol: ETH
```

## Deployment Process
Open terminal
Inside 
`export NETWORK=arbitrumSepolia`

Generate new set of addresses

`npx hardhat run scripts/generate-addresses.ts`

Example Output (don't copy these address because they are known. Your funds will be. Use the output you generated):
```
export ADMIN_PRIVATE_KEY=0x27e9cb3c0785aa6a9dbec237cf5f69894fba73f02daa25e86f274c773bdc3cb7
export ADMIN_ADDRESS=0x0641d933efbb003bf9e0ea333f03db59413014c5
export FAUCET_PRIVATE_KEY=0x03f141dc55bbec84871f5f575f338a28c1139c03e50205f12f48eada261c8c9c
export FAUCET_ADDRESS=0x4d1cf9499d3cef12231561cb7277677a8f27ea3f
export SOLVER_PRIVATE_KEY=0x8ed543dfea0381e658c74332e5aa0bfa622047e98b3680266ea9c6231e4c38b4
export SOLVER_ADDRESS=0x14365ec127bbc37aad6f2f8e122094688044730f
export MEDIATOR_PRIVATE_KEY=0x84391a9259e11ccbb22e2f7f965790ba40a367fd1355b1f034289eaaebca7035
export MEDIATOR_ADDRESS=0xf5aeeb89138f3c486c2124f607e597f8c55350b3
export RESOURCE_PROVIDER_PRIVATE_KEY=0x729e3ccc62549a1b768b3b3bbade172869d3bbc9f5fd4959182675c625ce5cc9
export RESOURCE_PROVIDER_ADDRESS=0x9471016953150e114dbc3e992030083f35a5935e
export JOB_CREATOR_PRIVATE_KEY=0x0ccc57d56abd8f54fc7a1e50f82e82e627bdc05c06b67fb869fbe6d3faf9de29
export JOB_CREATOR_ADDRESS=0xec7d44341655b91441ac819199a6b928ed9e662a
export DIRECTORY_PRIVATE_KEY=0xfba8780bb44ee5f67984f1ade392e58d93b9b6dda8baa35332826e3651613f3f
export DIRECTORY_ADDRESS=0x1fafffe930fcdd51462aee8f8b31b32451c5b07f
```
Copy your output and paste the whole block back in to terminal. Press **ENTER**

This will set your current environment variable need by **Deploy and Fund** Script later in the process.

You should also save these values in a secure place for future reference.

>If you are using this for local development and security is not critial you can append the exports in to `~/.bashrc`. 

Open `~/.bashrc` in the editor of your choice (for example `code ~/.bashrc` for VSCode).

Copy `ADMIN_ADDRESS` address value from YOUR output.
This address will be used during deploy process to pay for gas

Transfer **9** Eth from Metamask to YOUR `ADMIN_ADDRESS`
These funds will be used to fund other accounts.

Deploy Command:

`npx hardhat deploy`

Example Output (your addresses and transaction will be different):

```
admin: 0x0641d933efbb003bf9e0ea333f03db59413014c5
deploying "LilypadToken" (tx: 0x4a391c7baa7a7bb4d38c70e0eaf5362922f4c53644ee8868bb464e35e8d91308)...: deployed at 0x9268bB5C5f6403Ff02A89DCFf7ddBb07ff046F99 with 2234053 gas
deploying "LilypadPayments" (tx: 0xf1d12eebeaf449a9040fb00ed1a01427858ae76e29ac3d88a76519184413e33c)...: deployed at 0x22963655Fe6EbbfA46400E9F01012E54bAE543c4 with 2369600 gas
executing LilypadPayments.initialize (tx: 0xb7d380a87751df7f221f62a11ed9143eb3205a4687b470eca07b309caba81101) ...: performed with 59252 gas
executing LilypadToken.setControllerAddress (tx: 0x8f8275bb451aeec3f9d23a1b10b372cc2bda7015d4584590a5d69936b8678c54) ...: performed with 29555 gas
deploying "LilypadPayments" (tx: 0xf1d12eebeaf449a9040fb00ed1a01427858ae76e29ac3d88a76519184413e33c)...: deployed at 0x22963655Fe6EbbfA46400E9F01012E54bAE543c4 with 2369600 gas
executing LilypadStorage.initialize (tx: 0xe4f3afd5cac1868bf11368358fa7c1ddf309e71db566dc5b7cff3ad7141a4480) ...: performed with 28886 gas
deploying "LilypadPayments" (tx: 0xf1d12eebeaf449a9040fb00ed1a01427858ae76e29ac3d88a76519184413e33c)...: deployed at 0x22963655Fe6EbbfA46400E9F01012E54bAE543c4 with 2369600 gas
executing LilypadUsers.initialize (tx: 0xfbb25d14a71528b3e09b908343155ad83f3f8f800bbcc2333cec8de200342d86) ...: performed with 28817 gas
deploying "LilypadMediationRandom" (tx: 0x83fe1f3f662935e219749745e750be864bb8df18d04bb4845754db66693dfaf5)...: deployed at 0x104f5cc5d1593F1BA2a0eecf5882Be85e231acA9 with 1506321 gas
executing LilypadMediationRandom.initialize (tx: 0x5b8abe29dd253036866340c7bc862939eb177593ab10f40c29d536ff73db0e89) ...: performed with 28842 gas
deploying "LilypadOnChainJobCreator" (tx: 0xb6427e31fe4f83d486730976064cb0d8f9a58b16cab38fd4b63bd540bfac7548)...: deployed at 0x15F7FE757a582634E1Ed105ea68F07f6cF240b37 with 1755268 gas
deploying "ExampleClient" (tx: 0x59bd354405b735b93bca319402cb88be0734b1c93fdd9c14ef7898f5b03166e4)...: deployed at 0xe7e96cEF8812D2d24ADD2de592b5C786f915f64b with 1212230 gas
executing LilypadOnChainJobCreator.initialize (tx: 0xbf385a51558a5989c1a3b07c72f931992dd9e2a6d66bad9e5e513c09adec7019) ...: performed with 78498 gas
executing ExampleClient.initialize (tx: 0xd2582af972fa93e0fa730f9d433ca18c6c81bb905544d8ab1e513566ab4913be) ...: performed with 74237 gas
executing LilypadOnChainJobCreator.setControllerAddress (tx: 0xf6e56d1c638cde1f4a1f19a221caff652286e281c3af161122a77ae510a7fd9a) ...: performed with 29499 gas
deploying "LilypadController" (tx: 0xce20edf2cf36d846f263148285a2713c292c2ecbddaf2f36d538dbabdb831be2)...: deployed at 0x47b8399a8A3aD9665e4257904F99eAFE043c4F50 with 4910452 gas
executing LilypadController.initialize (tx: 0x0bf783e3fb41f869db63163b5984a53a4cfa6b511af10616cdba8546150b9d47) ...: performed with 210921 gas
executing LilypadStorage.setControllerAddress (tx: 0x02dc01236956bcf987191ae6d1a38ac6d5de525f158a3404669fe062b1956f6e) ...: performed with 29555 gas
executing LilypadPayments.setControllerAddress (tx: 0xf26882657e9283963a12d7ef6822d8068d0dd75cb690e2b733b1567164d56931) ...: performed with 29577 gas
executing LilypadMediationRandom.setControllerAddress (tx: 0xa13a09889accfee37f75e00cce7133d0877d5ad3c7360803047091f4ff7a85d1) ...: performed with 29555 gas
```
To fund eth for gas, run:

`npx hardhat run scripts/fund-services-ether.ts`

Example Output (your addresses and transaction will be different):

```
Moved 1000000000000000000 ETHER from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to faucet (0x9162b48910e12079c089477def4384312f0a6e00) - 0x657ceb29e3fda50009c11cca7d97dbde0915fa5a5e08ff84300ff9149affda58.
Moved 1000000000000000000 ETHER from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to solver (0xe05e1b71955da4934938a6b16f97e1db9de6b764) - 0x399214a2df4e586d7d8408c2c1d3d5ed49927fc5ada6000364ba0a6ebe992f88.
Moved 1000000000000000000 ETHER from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to mediator (0xec940f15ef63d1320e5d7a981e875892ec59a8a1) - 0x5c4cf1678066890768606168297391c1df79412be38be78b02edf440428cfdd1.
Moved 1000000000000000000 ETHER from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to resource_provider (0x59cc882e59263f60aac92f62ffba42c23e1061e8) - 0x91c4babda499a8602cc1dc72b69475a8fe006d257b9bc55fcc4845695ba3f31f.
Moved 1000000000000000000 ETHER from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to job_creator (0x8920ecb1150465d36c6a721521a815c567884646) - 0x0d29e3967bfffe43607e0e83743a59933d7baf72e5c12bf8a1ecec398fb6d95d.
Moved 1000000000000000000 ETHER from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to directory (0x851f226a872f97015ef59695b2d785c595b3a890) - 0xb6235aa3b878471c080f7d2a12a598a22cd3d2d869ea8fee305684d6308595c4.
Moved 1000000000000000000 ETHER from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to user (0x1da99b9e884C9e7B15361957577978c1fa66AfBb) - 0xe97f083ec0a82c370a02cd63fb259b33b35bdbd546a0025f2e743237377f0099.
```
To fund tokens, run:

`npx hardhat run scripts/fund-services-tokens.ts` 

Example Output (your addresses and transaction will be different):

```
Moved 100000000000000000000000 TOKENS from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to faucet (0x9162b48910e12079c089477def4384312f0a6e00) - 0x6a3215656ba25345c8a321d15a18bd39d6dc7b5e5ea11cf4a156249d5c0735e1.
Moved 100000000000000000000000 TOKENS from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to solver (0xe05e1b71955da4934938a6b16f97e1db9de6b764) - 0xb0d07a5a1e0ae3e3757ad32365328069df09fcb85ea8e32fe05cf4648e8d4b0c.
Moved 100000000000000000000000 TOKENS from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to mediator (0xec940f15ef63d1320e5d7a981e875892ec59a8a1) - 0x6ff9bea2baff2306cbdb30ce65fdc589568944d450d59f45ee620494288d684d.
Moved 100000000000000000000000 TOKENS from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to resource_provider (0x59cc882e59263f60aac92f62ffba42c23e1061e8) - 0xfb0fb60bfe05ea7458302687786f844e8901adc5fb5c9b783b4d3c5b28fc8e1f.
Moved 100000000000000000000000 TOKENS from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to job_creator (0x8920ecb1150465d36c6a721521a815c567884646) - 0xa3e66b64f87fa76b0755d4e9da99d5b3618a242ca16979d857f00b0e616ffa4f.
Moved 100000000000000000000000 TOKENS from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to directory (0x851f226a872f97015ef59695b2d785c595b3a890) - 0xca41803c232cd01a5903d23e49f1f0897f2d81fc6e9d376dc66d0a5f82122887.
Moved 100000000000000000000000 TOKENS from admin (0xd10d15cc705f7d2558352b1212a9b3685155d93d) to user (0x1da99b9e884C9e7B15361957577978c1fa66AfBb) - 0xcb416bdca70e29e68a2389a03c45fb0784a4e6756e316ff00b9ea674446244ae.
```

To verify accounts are funded, run:

`npx hardhat run scripts/balances.ts`

Example Output (your addresses and amounts will be different):

```
admin - 0xd10d15cc705f7d2558352b1212a9b3685155d93d
  * ETHER: 3.99784186831922
  * TOKENS: 999300000000000000000000000
faucet - 0x9162b48910e12079c089477def4384312f0a6e00
  * ETHER: 1.0
  * TOKENS: 100000000000000000000000
solver - 0xe05e1b71955da4934938a6b16f97e1db9de6b764
  * ETHER: 1.0
  * TOKENS: 100000000000000000000000
mediator - 0xec940f15ef63d1320e5d7a981e875892ec59a8a1
  * ETHER: 1.0
  * TOKENS: 100000000000000000000000
resource_provider - 0x59cc882e59263f60aac92f62ffba42c23e1061e8
  * ETHER: 1.0
  * TOKENS: 100000000000000000000000
job_creator - 0x8920ecb1150465d36c6a721521a815c567884646
  * ETHER: 1.0
  * TOKENS: 100000000000000000000000
directory - 0x851f226a872f97015ef59695b2d785c595b3a890
  * ETHER: 1.0
  * TOKENS: 100000000000000000000000
user - 0x1da99b9e884C9e7B15361957577978c1fa66AfBb
  * ETHER: 1.0
  * TOKENS: 100000000000000000000000
```
That's it. The contracts are deployed and required accounts are funded.

To get a list of deployed contracts, run:

`npx hardhat run scripts/balances.ts`

Example Output (yours will be different):
```
export WEB3_RPC_URL=https://sepolia-rollup.arbitrum.io/rpc
export WEB3_CONTROLLER_ADDRESS=0xF2fD1B9b262982F12446149A27d8901Ac68dcB59
export WEB3_TOKEN_ADDRESS=0x0AabEbf21Cc4591475bb64e105267b3d55C7e0f1
export WEB3_MEDIATION_ADDRESS=0x7B49d6ee530B0A538D26E344f3B02E79ACa96De2
export WEB3_JOBCREATOR_ADDRESS=0x691ddE4710baFA7D0d7c02da32FeFB513CC870fA
export WEB3_PAYMENTS_ADDRESS=0x6f982192B4D225D67ED52A31308a602438882D6b
export WEB3_STORAGE_ADDRESS=0xEB383625D7837aD49C1909f74d05a2db247CE0ca
export WEB3_USERS_ADDRESS=0x8FBAC8549fe9cb0630f7E30929C8256b9b67E68f
```

Be sure to update your `WEB3_RPC_URL` with Arbitrum Sepolia RPC end point.

The exports above are intended to be used by Lilypad Services, so make sure they are set wherever you will be running them. For local testing, paste them back into the terminal and append the lines to `~/.bashrc`.

We set contract addresses for `dev` (local development), `devnet`, and `testnet` in configuration files that are embedded in the binary: https://github.com/Lilypad-Tech/lilypad/tree/main/pkg/options/configs. Update these files if the deployed contracts will be used in one of these environments.

This environment variable is needed by Lilypad services, so be sure it's set as well. 

```export WEB3_CHAIN_ID=421614```


