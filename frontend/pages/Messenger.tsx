import { useEffect, useState } from "react";
import { JWTatom } from "../Jotai/Auth";
import { useAtom } from "jotai";
import { Link } from "react-router-dom";
function Messenger() {
    const [created, setCreated] = useState(false)
    const [room, setRoom] = useState<string>("")
    const [rooms, setRooms] = useState<string[]>([])

    const [JWT, _] = useAtom(JWTatom);

    const enableCreation = () => created ? setCreated(false) : setCreated(true)

    function createRoom() {
        // send room name to the database.
        if (room === "") {
            alert("cannot create a room an empty text as RoomName")
        } else {
            const socket = new WebSocket(`ws://localhost:8000/ws?token=${JWT}`)
            const message = {
                type: 'newRoom',
                content: room,
            };
            socket.onopen = () => socket.send(JSON.stringify(message));
            socket.onmessage = (event) => setRooms([...rooms, event.data])
            socket.onclose = () => console.log("WebSocket disconnected")
        }

    }

    useEffect(() => {
        // fetch rooms
        const request = { method: 'GET', headers: { 'content-type': 'application/json', "Authorization": `Bearer ${JWT}` } };
        fetch("http://localhost:8000/getRooms", request).then(response => response.json()).then(data => setRooms(data))
    }, [])
    return (
        <>
            <section id="create-room">
                <button onClick={enableCreation}>create room</button>
                {created &&
                    <>
                        <input type="text" placeholder="Choose a name" onChange={(e) => setRoom(e.target.value)} />
                        <button onClick={createRoom}>create</button>
                    </>
                }
            </section>

            <section id='rooms-list'>
                <h1>rooms</h1>
                {rooms && rooms.map((room) => {
                    return (
                        <ul id='messages'>
                            <li id='message'><Link to={`/${room}`}>{room}</Link></li>
                        </ul>)
                }
                )
                }
            </section>
            <section id='search-for room'></section>

        </>
    )
}

export default Messenger
