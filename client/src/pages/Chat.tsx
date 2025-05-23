import { useState, useEffect } from "react";

interface Message {
  message: string;
}

const Chat = () => {
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState<Message[]>([]);
  const [ws, setWs] = useState<WebSocket | null>(null);

  useEffect(() => {
    const wsConnection = new WebSocket("ws://localhost:8080/ws");

    wsConnection.onopen = () => {
      console.log("Connected to WebSocket");
      setWs(wsConnection);
    };

    wsConnection.onmessage = (event) => {
      try {
        const msg = JSON.parse(event.data);
        setMessages((prevMessages) => [...prevMessages, msg]);
      } catch (error) {
        console.error("Failed to parse message:", error);
      }
    };

    wsConnection.onerror = (error) => {
      console.error("WebSocket error:", error);
      setWs(null);
    };

    wsConnection.onclose = () => {
      console.log("WebSocket closed");
      setWs(null);
    };

    return () => {
      if (wsConnection.readyState === WebSocket.OPEN) {
        wsConnection.close();
      }
    };
  }, []);

  const sendMessage = () => {
    if (ws?.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ message }));
      setMessage("");
    }
  };

  return (
    <div>
      <h1>Test ws connection</h1>
      <div>
        <input
          type="text"
          placeholder="message"
          value={message}
          onChange={(e) => setMessage(e.target.value)}
        />
        <button onClick={sendMessage}>Send!!</button>
      </div>
      <div>
        {messages.map((msg, index) => (
          <div key={index}>
            <p>{msg.message}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Chat;
