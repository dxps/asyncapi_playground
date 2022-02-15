## Streetlights - Using Go

This is based on the official [Streetlights](https://www.asyncapi.com/docs/tutorials/streetlights) tutorial, but it uses Go instead of Node.js.



<br/>

### Prereqs

- [AsyncAPI Generator](https://github.com/asyncapi/generator/) (`ag`)
  - Used `pnpm install -g @asyncapi/generator` that installed ver. 1.9.0 (available at the time of this writing).
- A local instance of RabbitMQ with.
  - To quickly start it, use `docker run -d -p 15672:15672 -p 5672:5672 rabbitmq:3-management`


<br/>

### Steps

The following steps were performed:

1. Created the specification (`asyncapi.yaml` file).
   
2. Generated the Go code using `ag asyncapi.yaml @asyncapi/go-watermill-template -o goapp -p moduleName=github.com/dxps/asyncapi_playground/streetlights_go/goapp`
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

3. Go app start
    ```shell
    ❯ cd goapp
    ❯ go mod tidy
    ❯ go run main.go
    ```

4. mm


<br/>
