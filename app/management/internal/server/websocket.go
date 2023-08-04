package server

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type WebSocketHandler struct {
	logger *log.Helper
	pool   *ConnectionPool
}

type ConnectionPool struct {
	pool chan *websocket.Conn
	mu   sync.Mutex // guards
}

func NewConnectionPool(capacity int) *ConnectionPool {
	return &ConnectionPool{
		pool: make(chan *websocket.Conn, capacity),
	}
}

func createConnection() (*websocket.Conn, error) {
	/* 根据实际情况创建 WebSocket 连接 */
	/* 使用 websocket.Dial 方法 */
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewWebSocketHandler(logger log.Logger, pool *ConnectionPool) *WebSocketHandler {
	return &WebSocketHandler{
		logger: log.NewHelper(logger),
		pool:   pool,
	}
}

func (p *ConnectionPool) Get() (*websocket.Conn, error) {
	select {
	case conn := <-p.pool:
		return conn, nil
	default:
		return nil, websocket.ErrBadHandshake
	}
}

func (p *ConnectionPool) Put(conn *websocket.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()
	select {
	case p.pool <- conn:
	default:
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}
}

func (h *WebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.logger.Error("upgrade websocket connection error: %v", err)
		return
	}

	/* 从连接池中获取连接 */
	poolConn, err := h.pool.Get()
	if err != nil {
		/* 连接池为空，需要重新创建连接或者返回错误响应给客户端 */
		poolConn, err = createConnection()
		if err != nil {
			_ = conn.Close()
			h.logger.Errorf("failed to create new connection: %v", err)
			return
		}
		h.pool.Put(poolConn)
	}

	go h.handleConnection(conn)
}

func (h *WebSocketHandler) handleConnection(conn *websocket.Conn) {

	defer func() {
		/* 处理完毕后将连接放回连接池 */
		h.pool.Put(conn)

		/* 关闭连接 */
		if err := conn.Close(); err != nil {
			h.logger.Errorf("error closing websocket connection: %v", err)
		}
	}()

	/* 设置 Ping\Pong 处理函数 */
	conn.SetPingHandler(func(appData string) error {
		return conn.WriteControl(websocket.PongMessage, []byte("Pong"), time.Now().Add(time.Second*10))
	})
	conn.SetPongHandler(func(appData string) error {
		return conn.WriteControl(websocket.PingMessage, []byte("Ping"), time.Now().Add(time.Second*10))
	})
	/* 设置关闭函数 */
	conn.SetCloseHandler(func(code int, text string) error {
		h.logger.Infof("code: %d, text: %s", code, text)
		return nil
	})
	/* 实现连接的逻辑: 通过conn.ReadMessage 和 conn.WriteMessage 方法读写数据 */
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			/* 正常关闭连接，无需报错 */
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				h.logger.Infof("WebSocket connection closed normal: %v", err)
				return
			}
			h.logger.Errorf("WebSocket connection closed unexpectedly: %v", err)
			return
		}
		h.logger.Infof("Rev: %s, messageType: %v", message, messageType)

		switch messageType {
		case websocket.TextMessage:
			/* 处理文本消息 */
			h.handleTextMessage(string(message), conn)
		case websocket.BinaryMessage:
		/* 处理二进制消息 */
		case websocket.CloseMessage:
			/* 收到关闭消息时，退出循环 */
			break
		default:
		}
	}
}

func (h *WebSocketHandler) handleTextMessage(message string, conn *websocket.Conn) {
	/* 基于具体的业务逻辑处理文本消息 */
	var req map[string]interface{}

	/* 解析数据 */
	if err := json.Unmarshal([]byte(message), &req); err != nil {
		h.logger.Errorf("unmarshal message from websocket error: %v", err)
	}
	err := conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		h.logger.Errorf("Error writing message to WebSocket: %v", err)
		return
	}
}
