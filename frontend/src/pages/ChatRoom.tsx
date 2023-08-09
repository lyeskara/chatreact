import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import Rooms from './Rooms'
import { JWTatom } from "../../utils/Auth";
import { useAtom } from "jotai";
import styles from '../../styles/chatroom.module.css'
import Message from '../../utils/MessagesInterface'
import Messages from '../components/Messages';
import dataType from '../../utils/dataTypeInterface'


export default function ChatRoom() {

    const [message, setMessage] = useState<string>("")
    const [messages, setMessages] = useState<Message[]>([])
    const [currentUser, setCurrentUser] = useState<string>("")
    const [roomSwitch, setRoomSwitch] = useState<boolean>(false)

    const [JWT, _] = useAtom(JWTatom);
    const roomName = useParams().room

    // environment variables for http api and websocket urls
    const api = import.meta.env.VITE_API_URL
    const socket_url = import.meta.env.VITE_SOCKET_URL

    const [socket, setSocket] = useState<WebSocket>()

    useEffect(() => {
        const request = { method: 'GET', headers: { 'content-type': 'application/json', "Authorization": `Bearer ${JWT}` } };
        try {
            fetch(`${api}/getMessages?roomName=${roomName}`, request)
                .then(response => response.json())
                .then((data: dataType) => {
                    setMessages(data.Messages);
                    setCurrentUser(data.currentUser)
                    roomSwitch ? setRoomSwitch(false) : setRoomSwitch(true)
                })
        } catch (error) {
            console.log(error)
        }
    }, [roomName])

    const Newmessage = {
        type: 'newMessage',
        msg: message,
        room: roomName
    };

    useEffect(() => {
        const ws = new WebSocket(`${socket_url}/ws?token=${JWT}`);

        ws.onopen = () => {
            console.log("WebSocket connected");
        };

        ws.onmessage = (event) => {
            const receivedMessage: Message = JSON.parse(event.data);
            if (receivedMessage.roomName && receivedMessage.roomName === roomName) {
                setMessages(prevMessages => {
                    return prevMessages ? [...prevMessages, receivedMessage] : [receivedMessage];
                });
            }
        };

        ws.onclose = () => {
            console.log("WebSocket disconnected");
        };

        setSocket(ws);

        return () => {
            ws.close();
        };

    }, [roomSwitch]);

    const createMessage = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        if (message === "") {
            alert("you need to write something!");
            return;
        }

        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify(Newmessage));
            setMessage('');
        } else {
            console.error("WebSocket is not open");
        }
    };


    return (
        <main className={styles.container}>
            <Rooms />
            <section id='chat-room' className={styles.chat}>
                <h1 className={styles.msgHeader}>{roomName}</h1>
                <Messages messages={messages} currentUser={currentUser} />
                <form onSubmit={createMessage} className={styles.Input}>
                    <input id='newmessage' type="text" placeholder="write message here" value={message} className={styles.inputt} onChange={(e) => setMessage(e.target.value)} />
                    <button id='submitmessage' type='submit'>send</button>
                </form>
            </section>
        </main>
    )
}

