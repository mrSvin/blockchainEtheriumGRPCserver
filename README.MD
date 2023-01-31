<h1 align="center">Реализация gRPC сервера для взаимодействия с  blockchain кошельками на базе контракта etherium</h1>

## Реализованные функции
- Контракт solidity для создания кошелька с n балансом, получение информации по созданным кошелькам, выполнение отправки валюты с одного кошелька на другой
- Сервер gRPC для взаимодействия с контрнактом для создания, кошелька, получение информации о нем, отправка валюты между кошшельками
- Плавная остановка сервера по завершению работы


## Download dependencies
```
go get all
```

## Создание билда и генерирование контракта для взаимодействия с go
Build abi, bin & auto-generate SimpleStorage.go file
```
solcjs --optimize --abi --bin ./SimpleStorage.sol -o build
mkdir api
abigen --bin=./build/Store_sol_Store.bin --abi=./build/Store_sol_Store.abi --pkg=api --out=./api/Store.go

```

## Ganache
Запустить Ganache Etherium

## Необходимо с Ganache взять переременнные и установить их в переменную среду app.env
```
gateway = <адрес с портом etherium>
accountPrivateKey = <Приватный ключ аккаунта etherium>
accountHexAddress = <Hex адресс используемого аккаунта>
smartContractHexAddress = <Hex адресс задеплоиного контракта>
```
## Взаимодействие с gRPC со стороны клиента
Репозиторий клиента для взаимодействие с сервером:
```
cd cmd/pkg/app
```