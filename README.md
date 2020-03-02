# Zendesk Code Challenge

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
