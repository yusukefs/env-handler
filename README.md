# env-handler

## Installation

```sh
go get github.com/yuuuuwwww/env-handler
go install github.com/yuuuuwwww/env-handler
```

## Usage

### Generate Sample Env File from Existing Env file

If you have .env file:

```sh
env-handler gen-sample
```

You also can specify filename:

```sh
env-handler gen-sample -e "different-filename.env"
```

### Add New Environment Variable

You can add a new environment variable, e.g., `key=value`:

```sh
env-handler add key value
```

appending new line `key=value` to `.env` and `.env.sample`.

You also can specify filename.

## Dev References

- [Golangでcliツール作るならこれで！](https://qiita.com/gatchan0807/items/4ffdf65f7affe8faec5a)
- [Golangで自分用のcli翻訳ツール作ってみた](https://qiita.com/gatchan0807/items/83ebf6fec4790d7a6130)

