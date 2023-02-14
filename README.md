# TOTP

## 準備

```shell
cp data/secret.sample.txt data/secret.txt
vi data/secret.txt
```

## 実行例

```shell
goenv exec go run main.go
```
もしくは
```shell
./script/exec.sh
```

### Docker 使用の場合

```shell
docker-compose run --rm totp go run main.go
```
もしくは
```shell
./script/exec_in_docker.sh
```

