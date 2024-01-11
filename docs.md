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
