import { useEffect,  useRef, useState } from "react";
import Styles from "./styles/chat.module.css"
import { Navigate } from "react-router-dom";
import { useAuth } from "./providers/authentication";

const Chat = () => {
  const ws = useRef<WebSocket | null>(null);
  const { isLoggedIn } = useAuth()
  const [messages, setMessages] = useState<string[]>([]);
  const [input, setInput] = useState("");


  useEffect(() => {
    ws.current = new WebSocket(`ws://localhost:5000/ws?room=${1}`);

    ws.current.onmessage = (event) => {
      setMessages((prev) => [...prev, event.data]);
    };

    return () => {
      ws.current?.close();
    };
  }, []);

  const sendMessage = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key != "Enter") {
        return 
    }
    if (ws.current && ws.current.readyState === WebSocket.OPEN) {
      ws.current.send(input);
      setInput("");
    }
  };

    return (
        <div className={Styles.chat_container}>
                {messages.map((msg, i) => (
                <div key={i}>{msg}</div>
                ))}
            <input className={Styles.chat_message} 
                onKeyDown={sendMessage}
                value={input}
                onChange={(e) => setInput(e.target.value)}
                placeholder="Type message..."></input>
        </div>
    )
}

export default Chat;