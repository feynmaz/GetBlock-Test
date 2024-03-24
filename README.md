# GetBlock-Test

Зарегистрироваться на GetBlock.io, используя API GetBlock создать сервис, выводящий адрес, баланс которого изменился (в любую сторону) больше остальных за последние сто блоков.

Получить номер последнего блока можно с помощью следующего метода:

https://getblock.io/docs/eth/json-rpc/eth_eth_blocknumber/

А данные блока вместе с транзакциями через

https://getblock.io/docs/eth/json-rpc/eth_eth_getblockbynumber/

> Важно: API ключи в репозитории не хранить, не использовать пакеты go-ethereum и подобные ему

## Алгоритм обращений к GetBlock.io API
> Реализовано в adapters/getblock/adapter.go
1. Получить номер последнего блока.
```js
{
    "jsonrpc": "2.0",
    "method": "eth_blockNumber",
    "params": [],
    "id": "getblock.io"
}
```
2. Получить хеш последнего блока.
```js
{
    "jsonrpc": "2.0",
    "method": "eth_getBlockByNumber",
    "params": [
        "<номер_блока>",
        false
    ],
    "id": "getblock.io"
}
```
хеш = "Result"."Hash".

3. N раз, где N - количество последних блоков, заданное пользователем (в cmd/main.go):
    1. Получить блок с транзакциями по хешу.
    ```js
    {
		"jsonrpc": "2.0",
		"method": "eth_getBlockByHash",
		"params": [
			"<хеш>",
			true
		],
		"id": "getblock.io"
	}
    ```
    2. хеш = "Result"."ParentHash".


Количество запросов к API равно
> Количество блоков + 2

## Алгоритм нахождения адреса с наибольшим изменением
> Реализовано в balance/service.go

В мапе хранятся адреса с балансом. Ключ - адрес, значение - баланс в Gwei.
```
map[string]*big.Int
```

1. Для каждой транзакции:
    1. В мапе с балансом уменьшить значение с ключом Transaction.From на Transaction.Value.
    2. В мапе с балансом увеличить значение с ключом Transaction.To на Transaction.Value.
2. Из мапы с балансом получить элемент с наибольшим модулем значения. Вернуть его ключ.
