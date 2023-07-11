import Link from 'next/link';
import styles from './page.module.css';
import Image from 'next/image';

export default function nav() {
  return (
      <nav className={styles.navs}>
        <header className={styles.div}>
          <ul className={styles['nav-links']}>
            <li>
              <Link href="/" className={styles['text-wrapper-2']}>Profile</Link>
            </li>
            <li>
              <Link href="#" className={styles['text-wrapper-2']}>Call Room</Link>
            </li>
            <li>
              <Link href="#" className={styles['text-wrapper-2']}>Log Out</Link>
            </li>
          </ul>
          <div className={styles['search-box']}>
            <Image className={styles.img} alt="Img" src="/searchinconsvg.svg" width={100} height={24} />
            <div className={styles['search-for-anything']}>
              <input
                type="text"
                className={styles['input']}
                placeholder="Search for anything"
              />
            </div>
          </div>
        </header>
      </nav>
  )
}
