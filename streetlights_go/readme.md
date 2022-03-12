## Streetlights - Using Go

This is based on the official [Streetlights](https://www.asyncapi.com/docs/tutorials/streetlights) tutorial, but it uses Go instead of Node.js.

The setup is minimal, consisting of:
- A RabbitMQ instance
  - A popular message broker that talks AMQP protocol.
- An AsyncAPI spec file
  - Using this and AsyncAPI Generator, a Go app is generated.
  - Generated Go app plays the role of an event subscriber,<br/>
    subscribing to a topic and consuming the events that are published there.
- An event publisher
  - A simple cURL (as client to RabbitMQ's HTTP API) is used.
```
  .-----------.         .-----------------.          .--------------.
  | Publisher |         |   AMQP Broker   |          |  Subscriber  |
  |  (curl)   |-------->|    (RabbitMQ)   |--------->|   (Go app)   |
  '-----------'         '-----------------'          '--------------'
```


<br/>

### Prereqs

- [AsyncAPI Generator](https://github.com/asyncapi/generator/) (`ag`) tool
  - Use `pnpm install -g @asyncapi/generator` (or `npm`, if you prefer) to install it.
    - Ver. 1.9.0 was used (aavailable at the time of this writing).
- A local running RabbitMQ instance.
  - To quickly start it, use `docker run -d -p 15672:15672 -p 5672:5672 rabbitmq:3-management`


<br/>

### Steps

The following steps were performed:

1. Created the queue.
   - Through RabbitMQ's HTTP API, `light/measured` queue has been created using cURL:<br/>
     ```shell
     curl --user guest:guest -X PUT http://localhost:15672/api/queues/%2f/light%2fmeasured \
          -H 'content-type: application/json' -d '{"auto_delete":false, "durable":true}'
     ```
     Going to RabbitMQ's [Mgmt UI](http://localhost:15672/), you can verify that the queue is already there.<br/>
     Note that the forward slash (`/`) character, part of `light/measured` queue name, has been URL encoded as `%2f`.<br/>
   
2. Created the specification (`asyncapi.yaml` file).
   
3. Generated the Go code using `ag asyncapi.yaml @asyncapi/go-watermill-template -o goapp -p moduleName=github.com/dxps/asyncapi_playground/streetlights_go/goapp` <br/>
   Used `--debug` here just for more verbosity.
   ```shell
   ❯ ag --debug asyncapi.yaml @asyncapi/go-watermill-template -o goapp -p moduleName=github.com/dxps/asyncapi_playground/streetlights_go/goapp
    Unable to resolve template location at undefined. Package is not available locally. Error: Cannot find module '@asyncapi/go-watermill-template/package.json'
    Require stack:
    - /home/dxps/apps/node/lib/node_modules/noop.js
        at Function.Module._resolveFilename (internal/modules/cjs/loader.js:902:15)
        at resolveFileName (/home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/node_modules/resolve-from/index.js:29:39)
        ...
        at /home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/cli.js:154:9 {
    code: 'MODULE_NOT_FOUND',
    requireStack: [ '/home/dxps/apps/node/lib/node_modules/noop.js' ]
    }
    Template installation started because the template cannot be found on disk.
    Template @asyncapi/go-watermill-template successfully installed in /home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/node_modules/@asyncapi/go-watermill-template.
    Version of used template is 0.1.37.


    Done! ✨
    Check out your shiny new generated files at /home/dxps/dev/dxps-gh/asyncapi_playground/streetlights_go/goapp.
    ❯
   ```

4. Started the generated Go app.
    ```shell
    ❯ cd goapp
    ❯ go mod tidy
    ❯ go run main.go
    ```
    The app startup should look like this:
    ```shell
    ❯ go run main.go 
    [watermill] 2022/02/16 09:55:17.943214 connection.go:98: 	level=INFO  msg="Connected to AMQP" 
    [watermill] 2022/02/16 09:55:17.943297 router.go:231: 	level=INFO  msg="Adding handler" handler_name=OnLightMeasured topic=light/measured 
    [watermill] 2022/02/16 09:55:17.944593 router.go:453: 	level=INFO  msg="Starting handler" subscriber_name=OnLightMeasured topic=light/measured 
    [watermill] 2022/02/16 09:55:17.945158 subscriber.go:166: 	level=INFO  msg="Starting consuming from AMQP channel" amqp_exchange_name= amqp_queue_name=light/measured topic=light/measured 

    ```

5. Published a message
   - to the default exchange
   - using `routing_key` with the same name as the previously created queue,<br/>
    based on which the message is routed to that queue.
   ```shell
   curl --user guest:guest -X POST http://localhost:15672/api/exchanges/%2f/amq.default/publish \
   -H 'content-type: application/json' \
   -d ' {
          "properties":{},
          "routing_key":"light/measured",
          "payload":"{ \"id\":1, \"lumens\":2, \"sentAt\": \"2022-02-16\" }",
          "payload_encoding":"string"
        }'
   ```
6. And the Go app should get the message, and as part of the consumption it prints to the standard output:
   ```
   2022/02/16 09:58:52 received message payload: { "id":1, "lumens":2, "sentAt": "2022-02-16" }
   ```


<br/>
