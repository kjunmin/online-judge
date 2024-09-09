import Link from "next/link"

export default function Home() {
  return (
    <div className="flex flex-col items-center gap-3">
      <section>
        Title
      </section>
      <section className="">
        <ul>
          <li><Link href="/problem/1">Problem 1</Link></li>
          <li><Link href="/problem/2">Problem 2</Link></li>
          <li><Link href="/problem/3">Problem 3</Link></li>
        </ul>
      </section>
    </div>
  );
}
