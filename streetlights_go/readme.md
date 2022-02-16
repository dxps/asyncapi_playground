## Streetlights - Using Go

This is based on the official [Streetlights](https://www.asyncapi.com/docs/tutorials/streetlights) tutorial, but it uses Go instead of Node.js.

The setup is minimal:
- a RabbitMQ instance that talks AMQP protocol
- an AsyncAPI spec file that generates the Go app through AsyncAPI Generator
- the generated Go app subscribes to an "exchange" (part of RabbitMQ messaging model)
- a publisher (simple curl as client to RabbitMQ's API)
```
  .-----------.         .-----------------.          .--------------.
  | Publisher |         |   AMQP Broker   |          | Subscriber   |
  |  (curl)   |-------->|    (RabbitMQ)   |--------->|  (Go app)    |
  '-----------'         '-----------------'          '--------------'
```


<br/>

### Prereqs

- [AsyncAPI Generator](https://github.com/asyncapi/generator/) (`ag`)
  - Used `pnpm install -g @asyncapi/generator` that installed ver. 1.9.0 (available at the time of this writing).
- A local instance of RabbitMQ with.
  - To quickly start it, use `docker run -d -p 15672:15672 -p 5672:5672 rabbitmq:3-management`


<br/>

### Steps

The following steps were performed:

1. Created the queue.
   - Through RabbitMQ's API, `light/measured` queue has been created using cURL:<br/>
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
        at resolveFrom (/home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/node_modules/resolve-from/index.js:43:9)
        at module.exports (/home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/node_modules/resolve-from/index.js:46:47)
        at utils.getTemplateDetails (/home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/lib/utils.js:190:30)
        at /home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/lib/generator.js:367:26
        at new Promise (<anonymous>)
        at Generator.installTemplate (/home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/lib/generator.js:360:12)
        at Generator.generate (/home/dxps/apps/node/pnpm-global/5/node_modules/.pnpm/@asyncapi+generator@1.9.0/node_modules/@asyncapi/generator/lib/generator.js:180:73)
        at processTicksAndRejections (internal/process/task_queues.js:95:5)
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

5. Published a message to the exchange.<br/>
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
6. And the app should get the message and print to the output:
   ```
   2022/02/16 09:58:52 received message payload: { "id":1, "lumens":2, "sentAt": "2022-02-16" }
   ```


<br/>
