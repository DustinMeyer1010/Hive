import React, { useEffect, useState, useRef } from "react";

interface Message {
  username: string;
  message: string;
}

function App() {
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [username, setUsername] = useState<string>("");
  const [message, setMessage] = useState<string>("");
  const [chat, setChat] = useState<Message[]>([]);
  const chatRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/ws");
    setWs(socket);

    socket.onmessage = (event) => {
      const msg: Message = JSON.parse(event.data);
      setChat((prevChat) => [...prevChat, msg]);
      scrollToBottom();
    };

    return () => {
      socket.close();
    };
  }, []);

  const sendMessage = () => {
    if (!username || !message || !ws) return;
    ws.send(JSON.stringify({ username, message }));
    setMessage("");
  };

  const scrollToBottom = () => {
    if (chatRef.current) {
      chatRef.current.scrollTop = chatRef.current.scrollHeight;
    }
  };

  return (
    <div style={{ maxWidth: 600, margin: "auto", padding: 20 }}>
      <h2>Live Chat</h2>
      <div
        ref={chatRef}
        style={{
          height: 300,
          overflowY: "scroll",
          border: "1px solid #ccc",
          padding: 10,
          marginBottom: 10,
        }}
      >
        {chat.map((msg, idx) => (
          <div key={idx}>
            <b>{msg.username}</b>: {msg.message}
          </div>
        ))}
      </div>
      <input
        placeholder="Your name"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
        style={{ width: "30%", marginRight: 10 }}
      />
      <input
        placeholder="Type a message"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        style={{ width: "50%", marginRight: 10 }}
        onKeyDown={(e) => e.key === "Enter" && sendMessage()}
      />
      <button onClick={sendMessage}>Send</button>
    </div>
  );
}

export default App;
