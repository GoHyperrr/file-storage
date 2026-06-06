# file-storage — Hyperrr Pluggable File Storage Module

[![Go Reference](https://pkg.go.dev/badge/github.com/GoHyperrr/file-storage.svg)](https://pkg.go.dev/github.com/GoHyperrr/file-storage)

This repository contains the pluggable file storage bucket adapters for the Hyperrr engine using standard Go Cloud Development Kit (`gocloud.dev/blob`) abstractions.

---

## 🛠️ Active Providers

* **S3 Storage Blob**: Pluggable storage adapter supporting AWS S3, Google Cloud Storage, Azure Blob, and local/minio block storage targets.

---

## 🛠️ Usage

This module is imported dynamically by the core Hyperrr engine during startup to support external file storage buckets.

To learn more about how to configure storage backends, see the [Hyperrr Developer Guide](https://github.com/GoHyperrr/hyperrr/blob/main/developer.md).
