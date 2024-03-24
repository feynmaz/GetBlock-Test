# GetBlock-Test

## Алгоритм обращений к GetBlock.io API
1. Получить номер последнего блока
```js
{
    "jsonrpc": "2.0",
    "method": "eth_blockNumber",
    "params": [],
    "id": "getblock.io"
}
```
2. Получить хеш последнего блока
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

3. N раз:
    1. Получить блок с транзакциями по хешу
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
    2. хеш = "Result"."ParentHash"