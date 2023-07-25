// pipeline\p1_find\application\queue_handler.go
package application

import "sync"

// QueueHandler é responsável por manipular a fila de nomes para processamento em paralelo.
type QueueHandler struct {
	queue []string   // Fila de nomes
	mutex sync.Mutex // Mutex para sincronização das operações na fila
}

// NewQueueHandler cria uma nova instância de QueueHandler.
func NewQueueHandler() *QueueHandler {
	return &QueueHandler{
		queue: make([]string, 0),
	}
}

// AddToQueue adiciona um nome à fila.
func (handler *QueueHandler) AddToQueue(name string) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	handler.queue = append(handler.queue, name)
}

// GetFromQueue obtém um nome da fila.
func (handler *QueueHandler) GetFromQueue() (string, bool) {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	if len(handler.queue) > 0 {
		name := handler.queue[0]
		handler.queue = handler.queue[1:]
		return name, true
	}

	return "", false
}

// IsQueueEmpty verifica se a fila está vazia.
func (handler *QueueHandler) IsQueueEmpty() bool {
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	return len(handler.queue) == 0
}
