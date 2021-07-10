# monolith
The whole back end of the Adalo app

## setup
You need go. It's a go app.

I'm not doing anything crazy so `go get` should work just fine.

I am using a postgresql database and migrate to set that up. You will need to install both psql and migrate for your operating system, as well as docker if you want to use the command in the makefile to start the database up.

I also included swagger documentation support for the api. If you want to edit or contribute to that, you'll need to install go-swagger. The `swagger.yml` file should allow you to create a client for whatever service you want to call this from - pretty much every major language has a swagger client generation tool nowadays so I hope this helps. 

To run, run `make run`. In order for it to be usable you will need environment variable configurations. These can come from a `.env` file in the root directory if you wish. Reach out to riley to get that. 

I can also provide the `launch.json` that I use with VS Code. 


## important note
the swagger file generated needs to be editted in order to have the correct server to point to. Please add (for the time being):

```
host: <the server url>
```
to the root configuration after generation. Yes, it's Swagger 2.0. Don't @ at me about that - when go-swagger ups it properly, I'll follow suit.