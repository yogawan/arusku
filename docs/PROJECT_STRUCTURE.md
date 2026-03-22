# Project Structure

Starter ini mengikuti pola project layout Go yang umum.

## Direktori utama

- `cmd/belajar-golang`: entrypoint executable.
- `internal/app/belajargolang`: flow aplikasi.
- `internal/pkg`: package privat reusable internal.
- `pkg`: package publik (boleh diimpor project lain).
- `configs`: template konfigurasi.
- `api`: spesifikasi OpenAPI.
- `scripts`: helper script build/run.
- `build`: konfigurasi packaging/CI.
- `deployments`: template deployment.
- `test`: test eksternal/integrasi.
