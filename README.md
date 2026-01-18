
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
>migrate create -ext sql -dir db/migration init_sechema

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
