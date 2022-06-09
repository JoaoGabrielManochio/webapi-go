Exemplo de request para verificaçao da api

User

Metodo addUser POST: http://localhost:8080/api/v1/user

{
	"name": "João Gabriel PJ",
	"email": "joao1234@teste.com.br",
	"password": "MTIzNDU2",
	"cpf_cnpj": "82.756.110/0001-42"
}

Metodo getUser GET: http://localhost:8080/api/v1/user/{{id}}

Metodo getUsers GET: http://localhost:8080/api/v1/user

Metodo updateUser PUT: http://localhost:8080/api/v1/user

{
	"id": 1,
	"name": "João Gabriel TESTADO2",
	"email": "joao@teste.com.br",
	"password": "NjU0MzIx",
	"cpf_cnpj": "108.744.500-04"
}

Metodo deleteUser DELETE: http://localhost:8080/api/v1/user

Wallet

Metodo addWallet POST: http://localhost:8080/api/v1/wallet

{
	"name": "Default PJ",
	"user_id": 3,
	"value": 50.00
}

Metodo getWallet GET: http://localhost:8080/api/v1/wallet/{{id}}

Metodo getWallets GET: http://localhost:8080/api/v1/wallet

Metodo updateWallet PUT: http://localhost:8080/api/v1/wallet

{
	"id": 1,
	"name": "Default2",
	"user_id": 10,
	"value": 10000.50
}

Metodo deleteWallet DELETE: http://localhost:8080/api/v1/wallet

Transaction

Metodo addTransaction POST: http://localhost:8080/api/v1/transaction

{
	"value":1.00,
	"payer_id": 1,
	"payer_receive_id": 2
}

Metodo getTransaction GET: http://localhost:8080/api/v1/transaction/{{id}}

Metodo getTransactions GET: http://localhost:8080/api/v1/transaction
