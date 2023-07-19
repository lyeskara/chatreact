import { useState } from 'react'
import { Link } from 'react-router-dom';
import styles from '../styles/signup.module.css'
import { AuthAtom, JWTatom } from '../Jotai/Auth'
import { useAtom } from 'jotai';
import { useNavigate } from 'react-router-dom';

function Login() {
    
    const [username, setUsername] = useState<string>();
    const [password, setPassword] = useState<string>();
    const [_, setAuth] = useAtom(AuthAtom)
    const [__, setJWT] = useAtom(JWTatom)
    const navigate = useNavigate()
    const SignUp = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        const request = { method: 'POST', headers: { 'content-type': 'application/json' }, body: JSON.stringify({ "username": username, "password":password }) }
        const response = await fetch("http://localhost:8000/login", request)
        if (response.status === 409) {
            const errorMessage = await response.text();
            alert(errorMessage);
        } else {
            let JWT = response.headers.get("Authorization")
            if (JWT != null) {
                let token = JWT.split(' ')[1];
                localStorage.setItem('bearer', token);
                setAuth(true);
                setJWT(token);
            }
            navigate('/') 
        }
    }

    return (
        <form className={styles.signin} onSubmit={SignUp}>
            <h2 className={styles.welcomeBack}>Welcome back</h2>
            <input
                type="text" 
                placeholder="add username"
                className={`${styles.signinChild} ${styles.input}`}
                onChange={(e) => setUsername(e.target.value)} 
            />
            <input
                type="text" 
                placeholder="add password"
                className={`${styles.signinChild} ${styles.input}`}
                onChange={(e) => setPassword(e.target.value)} 
            />
            <button className={styles.signinItem} type="submit">Login</button>
            <Link to={"/signup"} className={styles.redirect}>Don't have an account? sign up</Link>
        </form>
    );
}

export default Login
