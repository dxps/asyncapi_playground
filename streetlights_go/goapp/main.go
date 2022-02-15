
package main

import (
	"context"
  "log"
  "os"
  "os/signal"
  "syscall"
	"github.com/dxps/asyncapi_playground/streetlights_go/goapp/asyncapi"
)

func main() {
  router, err := asyncapi.GetRouter()
  if err != nil {
    log.Fatalf("error creating watermill router: %s", err)
  }

  
  amqpSubscriber, err := asyncapi.GetAMQPSubscriber(asyncapi.GetAMQPURI())
  if err != nil {
    log.Fatalf("error creating amqpSubscriber: %s", err)
  }

  asyncapi.ConfigureAMQPSubscriptionHandlers(router, amqpSubscriber)
  

  ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
  defer stop()
  if err = router.Run(ctx); err != nil {
    log.Fatalf("error running watermill router: %s", err)
  }
}

