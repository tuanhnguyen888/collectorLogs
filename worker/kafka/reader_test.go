package kafka

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewKafkaReader(t *testing.T) {
	reader := NewKafkaReader()

	require.NotNil(t, reader)
}

//func TestReader_FetchMessage(t *testing.T) {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	reader := kafka.NewKafkaReader()
//
//	messages := make(chan kafkago.Message)
//
//	go func() {
//		err := reader.FetchMessage(ctx, messages)
//		require.NoError(t, err)
//	}()
//
//	select {
//	case <-ctx.Done():
//		require.Fail(t, "timed out waiting for messages")
//	case message := <-messages:
//		require.NotNil(t, message)
//		require.NotEmpty(t, message.Value)
//		assert.IsType(t, message, kafkago.Message{})
//	}
//
//}
