import { useEffect, useState } from "react";
import { JWTatom } from "../../utils/Auth";
import { useAtom } from "jotai";
import { Link, useNavigate } from "react-router-dom";

import styles from '../../styles/messenger.module.css'

function Rooms() {

    const [created, setCreated] = useState(false)
    const [room, setRoom] = useState<string>("")
    const [rooms, setRooms] = useState<string[]>([])

    const [JWT, _] = useAtom(JWTatom);

    const navigate = useNavigate()

    const api = import.meta.env.VITE_API_URL
    const socket_url = import.meta.env.VITE_SOCKET_URL

    const [socket, setSocket] = useState<WebSocket>()

    const enableCreation = () => created ? setCreated(false) : setCreated(true)

    useEffect(() => {
        const ws = new WebSocket(`${socket_url}/ws?token=${JWT}`);

        ws.onopen = () => {
            console.log("WebSocket connected");
        };

        ws.onclose = () => {
            console.log("WebSocket disconnected");
        };

        setSocket(ws);

        // Cleanup on component unmount
        return () => {
            ws.close();
        };

    }, []);

    const message = {
        type: 'newRoom',
        content: room,
    };

    function createRoom() {
        if (room === "") {
            alert("cannot create a room with an empty text as RoomName");
            return;
        }

        // If the socket exists and is open, send the message
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify(message));
        } else {
            console.error("WebSocket is not open");
        }

        navigate(`/${room}`);
        rooms != null ? setRooms([...rooms, room]) : setRooms([room]);
        setCreated(false);
    }

    useEffect(() => {
        // fetch rooms
        const request = { method: 'GET', headers: { 'content-type': 'application/json', "Authorization": `Bearer ${JWT}` } };
        fetch(`${api}/getRooms`, request).then(response => response.json()).then(data => setRooms(data))
    }, [])

    const Logout = () => {
        localStorage.removeItem('bearer')
        localStorage.removeItem('expirationTime')

        window.location.reload()
    }

    return (
        <section className={styles.Rooms}>
            <section id="create-room" className={styles.top}>
                <h1 className={styles.roomsHeader}>Rooms</h1>
                {created ?
                    <div className={styles.addRoom} >
                        <input id="newRoomName" type="text" placeholder="Choose a name" onChange={(e) => setRoom(e.target.value)} />
                        <button id='submitNewRoom' onClick={createRoom} className={styles.createBtn}>create</button>
                    </div>
                    :
                    <button id="addroom" onClick={enableCreation} className={styles.addbn}>add</button>
                }
            </section>

            <section id='rooms-list'>
                <div className={styles.roomsBox}>
                    <ul id='rooms' className={styles.list} >
                        {rooms && rooms.map((room) => {
                            return (
                                <li id='room' className={styles.list_item}><Link to={`/${room}`} className={styles.link}><span className={styles.roomNam}>{room}</span></Link></li>
                            )
                        })}
                    </ul>
                </div>
            </section>
            <section id='search-for room'></section>
            <button className={styles.LogoutBtn} onClick={Logout}>Log out</button>

        </section>
    )
}

export default Rooms
