## Go Crawler

It is a bot to list a product on some eccomerces.

## Running

You may run it localy using the command `go run main.go` or build this and run the binary using:
```
go build main.go
```

Then use
```
./main
```

Also, you can dowload the binary [here](https://github.com/gabrieldebem/go-crawler/releases/tag/latest)

After that, a new application will be running on port `8080` on your localhost.

## Usage

To get a json listing a product you should access `localhost:8080/api/products`

And pass a query param named `product`.

Your Url should look like this `localhost:8080/api/products?product=notebook`

This route will return a list of products with `Name`, `Price` and the `Link` to buy it.

