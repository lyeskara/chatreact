import type { NextPage } from "next";
import styles from "./page.module.css";
import Link from "next/link";

const SignIn: NextPage = () => {
  return (
    <form className={styles.signin}>
      <h2 className={styles.welcomeBack}>Welcome Back</h2>
      <input type="text" placeholder="add username" className={`${styles.signinChild} ${styles.input}`} />
      <input type="password" placeholder="add password" className={`${styles.signinChild} ${styles.input}`} />
      <button className={styles.signinItem}>Log in</button>
      <Link href={""} className={styles.redirect}>Don't have an account? Sign up</Link>
    </form>
  );
};

export default SignIn;