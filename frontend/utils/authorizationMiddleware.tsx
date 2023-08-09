import { useEffect } from 'react';
import { AuthAtom, JWTatom } from '../utils/Auth'
import { useAtom } from 'jotai';
import { useNavigate } from 'react-router-dom';

export function AuthMiddleware({ children }: any) {

    const JWT = localStorage.getItem("bearer");
    const expiration = localStorage.getItem("expirationTime")
    const [Auth, setAuth] = useAtom(AuthAtom);
    const [_, setJWT] = useAtom(JWTatom);
    const navigate = useNavigate()
    let currentTime = (new Date()).toISOString()
    // check token is still valid   
    useEffect(() => {
        if (JWT && expiration && expiration > currentTime ) {
            setAuth(true);
            setJWT(JWT);
        } else {
            setAuth(false)
        }
    }, [])

    if (Auth) {
        return children
    } else {
        navigate('/login')
    }
}