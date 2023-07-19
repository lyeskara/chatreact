import { useEffect, useState } from "react";

interface ComponentProps {
    SDP: RTCSessionDescription;
    JWT: string;
}

interface resultType {
    username: string
}

export default function Search(props: ComponentProps) {
    const [query, setQuery] = useState<string>()
    const [results, setResults] = useState<resultType[]>()

    useEffect(() => {
        const request = { method: 'POST', headers: { 'content-type': 'application/json', "Authorization": `Bearer ${props.JWT}` }, body: JSON.stringify({ "search": query }) }
        fetch('http://localhost:8000/getUsers', request)
            .then(response => response.json())
            .then((data) => setResults(data))
    }, [query])

    const SendSDP = () => {

    }

    return (
        <>
            <input type="text"
                placeholder="search by username"
                onChange={(e) => setQuery(e.target.value)} />
            {results && results.map((user) => {
                return (
                    <ul>
                        <li ><button onClick={SendSDP}>{user.username} </button></li>

                    </ul>
                )

            })}
        </>
    )
}

