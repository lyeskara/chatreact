import styles from '../../styles/chatroom.module.css'
import Message from '../../utils/MessagesInterface'
interface MessagesProps {
  messages: Message[],
  currentUser: string
}
const Messages: React.FC<MessagesProps> = ({ messages, currentUser }) => {
  return (
    <div className={styles.messageBox}>
      <ul id='messages' className={styles.msgs}>
        {messages && messages.map((msg) =>
          <li className={msg.username === currentUser ? styles.YourMessage : styles.NotYourMessage}>
            {msg.username === currentUser ?
              <>
                <span className={styles.Message}>{msg.message}</span>
                <span className={styles.name}>{msg.username}</span>
              </>
              :
              <>
                <span className={styles.name}>{msg.username}</span>
                <span className={styles.Message}>{msg.message}</span>
              </>
            }
          </li>
        )}
      </ul>
    </div>
  )
}

export default Messages


