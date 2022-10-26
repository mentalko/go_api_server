# Finance API server

Реализует простые банковские операции как: пополнение, перевод средств и обмен валют.\


Implements simple banking operations such as: replenishment, transfer of funds and currency exchange.\

## Features
* registration by phone number
* show balance
* top up your account
* transferring money to another user by his login
* exchange money from dollars to rubles and vice versa


## TODO:
- [ ] add swagger


## How to use
1. Run docker-compose
```
docker-compose up --build 
```

## Avalable endpoints: 

GET,POST http://localhost/accounts

GET http://localhost/accounts/{id}


GET, POST http://localhost/transactions

GET http://localhost/topup

GET http://localhost/exchange
