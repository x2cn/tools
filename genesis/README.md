# Genesis

Helps you form a genesis file.

* [Running](#running)
* [API](#api)

## Running

```
genesis [-min-wait 10s] [-min-vals 10] [-template path-to-template]

Examples:
        genesis -min-vals 10

Flags:
  -min-vals int
        Form a genesis file after 10 validators joined.
  -min-wait string
        Form a genesis file after the given duration.
  -template
        Path to the genesis file template (default: "./genesis.json.tmpl")
```

## API

### Add a validator

```
POST /validator
```

Arguments:

|----------------------+--------+-------------------------------------------------------+------------------------------------|
| name                 | string | validator's name                                      | ec2-new-york-1                     |
| amount               | int    | validator's amount                                    | 10                                 |
| pub_key              | []byte | validator's public key                                |                                    |
| push_genesis_file_to | string | URL to send POST request with a ready genesis file to | http://localhost:46657/api/genesis |
|----------------------+--------+-------------------------------------------------------+------------------------------------|

Returns:

* 200 - success
* 405 - can't add because genesis file have been formed already

### Get genesis file

```
GET /genesis.json
```

Returns:

* 200 - success
* 404 - genesis file has not been formed yet
