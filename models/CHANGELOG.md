# Changelog

We document all notable changes to this project in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed

* Downgrades klog/logr. The upgrade was unintentional and breaks downstream
  projects that depend on relay-client-go/models (particularly
  controller-runtime projects).

## [1.0.4] - 2021-10-06

### Added

* Added an option to include a WorkflowExecutionSink in the tenant mapper.

## [1.0.3] - 2021-10-06

### Removed

* Removed legacy workflow handling used by workflow runs.

## [1.0.2] - 2021-09-21

### Added

* Initial workflow mapping support.

## [1.0.1] - 2021-08-24

### Fixed

* Finalize migration from [puppetlabs/relay-core](https://github.com/puppetlabs/relay-core).

## [1.0.0] - 2021-08-23

### Added

* Initial migration from [puppetlabs/relay-core](https://github.com/puppetlabs/relay-core).

[Unreleased]: https://github.com/puppetlabs/relay-client-go/compare/models/v1.0.4...HEAD
[1.0.4]: https://github.com/puppetlabs/relay-client-go/compare/models/v1.0.3...models/v1.0.4
[1.0.3]: https://github.com/puppetlabs/relay-client-go/compare/models/v1.0.2...models/v1.0.3
[1.0.2]: https://github.com/puppetlabs/relay-client-go/compare/models/v1.0.1...models/v1.0.2
[1.0.1]: https://github.com/puppetlabs/relay-client-go/compare/models/v1.0.0...models/v1.0.1
[1.0.0]: https://github.com/puppetlabs/relay-client-go/compare/dbd4bbfeab459f0f38cad0e56a76eefc0fe78be7...models/v1.0.0
