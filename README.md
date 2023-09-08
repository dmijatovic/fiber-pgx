# Fiber and Postgres 15

This project is created to benchmark postgres drivers with golang fiber.
The ideas is based on my experience during making todo-bench project where I noticed that postgres and the speed of the driver high likely plays important role in the overall results of my benchmark. In that project Golang did not score particulairly well compared to rust/actix. I have created 2 golang projects, one with Fiber and one with Mux. Both had quite comparable score. Looking at some other benchmarks the fiber performance should be higher than mux performance. One of my conclusion is that the performance of golang projects is limited by the pg driver I used.

Then I read this [article, which confirms my assumptions](https://levelup.gitconnected.com/fastest-postgresql-client-library-for-go-579fa97909fb#:~:text=By%20looking%20at%20the%20result,based%20on%20your%20own%20environment.)

Based on the scores in that article I decided to try to benchmark few pg drivers for golang, starting with the most performant pgx and widely used pg.

## Requirements

This project requires Postgres server running. Use docker compose file in the postgres folder

```bash
# navigate to postgres folder
cd postgres
# start postgres with required user and database
docker compose up -d
```
