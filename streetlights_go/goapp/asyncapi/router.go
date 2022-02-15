
package asyncapi

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// GetRouter returns a watermill router. 
func GetRouter() (*message.Router, error){
	logger := watermill.NewStdLogger(false, false)
	return message.NewRouter(message.RouterConfig{}, logger)
}


// ConfigureAMQPSubscriptionHandlers configures the router with the subscription handler.    
func ConfigureAMQPSubscriptionHandlers(r *message.Router, s message.Subscriber) {

  r.AddNoPublisherHandler(
    "OnLightMeasured",     // handler name, must be unique
    "light/measured",         // topic from which we will read events
    s,
    OnLightMeasured, 
  )

}    

  