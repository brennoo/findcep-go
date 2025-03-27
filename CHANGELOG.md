# Changelog

## 0.1.0-alpha.4 (2025-03-27)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/brennoo/findcep-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Bug Fixes

* **test:** return early after test failure ([#54](https://github.com/brennoo/findcep-go/issues/54)) ([03e075c](https://github.com/brennoo/findcep-go/commit/03e075cebb5b59168afdfd19edbbe618b6ab649f))


### Chores

* **docs:** improve security documentation ([#52](https://github.com/brennoo/findcep-go/issues/52)) ([c5d32d9](https://github.com/brennoo/findcep-go/commit/c5d32d97e82924d4d350dd8dcce667a0d1365952))
* fix typos ([#55](https://github.com/brennoo/findcep-go/issues/55)) ([fd1e8ee](https://github.com/brennoo/findcep-go/commit/fd1e8ee26d5d5ce0ec3d77f187e5871c94dbb3dd))

## 0.1.0-alpha.3 (2025-03-14)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/brennoo/findcep-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* add SKIP_BREW env var to ./scripts/bootstrap ([#46](https://github.com/brennoo/findcep-go/issues/46)) ([488e875](https://github.com/brennoo/findcep-go/commit/488e875936d1a18630d45128a78ce7fb837808a8))
* **client:** accept RFC6838 JSON content types ([#47](https://github.com/brennoo/findcep-go/issues/47)) ([4f76c19](https://github.com/brennoo/findcep-go/commit/4f76c1925f3543b10dffaeb3ab7e61911a2eb75f))
* **client:** allow custom baseurls without trailing slash ([#45](https://github.com/brennoo/findcep-go/issues/45)) ([df7e53e](https://github.com/brennoo/findcep-go/commit/df7e53ec69b596f1771aa089073c5ef71e96126a))
* **client:** improve default client options support ([#49](https://github.com/brennoo/findcep-go/issues/49)) ([8deff45](https://github.com/brennoo/findcep-go/commit/8deff4568dc6146539769a1c2048e118483d56d8))


### Chores

* **internal:** remove extra empty newlines ([#50](https://github.com/brennoo/findcep-go/issues/50)) ([5caa453](https://github.com/brennoo/findcep-go/commit/5caa453f06ab3ff064eaaf8477328b037842d04b))


### Documentation

* update URLs from stainlessapi.com to stainless.com ([#42](https://github.com/brennoo/findcep-go/issues/42)) ([6b0954d](https://github.com/brennoo/findcep-go/commit/6b0954dc0c2f291d4dd586ba45ba4d04499be99b))


### Refactors

* tidy up dependencies ([#48](https://github.com/brennoo/findcep-go/issues/48)) ([dfba88f](https://github.com/brennoo/findcep-go/commit/dfba88f0c456b8c4a09c20c593c3150fc46918db))

## 0.1.0-alpha.2 (2025-02-22)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/brennoo/findcep-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Bug Fixes

* **client:** don't truncate manually specified filenames ([#40](https://github.com/brennoo/findcep-go/issues/40)) ([b768020](https://github.com/brennoo/findcep-go/commit/b76802041dda18d22bdad29538dfd74d13d66ee9))
* do not call path.Base on ContentType ([#38](https://github.com/brennoo/findcep-go/issues/38)) ([bccaaaf](https://github.com/brennoo/findcep-go/commit/bccaaaf8ad6b216f493e85cfe925932a8d61ebbd))


### Chores

* **internal:** fix devcontainers setup ([#41](https://github.com/brennoo/findcep-go/issues/41)) ([83471a7](https://github.com/brennoo/findcep-go/commit/83471a7ef8ca1afd6a9949d595fa6e39f34e8bad))

## 0.1.0-alpha.1 (2025-02-08)

Full Changelog: [v0.0.1-alpha.8...v0.1.0-alpha.1](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.8...v0.1.0-alpha.1)

### Features

* **client:** send `X-Stainless-Timeout` header ([#33](https://github.com/brennoo/findcep-go/issues/33)) ([2f27c18](https://github.com/brennoo/findcep-go/commit/2f27c18816e67e73e8518061398d886460a19017))


### Bug Fixes

* fix early cancel when RequestTimeout is provided for streaming requests ([#35](https://github.com/brennoo/findcep-go/issues/35)) ([bf64dfa](https://github.com/brennoo/findcep-go/commit/bf64dfa5c7f92cca9e2084c615059722c08d57e1))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#34](https://github.com/brennoo/findcep-go/issues/34)) ([9d3499a](https://github.com/brennoo/findcep-go/commit/9d3499a9eb773b77d75d5f71d46e3f5353e13b82))
* **internal:** codegen related update ([#31](https://github.com/brennoo/findcep-go/issues/31)) ([c2db1b3](https://github.com/brennoo/findcep-go/commit/c2db1b38ea244d9c6b08f92ed3ed5e9a953982fe))
* minor change to tests ([#36](https://github.com/brennoo/findcep-go/issues/36)) ([2bd9a45](https://github.com/brennoo/findcep-go/commit/2bd9a45d65f90938751cd2ca21ef0c05b2cb0a6a))

## 0.0.1-alpha.8 (2025-02-01)

Full Changelog: [v0.0.1-alpha.7...v0.0.1-alpha.8](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.7...v0.0.1-alpha.8)

### Bug Fixes

* fix unicode encoding for json ([#27](https://github.com/brennoo/findcep-go/issues/27)) ([5984217](https://github.com/brennoo/findcep-go/commit/5984217380f2b02f3ceeaddb8a0d329b29822265))

## 0.0.1-alpha.7 (2025-01-29)

Full Changelog: [v0.0.1-alpha.6...v0.0.1-alpha.7](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.6...v0.0.1-alpha.7)

### Chores

* **internal:** codegen related update ([#23](https://github.com/brennoo/findcep-go/issues/23)) ([dceba7d](https://github.com/brennoo/findcep-go/commit/dceba7dfe8b50729e0df7c83d1ca241d28c7dd9d))
* refactor client tests ([#25](https://github.com/brennoo/findcep-go/issues/25)) ([96fbcc7](https://github.com/brennoo/findcep-go/commit/96fbcc77b9d305d1dbd12930b27730021b4ab917))

## 0.0.1-alpha.6 (2025-01-21)

Full Changelog: [v0.0.1-alpha.5...v0.0.1-alpha.6](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.5...v0.0.1-alpha.6)

### Chores

* **internal:** codegen related update ([#20](https://github.com/brennoo/findcep-go/issues/20)) ([ef5c14c](https://github.com/brennoo/findcep-go/commit/ef5c14ce644d5b5d70fec33db8d0200d9bbb57c6))

## 0.0.1-alpha.5 (2025-01-09)

Full Changelog: [v0.0.1-alpha.4...v0.0.1-alpha.5](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.4...v0.0.1-alpha.5)

### Chores

* **internal:** codegen related update ([#17](https://github.com/brennoo/findcep-go/issues/17)) ([c342e62](https://github.com/brennoo/findcep-go/commit/c342e62830135590ffa42c846e6e142b2fa69bc3))

## 0.0.1-alpha.4 (2025-01-08)

Full Changelog: [v0.0.1-alpha.3...v0.0.1-alpha.4](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.3...v0.0.1-alpha.4)

### Chores

* **internal:** codegen related update ([#15](https://github.com/brennoo/findcep-go/issues/15)) ([37719da](https://github.com/brennoo/findcep-go/commit/37719da1417436155e58cbca54509e78a2f22a14))
* **internal:** update examples ([#13](https://github.com/brennoo/findcep-go/issues/13)) ([836d889](https://github.com/brennoo/findcep-go/commit/836d88961b03698dc066e28a81efe4bfeaf07ebf))

## 0.0.1-alpha.3 (2025-01-02)

Full Changelog: [v0.0.1-alpha.2...v0.0.1-alpha.3](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.2...v0.0.1-alpha.3)

### Chores

* **internal:** codegen related update ([#11](https://github.com/brennoo/findcep-go/issues/11)) ([e0d49b7](https://github.com/brennoo/findcep-go/commit/e0d49b701c2f9d6a4c46a73074f69494f5c72bec))
* **internal:** codegen related update ([#9](https://github.com/brennoo/findcep-go/issues/9)) ([15f9814](https://github.com/brennoo/findcep-go/commit/15f9814476aab9f824f27ae614aef459c5b16145))

## 0.0.1-alpha.2 (2024-11-12)

Full Changelog: [v0.0.1-alpha.1...v0.0.1-alpha.2](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.1...v0.0.1-alpha.2)

### Chores

* rebuild project due to codegen change ([#5](https://github.com/brennoo/findcep-go/issues/5)) ([a5c8915](https://github.com/brennoo/findcep-go/commit/a5c89156bdd5b2ce8cdc782d82ae4e29ac1429cf))
* rebuild project due to codegen change ([#7](https://github.com/brennoo/findcep-go/issues/7)) ([c7e32c9](https://github.com/brennoo/findcep-go/commit/c7e32c9aed394fe54545362ac19019f1d5b6b480))

## 0.0.1-alpha.1 (2024-07-22)

Full Changelog: [v0.0.1-alpha.0...v0.0.1-alpha.1](https://github.com/brennoo/findcep-go/compare/v0.0.1-alpha.0...v0.0.1-alpha.1)

### Chores

* configure new SDK language ([99dde31](https://github.com/brennoo/findcep-go/commit/99dde31c9e2209bc3d7a506fa87deca2fd6d7605))
* go live ([#1](https://github.com/brennoo/findcep-go/issues/1)) ([7d9c522](https://github.com/brennoo/findcep-go/commit/7d9c52288811023b65270f13d3348108dbd7ad49))
