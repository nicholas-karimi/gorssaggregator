### About the Project

#### What is an RSS Feed
- A protocol that makes it easy to distribute things like podcasts and blog posts easy.

#### The Server
> It'll allow users to add different RSS feeds to its database. And then we will automatically collect all the posts from the feed, download them and save them in the database for later retrieval/viewing.


### Running
Step 1: Create an executable
`go build`

Build and run the server
`go build && ./gorssaggregator `

#### Make env available manually
When using .env, the secrets might not be available in your current shell session. To make it available
`export PORT=8000` 

To automatically get your configs use the `godotenv` package. 
`go get github.com/joho/godotenv`
then run
`go mod vendor` to have a local copy
 The `vendor` folder will be automatically created. Think of it as *node modules* folder in JS. Its advisable t commit it since in go there are no dependencies-many.
### Clean up imports
Run the `go mod tidy`

### Server Spinning
Use the `chi router` package
Install
`go get github.com/go-chi/chi`
`go get github.com/go-chi/cors`

### Database
Postgres
- install in local machine
#3rd party plugins to interact with database
1. sqlc
To handle queries.
`go install github.com/kyleconroy/sqlc/cmd/sqlc@latest` use `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
[SQLC DOC](https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html)
check version
`sqlc version`
2. goose
To handle migrations
`goose` is a database migration tool. Manage your database schema by creating incremental SQL changes and/or Go functions.
`go install github.com/pressly/goose/v3/cmd/goose@latest`

`goose -version`

#### Run Migrations
Navigate to the sql/schema directory and run the following command to create migrations.
`goose postgres postgres://karimi:@localhost:5432/rssagg up`

Down Migrations
`goose postgres postgres://admin:Incorrect@localhost:5432/rssagg down`


Run Queries
`sqlc generate`


DB Connection
`go get github.com/lib/pq`

UUID
`go get github.com/google/uuid` 

Clean up unsued imports and vendorize them
`go mod tidy, go mod vendor`