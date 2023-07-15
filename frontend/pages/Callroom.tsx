import { AuthAtom, JWTatom } from '../Jotai/Auth'
import { useAtom } from 'jotai';
function Callroom() {
  const [Auth, _] = useAtom(AuthAtom);
    const [JWT, __] = useAtom(JWTatom);

    console.log(Auth, JWT)
  return (
    <div>
      <h1>CALL ROOM</h1>
    </div>
  )
}

export default Callroom
