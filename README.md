
# Backend Master Class
https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=1

## [Backend #1] Design DB schema and generate SQL code with dbdiagram.io

https://dbdiagram.io/d


## [Backend #3] How to write & run database migration in Golang
Install migrate
>brew install sqlc
>migrate -verison
v4.11.0
>migrate help

Create files
>migrate create -ext sql -dir db/migration2 -seq init_sechema

Shell to postgres db.
>docker exec -it db-simplebank /bin/sh
>createdb --username=postgres --owner=postgres simple_bank

>psql -U postgres simple_bank
>\q  #quite
>dropdb --username=postgres simple_bank

Run without shell
>docker exec -it db-simplebank createdb --username=postgres --owner=postgres simple_bank
>docker exec -it db-simplebank psql -U postgres simple_bank
>docker exec -it db-simplebank dropdb -U postgres simple_bank


Migrate
>migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"

>migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up


## [Backend #4] Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc
https://docs.sqlc.dev/en/latest/overview/install.html


#### sqlc

https://docs.sqlc.dev/en/latest/overview/install.html

 Install sqlc
 >brew install sqlc
 >sqlc version
 >sqlc help

 sqlc initial file.
 >sqlc init

 Generate go 
 https://docs.sqlc.dev/en/latest/howto/insert.html


Go init
 >go mod init github.com/yuttana76/simbplebank
 >go mod tidy

 sqlc generate command
 >sqlc generate

 make file
 >make sqlc

## [Backend #5] Write Golang unit tests for database CRUD with random data

Install postgrest connection drisver

https://github.com/lib/pq
>go get github.com/lib/pq@latest

After import package to run
>go mod tidy

### Test package
Install
>go get github.com/stretchr/testify

### Run test use make
>make test

## [Backend #6] A clean way to implement database transaction in Golang

Transaction ?
ACID
1.Atomicity(A)
Either all operations complsete successfully or the transaction fails and the db is unchanged.

2.Consistency(C)
The db state must be valid after the transaction.All constraints must be satisfield

3.Isolation(I)
Con current transactions must not affect each other.

4.Durability(D)
Data written by a successful transaction must be recorded in persisten storage.

### UDEMY 9: DB transaction
>docker exec -it <postgrest name> psql -U root -d simple_bank

How handle deadlock
1.ADD quey\account.sql 
-- name: GetAccountforupdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR UPDATE;

2.Run make file to generate sqlc code.
>make sqlc

postgrest lock monitoring

https://wiki.postgresql.org/wiki/Lock_Monitoring


'''
SELECT blocked_locks.pid     AS blocked_pid,
         blocked_activity.usename  AS blocked_user,
         blocking_locks.pid     AS blocking_pid,
         blocking_activity.usename AS blocking_user,
         blocked_activity.query    AS blocked_statement,
         blocking_activity.query   AS current_statement_in_blocking_process
   FROM  pg_catalog.pg_locks         blocked_locks
    JOIN pg_catalog.pg_stat_activity blocked_activity  ON blocked_activity.pid = blocked_locks.pid
    JOIN pg_catalog.pg_locks         blocking_locks 
        ON blocking_locks.locktype = blocked_locks.locktype
        AND blocking_locks.database IS NOT DISTINCT FROM blocked_locks.database
        AND blocking_locks.relation IS NOT DISTINCT FROM blocked_locks.relation
        AND blocking_locks.page IS NOT DISTINCT FROM blocked_locks.page
        AND blocking_locks.tuple IS NOT DISTINCT FROM blocked_locks.tuple
        AND blocking_locks.virtualxid IS NOT DISTINCT FROM blocked_locks.virtualxid
        AND blocking_locks.transactionid IS NOT DISTINCT FROM blocked_locks.transactionid
        AND blocking_locks.classid IS NOT DISTINCT FROM blocked_locks.classid
        AND blocking_locks.objid IS NOT DISTINCT FROM blocked_locks.objid
        AND blocking_locks.objsubid IS NOT DISTINCT FROM blocked_locks.objsubid
        AND blocking_locks.pid != blocked_locks.pid

    JOIN pg_catalog.pg_stat_activity blocking_activity ON blocking_activity.pid = blocking_locks.pid
   WHERE NOT blocked_locks.granted;
'''

'''
SELECT blocked_locks.pid     AS blocked_pid,
         blocked_activity.usename  AS blocked_user,
         blocking_locks.pid     AS blocking_pid,
         blocking_activity.usename AS blocking_user,
         blocked_activity.query    AS blocked_statement,
         blocking_activity.query   AS current_statement_in_blocking_process,
         blocked_activity.application_name AS blocked_application,
         blocking_activity.application_name AS blocking_application
   FROM  pg_catalog.pg_locks         blocked_locks
    JOIN pg_catalog.pg_stat_activity blocked_activity  ON blocked_activity.pid = blocked_locks.pid
    JOIN pg_catalog.pg_locks         blocking_locks 
        ON blocking_locks.locktype = blocked_locks.locktype
        AND blocking_locks.DATABASE IS NOT DISTINCT FROM blocked_locks.DATABASE
        AND blocking_locks.relation IS NOT DISTINCT FROM blocked_locks.relation
        AND blocking_locks.page IS NOT DISTINCT FROM blocked_locks.page
        AND blocking_locks.tuple IS NOT DISTINCT FROM blocked_locks.tuple
        AND blocking_locks.virtualxid IS NOT DISTINCT FROM blocked_locks.virtualxid
        AND blocking_locks.transactionid IS NOT DISTINCT FROM blocked_locks.transactionid
        AND blocking_locks.classid IS NOT DISTINCT FROM blocked_locks.classid
        AND blocking_locks.objid IS NOT DISTINCT FROM blocked_locks.objid
        AND blocking_locks.objsubid IS NOT DISTINCT FROM blocked_locks.objsubid
        AND blocking_locks.pid != blocked_locks.pid
 
    JOIN pg_catalog.pg_stat_activity blocking_activity ON blocking_activity.pid = blocking_locks.pid
   WHERE NOT blocked_locks.GRANTED;
'''

