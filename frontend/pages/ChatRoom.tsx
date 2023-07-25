import { useEffect, useState, useCallback } from 'react'
import { useParams } from 'react-router-dom'
import Messenger from './Messenger'
import { JWTatom } from "../Jotai/Auth";
import { useAtom } from "jotai";

interface Message {
    message: String
    roomName: String
    username: String
}

function ChatRoom() {
    const [message, setMessage] = useState<string>("")
    const [messages, setMessages] = useState<Message[]>([])

    const [MemoMessages, SetMemoMessages] = useState<Set<Message>>()

    const [JWT, _] = useAtom(JWTatom);
    const roomName = useParams().room

    useEffect(() => {
        const request = { method: 'GET', headers: { 'content-type': 'application/json', "Authorization": `Bearer ${JWT}` } };
        fetch(`http://localhost:8000/getMessages?roomName=${roomName}`, request).then(response => response.json()).then(data => setMessages(data))
    }, [roomName])

    const createMessage = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        if (message === "") {
            alert("you need to write something!")
        } else {
            // send message to database. stored with timestamp. room id. user id
            const socket = new WebSocket(`ws://localhost:8000/ws?token=${JWT}`)
            const Newmessage = {
                type: 'newMessage',
                msg: message,
                room: roomName
            };
            socket.onopen = () => socket.send(JSON.stringify(Newmessage));
            socket.onmessage = (event) => setMessages([...messages, JSON.parse(event.data)]);
            socket.onclose = () => console.log("WebSocket disconnected")
        }

    }
    return (
        <>
            <section>
                <Messenger />
            </section>
            <section id='chat-room'>
                <h1>messages</h1>
                <ul id='messages'>
                    {messages && messages.map((msg) => <li>{msg.message}</li>)}
                </ul>

                <form onSubmit={createMessage}>
                    <input type="text" placeholder="write message here" value={message} onChange={(e) => setMessage(e.target.value)} />
                    <button type='submit'>send</button>
                </form>
            </section>
        </>
    )
}

export default ChatRoom
