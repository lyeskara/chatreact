import { useState, useEffect } from "react"

import { JWTatom } from "../Jotai/Auth";
import { useAtom } from "jotai";

export default function Userdata() {
    // const [username, setUsername] = useState<string>()
    // const [password, setPassword] = useState<string>()
    const [JWT, _] = useAtom(JWTatom)
    console.log(JWT)
    const request = { method: 'GET', headers: { 'content-type': 'application/json', "Authorization": `Bearer ${JWT}` } };

    useEffect(() => {

        fetch("http://localhost:8000/userdata", request).then(response => {
            console.log(response.json())
        })

    }, [])

    return (
        <div>
            <section id="username"></section>
            <section id="password"></section>
        </div>
    )
}
