# Changelog

## [2.10.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.9.1...v2.10.0) (2024-12-02)


### Features

* Add base services compose file ([#452](https://github.com/Lilypad-Tech/lilypad/issues/452)) ([86a3bc6](https://github.com/Lilypad-Tech/lilypad/commit/86a3bc6fa13fbe636e2492ddc2573748da6e3f33))
* Add solver metrics ([#435](https://github.com/Lilypad-Tech/lilypad/issues/435)) ([bff3def](https://github.com/Lilypad-Tech/lilypad/commit/bff3deff64a625ed078ff5a8148f059dd4f52b62))
* Demonet network config ([#453](https://github.com/Lilypad-Tech/lilypad/issues/453)) ([732b5ad](https://github.com/Lilypad-Tech/lilypad/commit/732b5ad1990d06ff693b78f8de528c0881389e90))


### Bug Fixes

* Run bacalhau in a separate container ([#438](https://github.com/Lilypad-Tech/lilypad/issues/438)) ([25083a7](https://github.com/Lilypad-Tech/lilypad/commit/25083a70623e7f6c70ef3539deeff5b83ff49080))
* Silence web3 event listener cancellation errors ([#456](https://github.com/Lilypad-Tech/lilypad/issues/456)) ([d679ea3](https://github.com/Lilypad-Tech/lilypad/commit/d679ea3c7890abd7d7fb0c3f382f79dda94c88bd))

## [2.9.1](https://github.com/Lilypad-Tech/lilypad/compare/v2.9.0...v2.9.1) (2024-11-18)


### Bug Fixes

* Solver connection errors ([#427](https://github.com/Lilypad-Tech/lilypad/issues/427)) ([52de132](https://github.com/Lilypad-Tech/lilypad/commit/52de132947a60f45fecb6229011b5c22baf54450))

## [2.9.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.8.0...v2.9.0) (2024-11-05)


### Features

* Add custom retryablehttp error handler ([#423](https://github.com/Lilypad-Tech/lilypad/issues/423)) ([711ae57](https://github.com/Lilypad-Tech/lilypad/commit/711ae579dffba5701f9e5ce8faaae25928fcf363))
* Add deal making trace ([#336](https://github.com/Lilypad-Tech/lilypad/issues/336)) ([a8febc0](https://github.com/Lilypad-Tech/lilypad/commit/a8febc04f55f4ed185f0e1010551a54e638988df))
* Add docker compose for RPs ([#404](https://github.com/Lilypad-Tech/lilypad/issues/404)) ([5b6e598](https://github.com/Lilypad-Tech/lilypad/commit/5b6e5989f2ddf9f528b2ec513c048a5b6bc0d4b9))
* Add job creator CLI run trace ([#418](https://github.com/Lilypad-Tech/lilypad/issues/418)) ([f9ce8a7](https://github.com/Lilypad-Tech/lilypad/commit/f9ce8a78311d9935308a0411e0df7511a73dea2e))
* Add job creator tracer ([#411](https://github.com/Lilypad-Tech/lilypad/issues/411)) ([b467d0f](https://github.com/Lilypad-Tech/lilypad/commit/b467d0fa36f9720038bf5714d7ce1409a375d2a8))
* Add Lilypad version HTTP header ([#408](https://github.com/Lilypad-Tech/lilypad/issues/408)) ([b4e56e3](https://github.com/Lilypad-Tech/lilypad/commit/b4e56e34a4f6b6b3979e9523731b8952162ad143))
* Add solver rate limiter ([#419](https://github.com/Lilypad-Tech/lilypad/issues/419)) ([d30b6d8](https://github.com/Lilypad-Tech/lilypad/commit/d30b6d81e7f315b37d8a2bf61210c6ef19a9877e))
* Upgrade binary to support v4 cardinfof schema ([#412](https://github.com/Lilypad-Tech/lilypad/issues/412)) ([d93c821](https://github.com/Lilypad-Tech/lilypad/commit/d93c82174a152fcad0f3d2d290d137efdb5334b2))


### Bug Fixes

* Add version info to RP Docker build ([#415](https://github.com/Lilypad-Tech/lilypad/issues/415)) ([d0b40cc](https://github.com/Lilypad-Tech/lilypad/commit/d0b40cc41f0cc12b24f8c3e17e38436c7f3c99fe))
* Handle accept result and download result errors ([#416](https://github.com/Lilypad-Tech/lilypad/issues/416)) ([c2c4a11](https://github.com/Lilypad-Tech/lilypad/commit/c2c4a116a273ac9bfc26a7189f8d104492d24a26))

## [2.8.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.7.0...v2.8.0) (2024-10-14)


### Features

* Add get_ethclient trace ([#390](https://github.com/Lilypad-Tech/lilypad/issues/390)) ([edb8fe8](https://github.com/Lilypad-Tech/lilypad/commit/edb8fe8201d09db75a145513528830fec8d701f1))
* Add IPFS client to retrieve results ([#400](https://github.com/Lilypad-Tech/lilypad/issues/400)) ([85f92b9](https://github.com/Lilypad-Tech/lilypad/commit/85f92b9fea2531ecc6f448f4f7122f076f7c212c))


### Bug Fixes

* Lower default pricing and check LP balance ([#403](https://github.com/Lilypad-Tech/lilypad/issues/403)) ([cccc033](https://github.com/Lilypad-Tech/lilypad/commit/cccc033d9d4d936847c55fda8e98ac976bf89794))
* Re-register resource offers after runJob ([#396](https://github.com/Lilypad-Tech/lilypad/issues/396)) ([1cc64f2](https://github.com/Lilypad-Tech/lilypad/commit/1cc64f245b6d9f44d6430953197335215ccdac66))

## [2.7.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.6.0...v2.7.0) (2024-09-26)


### Features

* Add solver HTTP tracing ([#388](https://github.com/Lilypad-Tech/lilypad/issues/388)) ([88fe924](https://github.com/Lilypad-Tech/lilypad/commit/88fe924c3c73016aaa4d94f96666bc1696c561d2))
* Fetch machine specs from bacalhau ([#371](https://github.com/Lilypad-Tech/lilypad/issues/371)) ([d867292](https://github.com/Lilypad-Tech/lilypad/commit/d867292b816fb9665bb8d6d9a30e5da09bf7e9d3))


### Bug Fixes

* Resourceoffer updates ([#381](https://github.com/Lilypad-Tech/lilypad/issues/381)) ([6353895](https://github.com/Lilypad-Tech/lilypad/commit/6353895e5b5103502c49277dbe749d80a0df990b))
* Wrong balance err check ([#382](https://github.com/Lilypad-Tech/lilypad/issues/382)) ([aaf5d9a](https://github.com/Lilypad-Tech/lilypad/commit/aaf5d9a6cb3ea95ede239d58f1d80c12dfd481f8))

## [2.6.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.5.0...v2.6.0) (2024-09-17)


### Features

* Added public addresses with failover ([#365](https://github.com/Lilypad-Tech/lilypad/issues/365)) ([973f396](https://github.com/Lilypad-Tech/lilypad/commit/973f396521f9ed17d5009d5368640e1132b16d6a))
* Card-api-update ([#366](https://github.com/Lilypad-Tech/lilypad/issues/366)) ([0747cef](https://github.com/Lilypad-Tech/lilypad/commit/0747cefecbe2e2bc99d5f02ce12dcfc21337df73))


### Bug Fixes

* Exclude canceled job offers ([#363](https://github.com/Lilypad-Tech/lilypad/issues/363)) ([8b2784f](https://github.com/Lilypad-Tech/lilypad/commit/8b2784f0fc5a5ec8a3b426a098a9d11ab1864763))

## [2.5.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.4.0...v2.5.0) (2024-09-11)


### Features

* Allow address targeting ([#307](https://github.com/Lilypad-Tech/lilypad/issues/307)) ([2d49cc4](https://github.com/Lilypad-Tech/lilypad/commit/2d49cc47a27d0f6c7d866b0bf242688bfc9e92ca))
* Changed rpc to Quicknode  ([#348](https://github.com/Lilypad-Tech/lilypad/issues/348)) ([6db6ec9](https://github.com/Lilypad-Tech/lilypad/commit/6db6ec93d884fa50b9f6f5424278ac940caf4da2))
* Pow post binary ([#316](https://github.com/Lilypad-Tech/lilypad/issues/316)) ([e37df74](https://github.com/Lilypad-Tech/lilypad/commit/e37df74a3c874ada1908ee30e86e4816cf089b85))
* Preflight check for bacalhau version ([#338](https://github.com/Lilypad-Tech/lilypad/issues/338)) ([c9a349c](https://github.com/Lilypad-Tech/lilypad/commit/c9a349c8d834f7f2620aecd0f91c11a1c19c3471))

## [2.4.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.3.1...v2.4.0) (2024-09-04)


### Features

* Add run job traces ([#292](https://github.com/Lilypad-Tech/lilypad/issues/292)) ([bff7bcf](https://github.com/Lilypad-Tech/lilypad/commit/bff7bcf825b819b35efbbd8c0da23c2b4eee5fcd))

## [2.3.1](https://github.com/Lilypad-Tech/lilypad/compare/v2.3.0...v2.3.1) (2024-08-26)


### Bug Fixes

* Gpu support for rp Dockerfile ([#299](https://github.com/Lilypad-Tech/lilypad/issues/299)) ([d830232](https://github.com/Lilypad-Tech/lilypad/commit/d83023239a564c0bdb0abb78242e4212d8c8ef6c))

## [2.3.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.8...v2.3.0) (2024-08-25)


### Features

* Add ETH Balance Check on RP Matching ([#298](https://github.com/Lilypad-Tech/lilypad/issues/298)) ([561a64a](https://github.com/Lilypad-Tech/lilypad/commit/561a64acbc007a210abea3aac2d216955fb3c6c9))

## [2.2.8](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.7...v2.2.8) (2024-08-16)


### Bug Fixes

* Remove ntfs zone files ([#286](https://github.com/Lilypad-Tech/lilypad/issues/286)) ([6b7b2dc](https://github.com/Lilypad-Tech/lilypad/commit/6b7b2dcc3c0e94607e20ee2646d78513ff907cfc))

## [2.2.7](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.6...v2.2.7) (2024-08-07)


### Reverts

* Removes gpc options check ([#282](https://github.com/Lilypad-Tech/lilypad/issues/282)) ([09c5518](https://github.com/Lilypad-Tech/lilypad/commit/09c5518fa8bed7b98881bcde39c0ebfcaad43d7f))

## [2.2.6](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.5...v2.2.6) (2024-08-05)


### Bug Fixes

* Reconnect sc event loop ([#266](https://github.com/Lilypad-Tech/lilypad/issues/266)) ([436363c](https://github.com/Lilypad-Tech/lilypad/commit/436363cbd4a938cc5a45e1c21e0418624975f2f0))

## [2.2.5](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.4...v2.2.5) (2024-07-30)


### Bug Fixes

* Remove auto-updater script ([#270](https://github.com/Lilypad-Tech/lilypad/issues/270)) ([605bfd2](https://github.com/Lilypad-Tech/lilypad/commit/605bfd2fdb407538a39f36955b32428f6623c2d3))

## [2.2.4](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.3...v2.2.4) (2024-07-30)


### Bug Fixes

* Remove DisablePow from PersistentFlags ([69e8e03](https://github.com/Lilypad-Tech/lilypad/commit/69e8e039323def5a21b532e931af45dd3dfcf867))


### Reverts

* "chore: Remove PR title labeling ([#267](https://github.com/Lilypad-Tech/lilypad/issues/267))" ([#268](https://github.com/Lilypad-Tech/lilypad/issues/268)) ([ca76faf](https://github.com/Lilypad-Tech/lilypad/commit/ca76fafef3f4085bf24a913babedd9678b6dad33))

## [2.2.3](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.2...v2.2.3) (2024-07-23)


### Bug Fixes

* Correct pow flag handling in Cobra ([#261](https://github.com/Lilypad-Tech/lilypad/issues/261)) ([8b1f44f](https://github.com/Lilypad-Tech/lilypad/commit/8b1f44ff9d08c289e04bb53ca83b4f87e9041e6b))

## [2.2.2](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.1...v2.2.2) (2024-07-15)


### Bug Fixes

* Websocket disconnect and goroutine leak ([#220](https://github.com/Lilypad-Tech/lilypad/issues/220)) ([cca7510](https://github.com/Lilypad-Tech/lilypad/commit/cca75105df1dbb4523750338493b240902c01988))

## [2.2.1](https://github.com/Lilypad-Tech/lilypad/compare/v2.2.0...v2.2.1) (2024-07-11)


### Bug Fixes

* Set api host on embedded configs ([#253](https://github.com/Lilypad-Tech/lilypad/issues/253)) ([0aa29f3](https://github.com/Lilypad-Tech/lilypad/commit/0aa29f37b037d51bfcf86fe47562f2249a1ebecf))

## [2.2.0](https://github.com/Lilypad-Tech/lilypad/compare/v2.1.1...v2.2.0) (2024-07-10)


### Features

* Metrics hashrate ([#237](https://github.com/Lilypad-Tech/lilypad/issues/237)) ([50d439e](https://github.com/Lilypad-Tech/lilypad/commit/50d439e45113c92f0c1aa695be1e45b50ebba0f3))

## [2.1.1](https://github.com/Lilypad-Tech/lilypad/compare/v2.1.0...v2.1.1) (2024-07-08)


### Bug Fixes

* Clean openzeplin proxy deploy metadata ([#218](https://github.com/Lilypad-Tech/lilypad/issues/218)) ([9798a09](https://github.com/Lilypad-Tech/lilypad/commit/9798a09058ab98e9f2a5cfe3e3bbd4e1bc94a614))
