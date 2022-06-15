# Changelog

We document all notable changes to this project in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.5.0] - 2022-06-15

### Changed

* Updated to building with v6.0.0 of openapi-generator, plus https://github.com/OpenAPITools/openapi-generator/pull/11837 which fixes nested OneOf/AllOf declarations needed by many of our clients.
* API client names have generally all changed in a backwards-incompatible way to avoid naming conflicts. Model names have mostly remained the same.

### Added

* Access control api client and models
* Connections api client and models
* Content api client and models
* Folders api client and models
* Workflows api client and models

## [0.4.6] - 2022-03-03

### Added
* Added step messages.
* Added billing plans and prices.

### Build
* Fixed build under Go 1.17.
* Automatically formatted code and slightly improved the `oneOf` template to support a few more models and APIs.

## [0.4.5] - 2021-11-15

### Added
* Added workflow step output sensitivity.

## [0.4.4] - 2021-10-28

### Added
* Added reference URLs to workflow run POST endpoint response

## [0.4.3] - 2021-10-21

### Added
* Added step decorator schemas.

## [0.4.2] - 2021-09-01

### Changed
* Output fields `name` and `value` are now `required`.

## [0.4.1] - 2021-09-01

### Fixed
* Use `--additional-properties useOneOfDiscriminatorLookup=true` for improved `oneOf` handling.

## [0.4.0] - 2021-08-31

### Fixed
* Added `required` fields to workflow run creator models.
* Fix `oneOf` generation using forked build.

### Added
* Added new revision source models.

### Changed
* Updated [openapi-generator](https://github.com/OpenAPITools/openapi-generator) to [v5.2.1](https://github.com/OpenAPITools/openapi-generator/releases/tag/v5.2.1).

## [0.3.0] - 2021-08-30

### Added
* Add client generation for updated events.

### Fixed
* Added event key to event payload.

## [0.2.0] - 2021-08-26

### Added

* Added client generation for updated workflow runs.
* Added client generation for connection provider models.
* Added client generation for updated workflow secrets.
* Expanded client generation for existing specification.

## [0.1.0] - 2021-08-23

### Added

* Client generation for workflow revision specification changes.
* Client generation for workflow view specification changes.
* Initial migration from [puppetlabs/relay](https://github.com/puppetlabs/relay).

[Unreleased]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.5.0...HEAD
[0.5.0]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.4.6...client/v0.5.0
[0.4.6]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.4.6...client/v0.4.6
[0.4.5]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.4.4...client/v0.4.5
[0.4.4]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.4.3...client/v0.4.4
[0.4.3]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.4.2...client/v0.4.3
[0.4.2]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.4.1...client/v0.4.2
[0.4.1]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.4.0...client/v0.4.1
[0.4.0]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.3.0...client/v0.4.0
[0.3.0]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.2.0...client/v0.3.0
[0.2.0]: https://github.com/puppetlabs/relay-client-go/compare/client/v0.1.0...client/v0.2.0
[0.1.0]: https://github.com/puppetlabs/relay-client-go/compare/dbd4bbfeab459f0f38cad0e56a76eefc0fe78be7...client/v0.1.0
