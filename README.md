# QueueLocal 

# :jigsaw: THIS CLI IS WORK IN PROGRESS...

A CLI to run local commands to test specific URLs (For example calling lambda code locally to test SNS or SQS calls.)

## Download

Go to [releases](https://github.com/PrashamTrivedi/QueueLocal/releases) and download the latest binary according to your platform. You can see changelog and discuss about the releases [here](https://github.com/PrashamTrivedi/QueueLocal/discussions). 

## All options

`taskConnector --help`

Running above command will list all subcommands and their description.

## Getting started.

By default this CLI looks for `taskConnector.json` file in the directory where you are running the CLI. And runs commands described in the JSON file. Optionally you can pass your own JSON file using `-c` or `--configFile` option in any subcommands.

### Command format 

The JSON file expects key value format. Where the key will be your URL and the value will be command to run against the URL. For example below is an expected JSON.

```json
{
    "localhost:4000":"node -v",
    "https://sqs.eu-cdef-1.amazonaws.com/ZZZZZZZZZ/sqs-xyz-qwert":"go run xyz",
    "https://sqs.eu-cdef-1.amazonaws.com/ZZZZZZZZZ/sqs-xyz-asdf":"node -e asdf.js",
    "arn:aws:sns:eu-west-1:ZZZZZZZZ:testTopic":"sh test.sh"
}
```

## Running CLI and code changes :memo:

Start the CLI using `taskConnector startServer` command. By default it will listen to port `5000`, in case this port is busy, or you want it to listen to another port, run `taskConnector startServer -p 80` command. 

In your code, specially in local environment, change the environment variables for your SQS or SNS from 
`https://sqs.eu-cdef-1.amazonaws.com/ZZZZZZZZZ/sqs-xyz-asdf` to `http://localhost:500?key=https://sqs.eu-cdef-1.amazonaws.com/ZZZZZZZZZ/sqs-xyz-asdf` or `arn:aws:sns:eu-west-1:ZZZZZZZZ:testTopic` to `http://localhost:80?key=arn:aws:sns:eu-west-1:ZZZZZZZZ:testTopic`.

After this when you run the `startServer` command, the CLI will find the command associated with key and will run the command.

# TODO :wrench:

- [ ] Allowing body, header and query params in command.
- [ ] Testing 
