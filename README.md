### listen2rabbit
### Ru

Демо пакеты go-модуля "прослушивателя" `listen2rabbit4dirx` данных в очереди RabbitMQ.   
Если данные появились в нашей очереди, go-модулем `call2handler`, вызывается метод обработчика Directum RX, через сформированную гиперссылку `Hyperlink`.    
Интеграционный метод обработчика, который выполняется при переходе по гиперссылке, должен существовать в Directum RX.  

Для прослушивания сообщений, запустить модуль:    
 
	listen2rabbit4dirx  

Демо модуль для публикаций сообщений (`SAP_A, SAP_B, SAP_C`) в очереди:    
 
	publisher2listener SAP_A


***Схема обмена данными (scheme exchange of data):***
			
```mermaid
graph TB

  SubGraph1Flow
  subgraph "RabbitMQ"
  SubGraph1Flow(Queue)
  SubGraph1Flow -- Listen --> RoutingKey

  end
 
  SubGraph2
  subgraph "Directum RX"
  SubGraph2Flow(Method of handler data a queue)
  SubGraph2Flow -- GetMessages --> SubGraph1Flow
  SubGraph2Flow -- SendResponse --> SubGraph1Flow
  end

  subgraph "Listener of queue"
  Node1[Go module `listen2rabbit4dirx`] -- Listen --> SubGraph1Flow
  Node1 --> Node2[Go module `call2handler`] --> SubGraph2[Hyperlink] -- Transition to link --> SubGraph2Flow  

end
```  

### En

Demo packages of module `listen2rabbit4dirx` it "listener" of RabbitMQ queue data.    
If data appeared in the our queue, the Directum RX method of handler is called the generated `Hyperlink`, via go module `call2handler`.   
Handler method of integration, while executed when a hyperlink is followed must exist in Directum RX.       

To listen messages, run the module:  
 
	listen2rabbit4dirx

Demo module for publishing messages (`SAP_A, SAP_B, SAP_C`) to to a queue:  
 
	publisher2listener SAP_A
