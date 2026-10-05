# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Calendar Versioning](https://calver.org/).

## [Unreleased]

## [2026.10.0] - 2026-10-05

### Updated

- go to 1.25
- prometheus/client_golang to v1.24.1
- sirupsen/logrus to v1.10.2
- spf13/cobra to v1.10.2
- spf13/viper to v1.21.0
- indirect go dependencies to latest and greatest

### Changed

- versioning scheme from SemVer to CalVer

## [0.1.4] - 2024-03-08

### Updated

- prometheus/client_golang to v1.19.0
- indirect go dependencies to latest and greatest

### Changed

- project now uses `task` (<https://taskfile.dev>) instead of `make`

## [0.1.3] - 2024-01-05

### Updated

- prometheus/client_golang to v1.18.0
- spf13/cobra v1.8.0
- spf13/viper to v1.18.2
- indirect go dependencies to latest and greatest

## [0.1.2] - 2023-10-16

### Updated

- prometheus/client_golang to v1.17.0
- spf13/viper to v1.17.0
- indirect go dependencies to latest and greatest

## [0.1.1] - 2023-07-18

### Added

- integration test to test the docker image

### Fixed

- connection error `tls: failed to verify certificate: x509: certificate signed by unknown authority`

## [0.1.0] - 2023-07-18

### Updated

- go from 1.18 to 1.20
- go dependencies to latest and greatest

### Changed

- changed base image from alpine to scratch

## [0.0.4] - 2023-06-06

### Added

- Docker container build published to docker hub, see https://hub.docker.com/r/cryptastic/chia-price-exporter

### Updated

- go from 1.18 to 1.20
- go dependencies to latest and greatest

## [0.0.3] - 2022-06-25

### Changed

- changed exporter port to `9952` as default port allocation,
  see https://github.com/prometheus/prometheus/wiki/Default-port-allocations

## [0.0.2] - 2022-06-21

### Added

- Added github workflow to generate go binaries for different architectures

## [0.0.1] - 2022-06-20

### Added

- This is the first release of the chia price exporter