### 11.Deeply understand transaction isolation levels & read
(!Important must keep understand)


### 12. Setup Github Actions for Golang + Postgres to run automated tests
Add progrest service

Install golang migate cli
*https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
** Release Downloads

>curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-amd64.tar.gz | tar xvz

### 13 Implement RESTful HTTP API in Go using Gin

Install gin
go get github.com/gin-gonic/gin

document here
https://github.com/gin-gonic/gin?tab=readme-ov-file


sqlc / validation
https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html#query-verification

Start gin by use make 
>make server

### 14. Load config from file & environment variables in Go with Viper

Use viper  
https://github.com/spf13/viper
Install
>go get github.com/spf13/viper

### 15. Mock DB for testing HTTP API in Go and achieve 100% coverage

gomock
install package
>go install github.com/golang/mock/mockgen@v1.6.0

>which mocgen
mocgen not found

>vi ~/.zshrc
OR >vi ~/.bash_profile
click i insert mode
click esc 
press :wq

>source ~/.zshrc
>which mockgen

User mockgen
>mockgen -help

### Generate code for Store interface to db/mock/store.go
>mockgen -package mockdb -destination db/mock/store.go github.com/yuttana76/simbplebank/db/sqlc Store

### 17. Add users table with unique & foreign key constraints in PostgreSQL
(This for add new model or change model TODO)
1. create dbdiagram  https://dbdiagram.io/d
2. export to postgrestdb
3. create migrate file sqlc
Create files
>migrate create -ext sql -dir db/migration -seq add_users

4.put sql script only added. to next seq file(generated(Step.3))

5.migrate to db
5.1 Run command 
    >make migrateup1 
5.2 Incase has some thing wrong haveto restore back 1 step
>make migratedown1 
(see table schema_migrations. should be version you want to stpe down. And dirty column must be NULL)

### 18. How to handle DB errors in Golang correctly
1.create file db.query.user.sql
2.run >make sqlc 
3.check file.
    (Modify)db.sqlc.models.go
    (Create)db.sqlc.user.sql
4.create user_test.sql

implement api test (mock)
5.run >make mock
6.run >make test

### 19. How to securely store passwords? Hash password in Go with Bcrypt!
(has resources files)

api validator
https://github.com/go-playground/validator

### 20. How to write stronger unit tests with a custom gomock matcher
# Refer to lecture 13
-jwt-> HS256,RSA
-PASETO (Platform-Agnostic SEcurity TOkens) ->AES256

### 22. How to create and verify JWT & PASETO token in Golang

Install jwt
https://github.com/golang-jwt/jwt

>go get github.com/golang-jwt/jwt/v5

PASETO
https://github.com/o1egl/paseto?tab=readme-ov-file#installation
>go get -u github.com/o1egl/paseto

### 23. Implement login user API that returns PASETO or JWT access token in Go

Code error !! Restor to previos commit
>git reset --hard <commit-hash>
>git reset --hard 7d3fcc5a3f469adf921ae0f08799c7f7a8e3d13d

Refer lecture 12
Add app.env
Update /util/config.go

run server
>make server

Test postman 
http://localhost:8080/users/login
{
    "username":"yuttana",
    "password":"password"
}

Modify /db/query/account.sql
    -- name: ListAccounts :many
    SELECT * FROM accounts
    WHERE owner = $1  // Add here
    ORDER BY id
    LIMIT $2 OFFSET $3;

Run sqlc to generate code
>make sqlc

Run mock to generate mock
>make mock

### 24. Implement authentication middleware and authorization rules in Golang using Gin

### 25. How to build a small Golang Docker image with a multistage Dockerfile
create new branch(ft=feature)
>git checkout -b ft/docker


Summary of Commands
Action 	Command
Delete local branch (safe, merged only)	
>git branch -d <branch-name>

Delete local branch (force)	
>git branch -D <branch-name>

Delete remote branch	
>git push origin --delete <branch-name>


update go lang to version=1.16.3
go.mod

Docker hub
https://hub.docker.com/

Command delete images <none>
>docker rmi $(docker images -f "dangling=true" -q)

Delete container
>docker rm <container name>
>docker rm simplebank

Docker build
>docker build -t simplebank:latest .
>dockage images

### 26. How to use docker network to connect 2 stand-alone containers

Test run simplebank image
Ron development mode
>docker run --name simplebank -p 8080:8080 simplebank:latest

Run with production mode.
>docker run --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest

Run with connect db container
>docker run --name simplebank -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://postgres:postgres@172.17.0.2:5432/simple_bank?sslmode=disable" simplebank:latest



>docker container inspect db-simplebank

> docker network ls
>docker network inspect go-simplebank_default

"ConfigOnly": false,
        "Containers": {
            "41afbdc634afec7cbd89cc623343d29132fcbbd3117bd5a915be739d11f5a9e3": {
                "Name": "db-simplebank",
                "EndpointID": "f87d2c02fe0d010e816f12a21445ae308ef4b086740e554c885b0eeae87f04be",
                "MacAddress": "02:42:ac:1a:00:02",
                "IPv4Address": "172.26.0.2/16",
                "IPv6Address": ""
            }
        },

Create bank network
>docker network create bank-network
>docker network connect bank-network  db-simplebank


Run with connect db container
>docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://postgres:postgres@db-simplebank:5432/simple_bank?sslmode=disable" simplebank:latest
