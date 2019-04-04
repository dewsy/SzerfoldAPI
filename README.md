# SzerfoldGO
A REST API server written in go, providing the backend for the upcoming Szeretetfoldje app (also written by me).

The bulk of the code is generated via Swagger codegen (see .yml for structure), however, the business logic is implemented by me.

I chose GO for it's felxibility, scalability, and easy to understand syntax.

Critical routes of the API are protected with token authentication.

Behind the server is a PostgresQL database. 
