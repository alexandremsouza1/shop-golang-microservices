package consumers

import (
	"github.com/meysamhadeli/shop-golang-microservices/internal/services/product_service/shared/delivery"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

func HandleConsumeCreateMedia(queue string, msg amqp.Delivery, productDeliveryBase *delivery.MediaDeliveryBase) error {

	log.Infof("Message received on queue: %s with message: %s", queue, string(msg.Body))
	return nil
}
