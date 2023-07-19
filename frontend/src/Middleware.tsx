import { useEffect } from 'react';
import { AuthAtom, JWTatom } from '../Jotai/Auth'
import { useAtom } from 'jotai';
import { useNavigate } from 'react-router-dom';

export function AuthMiddleware({ children }: any) {

    const JWT = localStorage.getItem("bearer");
    const [Auth, setAuth] = useAtom(AuthAtom);
    const [_, setJWT] = useAtom(JWTatom);
    const navigate = useNavigate()
    
    useEffect(() => {
        if (JWT != null) {
            setAuth(true);
            setJWT(JWT);
        }
    }, [])

    if (Auth) {
        return children
    } else {
        navigate('/login')
    }
}