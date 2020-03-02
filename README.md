# Zendesk Code Challenge

## How does it application work



## How to

### How to run

```bash

./auto/run

```

Or without Docker

```bash

go run main.go

```

Please use your arrow keys to select search data type and field name from all fields, then type the value for that field.

If there are something matched, the application will printed out all the information, including linked `User` and `Organisation` information.

Otherwise you should expect a error message says `Cannot find any matched result`

    Notice: Due to the limited time, currently you can't go back in the CLI, if you want to skip the current search then just type `Enter`.



###  How to test

```bash

./auto/test

```

Or without Docker

```bash

go test ./... -cover

```

### How to build Docker image

```bash

GITHUB_RUN_NUMBER=1 ./auto/build

```

## If I had more time / What can I improve

- Better object printing instead of using `spew.Dump()`
- Reuse more shard logic especially in `search/index.go`, but due to Go missing generics concept it's different when compare to `Ruby`, `Java` and some other popular languages.
- Add more tests, due to this code test costs me a lot of time to complete, so I didn't cover all the code currently, but I am keen to add more tests if possible.
- Better user experience in CLI application, current the cli app is very basic and go can't go back to previous menu.
