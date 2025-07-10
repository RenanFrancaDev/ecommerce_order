package rabbitmq

import (
	"encoding/json"
	"log"

	"ecommerce_order/internal/domain"

	"github.com/streadway/amqp"
)

type Publisher struct{}

func (p Publisher) Publish(order domain.Order) error {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		log.Println("Erro ao conectar no RabbitMQ:", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println("Erro ao abrir canal:", err)
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("orders", false, false, false, false, nil)
	if err != nil {
		log.Println("Erro ao declarar fila:", err)
		return err
	}

	body, err := json.Marshal(order)
	if err != nil {
		log.Println("Erro ao serializar ordem:", err)
		return err
	}

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		log.Println("Erro ao publicar mensagem:", err)
		return err
	}

	return nil
}
