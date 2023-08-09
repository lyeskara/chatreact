import { useState } from 'react'
import { Link } from 'react-router-dom';
import styles from '../../styles/signup.module.css'
import { useNavigate } from 'react-router-dom';

function Signup() {

    const navigate = useNavigate()
    const [username, setUsername] = useState<string>();
    const [password, setPassword] = useState<string>();
    const api = import.meta.env.VITE_API_URL

    const SignUp = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        const request = {
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: JSON.stringify({ "username": username, "password": password })
        };

        const response = await fetch(`${api}/signup`, request)

        if (response.status === 409) {
            const errorMessage = await response.text();
            alert(errorMessage);
        } else {
            let JWT = response.headers.get("Authorization")
            const token = JWT?.split(' ')[1];
            token && localStorage.setItem('bearer', token);
            let expirationTime = new Date();
            expirationTime.setHours(expirationTime.getHours() + 7);
            navigate(`/select a room`)
        }

    }

    return (
        <main className={styles.container}>
            <form className={styles.signin} onSubmit={SignUp}>
                <h2 className={styles.welcomeBack}>Create a profile</h2>
                <input
                    type="text"
                    placeholder="add username"
                    className={` ${styles.input}`}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <input
                    type="text"
                    placeholder="add password"
                    className={`${styles.input}`}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button className={styles.signinItem} type="submit">Sign up</button>
                <Link to="/login" className={styles.redirect}>already have an account? log in</Link>
            </form>
        </main>
    );
}

export default Signup
