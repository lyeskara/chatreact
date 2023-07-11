import type { NextPage } from "next";
import styles from "./page.module.css";
import Link from "next/link";

const SignIn: NextPage = () => {
  return (
    <form className={styles.signin}>
      <h2 className={styles.welcomeBack}>Create a profile</h2>
      <input type="text" placeholder="add username" className={`${styles.signinChild} ${styles.input}`} />
      <button className={styles.signinItem}>Sign up</button>
      <Link href={""} className={styles.redirect}>already have an account? log in</Link>
    </form>
  );
};

export default SignIn;

