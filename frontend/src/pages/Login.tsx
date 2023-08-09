import { useState } from 'react'
import { Link } from 'react-router-dom';
import styles from '../../styles/signup.module.css'
import { useNavigate } from 'react-router-dom';

function Login() {

    const [username, setUsername] = useState<string>();
    const [password, setPassword] = useState<string>();
    const navigate = useNavigate()

    const api = import.meta.env.VITE_API_URL

    const SignUp = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        const request = {
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: JSON.stringify({ "username": username, "password": password })
        }

        const response = await fetch(`${api}/login`, request)

        if (response.status === 409) {
            const errorMessage = await response.text();
            alert(errorMessage);
        } else {
            let JWT = response.headers.get("Authorization")
            const token = JWT?.split(' ')[1];
            token && localStorage.setItem('bearer', token);
            let expirationTime = new Date();
            expirationTime.setHours(expirationTime.getHours() + 6);
            localStorage.setItem('expirationTime', expirationTime.toISOString())
            navigate(`/select a room`)
        }
    }

    return (
        <main className={styles.container}>
            <form className={styles.signin} onSubmit={SignUp}>
                <h2 className={styles.welcomeBack}>Welcome back</h2>
                <input
                    type="text"
                    id='username'
                    placeholder="add username"
                    className={` ${styles.input}`}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <input
                    type="password"
                    id='password'
                    placeholder="add password"
                    className={` ${styles.input}`}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button id='loginBtn' className={styles.signBtn} type="submit">Login</button>
                <Link to={"/signup"} className={styles.redirect}>Don't have an account? sign up</Link>
            </form>
        </main>
    );
}

export default Login
