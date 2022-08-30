# firesert examples

Examples demonstrating how to use [firesert](https://github.com/JonnyOrman/firesert).

It contains the following examples:
- `example1` - inserts any data received
- `example2` - uses a struct to define the data to be inserted into Firestore
- `example3` - config retrieved from the environment
- `example4` - config retrieved from both the config file and the environment

# How to run

With Docker installed, clone the repo and run
```
docker-compose up --build
```

To use each of the examples, `POST` to the following:
- `example1` - `localhost:3001`
- `example2` - `localhost:3002`
- `example3` - `localhost:3003`
- `example4` - `localhost:3004`

Include the `Content-Type` header set to `application/json` and a body that looks like this:
```
{
    "message": {
        "data": "ewogICAgICAgICJwcm9wMSI6ICJhYmMiLAogICAgICAgICJwcm9wMiI6IDEyMwp9"
    }
}
```

The `data` value must be the object you want to insert to Firestore in Base64 format. The above example `data` will insert the following into Firestore:
```
{
    "prop1": "abc",
    "prop2": 123,
}
```

View the Firestore emulator at `localhost:4000/firestore/data`