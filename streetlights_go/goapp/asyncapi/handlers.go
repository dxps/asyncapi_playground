
package asyncapi

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
)

// OnLightMeasured subscription handler for light/measured.        
func OnLightMeasured(msg *message.Message) error {
    log.Printf("received message payload: %s", string(msg.Payload))

    var lm LightMeasured
    err := json.Unmarshal(msg.Payload, &lm)
    if err != nil {
        log.Printf("error unmarshalling message: %s, err is: %s", msg.Payload, err)
    }
    return nil
}

